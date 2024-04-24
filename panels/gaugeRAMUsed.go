package panels

import (
	"fmt"

	"github.com/perses/perses/go-sdk/panel"
	"github.com/perses/perses/go-sdk/prometheus/query"
)

func GaugeRAMUsedOptions(filter string) []panel.Option {
	return []panel.Option{
		panel.Description("Non available RAM"),
		CommonGaugeChart(),
		panel.AddQuery(
			query.PromQL(fmt.Sprintf(`
				(
					(
						avg_over_time(node_memory_MemTotal_bytes{%s}[$__rate_interval])
						-
						avg_over_time(node_memory_MemFree_bytes{%s}[$__rate_interval])
					)
					/
					(
						avg_over_time(node_memory_MemTotal_bytes{%s}[$__rate_interval])
					)
				) * 100
            `, filter, filter, filter),
				query.Datasource("argos-world")),
			query.PromQL(fmt.Sprintf(`
				100 - (
					(
						avg_over_time(node_memory_MemAvailable_bytes{%s}[$__rate_interval])
						*
						100
					) 
					/
					avg_over_time(node_memory_MemTotal_bytes{%s}[$__rate_interval])
				)
            `, filter, filter),
				query.Datasource("argos-world")),
		),
	}
}
