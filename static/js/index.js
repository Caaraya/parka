class Picker {
    constructor(target, width, height) {
        this.target = target;
        this.width = width;
        this.height = height;
        this.target.width = width;
        this.target.height = height;
        //Get context 
        this.context = this.target.getContext('2d');
        //Circle (Color Selector Circle)
        this.pickerCircle = { x: 10, y: 10, width: 7, height: 7 };
        this.listenForEvents()
    }
    draw() {
        this.build()
    }
    build() {
        //Create a Gradient Color (colors change on the width)
        let gradient = this.context.createLinearGradient(0, 0, this.width, 0);
        //Add Color Stops to the Gradient (from 0 to 1)
        gradient.addColorStop(0, 'rgb(255, 0, 0)');
        gradient.addColorStop(0.15, 'rgb(255, 0, 255)');
        gradient.addColorStop(0.33, 'rgb(0, 0, 255)');
        gradient.addColorStop(0.49, 'rgb(0, 255, 255)');
        gradient.addColorStop(0.67, 'rgb(0, 255, 0)');
        gradient.addColorStop(0.84, 'rgb(255, 255, 0)');
        gradient.addColorStop(1, 'rgb(255, 0, 0)');
        //Add color picker colors (red, green, blue, yellow...)
        //Render the Color Gradient from the 0's position to the full width and height
        this.context.fillStyle = gradient; ///, set it's style to be the color gradient
        this.context.fillRect(0, 0, this.width, this.height);
        //Apply black and white (on the height dimension instead of the width)
        gradient = this.context.createLinearGradient(0, 0, 0, this.height);
        //We have two colors so 0, 0.5 and 1 needs to be used.
        gradient.addColorStop(0, 'rgba(255, 255, 255, 1)');
        gradient.addColorStop(0.5, 'rgba(255, 255, 255, 0)');
        gradient.addColorStop(0.5, 'rgba(0, 0, 0, 0)');
        gradient.addColorStop(1, 'rgba(0, 0, 0, 1)');
        //set style and render it.
        this.context.fillStyle = gradient;
        this.context.fillRect(0, 0, this.width, this.height);

    }
    listenForEvents() {
        const onMouseDown = (e) => {
          let currentX = e.clientX - this.target.offsetLeft;
          let currentY = e.offsetY;
          this.pickerCircle.y = currentY;
          this.pickerCircle.x = currentX;
        }
        this.target.addEventListener('mousedown', onMouseDown);
        //Mouse move event on the canvas, call callback passing it the current color 
        this.target.addEventListener('mousedown', () => this.onChangeCallback(this.getPickedColor()));
      }
    onChange(callback) {
        //Save Callback function reference on the class
        this.onChangeCallback = callback;
    }
    getPickedColor() {
        //Get the Image Data (pixel value) pointed by the circle by using it's current position
        //getImageData returns an object that has the pixel data (1, 1) is for getting only one pixel.
        let imageData = this.context.getImageData(this.pickerCircle.x, this.pickerCircle.y, 1, 1);
        //Return back an object has the RGB color value of the pointed pixel.
        //The data is an array holds the red, green, blue and alpha values of the current pixel 
        return { r: imageData.data[0], g: imageData.data[1], b: imageData.data[2] };
    }

}
class Animation {
  constructor(animation) {
    this.animationjson = animation
    this.Name = animation.Name
    this.Frames = {}
    this.Frames.collection = animation.Frames
    this.Frames.height = animation.FrameHeight
    this.Frames.width = animation.FrameWidth
  }

  setImage(id, image) {
    this.Frames.collection.forEach(function(frame) {
      frame.Units.forEach(function (unit) {
        if (unit.Image.ImageId == id) {
          unit.Image.Image = image
        }
      })
    })
    this.animationjson.Frames = this.Frames.collection
  }

  getImages() {
    let images = []
    let ids = []
    this.Frames.collection.forEach(function(frame) {
      frame.Units.forEach(function (unit) {
        if (ids.indexOf(unit.Image.ImageId) < 0) {
          ids.push(unit.Image.ImageId)
          images.push(unit.Image)
        }
      })
    })
    return images
  }

