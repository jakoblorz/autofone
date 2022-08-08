package httpwriter

import (
	"context"

	"github.com/jakoblorz/autofone/pkg/privateapi"
)

type fileWriter struct {
	privateapi.Client
}

func New(client privateapi.Client) *fileWriter {
	return &fileWriter{
		Client: client,
	}
}

func (f *fileWriter) WriteFrom(ctx context.Context, c <-chan string) {
	for name := range c {
		f.write(ctx, name)
	}
}

func (f *fileWriter) write(ctx context.Context, name string) {
	err := f.Snapshots().CreateAndUpload(name)
	if err != nil {
		panic(err)
	}
}
