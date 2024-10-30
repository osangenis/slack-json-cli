package slack

import "io"

type OutputFormat string

const OuputFormat_PlainText = "plain"

func (me *MessageExport) Write(format OutputFormat, w io.Writer) {

}
