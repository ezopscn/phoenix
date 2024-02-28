package utils

import (
	"crypto/tls"
	"github.com/itrepablik/gomail"
	"phoenix/dto"
	"phoenix/pkg/logx"
)

// 发送邮件
func SendMail(mailConfig *dto.MailConfig, mailAddress *dto.MailAddress, mailContent *dto.MailContent) (err error) {
	// 创建 Message 对象
	m := gomail.NewMessage()

	// m.SetHeader("From", smtpConfig.Username) // 发件人
	m.SetHeader("From", mailConfig.DisplayName+"<"+mailConfig.Username+">") // 增加发件人别名
	m.SetHeader("To", mailAddress.To...)                                    // 收件人，可以多个收件人，但必须使用相同的 SMTP 连接
	m.SetHeader("Cc", mailAddress.CC...)                                    // 抄送，可以多个
	m.SetHeader("Bcc", mailAddress.BCC...)                                  // 暗送，可以多个

	// 邮件主题
	m.SetHeader("Subject", mailContent.Subject)

	// text/html: 浏览器在获取到这种文件时会自动调用 html 的解析器对文件进行相应的处理
	// text/plain: 将文件设置为纯文本的形式，浏览器在获取到这种文件时并不会对其进行处理
	m.SetBody(mailContent.MailType, mailContent.Body)

	// 附件
	//m.Attach("a.sh")  // 附件文件，可以是文件，照片，视频等等
	//m.Attach("a.mp4") // 视频
	//m.Attach("a.jpg") // 照片

	// 配置服务器信息
	d := gomail.NewDialer(
		mailConfig.Host,
		mailConfig.Port,
		mailConfig.Username,
		mailConfig.Password,
	)

	// 关闭 SSL 协议认证
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// 发送邮件
	logx.INFO("发送邮件给用户:", mailAddress.To)
	err = d.DialAndSend(m)
	if err != nil {
		logx.ERROR(err)
	}
	return err
}
