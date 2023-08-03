package utils

import (
	"log"
	"marketplace/internal/models"
	"net/smtp"

	"github.com/spf13/viper"
)

func ActivationSendByEmailForget(user models.User, hashForget string) error {
	from := viper.GetString("email.from")
	subject := "Subject: " + viper.GetString("email.subject") + "!\n"
	pass := viper.GetString("email.password")

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := "<html><body><h2>Ваш пароль был изменен, новый пароль!</h2><h1>" + hashForget + " <h1></body></html>"
	msg := []byte("From: " + from + "\n" +
		"To: " + user.Email + "\n" + subject + mime + body)

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{user.Email}, []byte(msg))

	if err != nil {
		log.Fatalf("Error: %s", err)
		return err
	}

	return nil
}

func ActivationSendByEmail(user *models.User, hashForget string) error {
	from := viper.GetString("email.from")
	subject := "Subject: " + viper.GetString("email.subject") + "!\n"
	pass := viper.GetString("email.password")

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := "<html><body><h2>Ваш код подтверждение,НИКОМУ НЕ ПЕРЕДАВАЙТЕ!</h2><h1>" + hashForget + " <h1></body></html>"
	msg := []byte("From: " + from + "\n" +
		"To: " + user.Email + "\n" + subject + mime + body)

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{user.Email}, []byte(msg))

	if err != nil {
		log.Fatalf("Error: %s", err)
		return err
	}

	return nil
}

func ActivationSendByEmailMessage(user models.User, hashForget string, rout string) error {

	from := viper.GetString("email.from")
	subject := "Subject: " + viper.GetString("email.subject") + "\n"
	pass := viper.GetString("email.password")
	msg := "From: " + from + "\n" +
		"To: " + user.Email + "\n" +
		subject + "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		`<html>
<body>
    <h3>Откройте ссылку для активации вашего профиля, ссылка будет работать 12 часов. Иначе, необходимо сбросить пароль !!!</h3>
    <h3>Если вы не регистрировались на сайте <a href="https://stock.alif.tj">stock.alif.tj</a> игнорируйте данное письмо.</h3>
    <div style="margin: 24px; text-align: center;"><a style="display: inline-block; background: #2196f3; color: #fff; padding: 20px 50px 20px 50px; border-radius: 5px; text-decoration: none; font-family: Tahoma; font-size: 25px; line-height: 1; font-weight: 100;" href="` + rout + hashForget + `">Активировать</a></div>
</body>
</html>
`

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{user.Email}, []byte(msg))

	if err != nil {
		log.Fatalf("Error: %s", err)
		return err
	}

	return nil
}
