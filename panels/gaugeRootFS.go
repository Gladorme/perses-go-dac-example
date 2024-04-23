package panels

import (
	"fmt"

	"github.com/perses/perses/go-sdk/panel"
	"github.com/perses/perses/go-sdk/prometheus/query"
)

func GaugeRootFSOptions(filter string) []panel.Option {
	return []panel.Option{
		panel.Description("Used Root FS"),
		CommonGaugeChart(),
		panel.AddQuery(
			query.PromQL(fmt.Sprintf(`
				100 - (
                        (
                            avg_over_time(node_filesystem_avail_bytes{%s,mountpoint="/",fstype!="rootfs"}[$__rate_interval])
                            *
                            100
                        )
                        /
                        avg_over_time(node_filesystem_size_bytes{%s,mountpoint="/",fstype!="rootfs"}[$__rate_interval])
                    )
            `, filter, filter),
				query.Datasource("argos-world")),
		),
	}
}
