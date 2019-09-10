<template>
    <div class="container">
        <div class="grid">
            <div class="row">
                <div class="row" v-for="(row, y) in matrix">
                    <div v-if="matrix[y][x] != 0" class="tile number" v-for="(_, x) in row">
                        {{ matrix[y][x] }}
                    </div>
                    <div v-else class="tile"></div>
                </div>
            </div>
        </div>
        <div class="button">
            <b-button
                    v-if="matrix != null && path === ''"
                    type="submit"
                    variant="primary"
                    :disabled="disabled"
                    @click="solve"
            >
                Solve
            </b-button>
            <b-row v-else-if="matrix != null">
                <b-col><b-button :disabled="pathIndex === 0" @click="previousMove">Prev</b-button></b-col>
                <b-col>
                    <b-button :disabled="true" class="bg-transparent move">
                        {{ pathIndex }} / {{ moves }}
                    </b-button>
                </b-col>
                <b-col><b-button :disabled="pathIndex === moves" @click="nextMove">Next</b-button></b-col>
                <b-col><b-button class="bg-danger" @click="resetSolution">Reset</b-button></b-col>
            </b-row>
        </div>
    </div>
</template>

<script>
export default {
    name: 'Grid',
    data () {
        return {
            pathIndex: 0
        }
    },
    methods: {
        nextMove (event) {
            this.updateGrid(this.path[this.pathIndex]);
            this.pathIndex += 1;
        },
        previousMove (event) {
            this.pathIndex -= 1;
            const moves = ['U', 'D', 'R', 'L'];
            const reverseMoves = ['D', 'U', 'L', 'R'];
            this.updateGrid(reverseMoves[moves.indexOf(this.path[this.pathIndex])]);
        },
        resetSolution (event) {
            this.$emit('reset');
        },
        solve (event) {
            this.$emit('solve');
        },
        updateGrid (move) {
            console.log(move);

            let zx = -1, zy = -1;
            for (let y = 0; y < this.matrix.length && zy === -1; y++) {
                for (let x = 0; x < this.matrix.length && zx === -1; x++) {
                    if (!this.matrix[y][x]) {
                        zx = x;
                        zy = y;
                    }
                }
            }

            switch (move) {
                case "R":
                    this.matrix[zy][zx] = this.matrix[zy][zx + 1];
                    this.matrix[zy][zx + 1] = 0;
                    break;
                case "L":
                    this.matrix[zy][zx] = this.matrix[zy][zx - 1];
                    this.matrix[zy][zx - 1] = 0;
                    break;
                case "U":
                    this.matrix[zy][zx] = this.matrix[zy - 1][zx];
                    this.matrix[zy - 1][zx] = 0;
                    break;
                case "D":
                    this.matrix[zy][zx] = this.matrix[zy + 1][zx];
                    this.matrix[zy + 1][zx] = 0;
            }
        },
    },
    props: ['disabled', 'matrix', 'moves', 'path'],
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    .button {
        display: flex;
        justify-content: center;
        width: 25vw;
        margin: 2vh;
    }

    .container {
        background-color: #fafafa;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    }

    .grid {
        border-radius: 20px;
        display: flex;
        justify-content: center;
        width: 25vw;
        height: 25vw;
        margin-top: 2vw;
    }

    .move {
        color: black !important;
        border-color: transparent !important;
        font-size: large;
    }

    .number {
        background-color: #e1e1e1;
    }

    .row {
        justify-content: center;
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        width: 100%;
    }

    .tile {
        border: 2px solid white;
        border-radius: 5px;
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
        flex: 1;
        font-size: 3vh;
        font-family: AkrutiMal1;
    }
</style>
