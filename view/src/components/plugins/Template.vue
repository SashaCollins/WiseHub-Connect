<template>
  <div class="card">
    <div class="card-header">
      <h5
          class="text-center"
          style="background-color: #464646; color: white; border-radius: 3px; padding: 15px">
        Template
      </h5>
    </div>
    <div class="card-body text-left">
      <p>Template for another plugin</p>
      <p>Add here your data that is fetched from your plugin, to be displayed in general view</p>
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