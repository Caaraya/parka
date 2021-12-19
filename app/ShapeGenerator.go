package app

import (
	"fmt"
	"os/exec"
	"strconv"
)

type Color struct {
	hex     string
	opacity float64
}

type Shape struct {
	stroke_thickness float64
	points           int
	fill             Color
	stroke           Color
	path             string
	min_rad          float64
	size_con         SizeConstraint
}

type SizeConstraint struct {
	width       float64
	height      float64
	pixel_scale int
}

func (s Shape) generate() string {

	cmd := exec.Command(
		"python3",
		"randomshapegen.py",
		"path:"+s.path,
		"stroke_thickness:"+strconv.FormatFloat(s.stroke_thickness, 'f', 2, 32),
		"points:"+strconv.Itoa(s.points),
		"fill:"+s.fill.hex,
		"fill_opacity:"+strconv.FormatFloat(s.fill.opacity, 'f', 2, 32),
		"stroke:"+s.stroke.hex,
		"stroke_opacity:"+strconv.FormatFloat(s.stroke.opacity, 'f', 2, 32),
		"min_rad:"+strconv.FormatFloat(s.min_rad, 'f', 2, 32),
		"width:"+strconv.FormatFloat(s.size_con.width, 'f', 2, 32),
		"height:"+strconv.FormatFloat(s.size_con.height, 'f', 2, 32),
		"pixel_scale:"+strconv.Itoa(s.size_con.pixel_scale),
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	return string(out)
}
