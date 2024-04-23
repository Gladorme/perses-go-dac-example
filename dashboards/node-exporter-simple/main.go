package main

import (
	"flag"

	"dac/panels"
	"github.com/perses/perses/go-sdk"
	"github.com/perses/perses/go-sdk/dashboard"
	panelgroup "github.com/perses/perses/go-sdk/panel-group"
	"github.com/perses/perses/go-sdk/prometheus/variable/promql"
	variablegroup "github.com/perses/perses/go-sdk/variable-group"
	listvariable "github.com/perses/perses/go-sdk/variable/list-variable"
	textvariable "github.com/perses/perses/go-sdk/variable/text-variable"
)

func main() {
	flag.Parse()
	exec := sdk.NewExec()

	// DATASOURCE => already registered on the instance at global level

	variableGroup := dashboard.AddVariableGroup(
		variablegroup.AddVariable("stack",
			textvariable.Text("erd4",
				textvariable.Constant(true),
				textvariable.Hidden(true),
			),
		),
		variablegroup.AddVariable("prometheus",
			textvariable.Text("system",
				textvariable.Constant(true),
				textvariable.Hidden(true),
			),
		),
		variablegroup.AddVariable("job",
			textvariable.Text("cmdbrtu-custom-sd",
				textvariable.Constant(true),
			),
		),
		variablegroup.AddVariable("instance",
			listvariable.List(
				promql.PrometheusPromQL("group by (instance) (node_uname_info{stack=\"$stack\",prometheus=\"$prometheus\",job=\"$job\"})",
					promql.LabelName("instance"),
					promql.Datasource("argos-world"),
				),
				listvariable.DisplayName("Host"),
			),
		),
	)

	filter := "stack=\"$stack\",prometheus=\"$prometheus\",job=\"$job\",instance=\"$instance\""

	// DASHBOARD
	builder, buildErr := dashboard.New("node-exporter-simple",
		dashboard.Name("Node Exporter Simple"),

		dashboard.ProjectName("gladorme"),

		// VARIABLES
		variableGroup,

		// PANELS
		dashboard.AddPanelGroup("Quick CPU / Mem / Disk",
			panelgroup.PanelsPerLine(7),
			panelgroup.AddPanel("Sys Load (5m avg)", panels.GaugeSysLoadOptions(filter)...),
			panelgroup.AddPanel("RAM Used", panels.GaugeRAMUsedOptions(filter)...),
			panelgroup.AddPanel("RootFS Total", panels.GaugeRootFSOptions(filter)...),
			panelgroup.AddPanel("CPU Cores", panels.StatCPUCoreOptions(filter)...),
			panelgroup.AddPanel("RAM Total", panels.StatRAMTotalOptions(filter)...),
			panelgroup.AddPanel("RootFS Total", panels.StatRootFSOptions(filter)...),
			panelgroup.AddPanel("System uptime", panels.StatUptimeOptions(filter)...),
		),
		dashboard.AddPanelGroup("Basic CPU / Mem / Net / Disk",
			panelgroup.PanelsPerLine(2),
			panelgroup.PanelHeight(7),
			panelgroup.AddPanel("CPU Basic", panels.TimeseriesCPUBasicOptions(filter)...),
			panelgroup.AddPanel("Memory Basic", panels.TimeseriesMemoryBasicOptions(filter)...),
			panelgroup.AddPanel("Network Traffic Basic", panels.TimeseriesNetworkTrafficBasicOptions(filter)...),
			panelgroup.AddPanel("Disk Space Used Basic", panels.TimeseriesDiskSpaceUsedBasicOptions(filter)...),
		),
	)

	exec.BuildDashboard(builder, buildErr)
}
