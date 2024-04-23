package panels

import (
	"fmt"

	"github.com/perses/perses/go-sdk/common"
	"github.com/perses/perses/go-sdk/panel"
	"github.com/perses/perses/go-sdk/prometheus/query"
)

func TimeseriesMemoryBasicOptions(filter string) []panel.Option {
	return []panel.Option{
		panel.Description("Basic memory usage"),
		CommonTimeSeriesPlugin(common.BytesUnit),
		panel.AddQuery(
			query.PromQL(fmt.Sprintf(`
				node_memory_MemTotal_bytes{%s}
            `, filter),
				query.SeriesNameFormat("RAM Total"),
				query.Datasource("argos-world")),
		),
		panel.AddQuery(
			query.PromQL(fmt.Sprintf(`
				node_memory_MemTotal_bytes{%s}
				- node_memory_MemFree_bytes{%s}
				- (
					node_memory_Cached_bytes{%s}
					+ node_memory_Buffers_bytes{%s}
					+ node_memory_SReclaimable_bytes{%s}
				)
            `, filter, filter, filter, filter, filter),
				query.SeriesNameFormat("RAM Used"),
				query.Datasource("argos-world")),
		),
		panel.AddQuery(
			query.PromQL(fmt.Sprintf(`
				node_memory_Cached_bytes{%s}
				+ node_memory_Buffers_bytes{%s}
				+ node_memory_SReclaimable_bytes{%s}
            `, filter, filter, filter),
				query.SeriesNameFormat("RAM Cache + Buffer"),
				query.Datasource("argos-world")),
		),
		panel.AddQuery(
			query.PromQL(fmt.Sprintf(`
				node_memory_MemFree_bytes{%s}
            `, filter),
				query.SeriesNameFormat("RAM Free"),
				query.Datasource("argos-world")),
		),
		panel.AddQuery(
			query.PromQL(fmt.Sprintf(`
				node_memory_SwapTotal_bytes{%s} - node_memory_SwapFree_bytes{%s}
            `, filter, filter),
				query.SeriesNameFormat("SWAP used"),
				query.Datasource("argos-world")),
		),
	}
}
