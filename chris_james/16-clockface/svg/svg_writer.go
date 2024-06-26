package svg

import (
	"fmt"
	"io"
	"time"

	cf "github.com/DarrelASandbox/playground-go/chris_james/16-clockface"
)

const secondHandLength = 90
const minuteHandLength = 80
const hourHandLength = 50
const clockCentreX = 150
const clockCentreY = 150

// SVGWriter writes an SVG representation of an
func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	hourHand(w, t)
	io.WriteString(w, svgEnd)
}

/*
1. Scale it to the length of the hand
2. Flip it over the X axis to account for the SVG having an origin in the top left hand corner
3. Translate it to the right position (so that it's coming from an origin of (150,150))

- The SecondHand function is heavily dependent on being an SVG without explicitly referencing SVGs or generating an SVG.
- There is no testing of the actual SVG code within the SecondHand function.
*/
func secondHand(w io.Writer, t time.Time) {
	p := makeHand(cf.SecondHandPoint(t), secondHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func minuteHand(w io.Writer, t time.Time) {
	p := makeHand(cf.MinuteHandPoint(t), minuteHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func hourHand(w io.Writer, t time.Time) {
	p := makeHand(cf.HourHandPoint(t), hourHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func makeHand(p cf.Point, length float64) cf.Point {
	p = cf.Point{X: p.X * length, Y: p.Y * length}                // scale
	p = cf.Point{X: p.X, Y: -p.Y}                                 // flip
	return cf.Point{X: p.X + clockCentreX, Y: p.Y + clockCentreY} // translate
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
