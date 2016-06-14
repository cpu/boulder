package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	netmail "net/mail"
	"os"
	"strings"
	"time"

	"github.com/jmhodges/clock"
	"gopkg.in/gorp.v1"

	"github.com/letsencrypt/boulder/cmd"
	blog "github.com/letsencrypt/boulder/log"
	"github.com/letsencrypt/boulder/mail"
	"github.com/letsencrypt/boulder/sa"
)

type mailer struct {
	clk           clock.Clock
	log           blog.Logger
	dbMap         *gorp.DbMap
	mailer        mail.Mailer
	subject       string
	emailTemplate string
	destinations  []string
	checkpoint    interval
	sleepInterval time.Duration
}

type interval struct {
	start int
	end   int
}

func (i *interval) isSane() error {
	if i.start < 0 || i.end < 0 {
		return errors.New(fmt.Sprintf(
			"interval start (%d) and end (%d) must both be positive integers",
			i.start, i.end))
	}

	if i.start > i.end && i.end != 0 {
		return errors.New(fmt.Sprintf(
			"interval start value (%d) is greater than end value (%d)",
			i.start, i.end))
	}

	return nil
}

func (m *mailer) run() error {
	// Do not allow a start larger than the # of destinations
	if m.checkpoint.start > len(m.destinations) {
		return errors.New(fmt.Sprintf(
			"interval start value (%d) is greater than number of destinations (%d)",
			m.checkpoint.start,
			len(m.destinations)))
	}
	// Do not allow a negative sleep interval
	if m.sleepInterval < 0 {
		return errors.New(fmt.Sprintf(
			"sleep interval (%d) is < 0", m.sleepInterval))
	}
	// If there is no endpoint specified, use the total # of destinations
	if m.checkpoint.end == 0 {
		m.checkpoint.end = len(m.destinations)
	}
	for i, dest := range m.destinations {
		if i < m.checkpoint.start || i >= m.checkpoint.end {
			continue
		}
		if strings.TrimSpace(dest) == "" {
			continue
		}
		err := m.mailer.SendMail([]string{dest}, m.subject, m.emailTemplate)
		if err != nil {
			return err
		}
		m.clk.Sleep(m.sleepInterval)
	}
	return nil
}

func main() {
	var from = flag.String("from", "", "From header for emails. Must be a bare email address.")
	var subject = flag.String("subject", "", "Subject of emails")
	var toFile = flag.String("toFile", "", "File containing a list of email addresses to send to, one per file.")
	var bodyFile = flag.String("body", "", "File containing the email body in plain text format.")
	var dryRun = flag.Bool("dryRun", true, "Whether to do a dry run.")
	var sleep = flag.Duration("sleep", 60*time.Second, "How long to sleep between emails.")
	var start = flag.Int("start", 0, "Line of input file to start from.")
	var end = flag.Int("end", 99999999, "Line of input file to end before.")
	type config struct {
		NotifyMailer struct {
			cmd.DBConfig
			cmd.PasswordConfig
			cmd.SMTPConfig
		}
	}
	var configFile = flag.String("configFile", "", "File containing a JSON config.")

	flag.Parse()
	if from == nil || subject == nil || bodyFile == nil || configFile == nil {
		flag.Usage()
		os.Exit(1)
	}

	seven := 7
	_, log := cmd.StatsAndLogging(cmd.StatsdConfig{}, cmd.SyslogConfig{StdoutLevel: &seven})

	configData, err := ioutil.ReadFile(*configFile)
	cmd.FailOnError(err, fmt.Sprintf("Reading %s", *configFile))
	var cfg config
	err = json.Unmarshal(configData, &cfg)
	cmd.FailOnError(err, "Unmarshaling config")

	dbURL, err := cfg.NotifyMailer.DBConfig.URL()
	cmd.FailOnError(err, "Couldn't load DB URL")
	dbMap, err := sa.NewDbMap(dbURL, 10)
	cmd.FailOnError(err, "Could not connect to database")

	// Load email body
	body, err := ioutil.ReadFile(*bodyFile)
	cmd.FailOnError(err, fmt.Sprintf("Reading %s", *bodyFile))

	address, err := netmail.ParseAddress(*from)
	cmd.FailOnError(err, fmt.Sprintf("Parsing %s", *from))

	toBody, err := ioutil.ReadFile(*toFile)
	cmd.FailOnError(err, fmt.Sprintf("Reading %s", *toFile))
	destinations := strings.Split(string(toBody), "\n")

	checkpointRange := interval{
		start: *start,
		end:   *end,
	}

	checkpointErr := checkpointRange.isSane()
	cmd.FailOnError(checkpointErr, "Building checkpoint range")

	var mailClient mail.Mailer
	if *dryRun {
		mailClient = mail.NewDryRun(address.Address, log)
	} else {
		smtpPassword, err := cfg.NotifyMailer.PasswordConfig.Pass()
		cmd.FailOnError(err, "Failed to load SMTP password")
		mailClient = mail.New(
			cfg.NotifyMailer.Server,
			cfg.NotifyMailer.Port,
			cfg.NotifyMailer.Username,
			smtpPassword,
			address.Address)
	}
	err = mailClient.Connect()
	cmd.FailOnError(err, fmt.Sprintf("Connecting to %s:%s",
		cfg.NotifyMailer.Server, cfg.NotifyMailer.Port))
	defer func() {
		err = mailClient.Close()
		cmd.FailOnError(err, "Closing mail client")
	}()

	m := mailer{
		clk:           cmd.Clock(),
		log:           log,
		dbMap:         dbMap,
		mailer:        mailClient,
		subject:       *subject,
		destinations:  destinations,
		emailTemplate: string(body),
		checkpoint:    checkpointRange,
		sleepInterval: *sleep,
	}

	err = m.run()
	cmd.FailOnError(err, "mailer.send returned error")
}
