package main

import (
	"flag"
	"fmt"
	"github.com/jbillote/YAB/config"
	"github.com/jbillote/YAB/util/logger"
	"gopkg.in/yaml.v2"

	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var (
	d *discordgo.Session
	uid string
	log logger.Logger
)

func onReady(s *discordgo.Session, e *discordgo.Ready) {
	log.Info("ToukaBot started")
	err := s.UpdateStatus(0, "")

	if err != nil {
		log.Error(err)
		log.Error("Unable to change status")
	}

	uid = e.User.ID
}

func onMessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	go ParseMessage(session, message)
}

func main() {
	log = logger.GetLogger("ToukaBot V2 Main")

	configLocation := flag.String("config", "./config.yml", "Configuration file")
	flag.Parse()

	configFile, err := os.Open(*configLocation)
	if err != nil {
		log.Fatal(err)
		log.Fatal(fmt.Sprintf("Unable to open config file at %s", *configLocation))
		os.Exit(-1)
	}

	c := config.Config{}
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&c)
	if err != nil {
		log.Fatal("Unable to decode config file")
		log.Fatal(err)
		os.Exit(-1)
	}

	config.SetConfig(&c)

	log.Info("Starting Discord session...")
	d, err = discordgo.New(c.BotToken)
	if err != nil {
		log.Fatal(err)
		log.Fatal("Failed to create Discord session")
		os.Exit(-1)
	}

	d.AddHandler(onReady)
	d.AddHandler(onMessageCreate)

	err = d.Open()
	if err != nil {
		log.Fatal(err)
		log.Fatal("Failed to create Discord websocket connection")
		os.Exit(-1)
	}

	// Start thread to change the status
	go RotateStatuses(d, c.Statuses, c.StatusInterval)

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, os.Kill)
	<-s

	return
}