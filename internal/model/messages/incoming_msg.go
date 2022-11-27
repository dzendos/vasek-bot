package messages

import (
	"fmt"
	"vasek-bot/internal/model/state"
)

type MessageSender interface {
	SendMessage(text string, userID int64) error
}

type Model struct {
	tgClient MessageSender
}

func New(tgClient MessageSender) *Model {
	return &Model{
		tgClient: tgClient,
	}
}

type Message struct {
	Text      string
	ChatID    int64
	UserID    int64
	MessageID int
	Video     bool
}

func (s *Model) IncomingMessage(msg *Message) error {
	// Trying to recognize the command.
	switch msg.Text {
	case "/start":
		return s.tgClient.SendMessage("Я Васек Великий и Всемогущий! Яблокособирающий и в деревню уезжающий! Отправляй мне видео того как ты занимаешься спортом, не то быть тебе васьком как мне всю жизнь!", msg.ChatID)
	case "/stat":
		text := fmt.Sprintf("Статистика:\n\nРомка Васек %d раз\nЖеня Васек %d раз\nМама Васек %d раз", state.RomkaState.Score, state.EvgState.Score, state.MomState.Score)

		if state.RomkaState.Score > state.EvgState.Score && state.RomkaState.Score > state.MomState.Score {
			text += "\n\n Ва-ха-ха, Ромка - главный Васек"
		}

		return s.tgClient.SendMessage(text, msg.ChatID)
	}

	if msg.Video {
		name := ""
		switch msg.UserID {
		case state.MomState.UserID:
			state.MomState.WasSent = true
			name = state.MomState.Name
		case state.EvgState.UserID:
			state.EvgState.WasSent = true
			name = state.EvgState.Name
		case state.RomkaState.UserID:
			state.RomkaState.WasSent = true
			name = state.RomkaState.Name
		}

		text := fmt.Sprintf("%s - молодец\n\nКандидаты в васьки:\n", name)
		if !state.RomkaState.WasSent {
			text += state.RomkaState.Name + "\n"
		}
		if !state.EvgState.WasSent {
			text += state.EvgState.Name + "\n"
		}
		if !state.MomState.WasSent {
			text += state.MomState.Name + "\n"
		}

		return s.tgClient.SendMessage(text, msg.ChatID)
	}

	return nil
}
