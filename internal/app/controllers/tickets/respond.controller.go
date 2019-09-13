package tickets

import "net/smtp"

func (c *Controller) Respond(ticketID string, authorID string, message string) error {
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		smtp.PlainAuth("", "aeksapsy", "PI427233897n+3", "smtp.gmail.com"),
		"aeksapsy@gmail.com",
		[]string{"michael.tolis@gmail.com"},
		[]byte(message),
	)

	return err
}
