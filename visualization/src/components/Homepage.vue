<template>
    <div class="container">
        <b-container fluid="true" class="container-fluid">
            <h1>WiseHub-Connect</h1>
        </b-container>
        <b-container fluid="true" class="container-fluid">

        </b-container>

    </div>
</template>

<script>

    export default {
        name: "Homepage",

    }
</script>

<style scoped lang="scss">
    h1 {
        color: #008B8B;
        width: 100%;
        margin-top: 15px;
        border-radius: 10px;
        text-align: center;
        background: #F5FFFA;
    }
    .container-fluid {
        width: 75%;
        height: fit-content;
        margin-top: 35px;
        padding: 20px;
        border-radius: 10px;
        background: #F5FFFA;

    }
    @media screen and (max-width: 800px) {
        .container-fluid {
            width: 100%;
            padding: 0;
        }
    }
</style>