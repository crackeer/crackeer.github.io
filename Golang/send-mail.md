# 发送邮件

github地址：https://github.com/xhit/go-simple-mail/v2

代码使用示例：

```go
package test

import (
	"crypto/tls"
	"errors"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

// Mailer
type Mailer struct {
	SendList []string
	Subject  string
	client   *mail.SMTPClient
}

// NewMailer
//  @param users
//  @param subject
//  @return *Mailer
//  @return error
func NewMailer(users []string, subject string) (*Mailer, error) {

	if len(users) < 1 {
		return nil, errors.New("empty send list")
	}

	server := mail.NewSMTPClient()

	// SMTP Server
	server.Host = "mail.host.com"
	server.Port = 25
	server.Username = "yourname@mail.host.com"
	server.Password = "your-password"
	server.Encryption = mail.EncryptionSTARTTLS

	server.KeepAlive = false

	// Timeout for connect to SMTP Server
	server.ConnectTimeout = 10 * time.Second

	// Timeout for send the data and wait respond
	server.SendTimeout = 10 * time.Second

	// Set TLSConfig to provide custom TLS configuration. For example,
	// to skip TLS verification (useful for testing):
	server.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// SMTP client
	smtpClient, err := server.Connect()
	if err != nil {
		return nil, err
	}

	return &Mailer{
		client:   smtpClient,
		SendList: users,
		Subject:  subject,
	}, nil
}

// SendHTML
//  @receiver mailer
//  @param content
func (mailer *Mailer) SendHTML(content string) error {
	// New email simple html with inline and CC
	email := mail.NewMSG()
	email.SetFrom("【Subject】"+mailer.Subject+"<your-name@mail.host.name>").
		AddTo(mailer.SendList...).
		SetSubject("【通知】<your-name@mail.host.name>").SetBody(mail.TextHTML, content)
	if email.Error != nil {
		return email.Error
	}
	return email.Send(mailer.client)
}
