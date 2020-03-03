package main

import tb "gopkg.in/tucnak/telebot.v2"

func (srv *Server) handlers() {
	srv.bot.Handle("/start", srv.handleStart)
	srv.bot.Handle(tb.OnAddedToGroup, srv.handleGroupAdd)
}
