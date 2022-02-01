package app

import (
	"fmt"
	"os/exec"
	"strconv"
)

type Color struct {
	Hex     string
	Opacity float64
}

type Shape struct {
	StrokeThickness float64
	Points          int
	Fill            Color
	Stroke          Color
	Path            string
	MinRad          float64
	SizeCon         SizeConstraint
}

type SizeConstraint struct {
	Width      float64
	Height     float64
	PixelScale int
}

func (s Shape) Generate() string {

	cmd := exec.Command(
		"python3",
		"randomshapegen.py",
		"path:"+s.Path,
		"stroke_thickness:"+strconv.FormatFloat(s.StrokeThickness, 'f', 2, 32),
		"points:"+strconv.Itoa(s.Points),
		"fill:"+s.Fill.Hex,
		"fill_opacity:"+strconv.FormatFloat(s.Fill.Opacity, 'f', 2, 32),
		"stroke:"+s.Stroke.Hex,
		"stroke_opacity:"+strconv.FormatFloat(s.Stroke.Opacity, 'f', 2, 32),
		"min_rad:"+strconv.FormatFloat(s.MinRad, 'f', 2, 32),
		"width:"+strconv.FormatFloat(s.SizeCon.Width, 'f', 2, 32),
		"height:"+strconv.FormatFloat(s.SizeCon.Height, 'f', 2, 32),
		"pixel_scale:"+strconv.Itoa(s.SizeCon.PixelScale),
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	return string(out)
}