  getAnimation() {
    return this.animations
  }
}
class Animations {
  constructor(animations) {
    let anims = JSON.parse(animations)
    animations = []
    anims.forEach(function(anim){
      animations.push(new Animation(anim))
    })
    this.animations = animations
  }
  getAnimationNames() {
    let animnames = []
    this.animations.forEach(function(anim) {
      animnames.push(anim.Name)
    })
    return animnames
  }
  getAnimationFromName(name) {
    let returned = undefined
    this.animations.forEach(function(anim) {
      if(anim.Name == name) {
        returned = anim
      }
    })
    return returned
  }
}
document.addEventListener('DOMContentLoaded', function(event){
    //Create an instance passing it the canvas, width and height
    let picker = new Picker(document.getElementById('color-picker'), 250, 220);

    getAnimations()

    //Draw 
    picker.draw();
    //On Circle position change 
    picker.onChange((color) => {
    //Get the preview DOM element
    let selected_color_picker = document.getElementsByName('select-color')[0];
    let select_fill = document.getElementsByClassName('select-fill-color')[0];
    let select_stroke = document.getElementsByClassName('select-stroke-color')[0];
    //Change it's backagroundColor to the current color (rgb CSS function)
    if(selected_color_picker.value == 'Fill'){
      select_fill.style.backgroundColor = `rgb(${color.r}, ${color.g}, ${color.b})`;
    } else {
      select_stroke.style.backgroundColor = `rgb(${color.r}, ${color.g}, ${color.b})`;
    }
      
  });
  var coll = document.getElementsByClassName('collapsible');
  var i;

  for (i = 0; i < coll.length;   i++) {
    coll[i].addEventListener('click', function() {
      this.classList.toggle('active');
      var content = this.nextElementSibling;
      if (content.style.display === 'block') {
        content.style.display = 'none';
      } else {
        content.style.display = 'block';
      }
    });
  }
  //NOTE: Remeber we are return a color object that has a three properties(Red, Green and Blue)
});

function getAnimations() {
  if (window.XMLHttpRequest) {
    // code for modern browsers
    xmlhttp = new XMLHttpRequest();
 } else {
    // code for old IE browsers
    xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
  }
  xmlhttp.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
        // Typical action to be performed when the document is ready:
        animations = new Animations(this.responseText)
        fillAnimations(animations)
      }
  };
  xmlhttp.open("GET", "/animations", true);
  xmlhttp.responseType = 'Text'
  xmlhttp.setRequestHeader('Content-Type', 'application/json');
  xmlhttp.send(JSON.stringify(this.collectSelectedObject()));
}

function collectSelectedObject() {
  inputs = new Map([
    ['select-stroke-thickness', 'StrokeThickness'],
    ['select-points', 'Points'],
    ['select-minimum-radius', 'MinRad'],
    ['select-stroke-color', 'Stroke.Hex'],
    ['select-stroke-opacity', 'Stroke.Opacity'],
    ['select-fill-color', 'Fill.Hex'],
    ['select-fill-opacity', 'Fill.Opacity'],
    ['select-width', 'SizeCon.Width'],
    ['select-height', 'SizeCon.Height'],
    ['select-pixel-scale', 'SizeCon.PixelScale'],
    ['select-path', 'Path'],
  ])
  result = {};
  inputs.forEach (function(value, key) {
    input = document.getElementById(key)
    var value_inputs = ['input', 'select'];
    let element
    if(input && value_inputs.indexOf(input.tagName.toLowerCase()) !== -1) {
      number = ['select-points', 'select-pixel-scale'].indexOf(key) !== -1 ? parseInt : parseFloat
      element = input.type.toLowerCase()=="number" ? number(input.value): input.value
     
    } else if (input && input.tagName.toLowerCase() == 'span') {
      element = input.style.backgroundColor;
    }
    var result_obj = value.split('.')
    var nested = result_obj.length > 1;
    if (nested) {
      if (result[result_obj[0]]){
        result[result_obj[0]][result_obj[1]] = element
      } else {
        r = {}
        r[result_obj[1]] = element
        result[result_obj[0]] = r
      }
    } else {
      result[value] = element
    }
  })
  return result
}

