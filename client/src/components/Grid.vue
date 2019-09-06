<template>
    <table>
        <tbody class="matrix">
        <tr v-for="(row, y) in matrix">
            <td class="tile" v-for="(_, x) in row">
                {{ matrix[y][x] }}
            </td>
        </tr>
        </tbody>
    </table>
</template>

<script>
export default {
    name: 'Grid',
    props: ['matrix'],
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
