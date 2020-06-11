<template>
	<div class="bg-img">
		<div class="container">
			<form @submit.prevent="handleSubmit">
				<div class="row">
					<h3 style="text-align:center">Login</h3>
					<div class="vl">
						<span class="vl-innertext">or</span>
					</div>
<!--TODO-->
					<div class="col">
						<a href="#" class="github btn">
							<i class="fas fa-github fa-fw"></i> Login with GitHub
						</a>
						<a href="#" class="gitlab btn">
							<i class="fas fa-gitlab fa-fw"></i> Login with GitLab
						</a>
					</div>

					<div class="col">
						<div class="hide-md-lg">
							<p>Or sign in manually:</p>
						</div>

						<input type="text" name="username" placeholder="Username" required v-model="user.name">
						<input type="password" name="password" placeholder="Password" required v-model="user.password">
						<input type="submit" class="btn" value="Login">
					</div>

				</div>
			</form>
			<div class="bottom-container">
				<div class="row">
					<div class="col">
						<router-link to="/signup" style="color:white" class="btn">Sign up</router-link>
					</div>
					<div class="col">
						<router-link to="/forgot" style="color:white" class="btn">Forgot password?</router-link>
					</div>
				</div>
			</div>
		</div>

	</div>
</template>

<script>
	import User from '../model/user'
    export default {
        name: "LogIn",
        data () {
          return {
            user: new User("", ""),
          }
        },
        methods: {
          handleSubmit: function() {
			console.log(this.user);
          },
        }
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
		max-width: 767px;
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

	/* add appropriate colors to fb, twitter and google buttons */
	.github {
		background-color: #303030;
		color: white;
	}

	.gitlab {
		background-color: #dd4b39;
		color: white;
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

	/* Two-column layout */
	.col {
		float: left;
		width: 50%;
		margin: auto;
		padding: 0 50px;
		margin-top: 6px;
	}

	/* Clear floats after the columns */
	.row:after {
		content: "";
		display: table;
		clear: both;
	}

	/* vertical line */
	.vl {
		position: absolute;
		left: 50%;
		transform: translate(-50%, 30%);
		border: 2px solid #ddd;
		height: 155px;
	}

	/* text inside the vertical line */
	.vl-innertext {
		position: absolute;
		top: 50%;
		transform: translate(-50%, -50%);
		background-color: #f1f1f1;
		border: 1px solid #ccc;
		border-radius: 50%;
		padding: 8px 10px;
	}

	/* hide some text on medium and large screens */
	.hide-md-lg {
		display: none;
	}

	/* bottom container */
	.bottom-container {
		text-align: center;
		/*background-color: #469898;*/
		border-radius: 0px 0px 4px 4px;
	}

	/* Responsive layout - when the screen is less than 650px wide, make the two columns stack on top of each other instead of next to each other */
	@media screen and (max-width: 767px) {
		.col {
			width: 100%;
			margin-top: 0;
		}
		/* hide the vertical line */
		.vl {
			display: none;
		}
		/* show the hidden text on small screens */
		.hide-md-lg {
			display: block;
			text-align: center;
		}
	}
</style>