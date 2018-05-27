package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	BotID   string
	Session *discordgo.Session
)

func Connect(token string) {
	// Create a new Discord session using the provided bot token.
	var err error
	Session, err = discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println("!Hiba,", err)
		return
	}

	// Get the account information.
	u, err := Session.User("@me")
	if err != nil {
		fmt.Println("sikertelen lekérdezés,", err)
	}

	// Store the account ID for later use.
	BotID = u.ID

	fmt.Println("Kapcsolódva")
}

func Start() {
	// Open the websocket and begin listening.
	err := Session.Open()
	if err != nil {
		fmt.Println("Sikertelen kapcsolódás,", err)
		return
	}

	fmt.Println("A ParrotDise bot sikeresen elindult.")

	return
}

func AddHandler(handler interface{}) {
	Session.AddHandler(handler)
}
