package main

import (
	"crypto/tls"
	"fmt"
	"net/mail"
	"os"

	"github.com/namsral/flag"

	"github.com/matcornic/hermes/v2"
	"gopkg.in/gomail.v2"
)

var addr string

func main() {
	flag.StringVar(&addr, "addr", "unkown", "xxxx")
	flag.Parse()
	fmt.Println(addr)
	fmt.Println(os.Getenv("cos_addr"))
	h := hermes.Hermes{
		Product: hermes.Product{
			Name: "FakeNotion",
			Link: "https://FakeNotion.com/",
			Logo: "https://github.com/matcornic/hermes/blob/master/examples/gopher.png?raw=true",
		},
	}
	email := hermes.Email{
		Body: hermes.Body{
			Name: "Jon Snow",
			Intros: []string{
				"Welcome to Hermes! We're very excited to have you on board.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "To get started with Hermes, please click here:",
					Button: hermes.Button{
						Color: "#22BC66", // Optional action button color
						Text:  "Confirm your account",
						Link:  "https://hermes-example.com/confirm?token=d9729feb74992cc3482b350163a1a010",
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}

	// Generate an HTML email with the provided contents (for modern clients)
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	// Generate the plaintext version of the e-mail (for clients that do not support xHTML)
	emailText, err := h.GeneratePlainText(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	// https://ethereal.email/create // 生成email账号密码

	d := gomail.NewDialer("smtp.ethereal.email", 587, "xjzplaul5t2rsoex@ethereal.email", "TfKTNzKErPMd1nutUQ")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	m := gomail.NewMessage()
	from := mail.Address{
		Name:    "FakeNotion",
		Address: "xjzplaul5t2rsoex@ethereal.email",
	}
	m.SetHeader("From", from.String())
	m.SetHeader("To", "clylia217@gmail.com")
	m.SetHeader("Subject", "testing send email")

	m.SetBody("text/plain", emailText)
	m.AddAlternative("text/html", emailBody)
	err = d.DialAndSend(m)
	if err != nil {
		panic(err)
	}

	SendEmail()
}

// SendEmail will send email to given address
func SendEmail() {
	m := gomail.NewMessage()
	m.SetHeader("From", "alex@example.com")
	m.SetHeader("To", "522383345@qq.com", "clylia217@gmail.com")
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

	d := gomail.NewDialer("smtp.ethereal.email", 587, "xjzplaul5t2rsoex@ethereal.email", "TfKTNzKErPMd1nutUQ")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
