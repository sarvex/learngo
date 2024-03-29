package pipe

import (
	"bufio"
	"io"
)

// TextLog parses text based log lines.
type TextLog struct {
	reader io.Reader
}

// NewTextLog creates a text parser.
func NewTextLog(r io.Reader) *TextLog {
	return &TextLog{reader: r}
}

// Each yields records from a text log.
func (p *TextLog) Each(yield func(Record) error) error {
	defer readClose(p.reader)

	// Use the same record for unmarshaling.
	var r Record

	in := bufio.NewScanner(p.reader)

	for in.Scan() {
		if err := r.UnmarshalText(in.Bytes()); err != nil {
			return err
		}

		if err := yield(r); err != nil {
			return err
		}
	}

	return in.Err()
}
