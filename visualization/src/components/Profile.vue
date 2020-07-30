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
              <h3 class="text-lg-center">Account</h3>
              <form>
                <div class="row form-group">
                  <label
                      for="staticName"
                      class="col-sm-2 col-form-label-lg">
                    Name:
                  </label>
                  <div class="col-sm">
                    <input
                        type="text"
                        readonly
                        class="form-control-plaintext form-control-lg"
                        id="staticName"
                        value="Heinrich Pumpernickl">
                  </div>
                  <div class="col-sm-1">
                    <button
                        type="button"
                        class="btn btn-outline-primary btn-sm">
                      Edit
                    </button>
                  </div>
                </div>
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
                        value="email@example.com">
                  </div>
                  <div class="col-sm-1">
                    <button
                        type="button"
                        class="btn btn-outline-primary btn-sm">
                      Edit
                    </button>
                  </div>
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
          <div class="row">
            <div class="col">
              <div class="card">
                <h2>About Me</h2>
                <div class="fakeimg" style="height:100px;">Image</div>
                <p>Some text about me in culpa qui officia deserunt mollit anim..</p>
              </div>
            </div>
            <div class="col">
              <div class="card">
                <h3>Popular Post</h3>
                <div class="fakeimg">Image</div><br>
                <div class="fakeimg">Image</div><br>
                <div class="fakeimg">Image</div>
              </div>
            </div>
            <div class="col">
              <div class="card">
                <h3>Follow Me</h3>
                <p>Some text..</p>
              </div>
            </div>
          </div>
        </div>
    </div>
</template>

<script>
    import { mapState, mapActions } from 'vuex'

    export default {
        name: "Profile",
        computed: {

        },

        methods: {

        }
    }
</script>

<style scoped lang="scss">

    .container {
      //font-family: Arial;
      background: #f1f1f1;
    }

    .picture {
      text-align: center;
      padding-top: 30px;
      padding-bottom: 30px;
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
      background-color: #9a2020;
      padding: 20px;
      margin-top: 20px;
    }
    /* Add a card effect for articles */
    .card {
      float: inside;
      background-color: white;
      padding: 20px;
      margin-top: 20px;
    }

    /* Clear floats after the columns */
    .container-fluid:after {
      content: "";
      display: table;
      clear: both;
    }

    @media screen and (max-width: 800px) {
      .general, .plugins {
        width: 100%;
        padding: 0;
      }
    }
</style>