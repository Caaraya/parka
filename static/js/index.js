class Picker {
    constructor(target, width, height) {
        this.target = target;
        this.width = width;
        this.height = height;
        this.target.width = width;
        this.target.height = height;
        //Get context 
        this.context = this.target.getContext("2d");
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
        gradient.addColorStop(0, "rgb(255, 0, 0)");
        gradient.addColorStop(0.15, "rgb(255, 0, 255)");
        gradient.addColorStop(0.33, "rgb(0, 0, 255)");
        gradient.addColorStop(0.49, "rgb(0, 255, 255)");
        gradient.addColorStop(0.67, "rgb(0, 255, 0)");
        gradient.addColorStop(0.84, "rgb(255, 255, 0)");
        gradient.addColorStop(1, "rgb(255, 0, 0)");
        //Add color picker colors (red, green, blue, yellow...)
        //Render the Color Gradient from the 0's position to the full width and height
        this.context.fillStyle = gradient; ///, set it's style to be the color gradient
        this.context.fillRect(0, 0, this.width, this.height);
        //Apply black and white (on the height dimension instead of the width)
        gradient = this.context.createLinearGradient(0, 0, 0, this.height);
        //We have two colors so 0, 0.5 and 1 needs to be used.
        gradient.addColorStop(0, "rgba(255, 255, 255, 1)");
        gradient.addColorStop(0.5, "rgba(255, 255, 255, 0)");
        gradient.addColorStop(0.5, "rgba(0, 0, 0, 0)");
        gradient.addColorStop(1, "rgba(0, 0, 0, 1)");
        //set style and render it.
        this.context.fillStyle = gradient;
        this.context.fillRect(0, 0, this.width, this.height);

    }
    listenForEvents() {
        let isMouseDown = false;
        const onMouseDown = (e) => {
          let currentX = e.clientX - this.target.offsetLeft;
          let currentY = e.clientY - this.target.offsetTop;
          if(currentY > this.pickerCircle.y && currentY < this.pickerCircle.y + this.pickerCircle.width && currentX > this.pickerCircle.x && currentX < this.pickerCircle.x + this.pickerCircle.width) {
            isMouseDown = true;
          } else {
            this.pickerCircle.x = currentX;
            this.pickerCircle.y = currentY;
          }
        }
        const onMouseMove = (e) => {
          if(isMouseDown) {
           let currentX = e.clientX - this.target.offsetLeft;
           let currentY = e.clientY - this.target.offsetTop;
            this.pickerCircle.x = currentX;
            this.pickerCircle.y = currentY;
          }
        }
        const onMouseUp = () => {
          isMouseDown = false;
        }
        this.target.addEventListener("mousedown", onMouseDown);
        this.target.addEventListener("mousemove", onMouseMove);
        //Mouse move event on the canvas, call callback passing it the current color 
        this.target.addEventListener("mousemove", () => this.onChangeCallback(this.getPickedColor()));
        //Mouse up on the Document     
        document.addEventListener("mouseup", onMouseUp);
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
document.addEventListener("DOMContentLoaded", function(event){
    //Create an instance passing it the canvas, width and height
    let picker = new Picker(document.getElementById("color-picker"), 250, 220);

    //Draw 
    picker.draw();
    //On Circle position change 
    picker.onChange((color) => {
    //Get the preview DOM element
    let selected = document.getElementsByClassName("selected")[0];
    //Change it's backagroundColor to the current color (rgb CSS function)
    selected.style.backgroundColor = `rgb(${color.r}, ${color.g}, ${color.b})`;
    
  });
  var coll = document.getElementsByClassName("collapsible");
  var i;

  for (i = 0; i < coll.length; i++) {
    coll[i].addEventListener("click", function() {
      this.classList.toggle("active");
      var content = this.nextElementSibling;
      if (content.style.display === "block") {
        content.style.display = "none";
      } else {
        content.style.display = "block";
      }
    });
  }
  //NOTE: Remeber we are return a color object that has a three properties(Red, Green and Blue)
});


