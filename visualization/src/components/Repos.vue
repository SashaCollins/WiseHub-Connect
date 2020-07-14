<template>
    <div v-if="loggedIn">
	    Logged In WOOOOOHOOOOO
    </div>
</template>

<script>
    export default {
      name: "Repos",
      data () {
        return {}

      },
	  computed: {
		loggedIn() {
		  return this.$store.state.auth.status.loggedIn;
		}
	  },
      methods: {

      }
    }
</script>

<style scoped>

</style>