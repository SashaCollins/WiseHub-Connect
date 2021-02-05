<template>
  <div class="container-fluid">
    <div class="header">
      <h2>WiseHub-Courses</h2>
    </div>

    <div v-if="!message">
      <div v-if="!loading">
        <VersionManagerCards v-bind:vmCards="getVersionManagerCards('Github')"></VersionManagerCards>
      </div>
      <div v-else>
        <div class="loading" >
          <span class="sr-only">Loading...</span>
        </div>
      </div>
    </div>
    <div v-else>
      {{ this.message }}
    </div>
  </div>
</template>

<script>
import VersionManagerCards from "@/components/helper/VersionManagerCards";

export default {
  name: "VersionManager",
  components: {
    VersionManagerCards,
  },
  data() {
    return {
      clicked: false,
      message: null,
      loading: false,
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
    onItemClick(event, item) {
      this.clicked = true;
    },

    getVersionManagerCards(name) {
      for (let i = 0; i < this.getUser.plugins.length; i++) {
        let plugin = this.getUser.plugins[i];
        if (name === plugin.pluginName) {
          return plugin;
        }
      }
    },
  },
  //triggers fetchData when page is mounted
  mounted() {
    this.loading = true;

    this.$store.dispatch('user/fetchData', {
      option: "general",
      user: this.getUser,
    }).then(
        (onSuccess) => {
          if (onSuccess.data.success) {
            if (onSuccess.data.pluginData) {
              this.plugins = this.getUser.plugins;
              this.loading = false;
            }
          }
        },
        (onError) => {
          this.message = onError.toString() || onError.message;
          this.loading = false;
        }
    );

    this.$forceUpdate()
    console.log("loading:")
    console.log(this.loading)
  },
}
</script>

<style scoped lang="scss">
.row {
  width: 100%;
  padding-left: 10px;
  //border-radius: 10px;
}

.row:after {
  content: "";
  display: table;
  clear: both;
}

.col {
  //border-radius: 10px;
  //margin-top: 20px;
}

/* Add a card effect for articles */
.card {
  float: top;
  background-color: white;
  //padding: 20px;
  //margin-top: 20px;
}

/* If you want the content centered horizontally and vertically */
.centered {
  position: relative;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  text-align: center;
}

/* Absolute Center Spinner */
.loading {
  position: fixed;
  z-index: 999;
  overflow: show;
  margin: auto;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
  width: 50px;
  height: 50px;
}

/* Transparent Overlay */
.loading:before {
  content: '';
  display: block;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(255,255,255,0.5);
}

/* :not(:required) hides these rules from IE9 and below */
.loading:not(:required) {
  /* hide "loading..." text */
  font: 0/0 a;
  color: transparent;
  text-shadow: none;
  background-color: transparent;
  border: 0;
}

.loading:not(:required):after {
  content: '';
  display: block;
  font-size: 10px;
  width: 50px;
  height: 50px;
  margin-top: -0.5em;

  border: 15px solid rgb(0, 139, 139);
  border-radius: 100%;
  border-bottom-color: transparent;
  -webkit-animation: spinner 1s linear 0s infinite;
  animation: spinner 1s linear 0s infinite;


}

/* Animation */

@-webkit-keyframes spinner {
  0% {
    -webkit-transform: rotate(0deg);
    -moz-transform: rotate(0deg);
    -ms-transform: rotate(0deg);
    -o-transform: rotate(0deg);
    transform: rotate(0deg);
  }
  100% {
    -webkit-transform: rotate(360deg);
    -moz-transform: rotate(360deg);
    -ms-transform: rotate(360deg);
    -o-transform: rotate(360deg);
    transform: rotate(360deg);
  }
}
@-moz-keyframes spinner {
  0% {
    -webkit-transform: rotate(0deg);
    -moz-transform: rotate(0deg);
    -ms-transform: rotate(0deg);
    -o-transform: rotate(0deg);
    transform: rotate(0deg);
  }
  100% {
    -webkit-transform: rotate(360deg);
    -moz-transform: rotate(360deg);
    -ms-transform: rotate(360deg);
    -o-transform: rotate(360deg);
    transform: rotate(360deg);
  }
}
@-o-keyframes spinner {
  0% {
    -webkit-transform: rotate(0deg);
    -moz-transform: rotate(0deg);
    -ms-transform: rotate(0deg);
    -o-transform: rotate(0deg);
    transform: rotate(0deg);
  }
  100% {
    -webkit-transform: rotate(360deg);
    -moz-transform: rotate(360deg);
    -ms-transform: rotate(360deg);
    -o-transform: rotate(360deg);
    transform: rotate(360deg);
  }
}
@keyframes spinner {
  0% {
    -webkit-transform: rotate(0deg);
    -moz-transform: rotate(0deg);
    -ms-transform: rotate(0deg);
    -o-transform: rotate(0deg);
    transform: rotate(0deg);
  }
  100% {
    -webkit-transform: rotate(360deg);
    -moz-transform: rotate(360deg);
    -ms-transform: rotate(360deg);
    -o-transform: rotate(360deg);
    transform: rotate(360deg);
  }
}
</style>