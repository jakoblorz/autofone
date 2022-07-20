package writer

import "github.com/jakoblorz/autofone/packets/process"

type Writer interface {
	Write(m *process.M)
}
