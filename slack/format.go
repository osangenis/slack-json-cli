package slack

import "io"

type FormatMessageFunc func(m *Message, w io.Writer)
type Formatter struct {
	MessageFormat FormatMessageFunc
}

func PlainTextFormatter() *Formatter {
	return &Formatter{
		MessageFormat: formatPlainMessage,
	}
}

func formatPlainMessage(m *Message, w io.Writer) {

}
