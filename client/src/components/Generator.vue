<template>
    <div class="col-md-6">
        <b-input-group prepend="Size">
            <b-form-select :options="options" v-model="size"/>
            <b-input-group-append>
                <b-button type="submit" variant="primary" @click="generate">Generate</b-button>
            </b-input-group-append>
        </b-input-group>
    </div>
</template>

<script>
export default {
    name: 'Generator',
    data () {
        return {
            options: [3, 4, 5, 6, 7, 8, 9],
            size: 3
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
