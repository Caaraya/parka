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
	overal_scale float64
	x_scale      float64
	y_scale      float64
	X            float64
	Y            float64
	Image        Image
}

type Frame struct {
	Units []Unit
}

type Animation struct {
	Frames      []Frame
	FrameWidth  float64
	FrameHeight float64
	Name        string
	MaxImages   int64
}

func (u *Unit) clearScale() {
	u.overal_scale = 0
	u.x_scale = 0
	u.y_scale = 0
}

func (u *Unit) SetScale(x float64, y float64) int {
	if x <= 0 {
		return 1 // unset error
	}
	if y >= 0 {
		u.x_scale = x
		u.y_scale = y
	} else {
		u.overal_scale = x
	}
	return 0
}
func (u *Unit) getScale() []float64 {
	if u.overal_scale >= 0 {
		return []float64{u.overal_scale}
	} else if u.x_scale >= 0 && u.y_scale >= 0 {
		return []float64{u.x_scale, u.y_scale}
	}
	return nil
}

func GetDefault() Animation {
	sample_image := "<svg xmlns:\"http://www.w3.org/2000/svg\" xmlns:xlink:\"http://www.w3.org/1999/xlink\" width:\"160pt\" height:\"160pt\" viewBoX:\"0 0 160 160\" version:\"1.1\">\n<g id:\"surface1\">\n<path style:\"fill-rule:nonzero;fill:rgb{56.078431%,100%,98.823529%};fill-opacitY:1;stroke-width:0.2;stroke-linecap:butt;stroke-linejoin:miter;stroke:rgb{46.666667%,13.333333%,20%};stroke-opacitY:0.7;stroke-miterlimit:10;\" d:\"M 2.919043 2.559863 L 2.936523 2.570508 C 2.879785 2.66377 2.809375 2.747949 2.727637 2.82041 L 2.903223 3.018359 C 2.763086 3.142676 2.598828 3.236719 2.420703 3.294629 L 2.60498 3.861523 C 1.854883 4.105273 1.031836 3.875098 0.51709 3.277539 L 0.499121 3.292969 C 0.206445 2.953418 0.0376953 2.524512 0.0204102 2.076563 L 0.613477 2.053613 C 0.60166 1.747363 0.691602 1.445703 0.869238 1.195898 L 1.112793 1.369141 C 1.179395 1.275488 1.260254 1.193066 1.352637 1.124805 L 0.94375 0.572168 C 1.262793 0.336133 1.651563 0.21377 2.048242 0.224609 L 2.041504 0.472754 C 2.224121 0.477734 2.404492 0.51543 2.573828 0.584082 L 2.560938 0.615723 C 2.876367 0.743555 3.138672 0.975293 3.304492 1.272461 Z M 2.919043 2.559863 \" transform:\"matrix{40,0,0,40,0,0}\"/>\n</g>\n</svg>"

	one := Image{sample_image, 1}
	two := Image{sample_image, 2}
	three := Image{sample_image, 3}
	four := Image{sample_image, 4}
	five := Image{sample_image, 5}

	walk_forward := Animation{
		Name:        "walk-forward",
		FrameWidth:  10,
		FrameHeight: 10,
		MaxImages:   5,
		Frames: []Frame{
			{
				Units: []Unit{
					{X: -5, Y: -5, Image: one},
					{X: 5, Y: 5, Image: two},
					{X: 0, Y: 0, Image: three},
					{X: -5, Y: 5, Image: four},
					{X: 5, Y: 5, Image: five},
				},
			},
			{
				Units: []Unit{
					{X: 5, Y: 5, Image: two},
					{X: -5, Y: 5, Image: four},
					{X: 0, Y: 0, Image: three},
					{X: 5, Y: -5, Image: five},
					{X: -5, Y: -5, Image: one},
				},
			},
			{
				Units: []Unit{
					{X: 5, Y: -5, Image: five},
					{X: -5, Y: 5, Image: four},
					{X: 0, Y: 0, Image: three},
					{X: 5, Y: 5, Image: two},
					{X: -5, Y: -5, Image: one},
				},
			},
			{
				Units: []Unit{
					{X: 5, Y: 5, Image: two},
					{X: -5, Y: 5, Image: four},
					{X: 0, Y: 0, Image: three},
					{X: 5, Y: -5, Image: five},
					{X: -5, Y: -5, Image: one},
				},
			},
		},
	}

	for i, v := range walk_forward.Frames {
		if i%2 == 0 {
			v.Units[0].SetScale(0.2, 0)
			v.Units[1].SetScale(0.2, 0)
			v.Units[2].SetScale(0.5, 0)
			v.Units[3].SetScale(0.1, 0)
			v.Units[4].SetScale(0.1, 0)
		} else {
			v.Units[0].SetScale(0.15, 0)
			v.Units[1].SetScale(0.15, 0)
			v.Units[2].SetScale(0.5, 0)
			v.Units[3].SetScale(0.15, 0)
			v.Units[4].SetScale(0.15, 0)
		}
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
