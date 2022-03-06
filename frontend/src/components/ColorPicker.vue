<template>

  <div class="color-picker">
    <div class="color-picker__overlay" v-if="isVisible" v-on:click="hide"></div>
    <transition name="pop">
      <div class="color-picker__flyout" v-if="isVisible">
    <div class="color-chip" v-bind:style="{'background' : rgb}"/>
    <div class="color-picker__inner">
      <div class="control" v-bind:style="gradientR">
        <label class=slider-label>R</label>
        <input type="range" min="0" max="255" v-model="r" />
      </div>
      <div class="control" v-bind:style="gradientG">
        <label class=slider-label>G</label>
        <input type="range" min="0" max="255" v-model="g" />
      </div>
      <div class="control" v-bind:style="gradientB">
        <label class=slider-label>B</label>
        <input type="range" min="0" max="255" v-model="b" />
      </div>
      <div class="control" v-bind:style="gradientA">
        <label class=slider-label>A</label>
        <input type="range" min="0" max="100" v-model="a" />
      </div>
    </div>
  </div>
    </transition>
    <div class="swatch" v-bind:style="{'background': color}" v-on:click="toggle"></div>
  </div>
</template>

<script>
export default {
  name: 'ColorPicker',
  props:["change", color],
  data: function() {
    return {
      isVisible: true,
      r: 265,
      g: 80,
      b: 99,
      a: 100
    }
  },
  computed: {
    rgb: function() {
      var c = parseInt(this.r) + ", " + parseInt(this.g) + ", " + parseInt(this.b);
      var s = "rgb(" + c + ")"; 
      return s;
    },
    gradientR: function() {
      var stops = [];
      for (var i = 0; i < 5; i++) {
        var r = i * 50;
        
        var c = r + ", " + parseInt(this.g) + ", " + parseInt(this.b)
        stops.push("rgb(" + c + ")")
      }

      return {
        backgroundImage: "linear-gradient(to right, " + stops.join(', ') + ")"
      }
    },
    gradientG: function() {
      var stops = [];
      for (var i = 0; i < 5; i++) {
        var g = i * 50;
        
        var c = parseInt(this.r) + ", " + g + ", " + parseInt(this.b)
        stops.push("rgb(" + c + ")")
      }

      return {
        backgroundImage: "linear-gradient(to right, " + stops.join(', ') + ")"
      }
    },
    gradientB: function() {
      var stops = [];
      for (var i = 0; i < 5; i++) {
        var b = i * 50;
        
        var c = parseInt(this.r) + ", " + parseInt(this.g) + ", " + b
        stops.push("rgb(" + c + ")")
      }

      return {
        backgroundImage: "linear-gradient(to right, " + stops.join(', ') + ")"
      }
    },
    gradientA: function() {
      var stops = [];
      for (var i = 0; i < 1; i++) {
        var a = i * 20;
        
        var c = parseInt(this.r) + ", " + parseInt(this.g) + ", " + parseInt(this.b) + ", " + parseFloat(a/100)
        stops.push("rgba(" + c + ")")
      }

      return {
        backgroundImage: "linear-gradient(to right, " + stops.join(', ') + ")"
      }
    },
  },
  methods: {

    show: function() {
      this.isVisible = true;
    },
    hide: function() {
      this.isVisible = false;
    },
    toggle: function() {
      this.isVisible = !this.isVisible;
    }
  },
}
</script>

<style>
body {
  background: #f8f8f8;
  display: flex;
  width: 100%;
  height: 100vh;
  justify-content: center;
  align-items: center;
}

.color-picker {
  position: relative;
}

.color-picker__overlay {
  width: 100%;
  height: 100vh;
  position: fixed;
  top: 0px;
  left: 0;
  background: black;
  z-index: 0;
  opacity: 0;
}

.color-picker__flyout {
  width: 240px;
  border: 1px solid #eee;
  border-radius: 4px;
  background: white;
  box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.12);
  font-family: "Roboto", "Helvetica Neue", sans-serif;
  position: absolute;
  bottom: -170px;
  left: -100px;
  z-index: 2;
}

.color-picker__inner {
  padding: 1.5rem 1rem;
}

.color-chip {
  height: 260px;
  display: flex;
  justify-content: center;
  align-items: center;
  color: white;
  border-radius: 4px 4px 0 0;
}

.control {
  width: 100%;
  height: 12px;
  border-radius: 12px;
  border: 1px solid #ddd;
}

.control + .control {
  margin-top: 1rem;
}

.slider-label {
  text-align: right;
  font-family: 'Courier New', Courier, monospace;
  font-style: oblique;
}

.control input {
  width: 100%;
  margin: 0;
}

.control input[type=range] {
  -webkit-appearance: none;
  width: 100%;
  background: transparent;
}

.control input[type=range]:focus {
  outline: none;
}

.control input[type=range]::-ms-track {
  width: 100%;
  cursor: pointer;
  background: transparent;
  border-color: transparent;
  color: transparent;
}

.control input[type=range]::-webkit-slider-thumb {
  -webkit-appearance: none;
  border: 1px solid #ddd;
  height: 20px;
  width: 20px;
  border-radius: 50px;
  background: #ffffff;
  cursor: pointer;
  box-shadow: 0px 1px 2px rgba(0, 0, 0, 0.12);
  margin-top: -4px;
}

.swatch {
  width: 24px;
  height: 24px;
  margin: 1rem auto 0 auto;
  border: 4px solid white;
  box-shadow: 0px 0px 1px rgba(0, 0, 0, 0.3);
  cursor: pointer;
}

.swatch:hover {
  border: 4px solid white;
  opacity: 0.8;
  box-shadow: 0px 0px 1px rgba(0, 0, 0, 0.3);
}

.pop-enter-active,
.pop-leave-active {
  transition: transform .5s, opacity .5s;
  transition-timing-function: cubic-bezier(.8, .3, 0.25, 1.75);
  transform: scale(1);
}

.pop-enter,
.pop-leave-active {
  opacity: 0;
  transform: scale(0);
}
</style>
