export default class User {
  constructor(password, email, role){
    this.name = '';
	this.password = password;
	this.email = email;
	this.role = role;
	this.courses = [];
  };
}