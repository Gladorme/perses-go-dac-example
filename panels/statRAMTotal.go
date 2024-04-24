package panels

import (
	"fmt"

	"github.com/perses/perses/go-sdk/common"
	"github.com/perses/perses/go-sdk/panel"
	"github.com/perses/perses/go-sdk/prometheus/query"
)

func StatRAMTotalOptions(filter string) []panel.Option {
	return []panel.Option{
		panel.Description("Total RAM"),
		CommonStatPlugin(common.BytesUnit),
		panel.AddQuery(
			query.PromQL(fmt.Sprintf(`
				node_memory_MemTotal_bytes{%s}
            `, filter),
				query.Datasource("argos-world")),
		),
	}
}
