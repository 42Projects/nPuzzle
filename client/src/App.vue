<template>
    <div id="app">
        <div class="col-9 row">
            <div class="col-4">
                <h3>N-Puzzle solver</h3>
                <p>Server status: {{ serverOnline ? 'online' : 'offline' }}</p>
                <Options @updated="updateOptions" />
                <FileReader :serverOnline="serverOnline" @loading="loadMatrix"/>
                <Generator @click="updateMatrix"/>
                <b-button type="submit" variant="primary" :disabled="!serverOnline" @click="solve">
                    Solve
                </b-button>
            </div>
            <Grid :matrix="matrix"/>
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
        window.setInterval(ping, 5000);
    },
    data () {
        return {
            client: null,
            heuristic: 'hamming',
            matrix: [],
            pathIndex: -1,
            solutionPath: '',
            search: 'greedy',
            serverOnline: false
        }
    },
    methods: {
        loadMatrix (text) {
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
                        }
                    })
                }
            });
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
                this.client.solve(problem, {}, (err, res) => {
                    if (err) {
                        alert(err);
                    } else {

                    }
                })
            }
        },
        updateMatrix (matrix) {
            this.matrix = matrix
        },
        updateOptions (options) {
            this.heuristic = options.heuristic;
            this.search = options.search;
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
