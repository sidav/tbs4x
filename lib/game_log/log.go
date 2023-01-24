package game_log

import (
	"fmt"
	"strings"
)

type MessageType uint8

const (
	MSG_REGULAR = iota
	MSG_WARNING
)

type logMessage struct {
	Message string
	Count   int
	Type    MessageType
}

func (m *logMessage) getText() string {
	if m.Count > 1 {
		return fmt.Sprintf("%s (x%d)", m.Message, m.Count)
	} else {
		return m.Message
	}
}

type GameLog struct {
	Last_msgs     []*logMessage
	logWasChanged bool
}

func (l *GameLog) Init(length int) {
	l.Last_msgs = make([]*logMessage, length)
	for i := range l.Last_msgs {
		l.Last_msgs[i] = &logMessage{
			Message: "",
			Count:   1,
			Type:    MSG_REGULAR,
		}
	}
}

func (l *GameLog) AppendMessage(msg string) {
	msg = capitalize(msg)
	if l.Last_msgs[len(l.Last_msgs)-1].Message == msg {
		l.Last_msgs[len(l.Last_msgs)-1].Count++
	} else {
		for i := 0; i < len(l.Last_msgs)-1; i++ {
			l.Last_msgs[i] = l.Last_msgs[i+1]
		}
		l.Last_msgs[len(l.Last_msgs)-1] = &logMessage{Message: msg, Count: 1}
	}
	l.logWasChanged = true
}

func (l *GameLog) AppendMessagef(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	l.AppendMessage(msg)
}

func (l *GameLog) Warning(msg string) {
	l.AppendMessage(msg)
	l.Last_msgs[len(l.Last_msgs)-1].Type = MSG_WARNING
}

func (l *GameLog) Warningf(msg string, args ...interface{}) {
	l.AppendMessage(fmt.Sprintf(msg, args...))
	l.Last_msgs[len(l.Last_msgs)-1].Type = MSG_WARNING
}

func (l *GameLog) WasChanged() bool {
	was := l.logWasChanged
	l.logWasChanged = false
	return was
}

func capitalize(s string) string {
	if len(s) > 0 {
		return strings.ToUpper(string(s[0])) + s[1:]
	}
	return s
}