function fillAnimations(animations) {
  this.animations = animations
  select = document.getElementById("select-animation")
  this.animations.getAnimationNames().forEach(function(animname) {
    node = document.createElement("option")
    node.id = animname
    node.value = animname
    node.textContent = animname
    select.appendChild(node)
  })
  if(select.childElementCount > 0) {
    select.value = select.children[0].value
  }
}

function generateAnimation() {
  let select = document.getElementById("select-animation")
  if (!select) return
  let option = select.selectedOptions[0].value
  if (!option) return
  let shapes = document.getElementsByClassName('selected-output-shape')
  if (!shapes) return
  let animation = this.animations.getAnimationFromName(option)
  let images = animation.getImages()
  for(let y=0;y < images.length;y++){
    animation.setImage(images[y].ImageId,shapes[y % shapes.length].src)
  }
  if (window.XMLHttpRequest) {
    // code for modern browsers
    xmlhttp = new XMLHttpRequest();
  } else {
      // code for old IE browsers
      xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
  }
    xmlhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
          // Typical action to be performed when the document is ready:
          this.responseText
        }
    };
    xmlhttp.open("POST", "/animations", true);
    xmlhttp.responseType = 'Text'
    xmlhttp.setRequestHeader('Content-Type', 'application/json');
    xmlhttp.send(JSON.stringify(animation.animationjson));
}

function generateSingleItem() {
  if (window.XMLHttpRequest) {
    // code for modern browsers
    xmlhttp = new XMLHttpRequest();
 } else {
    // code for old IE browsers
    xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
}
  xmlhttp.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
        // Typical action to be performed when the document is ready:
        addOutputShape(this.responseText)
      }
  };
  xmlhttp.open("POST", "/shape", true);
  xmlhttp.responseType = 'Text'
  xmlhttp.setRequestHeader('Content-Type', 'application/json');
  xmlhttp.send(JSON.stringify(this.collectSelectedObject()));
}

function generateMultipleItems() {
  if (window.XMLHttpRequest) {
    // code for modern browsers
    xmlhttp = new XMLHttpRequest();
  } else {
    // code for old IE browsers
    xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
  }
  xmlhttp.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
        // Typical action to be performed when the document is ready:
        list = this.responseText.split("\n\n")
        list.forEach(function(item) {
          if(item.length > 0) {
            addOutputShape(item)
          }
        })
      }
  };
  xmlhttp.open("POST", "/shapes", true);
  xmlhttp.responseType = 'Text'
  xmlhttp.setRequestHeader('Content-Type', 'application/json')
  let q = document.getElementById('queue')
  let li = []
  q.childNodes.forEach(function(child){
    li.push(JSON.parse(child.innerHTML))
  })
  xmlhttp.send(JSON.stringify(li));
}

function addOutputShape(shape) {
  let q = document.getElementById('shape-output')
  if (q.classList.contains('no-output'))
    q.classList.remove('no-output')
  var img = document.createElement('img');
  img.classList.add('output-shape')
  var svg64 = btoa(shape);
  var b64Start = 'data:image/svg+xml;base64,';
  var image64 = b64Start + svg64;
  img.src = image64;
  q.appendChild(img)
  img.addEventListener("click", selectImage)
}

function selectImage() {
  if(this.classList.contains('selected-output-shape')) {
    this.classList.remove('selected-output-shape')
  } else {
    this.classList.add('selected-output-shape')
  }
  setAnimationDisablement()
}

function setAnimationDisablement() {
  let but = document.getElementById('generate-animation')
  let shapes = document.getElementsByClassName('selected-output-shape')
  if (shapes.length > 0){
    but.disabled = false
    return
  }
  but.disabled = true
}

function clearShapeOutput() {
  let q = document.getElementById('shape-output')
  if (!q.classList.contains('no-output'))
    q.classList.add('no-output')
  let but = document.getElementById('generate-animation')
  but.disabled = true
  q.innerHTML = ''
}


function clearQueue() {
    let q = document.getElementById('queue')
    q.innerHTML = ''
}

function addToQueue() {
    let q = document.getElementById('queue')
    let li = document.createElement('li')
    li.id = 'queue-item-shape'
    li.textContent = JSON.stringify(this.collectSelectedObject())
    q.appendChild(li)
}