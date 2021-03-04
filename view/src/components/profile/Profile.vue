<template>
    <div class="container">
        <div class="header">
            <h2>WiseHub-Profile</h2>
        </div>
        <div class="container-fluid">
          <div class="picture">
            <img src="../../assets/avatar.png" alt="Avatar" class="avatar">
          </div>
          <div class="row">
            <div class="col">
              <div class="card">
                <h3 class="text-lg-center">Account</h3>
                <form>
                  <div class="form-group row" @submit.prevent="updateEmail">
                    <label
                        for="staticEmail"
                        class="col-sm-4 col-form-label-lg">
                      Email:
                    </label>
                    <div class="col-sm">
                      <input
                          type="text"
                          readonly
                          class="form-control-plaintext form-control-lg"
                          id="staticEmail"
                          :value="getUser.email"
                          @input="newEmail">
                    </div>
  <!--                  <div class="col-sm-1">-->
  <!--                    <button-->
  <!--                        type="button"-->
  <!--                        class="btn btn-outline-primary btn-sm">-->
  <!--                      Edit-->
  <!--                    </button>-->
  <!--                  </div>-->
                  </div>
                  <div class="form-group row">
                    <label
                        class="col-sm-4 col-form-label-lg">
                      Password:
                    </label>
  <!--                  <div class="col-sm">-->
  <!--                    <input-->
  <!--                        type="password"-->
  <!--                        class="form-control-lg"-->
  <!--                        id="inputPassword"-->
  <!--                        placeholder="Password">-->
  <!--                  </div>-->
                    <div class="col-sm">
                      <button
                          type="button"
                          class="btn"
                          data-toggle="modal"
                          data-target="#updatePassword"
                          style="background-color: #008b8b; color: white">
                        Change password
                      </button>
                    </div>
                  </div>
                </form>
              </div>
            </div>
          </div>
          <div class="row">
            <div v-if="!error" v-for="(item, index) in plugins" :key="index" class="col-lg-4 col-md-6 col-sm-12">
              <div class="card" @submit.prevent="updateCredentials(item)">
                <h3 class="text-center">{{ item.PluginName }}</h3>
                <div class="card-body">
                  <label
                      :for="item.UsernameHost"
                      class="col-md">
                    Username / Host:
                  </label>
                  <div class="col-md col-form-label">
                    <input
                        type="text"
                        class="form-control"
                        :id="item.UsernameHost"
                        v-model="item.UsernameHost"
                        :disabled="!disabled"
                    >
                  </div>
                  <label
                      :for="item.Token"
                      class="col-md col-form-label">
                    Token:
                  </label>
                  <div class="col-md">
                    <input
                        type="text"
                        class="form-control"
                        :id="item.Token"
                        placeholder="*************"
                        v-model="item.Token"
                        :disabled="!disabled"
                    >
                  </div>
                </div>
                <div class="btn-group" role="group">
                  <button @click="updateCredentials(item)" type="submit" class="btn btn-danger" :disabled="!disabled">Submit</button>
                  <button @click="disabled = !disabled" type="button" class="btn" style="background-color: #008b8b; color: white">Edit</button>
                </div>
              </div>
            </div>
            <div v-else v-for="(item, index) in errors" :key="index" class="col-lg-4 col-md-6 col-sm-12">
              <div class="card">
                <h3 class="text-center">{{ item.Tag }}</h3>
                <div class="card-body">
                  <label
                      :for="item.Code"
                      class="col-md">
                    Status:
                  </label>
                  <div class="col-md col-form-label">
                    <label
                        :id="item.Code"
                        class="col-md">
                      {{ item.Code }}
                    </label>
                  </div>
                  <label
                      :for="item.Description"
                      class="col-md col-form-label">
                    Description:
                  </label>
                  <div class="col-md">
                    <textarea
                        :id="item.Description"
                        class="form-control"
                        disabled
                    >
                      {{ item.Description }}
                    </textarea>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      <!-- Modal -->
      <div class="modal fade" id="updatePassword" tabindex="-1" role="dialog" aria-labelledby="updatePasswordTitle" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered" role="document">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title" id="updatePasswordTitle">Update password</h5>
              <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="modal-body">
              <UpdatePassword v-if="getUser"/>
            </div>
          </div>
        </div>
      </div>
    </div>
</template>

<script>
    import UpdatePassword from "@/components/auth/UpdatePassword";

    export default {
        name: "Profile",
      components: {
        UpdatePassword
      },
      data() {
        return {
          disabled: false,
          newEmail: '',
          updatedCredentials: [],
          plugins: [],
          errors: [],
          error: false
        }
      },
      computed: {
        loggedIn() {
          return this.$store.state.auth.status.loggedIn;
        },
        getUser() {
          return this.$store.state.user.user;
        }
      },
      methods: {
        updateEmail: function() {
          this.$store.dispatch('user/updateEmail', {
            oldEmail: this.getUser.email,
            newEmail: this.newEmail
          }).then(
              (onSuccess) => {
                if (onSuccess.data.success) {
                  this.message = "Update successfully!";
                }
              },
              (onError) => {
                if (onError.status === 403 || onError.status === 401) {
                  this.$store.dispatch('auth/refresh');
                }
                this.error = true;
                this.errors = [{
                  'Tag': 'Error - Try again',
                  'Code': onError.status,
                  'Description': (onError.response && onError.response.data) || onError.message || onError.toString()
                }];
              }
          );
        },
        updateCredentials: function(plugin) {
          plugin.Updated = true;
          this.updatedCredentials.push(plugin);
          this.$store.dispatch('user/updateCredentials', {
            email: this.getUser.email,
            plugins: this.updatedCredentials
          }).then(
              (onSuccess) => {
                if (onSuccess.data.success) {
                  this.message = "Update successfully!";
                }
              },
              (onError) => {
                this.error = true;
                this.errors = [{
                  'Tag': 'Error - Try again',
                  'Code': onError.status,
                  'Description': (onError.response && onError.response.data) || onError.message || onError.toString()
                }];
              }
          );
          this.updatedCredentials = [];
          this.disabled = false;
        },
      },
      mounted() {
        this.$store.dispatch('user/fetchProfile', this.getUser).then(
            (onSuccess) => {
              if (onSuccess.data.success) {
                this.plugins = onSuccess.data.plugins;
              }
            },
            (onError) => {
              if (onError.status_code === 401) {
                this.$router.push('/logout');
              }
              this.error = true;
              this.errors = [{
                'Tag': 'Error - Try again',
                'Code': onError.status,
                'Description': (onError.response && onError.response.data) || onError.message || onError.toString()
              }];
            }
        );
      },
    }
</script>

<style scoped lang="scss">

    .picture {
      text-align: center;
      padding-top: 65px;
    }
    .avatar {
      vertical-align: middle;
      width: 100px;
      height: 100px;
      border-radius: 50%;
      background-color: #008B8B;
    }

    /* Create two unequal columns that floats next to each other */

    /* Right column */
    .row {
      width: 100%;
      padding-left: 20px;
    }
    .row:after {
      content: "";
      display: table;
      clear: both;
    }

    .col {
      padding: 20px;
      margin-top: 20px;
    }
    /* Add a card effect for articles */
    .card {
      float: top;
      background-color: white;
      padding: 20px;
      margin-top: 20px;
    }

    /* Clear floats after the columns */
    //.container-fluid:after {
    //  content: "";
    //  display: table;
    //  clear: both;
    //}

    .header {
      position: -webkit-sticky;
      position: sticky;
      top: 0;
    }

    @media screen and (max-width: 800px) {

    }
</style>