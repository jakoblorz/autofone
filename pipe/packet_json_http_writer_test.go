package pipe

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jakoblorz/f1-metrics-transformer/packets"
	"github.com/stretchr/testify/assert"
)

func TestPacketJSONHTTPWriter_handle(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sig := make(chan int)
	serv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var decoded packets.PacketMotionData
		err := json.NewDecoder(r.Body).Decode(&decoded)
		assert.NoError(t, err)
		assert.Equal(t, motionData, decoded)
		close(sig)
	}))
	defer serv.Close()

	writer := WritePacketToHTTP(ctx, serv.URL, JSONEncoding, StdoutResponseHandler)

	writeCh := make(chan interface{})
	writer.ReadFrom(writeCh)

	go writer.Process()
	writeCh <- &motionData

	<-sig
}
