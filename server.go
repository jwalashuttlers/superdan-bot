package main

import (
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

// Server represents the superdan-bot server
type Server struct {
	bot *tb.Bot
}

// NewServer returns a new server
func NewServer() (*Server, error) {
	b, err := tb.NewBot(
		tb.Settings{
			Token:  cfg.Telegram.APIKey,
			Poller: &tb.LongPoller{Timeout: 10 * time.Second},
		},
	)
	if err != nil {
		return nil, err
	}

	srv := &Server{
		bot: b,
	}

	return srv, nil
}

// Run starts the server
func (srv *Server) Run() {
	// Initialize handlers
	srv.handlers()

	// Start the bot
	srv.bot.Start()
}

// handleStart handles `/start` command
func (srv *Server) handleStart(m *tb.Message) {
	srv.bot.Send(m.Sender, "Welcome to SuperDanBot")
}

// handleGroupAdd will be called whenever the bot is added to a group
func (srv *Server) handleGroupAdd(m *tb.Message) {
	srv.bot.Send(m.Chat, "Thanks for adding to group")
}
