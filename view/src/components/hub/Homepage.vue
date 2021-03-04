<template>
  <div class="bg-img">

  </div>
</template>

<script>
    export default {
        name: "Homepage",

    }
</script>

<style scoped lang="scss">

  .bg-img {
    /* The image used */
    background-image: url("../../assets/wisehubLogoV.png");

    /* Control the height of the image */
    min-height: 100vh;

    /* Center and scale the image nicely */
    background-position: center;
    background-repeat: no-repeat;
    background-size: cover;
    position: relative;
  }
  h1 {
      color: #008B8B;
      width: 100%;
      margin-top: 15px;
      border-radius: 10px;
      text-align: center;
      background: #F5FFFA;
  }

  @media screen and (max-width: 767px) {

  }
</style>