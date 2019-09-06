<template>
    <div class="options">
        <div>
            <p class="mb-1 text-center">Heuristic function:</p>
            <b-form-select class="text-center" v-model="heuristic" :options="heuristicOptions" @change="onUpdate"/>
        </div>
        <div>
            <p class="mb-1 text-center">A* search variant:</p>
            <b-form-select class="text-center" v-model="search" :options="searchOptions" @change="onUpdate"/>
        </div>
    </div>
</template>

<script>
export default {
    name: 'Options',
    data () {
        return {
            heuristicOptions: [
                { value: 'hamming', text: 'Hamming Distance' },
                { value: 'manhattan', text: 'Manhattan Distance' },
                { value: 'manhattan + linear conflicts', text: 'Manhattan Distance + Linear Conflicts' }
            ],
            heuristic: 'hamming',
            searchOptions: [
                { value: 'greedy', text: 'Greedy Search' },
                { value: 'uniform-cost', text: 'Uniform-Cost Search' }
            ],
            search: 'greedy'
        }
    },
    methods: {
        onUpdate () {
            this.$emit('updated', {
                heuristic: this.heuristic,
                search: this.search
            })
        },
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
