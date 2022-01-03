import sys
class Animation:
    def __init__():
        self.argsParsed = {}
        for i in range(1,len(sys.argv)):
            [subject, value] = sys.argv[i].split(":", 1)
            if value is None:
                continue
            upper = subject.strip().upper()
            if upper == "FRAMES":
                FRAMES = []
                frames = value.split("|||")
                for f in range(len(frames)):
                    FRAMES[f] = []
                    units = frames[f].split("||")
                    for u in range(len(units)):
                        FRAMES[f][u] = {}
                        unit = units[u].split("|")
                        i=0
                        while i<len(unit):
                            FRAMES[f][u][unit[i].upper()] = unit[i+1]
                            i += 2
                self.argsParsed["FRAMES"] = FRAMES
            else:
                self.argsParsed[upper] = value
    
    def getFrames():
        return self.argsParsed["FRAMES"]

    class Svg(object):
        "svg files with a data object (the svg), width, height and coordinates"

        def __init__(self, data, dim, coords):
            self.data = data
            self.width = dim[0]
            self.height = dim[1]
            self.x = coords[0]
            self.y = coords[1]

        def scale_width_to_reference(self, reference_width):
            """Proportionally scale the image to a given width."""
            scalings_factor = reference_width / self.width
            self.data.moveto(0, 0, scale=scalings_factor)
            self.width = self.width * scalings_factor
            self.height = self.height * scalings_factor

        def scale_by_factor(self, scalings_factor):
            """Proportionally scale image by a scaling factor."""
            self.data.moveto(0, 0, scale=scalings_factor)
            self.width = self.width * scalings_factor
            self.height = self.height * scalings_factor

        def move(self, x, y):
            """Move the coordinates of an image."""
            self.data.moveto(x, y)
            self.x = x
            self.y = y