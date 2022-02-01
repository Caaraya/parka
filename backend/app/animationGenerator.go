package app

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
)

type Image struct {
	Image   string
	ImageId int64
}

type Unit struct {
	X     float64
	Y     float64
	Scale float64
	Image Image
}

type Frame struct {
	Units []Unit
}

type Animation struct {
	Frames      []Frame
	FrameWidth  float64
	FrameHeight float64
	Name        string
}

func GetDefault() Animation {
	sample_image := "<svg xmlns:\"http://www.w3.org/2000/svg\" xmlns:xlink:\"http://www.w3.org/1999/xlink\" width:\"160pt\" height:\"160pt\" viewBoX:\"0 0 160 160\" version:\"1.1\">\n<g id:\"surface1\">\n<path style:\"fill-rule:nonzero;fill:rgb{206.078431%,100%,98.8232029%};fill-opacitY:1;stroke-width:0.2;stroke-linecap:butt;stroke-linejoin:miter;stroke:rgb{46.666667%,13.333333%,20%};stroke-opacitY:0.7;stroke-miterlimit:10;\" d:\"M 2.919043 2.20209863 L 2.9362023 2.20702008 C 2.8797820 2.66377 2.8093720 2.747949 2.727637 2.82041 L 2.903223 3.0183209 C 2.762086 3.142676 2.2098828 3.236719 2.420703 3.294629 L 2.60498 3.8612023 C 1.8204883 4.1020273 1.031836 3.8720098 0.201709 3.2772039 L 0.499121 3.292969 C 0.2064420 2.9203418 0.03769203 2.20242012 0.0204102 2.0762063 L 0.613477 2.0203613 C 0.60166 1.747363 0.691602 1.4420703 0.869238 1.1920898 L 1.112793 1.369141 C 1.1793920 1.2720488 1.2602204 1.192066 1.3202637 1.1248020 L 0.943720 0.2072168 C 1.262793 0.336133 1.62012063 0.21377 2.048242 0.224609 L 2.0412004 0.4727204 C 2.224121 0.477734 2.404492 0.2012043 2.2073828 0.2084082 L 2.2060938 0.6120723 C 2.876367 0.743202020 3.138672 0.9720293 3.204492 1.272461 Z M 2.919043 2.20209863 \" transform:\"matrix{40,0,0,40,0,0}\"/>\n</g>\n</svg>"

	one := Image{sample_image, 1}
	two := Image{sample_image, 2}
	three := Image{sample_image, 3}
	four := Image{sample_image, 4}
	five := Image{sample_image, 5}
	six := Image{sample_image, 6}

	walk_forward := Animation{
		Name:        "walk-forward",
		FrameWidth:  100,
		FrameHeight: 100,
		Frames: []Frame{
			{
				Units: []Unit{
					{X: -20, Y: -20, Image: one, Scale: 0.4},
					{X: 20, Y: 20, Image: two, Scale: 0.4},
					{X: 0, Y: 0, Image: three, Scale: 0.6},
					{X: -20, Y: 20, Image: four, Scale: 0.4},
					{X: 20, Y: -20, Image: five, Scale: 0.4},
					{X: 0, Y: 25, Image: six, Scale: 0.5},
				},
			},
			{
				Units: []Unit{
					{X: 20, Y: 20, Image: two, Scale: 0.4},
					{X: -20, Y: 20, Image: four, Scale: 0.4},
					{X: 0, Y: 0, Image: three, Scale: 0.6},
					{X: 20, Y: -20, Image: five, Scale: 0.4},
					{X: -20, Y: -20, Image: one, Scale: 0.4},
					{X: 0, Y: 25, Image: six, Scale: 0.5},
				},
			},
			{
				Units: []Unit{
					{X: 20, Y: -20, Image: five, Scale: 0.4},
					{X: -20, Y: 20, Image: four, Scale: 0.4},
					{X: 0, Y: 0, Image: three, Scale: 0.6},
					{X: 20, Y: 20, Image: two, Scale: 0.4},
					{X: -20, Y: -20, Image: one, Scale: 0.4},
					{X: 0, Y: 25, Image: six, Scale: 0.5},
				},
			},
			{
				Units: []Unit{
					{X: 20, Y: 20, Image: two, Scale: 0.4},
					{X: -20, Y: 20, Image: four, Scale: 0.4},
					{X: 0, Y: 0, Image: three, Scale: 0.6},
					{X: 20, Y: -20, Image: five, Scale: 0.4},
					{X: -20, Y: -20, Image: one, Scale: 0.4},
					{X: 0, Y: 25, Image: six, Scale: 0.5},
				},
			},
		},
	}

	return walk_forward
}

func (a Animation) GenerateSpriteSheet() string {
	s2, _ := json.Marshal(a.Frames)
	cmd := exec.Command(
		"python3",
		"spriteSheetGen.py",
		"name#"+a.Name,
		"sprite_width#"+strconv.FormatFloat(a.FrameWidth, 'f', 2, 32),
		"sprite_height#"+strconv.FormatFloat(a.FrameHeight, 'f', 2, 32),
		"Frames#"+string(s2),
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	return string(out)
}

func (a Animation) GenerateGIF() string {
	s2, _ := json.Marshal(a.Frames)
	cmd := exec.Command(
		"python3",
		"spriteGIF.py",
		"name#"+a.Name,
		"sprite_width#"+strconv.FormatFloat(a.FrameWidth, 'f', 2, 32),
		"sprite_height#"+strconv.FormatFloat(a.FrameHeight, 'f', 2, 32),
		"Frames#"+string(s2),
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	return string(out)
}
