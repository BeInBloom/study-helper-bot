package main

import (
	"flag"

	"github.com/BeInBloom/study-helper-bot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	//test
	token := mustToken()
	tgClient = telegram.New(tgBotHost, token)

	//fetcher = fetcher.New()

	//processor = processor.New()

	// consumer.Start()
}

func mustToken() string {
	token := flag.String("t", "", "token for acces to telegram bot")

	flag.Parse()

	if *token == "" {
		panic("no tg token")
	}

	return *token
}
