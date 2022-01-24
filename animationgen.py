import base64
import json
import svgutils.transform as sg
from pathlib import Path
import sys

def float_to_int_string(f):
    return str(int(float(f)))

class Animation:
    def __init__(self):
        self.argsParsed = {}
        for i in range(1,len(sys.argv)):
            [subject, value] = sys.argv[i].split("#", 1)
            if value is None:
                continue
            upper = subject.strip().upper()
            if upper == "FRAMES":
                self.argsParsed[upper] = json.loads(value)
            else:
                self.argsParsed[upper] = value
    
    def getFrames(self):
        return self.argsParsed["FRAMES"]
    
    def getRenderedFrame(self, frame):
        composite = self.getBackgroundImage()
        for imagemeta in self.argsParsed["FRAMES"][frame]["Units"]:
            imageString = base64.b64decode(imagemeta["Image"]["Image"].partition(",")[2]).decode("utf-8")
            ImageOrder = imagemeta["Image"]["ImageId"]
            x = imagemeta["X"]
            y = imagemeta["Y"]
            scale = 0.2 # should get from api
            txt = imageString.replace('\n', '')
            data = sg.fromstring(txt)
            imageW = float(data.width.replace("pt", ''))
            imageH = float(data.height.replace('pt', ''))
            image = self.Svg(data.getroot(), [imageW,imageH ], [0,0])
            image.scale_width_to_reference(int(float(self.argsParsed["SPRITE_WIDTH"])))
            image.scale_by_factor(scale)
            image.move(x, y)
            composite.append([image.data])
            #maybe need overall scale and individual scale, make a composite image with these
        return composite.to_str()
    
    def getBackgroundImage(self):
        txt = Path('background.svg').read_text()
        w = float_to_int_string(self.argsParsed["SPRITE_WIDTH"])
        h = float_to_int_string(self.argsParsed["SPRITE_HEIGHT"])
        txt = txt.replace("<width>", w).replace("<height>", h)
        data = sg.fromstring(txt)
        return data

    class Svg(object):
        "svg files with a data objeSct (the svg), width, height and coordinates"

        def __init__(self, data, dim, coords):
            self.data = data
            self.width = dim[0]
            self.height = dim[1]
            self.x = coords[0]
            self.y = coords[1]

        def scale_width_to_reference(self, reference_width):
            """Proportionally scale the image to a given width."""
            scalings_factor = reference_width / self.width
            self.data.moveto(0, 0, scalings_factor)
            self.width = self.width * scalings_factor
            self.height = self.height * scalings_factor

        def scale_by_factor(self, scalings_factor):
            """Proportionally scale image by a scaling factor."""
            self.data.moveto(0, 0, scalings_factor)
            self.width = self.width * scalings_factor
            self.height = self.height * scalings_factor

        def move(self, x, y):
            """Move the coordinates of an image."""
            self.data.moveto(x, y)
            self.x = x
            self.y = y
