package utils
import (
	"regexp"
	"strings"
	"fmt"
	"net/smtp"
)

type Message struct {
	Email   string
	Content string
	Errors  map[string]string
}

func(msg *Message) Validate() bool {
	msg.Errors = make(map[string]string)

	re := regexp.MustCompile(".+@.+\\..+")
	matched := re.Match([]byte(msg.Email))
	if matched == false {
		msg.Errors["Email"] = "Please enter a valid email address"
	}
	if strings.TrimSpace(msg.Content) == "" {
		msg.Errors["Content"] = "Please write a message"
	}

	return len(msg.Errors) == 0
}

func(msg *Message) Deliver() error {
	body := fmt.Sprintf("Reply-To: %v\r\nSubject: New Message\r\n%v", msg.Email, msg.Content)

	username := "***"
	password := "***"
	auth := smtp.PlainAuth("", username, password, "smtp.gmail.com")

	return smtp.SendMail("smtp.gmail.com:587", auth, msg.Email, []string{msg.Email}, []byte(body))
}
