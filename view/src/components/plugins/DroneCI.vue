<template>
    <div class="card">
      <div class="card-header">
        <h5
            class="text-center"
            style="background-color: #008B8B; color: white; border-radius: 3px; padding: 15px">
          Drone CI
        </h5>
      </div>
      <div class="card-body text-left">
        <p>Branch: {{ drone.repositories.repoBranch }}</p>
        <p>Build: #{{ drone.repositories.build.buildNumber }}</p>
        <p>Status: {{ drone.repositories.build.buildStatus }}</p>
      </div>
    </div>
</template>

<script>
export default {
  name: "Drone CI",
  props: ["drone"],
}
</script>

<style scoped>

</style>