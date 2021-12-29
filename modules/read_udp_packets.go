package modules

import (
	"context"
	"fmt"
	"net"
	"reflect"
	"sort"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/constants"
	"github.com/jakoblorz/metrikxd/pipe"
	"github.com/jakoblorz/metrikxd/pkg/log"
	"github.com/jakoblorz/metrikxd/pkg/step"
	"github.com/jakoblorz/metrikxd/www"
	"github.com/jakoblorz/metrikxd/www/partials"
)

func uint8ToSPF(p uint8, name string) partials.SinglePacketFilter {
	return partials.SinglePacketFilter{fmt.Sprintf("packet-%d", p), fmt.Sprintf("packet_%s", strings.ReplaceAll(strings.ToLower(name), " ", "_")), false, fmt.Sprintf("Packet %d (%s)", p, name)}
}

var packetFilters = []partials.SinglePacketFilter{
	uint8ToSPF(constants.PacketMotion, "Motion"),
	uint8ToSPF(constants.PacketSession, "Session"),
	uint8ToSPF(constants.PacketLap, "Lap"),
	uint8ToSPF(constants.PacketEvent, "Event"),
	uint8ToSPF(constants.PacketParticipants, "Participants"),
	uint8ToSPF(constants.PacketCarSetup, "Car Setup"),
	uint8ToSPF(constants.PacketCarTelemetry, "Car Telemetry"),
	uint8ToSPF(constants.PacketCarStatus, "Car Status"),
	uint8ToSPF(constants.PacketFinalClassification, "Final Classification"),
	uint8ToSPF(constants.PacketLobbyInfo, "Lobby Information"),
	uint8ToSPF(constants.PacketCarDamage, "Car Damage"),
	uint8ToSPF(constants.PacketSessionHistory, "Session History"),
}

func getOnOffState(options *pipe.PacketReaderOptions) []partials.SinglePacketFilter {
	filters := packetFilters
	for _, id := range options.Filter {
		filters[id].Value = true
	}
	return filters
}

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
	sort.Sort(initialOptions)
	p.Page = www.Page{"game-setup", p.renderF1GamePage, p.renderF1GamePartial, www.EmptySSEHandler}
	return p
}

func (r *ReadUDPPackets) getSharedProps() partials.RenderGameSetupSharedProps {

	return partials.RenderGameSetupSharedProps{
		Host:    r.Host,
		Port:    r.Port,
		Packets: getOnOffState(r.options),
	}
}

func (r *ReadUDPPackets) renderF1GamePage(c *fiber.Ctx) error {
	return partials.RenderGameSetupPage(c, r.getSharedProps())
}

func (r *ReadUDPPackets) renderF1GamePartial(c *fiber.Ctx) error {
	return partials.RenderGameSetupPartial(c, r.getSharedProps())
}

type UpdateUDPReaderRequest struct {
	Host string `form:"host"`
	Port int    `form:"port"`

	PacketMotion              string `form:"packet_motion"`
	PacketSession             string `form:"packet_session"`
	PacketLap                 string `form:"packet_lap"`
	PacketEvent               string `form:"packet_event"`
	PacketParticipants        string `form:"packet_participants"`
	PacketCarSetup            string `form:"packet_car_setup"`
	PacketCarTelemetry        string `form:"packet_car_telemetry"`
	PacketCarStatus           string `form:"packet_car_status"`
	PacketFinalClassification string `form:"packet_final_classification"`
	PacketLobbyInfo           string `form:"packet_lobby_info"`
	PacketCarDamage           string `form:"packet_car_damage"`
	PacketSessionHistory      string `form:"packet_session_history"`
}

func (u *UpdateUDPReaderRequest) ToFilter() (values []uint8) {
	values = make([]uint8, 0)
	if len(u.PacketMotion) > 0 {
		values = append(values, constants.PacketMotion)
	}
	if len(u.PacketSession) > 0 {
		values = append(values, constants.PacketSession)
	}
	if len(u.PacketLap) > 0 {
		values = append(values, constants.PacketLap)
	}
	if len(u.PacketEvent) > 0 {
		values = append(values, constants.PacketEvent)
	}
	if len(u.PacketParticipants) > 0 {
		values = append(values, constants.PacketParticipants)
	}
	if len(u.PacketCarSetup) > 0 {
		values = append(values, constants.PacketCarSetup)
	}
	if len(u.PacketCarTelemetry) > 0 {
		values = append(values, constants.PacketCarTelemetry)
	}
	if len(u.PacketCarStatus) > 0 {
		values = append(values, constants.PacketCarStatus)
	}
	if len(u.PacketFinalClassification) > 0 {
		values = append(values, constants.PacketFinalClassification)
	}
	if len(u.PacketLobbyInfo) > 0 {
		values = append(values, constants.PacketLobbyInfo)
	}
	if len(u.PacketCarDamage) > 0 {
		values = append(values, constants.PacketCarDamage)
	}
	if len(u.PacketSessionHistory) > 0 {
		values = append(values, constants.PacketSessionHistory)
	}
	return
}

func (r *ReadUDPPackets) updateUDPReader(c *fiber.Ctx) error {
	d := new(UpdateUDPReaderRequest)
	if err := c.BodyParser(d); err != nil {
		log.Printf("%+v", err)
		return c.Redirect(r.Page.Slug)
	}

	filter := d.ToFilter()
	updateHostPort := r.Host != d.Host || r.Port != d.Port
	if updateHostPort || !reflect.DeepEqual(r.options.Filter, filter) {
		r.setState(func() error {
			if updateHostPort {
				r.Host = d.Host
				r.Port = d.Port
			}
			r.options.Filter = filter
			return nil
		})
	}
	return partials.RenderGameSetupPage(c, r.getSharedProps())
}

func (r *ReadUDPPackets) Mount(app *fiber.App) {
	r.Page.Mount(app)
	app.Post(fmt.Sprintf("/%s", r.Page.Slug), r.updateUDPReader)
}

func (r *ReadUDPPackets) setState(u func() error) error {
	defer r.conn.Close()
	return u()
}

func (r *ReadUDPPackets) Run() {
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
