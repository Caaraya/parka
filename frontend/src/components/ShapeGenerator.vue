<template>
  <div>
    <h2>Let's Create a Color Picker</h2>
    <button class="collapsible">inputs</button>
    <div class="container content">
      <canvas id="color-picker"></canvas>
      <select id="select-color" name="select-color">
        <option id="Fill">Fill</option>
        <option id="Stroke">Stroke</option>
      </select>
      <div class="inputs">
        <label for="select-stroke-thickness">Stroke Thickness:</label>
        <input type="number" id="select-stroke-thickness" :value="selected.StrokeThickness"/>
        <label for="select-points">Points:</label>
        <input type="number" id="select-points" :value="selected.Points"/>
        <label for="select-fill">Fill:</label>
        <span id="select-fill">
          <span id="select-fill-color" class="select-fill-color" v-bind:style="{'background-color': selected.Stroke.Hex}"></span>
          <input type="number" id="select-fill-opacity" :value="selected.Fill.Opacity"/>
        </span>
        <label for="select-stroke">Stroke:</label>
        <span id="select-stroke">
          <span id="select-stroke-color" class="select-stroke-color" v-bind:style="{'background-color': selected.Stroke.Hex}"></span>
          <input type="number" id="select-stroke-opacity" :value="selected.Stroke.Opacity"/>
        </span>
        <label for="select-path">Path:</label>
        <select id="select-path" name="select-path" :value="selected.Path">
          <option id="STRAIGHT">STRAIGHT</option>
          <option id="ARKED">ARKED</option>
          <option id="NEGATIVE_ARKED">NEGATIVE_ARKED</option>
        </select>
        <label for="select-minimum-radius">Minimum Radius:</label>
        <input type="number" id="select-minimum-radius" :value="selected.MinRad"/>
        <label for="select-width">Width:</label>
        <input type="number" id="select-width" :value="selected.SizeCon.Width"/>
        <label for="select-height">Height:</label>
        <input type="number" id="select-height" :value="selected.SizeCon.Height"/>
        <label for="select-pixel-scale">Pixel Scale:</label>
        <input type="number" id="select-pixel-scale" :value="selected.SizeCon.PixelScale"/>
    </div>
    <div class="outputs">
      <button id="single-item" onclick="generateSingleItem()">Generate</button>
      <button id="add-to-queue" onclick="addToQueue()">Add to queue</button>
      <ol id="queue"></ol>
      <button onclick="generateMultipleItems()" id="generate-from-queue">Generate from queue</button>
      <button onclick="clearQueue()" id="clear-queue">Clear Queue</button>
    </div>
    </div>
    <div  id="output">
      <div id="shape-output" class="no-output"></div>
      <button onclick="clearShapeOutput()" id="clear-shape-output">Clear Shape Output</button>
      <div class="inputs">
        <select id="select-animation" name="select-animation" ></select>
        <button onclick="generateAnimation()" disabled id="generate-animation">Genetate Animation</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ShapeGenerator',
  props: {
    selected: Object
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
