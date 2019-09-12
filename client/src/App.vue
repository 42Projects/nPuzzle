<template>
    <div id="app" class="container">
        <b-navbar variant="faded" type="light">
            <b-navbar-brand href="#">
                <img src="logo.png" alt="nPuzzle">
            </b-navbar-brand>
            <b-nav-text right></b-nav-text>
            <div class="row status-wrapper">
                <div id="status">Server status:</div>
                <b-badge v-if="serverOnline" class="d-flex" variant="success">Online</b-badge>
                <b-badge v-else class="d-flex" variant="danger">Offline</b-badge>
            </div>
        </b-navbar>
        <div class="row">

            <!-- Left side with controls -->
            <div class="col-md-6 p-0 left-wrapper">
                <div class="options-wrapper">
                    <Options @updated="updateOptions" />
                </div>
                <div class="row picker-wrapper">
                    <FileReader
                            class="col-md-5 text-left"
                            :disabled="disableUpdate"
                            :serverOnline="serverOnline"
                            @loading="loadMatrix"
                    />
                    <span>
                        <strong>OR</strong>
                    </span>
                    <Generator :disabled="disableUpdate" @click="updateMatrix"/>
                </div>
            </div>

            <!-- Right side with grid and solution -->
            <div class="col-md-6 p-0">
                <Grid
                        :disabled="!serverOnline"
                        :matrix="matrix"
                        :moves="moves"
                        :path="solutionPath"
                        :solving="solving"
                        @play="disableUpdate = true"
                        @reset="reset"
                        @solve="solve"
                        @stop="disableUpdate = false"
                        @unsolve="unsolve"
                />
            </div>

        </div>
    </div>
</template>

<script>
import { Matrix, Message, Problem } from './npuzzle_pb';
import { NpuzzleClient } from './npuzzle_grpc_web_pb'
import FileReader from './components/FileReader.vue'
import Generator from './components/Generator.vue'
import Grid from './components/Grid.vue'
import Options from './components/Options.vue'

export default {
    name: 'app',
    components: {
        FileReader,
        Generator,
        Grid,
        Options
    },
    created () {
        this.client = new NpuzzleClient('http://' + window.location.hostname + ':8080', null, null);
        const ping = () => {
            let message = new Message();
            message.setMessage('ping');
            this.client.greets(message, {}, err => this.serverOnline = !err);
        };
        ping();
        window.setInterval(ping, 1000);
    },
    data () {
        return {
            client: null,
            disableUpdate: false,
            heuristic: 'hamming',
            initialMatrix: [[1, 2, 3], [8, 0, 4], [7, 6, 5]],
            matrix: [[1, 2, 3], [8, 0, 4], [7, 6, 5]],
            moves: 0,
            solutionPath: '',
            solving: false,
            search: 'greedy',
            serverOnline: false
        }
    },
    methods: {
        loadMatrix (text) {
            this.moves = '';
            this.solutionPath = '';
            this.matrix = [];
            let message = new Message();
            message.setMessage('ping');
            this.client.greets(message, {}, err => {
                if (err) {
                    alert('error (' + err.code + "): server isn't responding")
                } else {
                    message.setMessage(text);
                    this.client.parse(message, {}, (err, res) => {
                        if (!res.getSuccess()) {
                            alert('error parsing file: ' + res.getError());
                        } else {
                            let matrix = [];
                            res.getRowsList().forEach(elem => matrix.push(elem.getNumList()));
                            this.matrix = matrix;
                            this.initialMatrix = matrix.map(row => row.slice());
                        }
                    })
                }
            });
        },
        reset () {
            this.matrix = this.initialMatrix.map(row => row.slice());
        },
        solve () {
            if (this.serverOnline) {
                let problem = new Problem();
                problem.setHeuristic(this.heuristic);
                problem.setSearch(this.search);
                let rows = [];
                this.matrix.forEach(row => {
                    let matrixRow = new Matrix.Row();
                    matrixRow.setNumList(row);
                    rows.push(matrixRow);
                });
                problem.setRowsList(rows);
                this.disableUpdate = true;
                this.solving = true;
                this.client.solve(problem, {}, (err, res) => {
                    this.disableUpdate = false;
                    this.solving = false;
                    if (err) {
                        alert(err);
                    } else if (res.getSuccess() == false) {
                        alert(res.getError());
                    } else {
                        this.moves = res.getMoves();
                        this.solutionPath = res.getPath();
                    }
                })
            }
        },
        unsolve () {
            this.matrix = this.initialMatrix.map(row => row.slice());
            this.moves = '';
            this.solutionPath = '';
        },
        updateMatrix (matrix) {
            this.moves = 0;
            this.solutionPath = '';
            this.matrix = matrix;
            this.initialMatrix = matrix.map(row => row.slice());
        },
        updateOptions (options) {
            this.heuristic = options.heuristic;
            this.search = options.search;
        }
    }
}
</script>

<style>
    .left-wrapper {
        display: flex;
        flex-direction: column;
        justify-content: center;
    }

    .options-wrapper {
        display: flex;
        justify-content: center;
        margin-bottom: 2vh;
        margin-top: 2vh;
    }

    .picker-wrapper {
        justify-content: space-around;
    }

    .status-wrapper {
        align-self: center;
        width: max-content;
    }

    #app {
        font-family: 'Avenir', Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        text-align: center;
        color: #2c3e50;
    }

    .navbar {
        margin-left: -20%;
        width: 140%;
    }

    #status {
        margin-right: 0.5vw;
    }

    img {
        margin-left: -15%;
        height: 6vh;
        width: auto;
    }

    span {
        align-items: center;
        justify-content: center;
        flex-grow: 1;
        display: flex;
    }
</style>
