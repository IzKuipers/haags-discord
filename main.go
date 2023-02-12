package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token  string
	prefix = "!h"
	dg     *discordgo.Session
	err    error
)

func defineVars() {
	flag.StringVar(&Token, "t", "", "Bot Token for authentication")
	flag.Parse()
}

func main() {
	fmt.Println("Attempting to start Haags bot...")

	defineVars()

	tdg, terr := discordgo.New("Bot " + Token)

	dg = tdg
	err = terr

	if err != nil {
		return
	}

	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()

	if err != nil {
		fmt.Println("Failed to start bot:\n -> ", err)
		return
	}

	fmt.Println("Bot online and listening for input!")

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)

	<-sc

	fmt.Println("Stopping Haags bot...")
	dg.Close()
	fmt.Println("Done.")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := string(m.Content)

	split := strings.Split(content, " ")

	pf := split[0]

	str := ""

	if pf != prefix || m.Message.Author.ID == dg.State.User.ID {
		return
	}

	for i := 1; i < len(split); i++ {
		str += split[i] + " "
	}

	fmt.Println(str)

	s.ChannelMessageSendReply(m.ChannelID, convert(str), m.Reference())
}
