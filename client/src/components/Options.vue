<template>
    <div class="options">
        <b-input-group prepend="Heuristic function">
            <b-form-select
                    class="text-center"
                    v-b-tooltip="test"
                    v-model="heuristic"
                    :options="heuristicOptions"
                    @change="onUpdate"
            />
        </b-input-group>
        <b-input-group prepend="A* search variant">
            <b-form-select class="text-center" v-model="search" :options="searchOptions" @change="onUpdate"/>
        </b-input-group>
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
            heuristic: 'manhattan + linear conflicts',
            searchOptions: [
                { value: 'greedy', text: 'Greedy Search' },
                { value: 'uniform-cost', text: 'Uniform-Cost Search' }
            ],
            search: 'uniform-cost'
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
    .options {
        width: 80%;
    }
</style>
