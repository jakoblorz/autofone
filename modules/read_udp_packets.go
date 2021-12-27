package modules

import (
	"context"
	"fmt"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/pipe"
	"github.com/jakoblorz/metrikxd/pkg/log"
	"github.com/jakoblorz/metrikxd/pkg/step"
	"github.com/jakoblorz/metrikxd/www"
	"github.com/jakoblorz/metrikxd/www/partials"
)

type ReadUDPPackets struct {
	context.Context
	www.Page

	Host string
	Port int

	conn   net.Conn
	reader *pipe.PacketReader

	options       *pipe.PacketReaderOptions
	applyThenWith step.Step
}

func NewUDPPacketReader(ctx context.Context, host string, port int, initialOptions *pipe.PacketReaderOptions) *ReadUDPPackets {
	p := &ReadUDPPackets{
		Context: ctx,

		options: initialOptions,

		Host: host,
		Port: port,
	}
	p.Page = www.Page{"f1-game", p.renderF1GamePage, p.renderF1GamePartial, www.EmptySSEHandler}
	return p
}

func (r *ReadUDPPackets) getSharedProps() partials.RenderF1GameSharedProps {
	return partials.RenderF1GameSharedProps{
		Host: r.Host,
		Port: r.Port,
	}
}

func (r *ReadUDPPackets) renderF1GamePage(c *fiber.Ctx) error {
	return partials.RenderF1GamePage(c, r.getSharedProps())
}

func (r *ReadUDPPackets) renderF1GamePartial(c *fiber.Ctx) error {
	return partials.RenderF1GamePartial(c, r.getSharedProps())
}

type HostPortUpdateRequest struct {
	Host string `form:"host"`
	Port int    `form:"port"`
}

func (r *ReadUDPPackets) updateHostPort(c *fiber.Ctx) error {
	d := new(HostPortUpdateRequest)
	if err := c.BodyParser(d); err != nil {
		log.Printf("%+v", err)
		return c.Redirect(r.Page.Slug)
	}
	r.setState(func() error {
		r.Host = d.Host
		r.Port = d.Port
		return nil
	})
	return partials.RenderF1GamePage(c, r.getSharedProps())
}

func (r *ReadUDPPackets) Mount(app *fiber.App) {
	r.Page.Mount(app)
	app.Post(fmt.Sprintf("/%s", r.Page.Slug), r.updateHostPort)
}

func (r *ReadUDPPackets) setState(u func() error) error {
	defer r.conn.Close()
	return u()
}

func (r *ReadUDPPackets) Run() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("%+v", err)
		}
	}()

	for {
		select {
		case <-r.Context.Done():
			return
		default:
			func() {
				conn, err := net.ListenUDP("udp", &net.UDPAddr{
					IP:   net.ParseIP(r.Host),
					Port: r.Port,
				})
				if err != nil {
					log.Printf("%+v", err)
					return
				}
				defer conn.Close()

				log.Printf("Listening for incoming packets on %s:%d", r.Host, r.Port)

				r.conn = conn
				r.reader = pipe.ReadUDPPackets(r.Context, conn, r.options)
				if r.applyThenWith != nil {
					r.reader.Then(r.applyThenWith)
				}

				r.reader.Process()
			}()

		}
	}
}

func (r *ReadUDPPackets) Then(s step.Step) step.Step {
	r.applyThenWith = s
	return s
}
