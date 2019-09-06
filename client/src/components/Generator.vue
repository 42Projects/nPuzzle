<template>
    <div>
        <b-form-input type="number" v-model="size"></b-form-input>
        <b-form-group>
            <b-form-radio v-model="picked" value="solvable">Solvable</b-form-radio>
            <b-form-radio v-model="picked" value="unsolvable">Unsolvable</b-form-radio>
        </b-form-group>
        <b-button type="submit" variant="primary" @click="generate">Generate</b-button>
    </div>
</template>

<script>
export default {
    name: 'Generator',
    data () {
        return {
            size: 3,
            picked: 'solvable'
        }
    },
    methods: {
        generate () {
            if (!isNaN(this.size) && this.size >= 3) {

                /* Generate end state */
                const size = Number(this.size);
                let endState = Array.from({length: size * size}, () => -1);
                let current = 1, x = 0, ix = 1, y = 0, iy = 0;
                while (true) {
                    endState[x + y * size] = current;
                    if (current === 0)
                        break;
                    current += 1;
                    if (x + ix === size || x + ix < 0 || (ix !== 0 && endState[x + ix + y * size] !== -1)) {
                        iy = ix;
                        ix = 0;
                    } else if (y + iy === size || y + iy < 0 || (iy !== 0 && endState[x + (y + iy) * size] !== -1)) {
                        ix = -iy;
                        iy = 0;
                    }
                    x += ix;
                    y += iy;
                    if (current === size * size)
                        current = 0;
                }

                /* Scramble puzzle */
                for (let k = 0; k < 10000; k++) {
                    const index = endState.indexOf(0);
                    let moves = [];
                    if (index % size > 0)
                        moves.push(index - 1);
                    if (index % size < size - 1)
                        moves.push(index + 1);
                    if (index / size > 0 && index - size >= 0)
                        moves.push(index - size);
                    if (index / size < size - 1)
                        moves.push(index + size);
                    const move = moves[Math.floor(Math.random() * moves.length)];
                    endState[index] = endState[move];
                    endState[move] = 0;
                }

                /* Solvability */
                if (this.picked === 'unsolvable') {
                    if (endState[0] === 0 || endState[1] === 0) {
                        [endState[-1], endState[-2]] = [endState[-2], endState[-1]]
                    } else {
                        [endState[0], endState[1]] = [endState[1], endState[0]]
                    }
                }

                /* Convert to 2D array */
                let index = 0;
                let matrix = Array.apply(null, Array(size)).map(() => {});
                matrix.forEach((_, y) => {
                    matrix[y] = Array.apply(null, Array(size)).map(() => {});
                    matrix[y].forEach((_, x) => {
                        matrix[y][x] = endState[index++]
                    })
                });

                this.$emit('click', matrix)
            }
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
