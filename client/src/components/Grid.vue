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
                    v-if="!solving && matrix && path === ''"
                    type="submit"
                    variant="warning"
                    :disabled="disabled"
                    @click="solve"
            >
                Solve
            </b-button>
            <b-spinner v-else-if="solving && matrix && path === ''" label="Solving..." variant="warning"/>
            <div v-else-if="matrix != null" class="container">
                <b-row class="button-row w-75">
                    <b-col>
                        <b-button
                                v-if="handle === null"
                                variant="warning"
                                :disabled="pathIndex === moves"
                                @click="play"
                        >
                            Play
                        </b-button>
                        <b-button v-else class="bg-warning" @click="stop">Stop</b-button>
                    </b-col>
                    <b-input-group class="col-md-6" prepend="Speed">
                        <b-form-select :disabled="handle" :options="options" v-model="speed"/>
                    </b-input-group>
                    <b-col v-if="pathIndex === 0">
                        <b-button class="bg-danger" @click="unsolve">Unsolve</b-button>
                    </b-col>
                    <b-col v-else>
                        <b-button class="bg-danger" @click="reset">Reset</b-button>
                    </b-col>
                </b-row>
                <b-row class="button-row w-auto">
                    <b-col><b-button :disabled="pathIndex === 0" @click="previousMove">Prev</b-button></b-col>
                    <b-col>
                        <b-button :disabled="true" class="bg-transparent move">
                            {{ pathIndex }} / {{ moves }}
                        </b-button>
                    </b-col>
                    <b-col><b-button :disabled="pathIndex === moves" @click="nextMove">Next</b-button></b-col>
                </b-row>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'Grid',
    data () {
        return {
            handle: null,
            options: ['x1', 'x2', 'x5', 'x10', 'x25'],
            pathIndex: 0,
            speed: 'x1'
        }
    },
    methods: {
        nextMove (event) {
            this.updateGrid(this.path[this.pathIndex]);
            this.pathIndex += 1;
        },
        play (event) {
            this.$emit('play');
            let delay = 1000;
            switch (this.speed) {
                case 'x2':
                    delay /= 2; break;
                case 'x5':
                    delay /= 5; break;
                case 'x10':
                    delay /= 10; break;
                case 'x25':
                    delay /= 25; break;
            }

            this.handle = setInterval(() => {
                this.nextMove();
                if (this.pathIndex === this.moves)
                    this.stop();
            }, delay)
        },
        previousMove (event) {
            this.pathIndex -= 1;
            const moves = ['U', 'D', 'R', 'L'];
            const reverseMoves = ['D', 'U', 'L', 'R'];
            this.updateGrid(reverseMoves[moves.indexOf(this.path[this.pathIndex])]);
        },
        reset (event) {
            this.stop();
            this.pathIndex = 0;
            this.$emit('reset');
        },
        solve (event) {
            this.$emit('solve');
        },
        stop (event) {
            clearInterval(this.handle);
            this.$emit('stop');
            this.handle = null;
        },
        unsolve (event) {
            this.stop();
            this.pathIndex = 0;
            this.$emit('unsolve');
        },
        updateGrid (move) {
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
    props: ['disabled', 'matrix', 'moves', 'path', 'solving'],
    watch: {
        moves: function() { this.pathIndex = 0; }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    .button {
        justify-content: center;
        width: 37vw;
        margin: 2vh;
    }

    .button-row {
        margin-bottom: 0.5vh;
        margin-top: 0.5vh;
    }

    .container {
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
        width: 8vw;
        flex-grow: 1;
        color: black !important;
        border-color: transparent !important;
        font-size: large;
    }

    .number {
        background: -moz-linear-gradient(-45deg, #ff794d 0%, #ffa64d 23%, #ffd24d 69%, #ffdb4d 100%); /* FF3.6-15 */
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
        font-family: TSCu_Comic;
    }
</style>
