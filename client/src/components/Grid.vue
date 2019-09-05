<template>
    <table>
        <tbody class="matrix">
        <tr v-for="(row, y) in matrix">
            <td class="tile" v-for="(tile, x) in row">
                <b-form-select class="text-center" :options="options" v-model="matrix[y][x]" @change="onChange" />
            </td>
        </tr>
        </tbody>
    </table>
</template>

<script>
export default {
    name: 'Grid',
    created() {
        this.updateOptions(this.size);
        this.updateSize(this.size)
    },
    data () {
        return {
            options: null,
            matrix: null
        }
    },
    methods: {
        onChange() {
            this.$emit('changed', this.matrix)
        },
        updateOptions(size) {
            let options = [];
            for (let k = 0; k < size * size; k++) {
                options.push(k)
            }
            this.options = options
        },
        updateSize(size) {
            let matrix = [];
            for (let k = 0; k < size; k++) {
                matrix.push([]);
                for (let p = 0; p < size; p++) {
                    matrix[k].push(k * size + p + 1)
                }
            }
            matrix[size - 1][ size - 1] = 0;
            this.matrix = matrix
        }
    },
    props: ['size'],
    watch: {
        size: function(newSize) {
            this.updateOptions(newSize);
            this.updateSize(newSize)
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    .matrix {
        height: 15vw;
        width: 15vw;
    }
    .tile {
        border: black solid 0.1vw;
    }
</style>
