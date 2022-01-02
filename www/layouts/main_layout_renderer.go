package layouts

import (
	"encoding/json"
	"expvar"
	"sort"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/pipe"
	"github.com/jakoblorz/metrikxd/www/rendering"
	"github.com/zserge/metric"
)

type RenderMainLayoutProps interface {
	GetTitle() string
	GetSlug() string
}

type MainLayoutProps struct {
	Title string
	Slug  string
}

func (p *MainLayoutProps) GetTitle() string {
	return p.Title
}

func (p *MainLayoutProps) GetSlug() string {
	return p.Slug
}

func getMetricsFor(snapshot func() map[string]metric.Metric) []map[string]interface{} {
	metrics := make([]map[string]interface{}, 0)
	for name, metric := range snapshot() {
		m := map[string]interface{}{}
		b, _ := json.Marshal(metric)
		json.Unmarshal(b, &m)
		m["name"] = name
		metrics = append(metrics, m)
	}
	sort.Slice(metrics, func(i, j int) bool {
		n1 := metrics[i]["name"].(string)
		n2 := metrics[j]["name"].(string)
		return strings.Compare(n1, n2) < 0
	})
	return metrics
}

func RenderMainLayout(template string, c *fiber.Ctx, props RenderMainLayoutProps) error {
	p := fiber.Map{
		"Title": props.GetTitle(),
		"Slug":  props.GetSlug(),
	}
	if a, ok := props.(rendering.AdditionalPropsAppender); ok {
		p = a.AppendAdditionalProps(p)
	}
	p["GameSetupMetrics"] = getMetricsFor(func() map[string]metric.Metric {
		return map[string]metric.Metric{
			pipe.PacketReaderExpVarRX: expvar.Get("random:gauge").(metric.Metric),
			pipe.PacketReaderExpVarTX: expvar.Get(pipe.PacketReaderExpVarTX).(metric.Metric),
		}
	})[0]
	p["SendingMetrics"] = getMetricsFor(func() map[string]metric.Metric {
		return map[string]metric.Metric{
			pipe.PacketWriterExpVarRX: expvar.Get("random:gauge").(metric.Metric),
			pipe.PacketWriterExpVarTX: expvar.Get(pipe.PacketWriterExpVarTX).(metric.Metric),
		}
	})[0]
	return c.Render(template, p, "layouts/main")
}
