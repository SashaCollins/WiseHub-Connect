<template>
  <div class="container-fluid">
    <div class="row">
      {{ template }}
    </div>
  </div>
</template>

<script>
export default {
  name: "Template",
  props: ["template"],
}
</script>

<style scoped>

</style>