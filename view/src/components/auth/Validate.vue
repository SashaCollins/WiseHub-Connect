<template>
  <div class="bg-img">
    <div class="container"  v-if="message">
      <div class="alert alert-danger" role="alert">
        <button type="button" class="close" data-dismiss="alert" aria-label="Close">
          <span aria-hidden="true">&times;</span>
        </button>
        <h4 class="alert-heading">Error</h4>
        <p>{{ message }}</p>
        <hr>
        <p class="mb-0">Please confirm the WiseHub EMail!</p>
      </div>
    </div>
  </div>
</template>

<script>
import User from '../../model/user';

export default {
  name: "ValidateToken",
  data () {
    return {
      message: '',
      user: new User('', '')
    }
  },
  computed: {
    loggedIn() {
      return this.$store.state.auth.status.loggedIn;
    }
  },
  mounted() {
    if (!this.loggedIn) {
      let token = window.location.search.replace('?token=', '')
      this.$store.dispatch('auth/validate', token).then(
          (onSuccess) => {
            if (onSuccess.data.success) {
              this.$router.push('/login');
            }
          }, (onError) => {
            this.message = (onError.response && onError.response.data) || onError.message || onError.toString();
          }
      );
      this.message = '';
      this.$forceUpdate();
    }
  }
}
</script>
