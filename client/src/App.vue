<template>
  <div id="app">
    <div class="col-9 row">
      <div class="col-4">
        <h3>N-Puzzle solver</h3>
        <Options @updated="updateOptions"/>
      </div>
      <Grid :size="size" @changed="updateMatrix"/>
    </div>
  </div>
</template>

<script>
import Grid from './components/Grid.vue'
import Options from './components/Options.vue'

export default {
  name: 'app',
  components: {
    Grid,
    Options
  },
  data () {
    return {
      matrix: null,
      heuristic: 'hamming',
      search: 'greedy',
      size: 3
    }
  },
  methods: {
    updateMatrix (matrix) {
      this.matrix = matrix
    },
    updateOptions (options) {
      this.heuristic = options.heuristic;
      this.search = options.search;
      this.size = options.size;
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
