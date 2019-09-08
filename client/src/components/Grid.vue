<template>
    <div class="container">
        <div class="m-4">
            <table>
                <tbody>
                <tr v-for="(row, y) in matrix">
                    <td v-for="(_, x) in row">
                        {{ matrix[y][x] || ""}}
                    </td>
                </tr>
                </tbody>
            </table>
        </div>
        <div class="button">
            <b-button v-if="matrix !== null" type="submit" variant="primary" :disabled="disabled" @click="solve">
                Solve
            </b-button>
        </div>
    </div>
</template>

<script>
export default {
    name: 'Grid',
    methods: {
        solve (event) {
            this.$emit('solve')
        }
    },
    props: ['disabled', 'matrix'],
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    table {
        display: block;
        background-color: yellow;
        max-width: 1000px;
        table-layout: fixed;
    }

    tbody {
        background-color: #444;
        color: #fff;
        border-radius: 25px;
        padding: 20px;
        font-size: 150%;
    }

    td {
        border: black solid 0.1vw;
    }

    .button {
        margin: 3vh;
    }
</style>
