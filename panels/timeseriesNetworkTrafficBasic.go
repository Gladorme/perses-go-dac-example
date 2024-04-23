package panels

import (
	"fmt"

	"github.com/perses/perses/go-sdk/common"
	"github.com/perses/perses/go-sdk/panel"
	"github.com/perses/perses/go-sdk/prometheus/query"
)

func TimeseriesNetworkTrafficBasicOptions(filter string) []panel.Option {
	return []panel.Option{
		panel.Description("Basic network info per interface"),
		CommonTimeSeriesPlugin(string(common.PercentUnit)),
		panel.AddQuery(
			query.PromQL(fmt.Sprintf(`
                irate(node_network_receive_bytes_total{%s}[$__rate_interval]) * 8
            `, filter),
				query.SeriesNameFormat("recv {{device}}"),
				query.Datasource("argos-world")),
		),
		panel.AddQuery(
			query.PromQL(fmt.Sprintf(`
                irate(node_network_transmit_bytes_total{%s}[$__rate_interval]) * 8
            `, filter),
				query.SeriesNameFormat("trans {{device}}"),
				query.Datasource("argos-world")),
		),
	}
}
