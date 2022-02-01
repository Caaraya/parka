import base64
import json
import svgutils.transform as sg
from pathlib import Path
import sys

def float_to_int_string(f):
    return str(int(float(f)))

def float_from_pt(f):
    return float(f.replace("pt", ''))

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
            minutescale = imagemeta["Scale"]
            x = imagemeta["X"]
            y = imagemeta["Y"]
            txt = imageString.replace('\n', '')
            data = sg.fromstring(txt)
            overallWidth = float(self.argsParsed["SPRITE_WIDTH"])
            imageW = float_from_pt(data.width)
            imageH = float_from_pt(data.height)
            image = data.getroot()
            scale = overallWidth/imageW * float(minutescale)
            newposx = float_from_pt(composite.width)/2 - x - scale*imageW/2
            newposy = float_from_pt(composite.height)/2 - y - scale*imageH/2
            image.moveto(newposx, newposy, scale)
            composite.append([image])
            #maybe need overall cale and individual scale, make a composite image with these
        return composite.to_str().decode()
    
    def getBackgroundImage(self):
        txt = Path('background.svg').read_text()
        w = float_to_int_string(self.argsParsed["SPRITE_WIDTH"])
        h = float_to_int_string(self.argsParsed["SPRITE_HEIGHT"])
        txt = txt.replace("<width>", w).replace("<height>", h)
        data = sg.fromstring(txt)
        return data