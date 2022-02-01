#!/usr/bin/env python

import math, random
import cairo
import sys
from enum import Enum, unique

@unique
class pathShape(Enum):
    STRAIGHT = 0
    ARKED = 1
    NEGATIVE_ARKED = 2

#set defaults
argSystem = {
    "PATH" : pathShape.STRAIGHT ,
    "STROKE_THICKNESS" : 0.04,
    "WIDTH" : 3.0,
    "HEIGHT" : 2.0,
    "PIXEL_SCALE" : 100,
    "POINTS" : 7,
    "MIN_RAD" : 0.5,
    "FILL" : "#ff0066",
    "STROKE": "#ffffff",
    "FILL_OPACITY": 1.0,
    "STROKE_OPACITY":1.0,
    "FILE": sys.stdout.buffer
}
for i in range(1,len(sys.argv)):
    [subject, value] = sys.argv[i].split(":")
    if value is None:
        continue
    upper = subject.strip().upper()
    if upper == "PATH":
        if value == str(1) or value.upper() == "ARKED":
            argSystem["PATH"] = pathShape.ARKED
        elif value == str(2) or value.upper() == "NEGATIVE_ARKED":
            argSystem["PATH"] = pathShape.NEGATIVE_ARKED
        else:
            argSystem["PATH"] = pathShape.STRAIGHT
    elif upper == "FILE" or upper == "FILL" or upper == "STROKE":
        argSystem[upper] = value
    elif upper == "POINTS":
        argSystem[upper] = int(value)
    else:
        argSystem[upper] = float(value)

#set defaults
PATH = argSystem["PATH"]
STROKE_THICKNESS = argSystem["STROKE_THICKNESS"]
WIDTH, HEIGHT, PIXEL_SCALE, POINTS, MIN_RAD= argSystem["WIDTH"], argSystem["HEIGHT"], argSystem["PIXEL_SCALE"], argSystem["POINTS"], argSystem["MIN_RAD"]
FILL, STROKE = argSystem["FILL"], argSystem["STROKE"]
FILL_OPACITY, STROKE_OPACITY = argSystem["FILL_OPACITY"], argSystem["STROKE_OPACITY"]
FILE = argSystem["FILE"]

def hex_to_rgb(value):
    value = value.lstrip('#')
    return list(int(value[i:i+2], 16) for i in (0, 2, 4))

def get_rgb(value):
    value = value.lstrip('rgb(')
    value = value.rstrip(')')
    value = value.split(',')
    return list(int(value[i].strip()) for i in range(3))

def get_max_radius_maximal(deg, center, min_width:int, max_width:int, min_height:int, max_height:int):
    maxlen = min([abs((max_width - center["x"])/math.cos(deg)),
    abs((min_width - center["x"])/math.cos(deg)),
    abs((max_height - center["y"])/math.sin(deg)),
    abs((min_height - center["y"])/math.sin(deg))])
    return maxlen

def get_max_radius_minimal(deg, center, min_width:int, max_width:int, min_height:int, max_height:int):
    return min([max_width - center["x"], center["x"] - min_width, max_height - center["y"], center["y"] - min_height])

def create_random_shape(points:int, min_width:int, max_width:int, min_height:int, max_height, min_radius = 0.5, max_radius_maximal = True):
    l = []
    degs = []
    for i in range(points):
        if i == 0:
            degs.append(random.uniform(math.pi/points, 2*math.pi/ points))
        else:
            degs.append(random.uniform(degs[i-1] + 0.1, 2*(i + 1)*math.pi/ points))

    degs.sort()
    center = { "x":(max_width-min_width)/ 2, "y":(max_height-min_height)/2 }
    #x=Cx+(cos(d/(180/PI))
    #y=Cy+(sin(d/(180/PI))
    last_deg = degs[points-1]
    for d in degs:
        if max_radius_maximal == True:
            maxlen = get_max_radius_maximal(d, center, min_width, max_width, min_height, max_height)
        else:
            maxlen = get_max_radius_minimal(d, center, min_width, max_width, min_height, max_height)
        len = random.uniform(min_radius, maxlen)
        l.append({
            "x":(center["x"] + len * math.cos(d)),
            "y":(center["y"] + len * math.sin(d)),
            "center": center,
            "radius": len,
            "end_rad": d,
            "start_rad": last_deg,
        })
        last_deg = d
    return l

def draw_random_straight_shape(ctx, points:int, min_width:int, max_width:int, min_height:int, max_height, min_radius):
    shape = create_random_shape(points, min_width, max_width, min_height, max_height, min_radius)
    ctx.move_to(shape[0]["x"], shape[0]["y"])
    for index in range(1, len(shape)):
        ctx.line_to(shape[index]["x"], shape[index]["y"])
    ctx.close_path()

def draw_random_arked_shape(ctx, points:int, min_width:int, max_width:int, min_height:int, max_height, min_radius):
    shape = create_random_shape(points, min_width, max_width, min_height, max_height,min_radius,False )
    ctx.move_to(shape[0]["x"], shape[0]["y"])
    for index in range(1, len(shape)):
        ctx.arc(shape[index]["center"]["x"], shape[index]["center"]["y"], shape[index]["radius"], shape[index]["start_rad"], shape[index]["end_rad"])
    ctx.close_path()

def draw_random_negative_arked_shape(ctx, points:int, min_width:int, max_width:int, min_height:int, max_height, min_radius):
    shape = create_random_shape(points, min_width, max_width, min_height, max_height,min_radius,False)
    ctx.move_to(shape[0]["x"], shape[0]["y"])
    for index in range(1, len(shape)):
        ctx.arc_negative(shape[index]["center"]["x"], shape[index]["center"]["y"], shape[index]["radius"], shape[index]["start_rad"], shape[index]["end_rad"])
    ctx.close_path()


surface = cairo.SVGSurface(FILE,
                             WIDTH*PIXEL_SCALE,
                             HEIGHT*PIXEL_SCALE)
ctx = cairo.Context(surface)
ctx.scale(PIXEL_SCALE, PIXEL_SCALE)

ctx.rectangle(0, 0, WIDTH, HEIGHT)
ctx.set_source_rgba(0,0,0,0)
ctx.fill()

# Drawing code
if PATH == pathShape.NEGATIVE_ARKED:
    draw_random_negative_arked_shape(ctx, POINTS, 0,WIDTH, 0,HEIGHT, MIN_RAD)
elif PATH == pathShape.ARKED:
    draw_random_arked_shape(ctx, POINTS, 0,WIDTH, 0,HEIGHT, MIN_RAD)
else:
    draw_random_straight_shape(ctx, POINTS, 0, WIDTH, 0, HEIGHT, MIN_RAD)

if ('rgb' in FILL):
    fill = get_rgb(FILL)
else:
    fill = hex_to_rgb(FILL)
ctx.set_source_rgba(fill[0]/255, fill[1]/255, fill[2]/255, FILL_OPACITY)
ctx.fill_preserve()

if ('rgb' in STROKE):
    stroke = get_rgb(STROKE)
else:
    stroke = hex_to_rgb(STROKE)
ctx.set_source_rgba(stroke[0]/255, stroke[1]/255, stroke[2]/255, STROKE_OPACITY)
ctx.set_line_width(STROKE_THICKNESS)
ctx.stroke()

# End of drawing code
surface.finish()
