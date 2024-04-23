package panels

import (
	"fmt"

	"github.com/perses/perses/go-sdk/common"
	"github.com/perses/perses/go-sdk/panel"
	"github.com/perses/perses/go-sdk/prometheus/query"
)

func TimeseriesDiskSpaceUsedBasicOptions(filter string) []panel.Option {
	return []panel.Option{
		panel.Description("Disk space used of all filesystems mounted"),
		CommonTimeSeriesPlugin(string(common.PercentUnit)),
		panel.AddQuery(
			query.PromQL(fmt.Sprintf(`
                100 - (
                    (node_filesystem_avail_bytes{%s,device!~'rootfs'} * 100)
                    /
                    node_filesystem_size_bytes{%s,device!~'rootfs'}
                )
            `, filter, filter),
				query.SeriesNameFormat("{{mountpoint}}"),
				query.Datasource("argos-world")),
		),
	}
}
