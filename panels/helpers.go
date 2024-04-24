package panels

import (
	"github.com/perses/perses/go-sdk/common"
	"github.com/perses/perses/go-sdk/panel"
	"github.com/perses/perses/go-sdk/panel/gauge"
	"github.com/perses/perses/go-sdk/panel/stat"
	timeseries "github.com/perses/perses/go-sdk/panel/time-series"
)

func CommonGaugeChart() panel.Option {
	return gauge.Chart(
		gauge.Calculation(common.LastNumberCalculation),
		gauge.Format(common.Format{
			Unit: string(common.PercentUnit),
		}),
		gauge.Max(100),
		gauge.Thresholds(common.Thresholds{
			Steps: []common.StepOption{
				{
					Value: 0,
					Color: "rgba(50, 172, 45, 0.97)",
				},
				{
					Value: 80,
					Color: "rgba(237, 129, 40, 0.89)",
				},
				{
					Value: 90,
					Color: "rgba(245, 54, 54, 0.9)",
				},
			},
		}),
	)
}

func CommonStatPlugin(unit string) panel.Option {
	return stat.Chart(
		stat.Calculation(common.LastNumberCalculation),
		stat.Thresholds(common.Thresholds{
			DefaultColor: "#7b7b7b",
		}),
		stat.Format(common.Format{Unit: unit}),
	)
}

func CommonTimeSeriesPlugin(unit string) panel.Option {
	return timeseries.Chart(
		timeseries.WithYAxis(timeseries.YAxis{
			Format: &common.Format{
				Unit: unit,
			},
		}),
		timeseries.WithLegend(
			timeseries.Legend{
				Position: timeseries.RightPosition,
				Mode:     timeseries.TableMode,
				Values: []common.Calculation{
					common.MinCalculation,
					common.MaxCalculation,
				},
			}),
		//timeseries.WithVisual(timeseries.Visual{
		//	AreaOpacity: 0.05,
		//}),
	)
}
