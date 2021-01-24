package devbotreviewers

import (
	"fmt"
	"github.com/sharovik/devbot/internal/helper"

	"github.com/sharovik/devbot/internal/log"

	"github.com/sharovik/devbot/internal/container"
	"github.com/sharovik/devbot/internal/dto"
)

const (
	//EventName the name of the event
	EventName = "devbotreviewers"

	//EventVersion the version of the event
	EventVersion = "1.0.0"

	helpMessage = "Ask me `show reviewers` and I will show you the list of current reviewers."

	//The migrations folder, which can be used for event installation or for event update
	migrationDirectoryPath = "./events/devbotreviewers/migrations"
)

//EventStruct the struct for the event object. It will be used for initialisation of the event in defined-events.go file.
type EventStruct struct {
	EventName string
}

//Event - object which is ready to use
var Event = EventStruct{
	EventName: EventName,
}

//Execute method which is called by message processor
func (e EventStruct) Execute(message dto.BaseChatMessage) (dto.BaseChatMessage, error) {
	isHelpAnswerTriggered, err := helper.HelpMessageShouldBeTriggered(message.OriginalMessage.Text)
	if err != nil {
		log.Logger().Warn().Err(err).Msg("Something went wrong with help message parsing")
	}

	if isHelpAnswerTriggered {
		message.Text = helpMessage
		return message, nil
	}

	if len(container.C.Config.BitBucketConfig.RequiredReviewers) > 0 {
		message.Text = "So, here is the list:\n"
		for _, reviewer := range container.C.Config.BitBucketConfig.RequiredReviewers {
			message.Text += fmt.Sprintf("<@%s> - his public bitbucket UUID is `%s`\n", reviewer.SlackUID, reviewer.UUID)
		}
	}

	return message, nil
}

//Install method for installation of event
func (e EventStruct) Install() error {
	log.Logger().Debug().
		Str("event_name", EventName).
		Str("event_version", EventVersion).
		Msg("Triggered event installation")

	return container.C.Dictionary.InstallEvent(
		EventName,      //We specify the event name which will be used for scenario generation
		EventVersion,   //This will be set during the event creation
		"show reviewers", //Actual question, which system will wait and which will trigger our event
		"Give me a sec", //Answer which will be used by the bot
		"(?i)(show)(?:.+)(reviewers)", //Optional field. This is regular expression which can be used for question parsing.
		"",                 //Optional field. This is a regex group and it can be used for parsing the match group from the regexp result
	)
}

//Update for event update actions
func (e EventStruct) Update() error {
	return container.C.Dictionary.RunMigrations(migrationDirectoryPath)
}
