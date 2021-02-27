<template>
  <div class="container">
    <div class="col-8">
      <div class="segment">

        <form class="form">
          <div class="field" :class="{ error: errors.has('password') }">
            <label>Current Password</label>
            <input type="password" name="password" v-model="password" placeholder="Current Password" data-vv-as="current password" v-validate="'required'">
            <span v-show="errors.has('password')" class="is-danger">{{ errors.first('password') }}</span>
          </div>

          <div class="field" :class="{ error: errors.has('newPassword') }">
            <label>New Password</label>
            <input type="password" name="newPassword" ref="newPassword" v-model="getUser.password" placeholder="New Password" data-vv-as="new password" v-validate="'required'">
            <span v-show="errors.has('newPassword')" class="is-danger">{{ errors.first('newPassword') }}</span>
          </div>

          <div class="field" :class="{ error: errors.has('confirmPassword') }">
            <label>New Password</label>
            <input type="password" name="confirmPassword" v-model="confirmPassword" placeholder="New Password, Again" data-vv-as="confirm password" v-validate="'required|confirmed:newPassword'">
            <span v-show="errors.has('confirmPassword')" class="is-danger">{{ errors.first('confirmPassword') }}</span>
          </div>
        </form>
      </div>
      <footer>
        <div>
          <button type="submit" @click="changePassword" class="btn btn-primary" :disabled="!isFormValid">Change password</button>
        </div>
        <div
            v-if="message"
            class="alert"
            :class="successful ? 'alert-success' : 'alert-danger'"
        >{{message}}</div>
      </footer>
    </div>
  </div>
</template>

<script>
  export default {
    name: "UpdatePassword",
    data () {
      return {
        password: '',
        newPassword: '',
        confirmPassword: '',
        message: '',
        successful: false
      }
    },
    computed: {
      loggedIn() {
        return this.$store.state.auth.status.loggedIn;
      },
      getUser() {
        return this.$store.state.user.user;
      },
      isFormValid () {
        return Object.keys(this.fields).every(key => this.fields[key].valid);
      }
    },
    methods: {
      //handle a change password request
      changePassword: function () {
        this.$validator.validate().then(isValid => {
          if (isValid) {
            this.$store.dispatch('user/updatePassword', this.getUser).then(() => {
              this.successful = true;
              this.message = "Update successfully!";
            }, error => {
              if (onError.status === 403 || onError.status === 401) {
                this.$store.dispatch('auth/refresh');
              }
              this.message = (error.response && error.response.data) || error.message || error.toString();
            });
            this.message = '';
            this.$forceUpdate();
          }
        });
      }
    },
    //when user is not logged in this component is unavailable
    mounted() {
      if (!this.loggedIn) {
        this.$router.push('/login');
      }
    }
  }
</script>

<style scoped>

</style>