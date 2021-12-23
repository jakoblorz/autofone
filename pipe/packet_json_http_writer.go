package pipe

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/jakoblorz/metrikxd/pkg/step"
)

type HTTPEncoding string

const (
	JSONEncoding HTTPEncoding = "json"
)

type HTTPResponseHandler func(*http.Response) interface{}

var (
	StdoutResponseHandler HTTPResponseHandler = func(res *http.Response) interface{} {
		_, err := io.Copy(os.Stdout, res.Body)
		if err != nil {
			panic(err)
		}
		return nil
	}
)

func WritePacketToHTTP(ctx context.Context, to string, encoding HTTPEncoding, responseHandler HTTPResponseHandler) step.Step {
	switch encoding {
	case JSONEncoding:
		p := &PacketJSONHTTPWriter{
			Pool: sync.Pool{
				New: func() interface{} {
					return &http.Client{}
				},
			},
			to:             to,
			handleResponse: responseHandler,
		}
		p.Step = step.Intermediate(ctx, p.handle)
		return p
	}
	return nil
}

type PacketJSONHTTPWriter struct {
	step.Step
	sync.Pool

	to             string
	handleResponse HTTPResponseHandler
}

func (u *PacketJSONHTTPWriter) getClient() *http.Client {
	return u.Pool.Get().(*http.Client)
}

func (u *PacketJSONHTTPWriter) putClient(c *http.Client) {
	u.Pool.Put(c)
}

func (u *PacketJSONHTTPWriter) handle(pack interface{}) interface{} {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("%+v", err)
			err = binary.Write(os.Stderr, binary.LittleEndian, pack)
			if err != nil {
				log.Printf("%+v", err)
			}
		}
	}()
	data, err := json.Marshal(pack)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", u.to, bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := u.getClient()
	defer u.putClient(client)

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	return u.handleResponse(res)
}
