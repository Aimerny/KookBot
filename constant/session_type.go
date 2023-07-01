package constant

const (
	ChannelGroup     string = "GROUP"
	ChannelPerson    string = "PERSON"
	ChannelBroadcast string = "BROADCAST"
)

// MessageType 1:文字消息, 2:图片消息，3:视频消息，4:文件消息， 8:音频消息，9:KMarkdown，10:card 消息，255:系统消息

const (
	Text      string = "_1"
	Picture   string = "_2"
	Video     string = "_3"
	File      string = "_4"
	Audio     string = "_8"
	KMarkDown string = "_9"
	Card      string = "_10"
	SYSTEM    string = "_255"
	ALL       string = "*"
)
