package util

func GenEventType(channelType string, messageType string) string {
	return channelType + messageType
}
