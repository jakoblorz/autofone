package privateapi

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"golang.org/x/net/websocket"
)

type SnapshotsClient struct {
	*client
}

type SnapshotCreateResponse struct {
	File      string `json:"file"`
	SignedURL string `json:"signed_url"`
	ExpiresAt string `json:"expires_at"`
}

func (s *SnapshotsClient) Health() error {
	s.Lock()
	var (
		wsBaseURL   = s.client.baseURL
		httpBaseURL = s.client.baseURL
		token       = s.client.token
	)
	s.Unlock()
	wsBaseURL = strings.Replace(wsBaseURL, "https://", "wss://", 1)
	wsBaseURL = strings.Replace(wsBaseURL, "http://", "ws://", 1)
	ws, err := websocket.Dial(fmt.Sprintf("%s/api/w1/devices/health", wsBaseURL), "", httpBaseURL)
	if err != nil {
		return err
	}
	websocket.Message.Send(ws, fmt.Sprintf("Bearer %s", token))
	for {
		select {
		case <-s.ctx.Done():
			return nil
		case <-time.After(1 * time.Second):
			var expectMsg = primitive.NewObjectID().Hex()
			err = websocket.Message.Send(ws, expectMsg)
			if err != nil {
				return err
			}
			var actualMsg string
			err = websocket.Message.Receive(ws, &actualMsg)
			if err != nil {
				return err
			}
			if actualMsg != expectMsg {
				return fmt.Errorf("unexpected message: %s", actualMsg)
			}

		}
	}
}

func (s *SnapshotsClient) Create() (r *SnapshotCreateResponse, err error) {
	r = new(SnapshotCreateResponse)
	err = s.JSON("POST", "/api/v1/snapshots", nil, &r)
	return
}

func (s *SnapshotsClient) CreateAndUpload(file string) (err error) {
	var r *SnapshotCreateResponse
	r, err = s.Create()
	if err != nil {
		return
	}

	var f_hdl *os.File
	f_hdl, err = os.Open(file)
	if err != nil {
		return
	}
	defer f_hdl.Close()

	var req *http.Request
	req, err = http.NewRequest(http.MethodPut, r.SignedURL, f_hdl)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/octet-stream")
	s.client.Lock()
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.client.token))
	s.client.Unlock()

	client := clientPool.Get().(*http.Client)
	res, err := client.Do(req)
	clientPool.Put(client)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to upload snapshot: %s", res.Status)
	}
	return
}

func (s *SnapshotsClient) CreateAndWrite(rdr io.Reader) (err error) {
	var r *SnapshotCreateResponse
	r, err = s.Create()
	if err != nil {
		return
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPut, r.SignedURL, rdr)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/octet-stream")
	s.client.Lock()
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.client.token))
	s.client.Unlock()

	client := clientPool.Get().(*http.Client)
	res, err := client.Do(req)
	clientPool.Put(client)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to upload snapshot: %s", res.Status)
	}
	return
}
