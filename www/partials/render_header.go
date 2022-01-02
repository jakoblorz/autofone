package partials

import (
	"encoding/json"
	"expvar"
	"sort"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/pipe"
	"github.com/jakoblorz/metrikxd/www/layouts"
	"github.com/zserge/metric"
)

type RenderHeaderSharedProps struct{}

func (r *RenderHeaderSharedProps) AppendAdditionalProps(m fiber.Map) fiber.Map {
	return m
}

type renderHeaderPartialProps struct {
	layouts.PartialLayoutProps
	RenderHeaderSharedProps
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

func appendMetrics(p fiber.Map) {
	p["GameSetupMetrics"] = getMetricsFor(func() map[string]metric.Metric {
		return map[string]metric.Metric{
			pipe.PacketReaderExpVarRX: expvar.Get(pipe.PacketReaderExpVarRX).(metric.Metric),
			pipe.PacketReaderExpVarTX: expvar.Get(pipe.PacketReaderExpVarTX).(metric.Metric),
		}
	})[0]
	p["SendingMetrics"] = getMetricsFor(func() map[string]metric.Metric {
		return map[string]metric.Metric{
			pipe.PacketWriterExpVarRX: expvar.Get(pipe.PacketWriterExpVarRX).(metric.Metric),
			pipe.PacketWriterExpVarTX: expvar.Get(pipe.PacketWriterExpVarTX).(metric.Metric),
		}
	})[0]
}

func renderHeader(slug string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		p := fiber.Map{
			"Slug": slug,
		}
		appendMetrics(p)

		return c.Render("partials/header", p)
	}
}
