<template>
	<div class="bg-img">

		<form @submit.prevent="handleSubmit" class="container">
			<h3>SignUp</h3>

<!--			<span-->
<!--					v-show="submitted && errors.has('name')"-->
<!--					class="alert-danger">-->
<!--				{{errors.first('name')}}-->
<!--			</span>-->
<!--			<input-->
<!--					type="text"-->
<!--					placeholder="Username"-->
<!--					name="name"-->
<!--					v-validate="'required|min:3|max:512'"-->
<!--					v-model="user.name">-->
<!--			<span-->
<!--					v-show="submitted && errors.has('email')"-->
<!--					class="alert-danger">-->
<!--				{{errors.first('email')}}-->
<!--			</span>-->
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
	  loggedIn() {
	    return this.$store.state.auth.status.loggedIn
	  }
    },
    methods: {
      handleSubmit: function() {
        this.submitted = true
        this.$validator.validate().then(isValid => {
          if (isValid) {
            console.log(this.user.password);
            this.$store.dispatch("auth/register", this.user).then(
              (onSuccess) => {
                console.log(onSuccess);
                this.$router.push({name: 'login'})
              },
              (onFailure) => {
                console.log(onFailure.response)
                this.message = onFailure.response.data;
                console.log(this.message)
                this.submitted = false;
              },
            )
          }
        })
      },
    },
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
		background-image: url("https://www.w3schools.com/howto/img_nature.jpg");

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