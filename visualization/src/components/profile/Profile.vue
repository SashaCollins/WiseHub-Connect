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
                  <div class="form-group row">
                    <label
                        for="staticEmail"
                        class="col-sm-2 col-form-label-lg">
                      Email:
                    </label>
                    <div class="col-sm">
                      <input
                          type="text"
                          readonly
                          class="form-control-plaintext form-control-lg"
                          id="staticEmail"
                          :value="getUser.email">
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
                        class="col-sm-2 col-form-label-lg">
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
                          class="btn btn-primary">
                        change Password
                      </button>
                    </div>
                  </div>
                </form>
              </div>
            </div>
          </div>
          <div class="row">
            <div v-for="(item, index) in plugins" :key="index" class="col-lg-4 col-md-6 col-sm-12">
              <div class="card">
                <h3 class="mb-3">{{ item.h3 }}</h3>
                <div class="card-body">
                  <label
                      :for="item.name"
                      class="col-md">
                    Username / Host:
                  </label>
                  <div class="col-md col-form-label">
                    <input
                        type="text"
                        disabled
                        class="form-control"
                        :id="item.name"
                        :value="item.name">
                  </div>
                  <label
                      :for="item.token"
                      class="col-md col-form-label">
                    Token:
                  </label>
                  <div class="col-md">
                    <input
                        type="text"
                        disabled
                        class="form-control"
                        :id="item.token"
                        :value="item.token">
                  </div>
                  <div class="btn-group mt-5" role="group">
                    <button type="button" class="btn btn-primary" disabled>Submit</button>
                    <button type="button" class="btn btn-danger">Edit</button>
                  </div>
                </div>
              </div>
            </div>




<!--            <div class="col-lg-4 col-md-6 col-sm-12">-->
<!--              <div class="card">-->
<!--                <h3 class="mb-3">GitHub</h3>-->
<!--                <label-->
<!--                    for="githubUsername"-->
<!--                    class="col-md">-->
<!--                  Username:-->
<!--                </label>-->
<!--                <div class="col-md col-form-label">-->
<!--                  <input-->
<!--                      type="text"-->
<!--                      disabled-->
<!--                      class="form-control"-->
<!--                      id="githubUsername"-->
<!--                      :value="github.name">-->
<!--                </div>-->
<!--                <label-->
<!--                    for="githubApiToken"-->
<!--                    class="col-md col-form-label">-->
<!--                  API Token:-->
<!--                </label>-->
<!--                <div class="col-md">-->
<!--                  <input-->
<!--                      type="text"-->
<!--                      disabled-->
<!--                      class="form-control"-->
<!--                      id="githubApiToken"-->
<!--                      :value="github.token">-->
<!--                </div>-->
<!--                <div class="btn-group mt-5" role="group">-->
<!--                  <button type="button" class="btn btn-primary" disabled>Submit</button>-->
<!--                  <button type="button" class="btn btn-danger">Edit</button>-->
<!--                </div>-->
<!--              </div>-->
<!--            </div>-->
<!--            <div class="col-lg-4 col-md-6 col-sm-12">-->
<!--              <div class="card">-->
<!--                <h3 class="mb-3">Drone CI</h3>-->
<!--                <label-->
<!--                    for="droneHost"-->
<!--                    class="col-md">-->
<!--                  Host:-->
<!--                </label>-->
<!--                <div class="col-md col-form-label">-->
<!--                  <input-->
<!--                      type="text"-->
<!--                      disabled-->
<!--                      class="form-control"-->
<!--                      id="droneHost"-->
<!--                  <div class="d-flex justify-content-between align-items-center">-->
<!--                    <div class="btn-group">-->
<!--                      <button type="button" @click="item.btn_function" class="btn btn-sm btn-outline-primary">{{ item.btn_option }}</button>-->
<!--                      <span v-show="loading" class="spinner-border spinner-border-sm"></span>-->
<!--                      <button type="button" @click="item.delete" :disabled="!(item.status  === 'exited')" class="btn btn-sm btn-outline-danger">Delete</button>-->
<!--                      <span v-show="loading_del" class="spinner-border spinner-border-sm"></span>-->
<!--                    </div>-->
<!--                  </div>-->
<!--                      :value="drone.name">-->
<!--                </div>-->
<!--                <label-->
<!--                    for="droneApiToken"-->
<!--                    class="col-md col-form-label">-->
<!--                  API Token:-->
<!--                </label>-->
<!--                <div class="col-md">-->
<!--                  <input-->
<!--                      type="text"-->
<!--                      disabled-->
<!--                      class="form-control"-->
<!--                      id="droneApiToken"-->
<!--                      :value="drone.token">-->
<!--                </div>-->
<!--                <div class="btn-group mt-5" role="group">-->
<!--                  <button type="button" class="btn btn-primary" disabled>Submit</button>-->
<!--                  <button type="button" class="btn btn-danger">Edit</button>-->
<!--                </div>-->
<!--              </div>-->
<!--            </div>-->
<!--            <div class="col-lg-4 col-md-6 col-sm-12">-->
<!--              <div class="card">-->
<!--                <h3 class="mb-3">Heroku</h3>-->
<!--                <label-->
<!--                    for="herokuUsername"-->
<!--                    class="col-md">-->
<!--                  Username:-->
<!--                </label>-->
<!--                <div class="col-md col-form-label">-->
<!--                  <input-->
<!--                      type="text"-->
<!--                      disabled-->
<!--                      class="form-control"-->
<!--                      id="herokuUsername"-->
<!--                      :value="heroku.name">-->
<!--                </div>-->
<!--                <label-->
<!--                    for="herokuApiToken"-->
<!--                    class="col-md col-form-label">-->
<!--                  API Token:-->
<!--                </label>-->
<!--                <div class="col-md">-->
<!--                  <input-->
<!--                      type="text"-->
<!--                      disabled-->
<!--                      class="form-control"-->
<!--                      id="herokuApiToken"-->
<!--                      :value="heroku.token">-->
<!--                </div>-->
<!--                <div class="btn-group mt-5" role="group">-->
<!--                  <button type="button" class="btn btn-primary" disabled>Submit</button>-->
<!--                  <button type="button" class="btn btn-danger">Edit</button>-->
<!--                </div>-->
<!--              </div>-->
<!--            </div>-->
          </div>
        </div>
    </div>
</template>

<script>
    import User from "../../model/user";

    export default {
        name: "Profile",
        data() {
          return {
            plugins: [{
              'h3': 'Github',
              'name': 'SashaCollins',
              'token': '098765432123456789jhdsfafjkl'
            },{
              'h3': 'Drone CI',
              'name': 'drone.maxuniverse.de',
              'token': '0987654dfklj;sdfsdfgsdfgsfd'
            },{
              'h3': 'Heroku',
              'name': 'SashaCollins',
              'token': 'jhgsdhfkadslkjlkgd9874198071397'
            }]
          }
        },
        computed: {
          getUser() {
            return this.$store.state.user.user;
          }
        },
        methods: {

        },
        mounted() {
          console.log(this.getUser);
          this.$store.dispatch('user/fetchProfile', this.getUser).then(
              (onSuccess) => {
                if (onSuccess.data.Success) {
                  this.plugins.push(onSuccess.data.Plugins);
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

    /* Fake image */
    .fakeimg {
      background-color: #aaa;
      width: 100%;
      padding: 20px;
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