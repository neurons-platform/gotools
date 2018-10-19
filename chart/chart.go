package chart

import (
	"bytes"
	"os"
	"io/ioutil"
	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart"
	U "github.com/jingminglang/gotools/utils"
)

func DrawTimeSeries(fileName string,title string,max float64, series []chart.Series) {
	ttf, err := ioutil.ReadFile("ttf/msyh.ttf")
	if err != nil {
		// panic(err)
		U.Throw(err)
	}

	ft, err := truetype.Parse(ttf)
	if err != nil {
		// panic(err)
		U.Throw(err)
	}

	graph := chart.Chart{
		Title: title,
		TitleStyle: chart.StyleShow(),
		Font:  ft,
		XAxis: chart.XAxis{
			Name:           "时间",
			Style: chart.StyleShow(),
			NameStyle:      chart.StyleShow(),
			// ValueFormatter: chart.TimeMinuteValueFormatter, //TimeHourValueFormatter,
			ValueFormatter: chart.TimeValueFormatterWithFormat("01-02 15:04"), //TimeHourValueFormatter,
			// ValueFormatter: "01-02 3:04PM",
		},
		YAxis: chart.YAxis{
			Name:      "值",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
			Range: &chart.ContinuousRange{
				Min: 0.0,
				Max: max,
				// Max: 100.0,
			},

		},
		// YAxisSecondary: chart.YAxis{
                //         Style: chart.StyleShow(),
                // },
		Series: series,
	}
	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}
	// graph.YAxis.Range.SetMin(0.0)
	// graph.YAxis.Range.SetMax(100.0)


	buffer := bytes.NewBuffer([]byte{})
	err = graph.Render(chart.PNG, buffer)
	if err != nil {
		// panic(err)
		U.Throw(err)
	}

	fo, err := os.Create(fileName)
	if err != nil {
		// panic(err)
		U.Throw(err)
	}

	if _, err := fo.Write(buffer.Bytes()); err != nil {
		// panic(err)
		U.Throw(err)
	}
}
