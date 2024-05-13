package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Hotmonth/discord_voice_recorder_bot/internal/bot"
	"github.com/Hotmonth/discord_voice_recorder_bot/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	// TODO: Init logger
	log := setupLogger(cfg.Env)
	log = log.With(slog.String("env", cfg.Env))

	log.Info("Starting the bot")

	bot.InitBot(cfg.BotToken, log)

	defer bot.CloseBot(log)

	log.Info("Bot is running")

	fmt.Println("Bot is running. Press Ctrl + C to exit.")

	// vcs := bot.GetAllVoiceChannels()
	// for _, c := range vcs {
	// 	fmt.Println(c)
	// }
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// fmt.Println(discord.GetAllVoiceChannelsByGuildID("1237839145796374621"))

	// TODO: Init storage

	// TODO: Init bot
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
