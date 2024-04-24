package panels

import (
	"fmt"

	"github.com/perses/perses/go-sdk/common"
	"github.com/perses/perses/go-sdk/panel"
	"github.com/perses/perses/go-sdk/prometheus/query"
)

func TimeseriesCPUBasicOptions(filter string) []panel.Option {
	return []panel.Option{
		panel.Description("Basic CPU info"),
		CommonTimeSeriesPlugin(string(common.PercentDecimalUnit)),
		panel.AddQuery(
			query.PromQL(fmt.Sprintf(`
				sum by (instance) (
					irate(node_cpu_seconds_total{%s, mode="system"}[$__rate_interval])
				)
				/ 
				on(instance) group_left sum by (instance) (
					irate(node_cpu_seconds_total{%s}[$__rate_interval])
				)
            `, filter, filter),
				query.SeriesNameFormat("System"),
				query.Datasource("argos-world")),
		),
		panel.AddQuery(
			query.PromQL(fmt.Sprintf(`
				sum by(instance) (
					irate(node_cpu_seconds_total{%s, mode="user"}[$__rate_interval])
				)
				/ on(instance) group_left sum by (instance) (
					irate(node_cpu_seconds_total{%s}[$__rate_interval])
				)
            `, filter, filter),
				query.SeriesNameFormat("User"),
				query.Datasource("argos-world")),
		),
		panel.AddQuery(
			query.PromQL(fmt.Sprintf(`
				sum by(instance) (
					irate(node_cpu_seconds_total{%s, mode="iowait"}[$__rate_interval])
				)
				/ on(instance) group_left sum by (instance) (
					irate(node_cpu_seconds_total{%s}[$__rate_interval])
				)
            `, filter, filter),
				query.SeriesNameFormat("IOWait"),
				query.Datasource("argos-world")),
		),
		panel.AddQuery(
			query.PromQL(fmt.Sprintf(`
				sum by(instance) (
					irate(node_cpu_seconds_total{%s, mode=~".*irq"}[$__rate_interval])
				) 
				/ on(instance) group_left sum by (instance) (
					irate(node_cpu_seconds_total{%s}[$__rate_interval])
				)
            `, filter, filter),
				query.SeriesNameFormat("IRQ & SOFTIRQ"),
				query.Datasource("argos-world")),
		),
		panel.AddQuery(
			query.PromQL(fmt.Sprintf(`
				sum by (instance) (
					irate(node_cpu_seconds_total{%s, mode!='idle',mode!='user',mode!='system',mode!='iowait',mode!='irq',mode!='softirq'}[$__rate_interval])
				)
				/ on(instance) group_left sum by (instance) (
					irate(node_cpu_seconds_total{%s}[$__rate_interval])
				)
            `, filter, filter),
				query.SeriesNameFormat("Other"),
				query.Datasource("argos-world")),
		),
		panel.AddQuery(
			query.PromQL(fmt.Sprintf(`
				sum by(instance) (
					irate(node_cpu_seconds_total{%s, mode="idle"}[$__rate_interval])
				)
				/ on(instance) group_left sum by (instance) (
					irate(node_cpu_seconds_total{%s}[$__rate_interval])
				)
            `, filter, filter),
				query.SeriesNameFormat("Idle"),
				query.Datasource("argos-world")),
		),
	}
}
