package plot


import (
	"github.com/golang/freetype/truetype"
	"io/ioutil"
	"log"
	"math/rand"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {

	ttf, err := ioutil.ReadFile("ttf/msyh.ttf")
	if err != nil {
		log.Fatal(err)
	}
	ft, err := truetype.Parse(ttf)
	if err != nil {
		log.Fatal(err)
	}

	vg.AddFont("Mincho", ft)

	rand.Seed(int64(0))

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "IM-时段咨询量(万)"
	p.X.Label.Text = "时间"
	p.Y.Label.Text = "调用量"

	font, err := vg.MakeFont("Mincho", 12)
	if err != nil {
		log.Fatal(err)
	}
	p.Title.Font = font
	p.X.Label.Font = font
	p.Y.Label.Font = font

	err = plotutil.AddLinePoints(p,
		"First", randomPoints(15),
		"Second", randomPoints(15),
		"Third", randomPoints(15))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

// randomPoints returns some random x, y points.
func randomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i-1].X + rand.Float64()
		}
		pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}

