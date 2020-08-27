export default class User {
  constructor(password, email, role){
	this.password = password;
	this.email = email;
	this.role = role;
	this.plugins = [];
  };
}