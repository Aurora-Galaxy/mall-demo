package utils

import (
	"gin_mall/conf"
	"github.com/jordan-wright/email"
	"net/smtp"
)

// SendMail 发送通知
func SendMail(text []byte, mail string) error {
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = conf.Config.Email.SmtpEmail
	// 设置接收方的邮箱
	e.To = []string{mail}
	//设置主题
	e.Subject = "gin_mall"
	//设置文件发送的内容
	e.Text = text
	//设置服务器相关的配置
	err := e.Send(conf.Config.Email.SmtpHost, smtp.PlainAuth("",
		conf.Config.Email.SmtpEmail, conf.Config.Email.SmtpPass, "smtp.qq.com"))
	if err != nil {
		return err
	}
	//kmibjpfvgamxfggh
	return nil
}
