package dto

// 邮箱信息结构体
type MailConfig struct {
	Host        string `json:"host"`         // 服务器地址
	Port        int    `json:"port"`         // 端口
	DisplayName string `json:"display_name"` // 显示名称
	Username    string `json:"username"`     // 用户名
	Password    string `json:"password"`     // 密码
}

// 邮箱发送信息结构体
type MailAddress struct {
	From string   `json:"from"` // 发件人
	To   []string `json:"to"`   // 收件人
	CC   []string `json:"cc"`   // 抄送
	BCC  []string `json:"bcc"`  // 暗送
}

// 邮箱内容结构体
type MailContent struct {
	MailType string `json:"mail_type"` // 邮件类型
	Subject  string `json:"subject"`   // 主体
	Body     string `json:"body"`      // 内容
}
