package bot

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/Hotmonth/discord_voice_recorder_bot/internal/discord"
	"github.com/Hotmonth/discord_voice_recorder_bot/internal/lib/logger/sl"
	"github.com/bwmarrin/discordgo"
)

var Session *discordgo.Session

func InitBot(BotToken string, log *slog.Logger) error {
	session, err := discord.InitSession(BotToken)
	if err != nil {
		log.Error("Failed to initialize the bot", sl.Err(err))
	}

	session.AddHandler(HandleVoiceStateUpdate)

	err = session.Open()
	if err != nil {
		log.Error("Failed to open the connection", sl.Err(err))
	}
	Session = session
	return nil

}

func CloseBot(log *slog.Logger) {
	err := Session.Close()
	if err != nil {
		log.Error("Failed to close the connection", sl.Err(err))
	}
}

func GetAllVoiceChannels() []*discordgo.Channel {
	channels, err := Session.GuildChannels("1237839145796374621")
	if err != nil {
		return nil
	}

	voiceChannels := make([]*discordgo.Channel, 0)
	for _, channel := range channels {
		if channel.Type == discordgo.ChannelTypeGuildVoice {
			voiceChannels = append(voiceChannels, channel)
		}
	}
	return voiceChannels
}

func HandleVoiceStateUpdate(s *discordgo.Session, vs *discordgo.VoiceStateUpdate) {
	// Check if session and voice state update are not nil
	if s == nil || vs == nil {
		fmt.Println("Error: Session or voice state update is nil.")
		return
	}

	// Check if a new voice channel was created
	if vs.ChannelID != "" && (vs.BeforeUpdate == nil || vs.ChannelID != vs.BeforeUpdate.ChannelID) {
		channel, err := s.Channel(vs.ChannelID)
		if err != nil {
			fmt.Println("Error retrieving channel information:", err)
			return
		}
		if channel == nil {
			fmt.Println("Error: Retrieved channel is nil.")
			return
		}

		// Check if the created channel is a voice channel
		if channel.Type == discordgo.ChannelTypeGuildVoice {
			// Join the voice channel and start capturing audio
			go JoinVoiceChannelAndRecord(channel.ID)
		}
	}
}

func JoinVoiceChannelAndRecord(channelID string) {
	// Join the voice channel
	vc, err := Session.ChannelVoiceJoin("1237839145796374621", channelID, true, false)
	if err != nil {
		fmt.Println("Error joining voice channel:", err)
		return
	}
	defer vc.Disconnect()

	// Start capturing audio
	// You can implement your audio recording logic here
	fmt.Printf("Started recording in voice channel %s\n", channelID)

	// Wait for a predefined time or until the voice channel becomes empty
	time.Sleep(5 * time.Minute) // Change this to your desired time limit
}
