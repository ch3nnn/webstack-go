package constant

const (
	ProtocolShell = 1
	ProtocolHTTP  = 2

	HttpMethodGet  = 1
	HttpMethodPost = 2

	NotifyStatusNo      = 1
	NotifyStatusFailed  = 2
	NotifyStatusStopped = 3
	NotifyStatusKeyword = 4

	NotifyTypeEmail   = 1
	NotifyTypeWebhook = 2

	IsUsedYES = 1
	IsUsedNo  = -1
)

var ProtocolText = map[int64]string{
	ProtocolShell: "SHELL",
	ProtocolHTTP:  "HTTP",
}

var HttpMethodText = map[int64]string{
	HttpMethodGet:  "GET",
	HttpMethodPost: "POST",
}

var NotifyStatusText = map[int64]string{
	NotifyStatusNo:      "不通知",
	NotifyStatusFailed:  "失败通知",
	NotifyStatusStopped: "结束通知",
	NotifyStatusKeyword: "结果关键字匹配通知",
}

var NotifyTypeText = map[int64]string{
	NotifyTypeEmail:   "邮件",
	NotifyTypeWebhook: "Webhook",
}

var IsUsedText = map[int64]string{
	IsUsedYES: "启用",
	IsUsedNo:  "禁用",
}
