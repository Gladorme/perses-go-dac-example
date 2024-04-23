package panels

import (
	"fmt"

	"github.com/perses/perses/go-sdk/panel"
	"github.com/perses/perses/go-sdk/prometheus/query"
)

func GaugeSysLoadOptions(filter string) []panel.Option {
	return []panel.Option{
		panel.Description("Busy state of all CPU cores together (5 min average)"),
		CommonGaugeChart(),
		panel.AddQuery(
			query.PromQL(fmt.Sprintf(`
				avg_over_time(node_load5{%s}[$__rate_interval]) * 100
                /
                on(instance) group_left sum by (instance) (
                    irate(node_cpu_seconds_total{%s}[$__rate_interval])
                )
            `, filter, filter),
				query.Datasource("argos-world")),
		),
	}
}
