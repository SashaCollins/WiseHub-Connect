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
<!--                  <div class="row form-group">-->
<!--                    <label-->
<!--                        for="staticName"-->
<!--                        class="col-sm-2 col-form-label-lg">-->
<!--                      Name:-->
<!--                    </label>-->
<!--                    <div class="col-sm">-->
<!--                      <input-->
<!--                          type="text"-->
<!--                          readonly-->
<!--                          class="form-control-plaintext form-control-lg"-->
<!--                          id="staticName"-->
<!--                          value="Heinrich Pumpernickl">-->
<!--                    </div>-->
<!--                    <div class="col-sm-1">-->
<!--                      <button-->
<!--                          type="button"-->
<!--                          class="btn btn-outline-primary btn-sm">-->
<!--                        Edit-->
<!--                      </button>-->
<!--                    </div>-->
<!--                  </div>-->
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
                          class="btn btn-primary"
                          @click="updatePassword">
                        Change Password
                      </button>
                    </div>
                  </div>
                </form>
              </div>
            </div>
          </div>
          <div class="row">
            <div v-for="(item, index) in plugins" :key="index" class="col-lg-4 col-md-6 col-sm-12">
              <div class="card" @submit.prevent="updatePlugin(item)">
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
                        v-model="item.UsernameHost">
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
                        v-model="item.Token">
                  </div>
                </div>
                <div class="btn-group" role="group">
<!--                  <button type="button" class="btn btn-primary" disabled>Submit</button>-->
                  <button @click="updatePlugin(item)" type="submit" class="btn btn-primary">Update</button>
                </div>
              </div>
            </div>
          </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: "Profile",
        data() {
          return {
            newEmail: '',
            updatedPlugins: [],
            //dummy only exists if connection fails
            plugins: [{
              'PluginName': 'Dummy',
              'UsernameHost': 'Test',
              'Token': 'testToken',
              'Description': '',
              'Updated': false,
            }]
          }
        },
        computed: {
          getUser() {
            return this.$store.state.user.user;
          }
        },
        methods: {
          updatePlugin: function(plugin) {
            plugin.Updated = true;
            console.log(plugin);
            // for (let i = 0; i < this.plugins.length; i++) {
            //   let plugin = this.plugins[i];
            //   if (plugin.Updated) {
            //     this.updatedPlugins.push(plugin);
            //   }
            // }
            this.updatedPlugins.push(plugin);
            this.$store.dispatch('user/updatePlugins', {
              email: this.getUser.email,
              plugins: this.updatedPlugins
            }).then(
                (onSuccess) => {
                  console.log("onSuccess in Update")
                  console.log(onSuccess)
                },
                (onError) => {
                  console.log("onError in Update")
                  console.log(onError)
                }
            )
            this.updatedPlugins = [];
          },
          updateEmail: function() {
            this.$store.dispatch('user/updateEmail', {
              oldEmail: this.getUser.email,
              newEmail: this.newEmail
            }).then(
                (onSuccess) => {
                  console.log("onSuccess in Update")
                  console.log(onSuccess)
                },
                (onError) => {
                  console.log("onError in Update")
                  console.log(onError)
                }
            )

          },
          updatePassword: function() {
            this.$store.dispatch('user/updatePassword', this.getUser).then(
                (onSuccess) => {
                  console.log("onSuccess in Update")
                  console.log(onSuccess)
                },
                (onError) => {
                  console.log("onError in Update")
                  console.log(onError)
                }
            )

          }
        },
        mounted() {
          console.log(this.getUser);
          this.$store.dispatch('user/fetchProfile', this.getUser).then(
              (onSuccess) => {
                if (onSuccess.data.success) {
                  this.plugins = onSuccess.data.plugins;
                }
              },
              (onError) => {
                this.message = onError.toString() || onError.message;
              }
          )
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