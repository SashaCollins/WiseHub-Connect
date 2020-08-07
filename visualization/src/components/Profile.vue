<template>
    <div class="container">
        <div class="header">
            <h2>WiseHub-Profile</h2>
        </div>
        <div class="container-fluid">
          <div class="picture">
            <img src="../assets/avatar.png" alt="Avatar" class="avatar">
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
            <div class="col-lg-4 col-md-6 col-sm-12">
              <div class="card">
                <h3 class="mb-3">GitHub</h3>
                <label
                    for="githubUsername"
                    class="col-md">
                  GitHub Username:
                </label>
                <div class="col-md col-form-label">
                  <input
                      type="text"
                      disabled
                      class="form-control"
                      id="githubUsername"
                      value="Username">
                </div>
                <label
                    for="githubApiToken"
                    class="col-md col-form-label">
                  GitHub API Token:
                </label>
                <div class="col-md">
                  <input
                      type="text"
                      disabled
                      class="form-control"
                      id="githubApiToken"
                      value="API Token">
                </div>
                <div class="btn-group mt-5" role="group">
                  <button type="button" class="btn btn-primary" disabled>Submit</button>
                  <button type="button" class="btn btn-danger">Edit</button>
                </div>
              </div>
            </div>
            <div class="col-lg-4 col-md-6 col-sm-12">
              <div class="card">
                <h3 class="mb-3">Drone CI</h3>
                <label
                    for="droneHost"
                    class="col-md">
                  Drone Host:
                </label>
                <div class="col-md col-form-label">
                  <input
                      type="text"
                      disabled
                      class="form-control"
                      id="droneHost"
                      value="Host">
                </div>
                <label
                    for="droneApiToken"
                    class="col-md col-form-label">
                  Drone API Token:
                </label>
                <div class="col-md">
                  <input
                      type="text"
                      disabled
                      class="form-control"
                      id="droneApiToken"
                      value="API Token">
                </div>
                <div class="btn-group mt-5" role="group">
                  <button type="button" class="btn btn-primary" disabled>Submit</button>
                  <button type="button" class="btn btn-danger">Edit</button>
                </div>
              </div>
            </div>
            <div class="col-lg-4 col-md-6 col-sm-12">
              <div class="card">
                <h3 class="mb-3">Heroku</h3>
                <label
                    for="herokuUsername"
                    class="col-md">
                  Heroku Username:
                </label>
                <div class="col-md col-form-label">
                  <input
                      type="text"
                      disabled
                      class="form-control"
                      id="herokuUsername"
                      value="Username">
                </div>
                <label
                    for="herokuApiToken"
                    class="col-md col-form-label">
                  Heroku API Token:
                </label>
                <div class="col-md">
                  <input
                      type="text"
                      disabled
                      class="form-control"
                      id="herokuApiToken"
                      value="API Token">
                </div>
                <div class="btn-group mt-5" role="group">
                  <button type="button" class="btn btn-primary" disabled>Submit</button>
                  <button type="button" class="btn btn-danger">Edit</button>
                </div>
              </div>
            </div>
          </div>
        </div>
    </div>
</template>

<script>
    import { mapState, mapActions } from 'vuex'
    import User from "@/model/user";

    export default {
        name: "Profile",
        data() {
          return {

          }
        },
        computed: {
          getUser() {
            console.log(this.$store.state.auth.status.user.email)
            return this.$store.state.auth.status.user;
          }
        },
        methods: {

        },
        mounted() {
          // function() {
          //   this.submitted = true;
          //   this.$validator.validate().then(isValid => {
          //     if (isValid) {
          //       this.$store.dispatch("auth/login", this.user).then(
          //           (onSuccess) => {
          //             console.log(onSuccess)
          //             this.$router.push("/repositories");
          //           },
          //           (onFailure) => {
          //             console.log(onFailure.response)
          //             this.message = (onFailure.response && onFailure.response.data) || onFailure.message || onFailure.toString();
          //             console.log(this.message)
          //             this.submitted = false;
          //           })
          //     }
          //   })
          //   console.log(this.user);
          // },
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