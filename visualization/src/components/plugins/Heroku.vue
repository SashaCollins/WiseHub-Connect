<template>
  <div class="container">
    <div class="row">
      {{ heroku }}
    </div>
  </div>
</template>

<script>
export default {
  name: "Heroku",
  props: ["heroku"],
}
</script>

<style scoped>

</style>