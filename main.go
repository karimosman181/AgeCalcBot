package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	//set env varaibles
	os.Setenv("SLACK_BOT_TOKEN", "")
	os.Setenv("SLACK_APP_TOKEN", "")

	//create bot
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	var ExamplesArray []string
	strings.Join(ExamplesArray, "my yob is  20") 

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "age calculator",
		Examples:    ExamplesArray,
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			yearOfBirth, err := strconv.Atoi(request.Param("year"))
			if err != nil {
				fmt.Println("error")
			}

			t := time.Now()
			currentYear := t.Year()

			age := currentYear - yearOfBirth

			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
