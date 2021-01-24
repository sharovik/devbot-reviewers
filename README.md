# Reviewers list event
The event which can help to understand which reviewers are available in the memory of the devbot.

## Installation guide
Go to devbot project directory and run the next in CMD:
``` 
git clone git@github.com:sharovik/devbotreviewers.git events/devbotreviewers
```
Then open the `defined-events.go` file and add there event
```go
package events

import (
//...
	"github.com/sharovik/devbot/events/devbotreviewers"
//...
)

//DefinedEvents collects all the events which can be triggered by the messages
var DefinedEvents = base.Events{}

func init() {
	DefinedEvents.Events = make(map[string]base.Event)
//...
	DefinedEvents.Events[devbotreviewers.EventName] = devbotreviewers.Event
//...
}

```
After that, to install the event please run 
``` 
make install
```


## Usage
Write in PM or tag the bot user in the channel with this message
```
show reviewers
```
As the result you will see the list of available bitbucket reviewers. These reviewers bot will use during the pull-request validation and creation.
