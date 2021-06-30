/**
 * Auth :   liubo
 * Date :   2021/6/30 17:03
 * Comment: 邮件
 */

package crazyapp

import (
	"github.com/badforlabor/gocrazy/crazy3rd/glog"
	"gopkg.in/gomail.v2"
	"strings"
)

type IEmail interface {
	SendEmailTo(to[]string, subject, body string) error
}
func NewEmail(host string, port int, account string, password string) IEmail {
	var e = &oneEmail{}
	e.Host = host
	e.Port = port
	e.Account = account
	e.Password = password

	return e
}
func NewEmailExample() IEmail {
	return NewEmail("smtp.qq.com", 465, "563568850@qq.com", "password-xxxxxx")
}

type oneEmail struct {
	Host     string
	Port     int
	Account  string
	Password string
}
func (self *oneEmail) SendEmailTo(to []string, subject, body string) error {

	glog.Infof("发送邮件:from=%v, to=%v, subject=%v, body=%v", self.Account, to, subject, body)

	var from = self.Account

	// 邮件中是HTML格式。需要字符转换
	body = strings.Replace(body, " ", "&nbsp", -1)
	body = strings.Replace(body, "\n", "<br>", -1)

	m := gomail.NewMessage()
	m.SetAddressHeader("From", from, from)  // 发件人
	m.SetHeader("To",  // 收件人
		to...,
	)
	m.SetHeader("Subject", subject)  // 主题
	m.SetBody("text/html", "<body><div>" + body + "</div></body>")  // 正文
	//m.Attach("/home/Alex/lolcat.jpg")

	//d := gomail.NewDialer("smtp.qq.com", 465, from, "xaodkjrskzzmbfec")
	d := gomail.NewDialer(self.Host, self.Port, self.Account, self.Password)
	var err = d.DialAndSend(m)

	if err != nil {
		glog.Warningf("发送邮件失败:%v", err.Error())
	}

	return err
}

/*
// example

var emailWorker = crazyapp.NewEmail("smtp.qq.com", 465, "563568850@qq.com", "---password---")
func sendEmail(subject, body string) bool {
	var to = []string{"505700330@qq.com"}
	var e = emailWorker.SendEmailTo(to, subject, body)
	if e != nil {
		//netLog.Warnln("发送邮件错误:", e.Error())
	}
	return e == nil
}

*/