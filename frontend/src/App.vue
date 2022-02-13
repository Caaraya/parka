<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <ShapeGenerator :selected="getSelection()"/>
  </div>
</template>

<script>
import ShapeGenerator from './components/ShapeGenerator.vue'

export default {
  name: 'App',
  components: {
    ShapeGenerator
  },
  methods: {
    getSelection(){
      return {
			StrokeThickness: 0.3,
			Points:          4,
			Fill:{
				Hex:     "#ccccccc",
				Opacity: 1.0,
			},
			Stroke:{
				Hex:     "#772233",
				Opacity: 0.7,
			},
			Path:   "ARKED",
			MinRad: 1.0,
			SizeCon:{
				Width:      4.0,
				Height:     4.0,
				PixelScale: 100,
			},
      }
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
