package panels

import (
	"fmt"

	"github.com/perses/perses/go-sdk/common"
	"github.com/perses/perses/go-sdk/panel"
	"github.com/perses/perses/go-sdk/prometheus/query"
)

func StatCPUCoreOptions(filter string) []panel.Option {
	return []panel.Option{
		panel.Description("Total number of CPU cores"),
		CommonStatPlugin(common.DecimalUnit),
		panel.AddQuery(
			query.PromQL(fmt.Sprintf(`
				count(count(node_cpu_seconds_total{%s}) by (cpu))
            `, filter),
				query.Datasource("argos-world")),
		),
	}
}
