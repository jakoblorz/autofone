package writer

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/jakoblorz/autofone/packets/process"
	"github.com/jakoblorz/autofone/pkg/log"
)

var (
	clientPool = &sync.Pool{
		New: func() interface{} {
			return &http.Client{}
		},
	}
)

type HTTP struct {
	*process.P

	URL     string
	Verbose bool
	LogJSON bool
}

func (ch *HTTP) Write(m *process.M) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("%+v", err)
			err = binary.Write(os.Stderr, binary.LittleEndian, m.Pack)
			if err != nil {
				log.Printf("%+v", err)
			}
		}
	}()
	data, err := json.Marshal(m.Pack)
	if err != nil {
		panic(err)
	}

	if ch.Verbose || ch.LogJSON {
		message := fmt.Sprintf("posting with len = %d bytes json payload", len(data))
		if ch.LogJSON {
			message = fmt.Sprintf("%s: %s", message, string(data))
		}
		log.Print(message)
	}

	if len(ch.URL) == 0 {
		return
	}

	req, err := http.NewRequest("POST", strings.ReplaceAll(ch.URL, "{{packetID}}", fmt.Sprintf("%d", m.Header.GetPacketID())), bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := clientPool.Get().(*http.Client)
	res, err := client.Do(req)
	clientPool.Put(client)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	_, err = io.Copy(os.Stdout, res.Body)
	if err != nil {
		panic(err)
	}
}
