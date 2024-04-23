package panels

import (
	"fmt"

	"github.com/perses/perses/go-sdk/common"
	"github.com/perses/perses/go-sdk/panel"
	"github.com/perses/perses/go-sdk/prometheus/query"
)

func StatUptimeOptions(filter string) []panel.Option {
	return []panel.Option{
		panel.Description("System uptime"),
		CommonStatPlugin(string(common.SecondsUnit)),
		panel.AddQuery(
			query.PromQL(fmt.Sprintf(`
				node_time_seconds{%s}
                -
                node_boot_time_seconds{%s}
            `, filter, filter),
				query.Datasource("argos-world")),
		),
	}
}
