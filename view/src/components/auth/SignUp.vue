<template>
	<div class="bg-img">
		<form @submit.prevent="handleSubmit" class="container">
			<h3>SignUp</h3>
			<input
					type="email"
			        placeholder="Email"
			        name="email"
			        v-validate="'required|email|max:50'"
			        v-model="user.email">
			<span
					v-show="submitted && errors.has('password')"
					class="alert-danger">
				{{errors.first('password')}}
			</span>
			<input
					type="password"
					placeholder="Password"
					ref="password"
					name="password"
					v-validate="'required|min:10|max:512'"
					v-model="user.password">
			<span
					v-show="submitted && errors.has('repeat')"
					class="alert-danger">
				{{errors.first('repeat')}}
			</span>
			<input
					type="password"
					placeholder="repeat Password"
					name="repeat"
					v-validate="'required|confirmed:password'"
					v-model="confirm">
			<button type="submit" class="btn">SignUp</button>
			<div
					v-if="message"
					class="alert-danger">
				{{message}}
			</div>
		</form>
	</div>
</template>

<script>
  import User from '../../model/user';

  export default {
	name: "SignUp",
	data () {
	  return {
	    password: "",
		confirm: "",
		user: new User("", "", "normal"),
	    submitted: false,
	    message: "",
	  }
	},
    computed: {
	  //check if user is loggedIn
	  loggedIn() {
	    return this.$store.state.auth.status.loggedIn
	  }
    },
    methods: {
	  //handle an user registration request
      handleSubmit: function() {
        this.submitted = true
        this.$validator.validate().then(isValid => {
          if (isValid) {
            this.$store.dispatch("auth/register", this.user).then(
              () => {
                this.$router.push("/login")
              },
              (onFailure) => {
                this.message = onFailure.response.data;
                this.submitted = false;
              },
            )
          }
        })
      },
    },
    //when user is logged in this component is unavailable
    mounted() {
      if (this.loggedIn){
        this.$router.push("/")
      }
    },
  }
</script>

<style scoped lang="scss">
	* {
		box-sizing: border-box;
	}

	.bg-img {
		/* The image used */
		background-image: url("../../assets/wisehubLogoV.png");

		/* Control the height of the image */
		min-height: 100vh;

		/* Center and scale the image nicely */
		background-position: center;
		background-repeat: no-repeat;
		background-size: cover;
		position: relative;
	}

	h3 {
		color: white;
		width: 100%;
		/*margin-top: 15px;*/
		border-radius: 10px;
		text-align: center;
		/*background: #F5FFFA;*/
	}

	/* Add styles to the form container */
	.container {
		position: absolute;
		right: 0;
		margin: 20px;
		margin-top: 30px;
		border-radius: 10px;
		max-width: 350px;
		padding: 16px;
		background-color: #002828;
	}

	/*!* Full-width input fields *!*/
	input[type=text], input[type=password], input[type=email] {
		width: 100%;
		padding: 15px;
		margin: 5px 0 22px 0;
		border: none;
		background: #f1f1f1;
	}

	input[type=text]:focus, input[type=password]:focus, input[type=email]:focus {
		background-color: #ddd;
		outline: none;
	}

	/* !*Set a style for the submit button*!*/
	.btn {
		background-color: #469898;
		color: white;
		padding: 16px 20px;
		cursor: pointer;
	}

	/* style inputs and link buttons */
	input,
	.btn {
		width: 100%;
		padding: 12px;
		border: none;
		border-radius: 4px;
		margin: 5px 0;
		opacity: 0.85;
		display: inline-block;
		font-size: 17px;
		line-height: 20px;
		text-decoration: none; /* remove underline from anchors */
	}

	input:hover,
	.btn:hover {
		opacity: 1;
	}

	/* style the submit button */
	input[type=submit] {
		background-color: #469898;
		color: white;
		cursor: pointer;
	}

	input[type=submit]:hover {
		background-color: #55afaf;
	}

</style>