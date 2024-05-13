package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func InitSession(token string) (*discordgo.Session, error) {
	var session *discordgo.Session
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new Discord session: %w", err)
	}

	// session.Identify.Intents = discordgo.IntentsAll
	session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildVoiceStates)

	return session, nil
}
