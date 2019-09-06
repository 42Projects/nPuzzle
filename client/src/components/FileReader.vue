<template>
    <b-form-file placeholder="Choose a file" :disabled="!serverOnline" @change="loadTextFromFile"/>
</template>
<script>
export default {
    name: 'FileReader',
    methods: {
        loadTextFromFile (event) {
            const file = event.target.files[0];
            const reader = new FileReader();

            reader.onload = e => this.$emit('loading', e.target.result);
            reader.readAsText(file);
        }
    },
    props: ['serverOnline']
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    .text-reader {
        position: relative;
        overflow: hidden;
        display: inline-block;
        padding: 8px 12px;
        cursor: pointer;
    }
</style>
