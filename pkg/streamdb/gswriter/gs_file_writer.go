package gswriter

import (
	"context"
	"io"
	"os"
	"path"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"cloud.google.com/go/storage"
)

type fileWriter struct {
	*storage.BucketHandle
	Path string
}

func New(bucketHandle *storage.BucketHandle, path string) *fileWriter {
	return &fileWriter{
		BucketHandle: bucketHandle,
		Path:         path,
	}
}

func (g *fileWriter) WriteFrom(ctx context.Context, c <-chan string) {
	for name := range c {
		g.write(ctx, name)
	}
}

func (g *fileWriter) write(ctx context.Context, name string) {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := g.Object(path.Join(g.Path, primitive.NewObjectID().Hex())).NewWriter(ctx)
	defer w.Close()

	if _, err := io.Copy(w, f); err != nil {
		panic(err)
	}
	if err := w.Close(); err != nil {
		panic(err)
	}
}
