package panels

import (
	"fmt"

	"github.com/perses/perses/go-sdk/common"
	"github.com/perses/perses/go-sdk/panel"
	"github.com/perses/perses/go-sdk/prometheus/query"
)

func StatRootFSOptions(filter string) []panel.Option {
	return []panel.Option{
		panel.Description("Total RootFS"),
		CommonStatPlugin(common.BytesUnit),
		panel.AddQuery(
			query.PromQL(fmt.Sprintf(`
				node_filesystem_size_bytes{%s,mountpoint="/",fstype!="rootfs"}
            `, filter),
				query.Datasource("argos-world")),
		),
	}
}
