package tg

import (
	"log"

	"vasek-bot/internal/model/messages"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

type tokenGetter interface {
	Token() string
}

type Client struct {
	client *tgbotapi.BotAPI
}

func New(tokenGetter tokenGetter) (*Client, error) {
	client, err := tgbotapi.NewBotAPI(tokenGetter.Token())
	if err != nil {
		return nil, errors.Wrap(err, "cannot create NewBotAPI")
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) SendMessage(text string, userID int64) error {
	_, err := c.client.Send(tgbotapi.NewMessage(userID, text))
	if err != nil {
		return errors.Wrap(err, "client.Send")
	}
	return nil
}

func (c *Client) ShowAlert(text string, messageID string) error {
	alert := tgbotapi.NewCallback(messageID, text)

	_, err := c.client.Send(alert)

	if err != nil {
		return errors.Wrap(err, "client.ShowAlert")
	}

	return nil
}

func (c *Client) ListenUpdates(msgModel *messages.Model) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.client.GetUpdatesChan(u)

	log.Println("listening for messages")

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			err := msgModel.IncomingMessage(&messages.Message{
				Text:      update.Message.Text,
				UserID:    update.Message.From.ID,
				ChatID:    update.Message.Chat.ID,
				MessageID: update.Message.MessageID,
				Video:     (update.Message.Video != nil),
			})
			if err != nil {
				log.Println("error processing message: ", err)
			}
		}
	}
}
