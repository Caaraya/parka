#!/usr/bin/env python

import math, random
from typing import List
import cairo
import sys
from enum import Enum, unique

@unique
class pathShape(Enum):
    STRAIGHT = 0
    CURVED = 1

#set defaults
PATH = pathShape.STRAIGHT
WIDTH, HEIGHT, PIXEL_SCALE, POINTS = 3, 2, 100, 4
FILL, STROKE = "#ff0066", "#ffffff"
FILL_OPACITY, STROKE_OPACITY = 1,1

def hex_to_rgb(value):
    value = value.lstrip('#')
    return list(int(value[i:i+2], 16) for i in (0, 2, 4))

def create_random_shape(points:int, min_width:int, max_width:int, min_height:int, max_height):
    l = []
    center = { "x":random.uniform(min_width, max_width), "y":random.uniform(min_height,max_height)}
    current_point = { "x":random.uniform(min_width, max_width), "y":random.uniform(min_height,max_height)}
    last_point = None
    for i in range(points):
        l.append(current_point)
        if last_point is None:
            current_point = { "x":random.uniform(min_width, max_width), "y":random.uniform(min_height,max_height)}
        else:
            minw, maxw, minh, maxh = min_width, max_width, min_height, max_height
            if last_point["x"] >= max_width/2:
                maxw = last_point["x"]
            else:
                minw = last_point["x"]
            if last_point["y"] >= max_height/2:
                maxh = last_point["y"]
            else:
                minh = last_point["y"]
            current_point = { "x":random.uniform(minw, maxw), "y":random.uniform(minh, maxh)}
        last_point = current_point
    return l

def draw_random_straight_shape(ctx, points:int, min_width:int, max_width:int, min_height:int, max_height):
    shape = create_random_shape(points, min_width, max_width, min_height, max_height)
    ctx.move_to(shape[0]["x"], shape[0]["y"])
    for index in range(1, len(shape)):
        ctx.line_to(shape[index]["x"], shape[index]["y"])
    ctx.close_path()

surface = cairo.ImageSurface(cairo.FORMAT_RGB24,
                             WIDTH*PIXEL_SCALE,
                             HEIGHT*PIXEL_SCALE)
ctx = cairo.Context(surface)
ctx.scale(PIXEL_SCALE, PIXEL_SCALE)

ctx.rectangle(0, 0, WIDTH, HEIGHT)
ctx.set_source_rgba(1.0, 1.0, 1.0, 0.0)
ctx.fill()

# Drawing code
draw_random_straight_shape(ctx, POINTS, 0,WIDTH, 0,HEIGHT)

ctx.set_source_rgb(1, 0.5, 0)
ctx.fill_preserve()

ctx.set_source_rgb(1, 1, 0)
ctx.set_line_width(0.04)
ctx.stroke()

# End of drawing code

surface.write_to_png("example.png")  # Output to PNG
