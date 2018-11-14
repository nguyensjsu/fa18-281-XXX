var express = require('express');
var axios = require('axios');
var path = require('path');
var XMLHttpRequest = require("xmlhttprequest").XMLHttpRequest;
var bodyParser = require('body-parser');
var csurf = require('csurf');
var session = require('express-session');
var cookieParser = require('cookie-parser');
var Client = require('node-rest-client').Client;
var secret = require('./config/secret');
var ejs = require('ejs');
var engine = require('ejs-mate');

var app = express();

app.use(session({
	secret: 'mysupersecret',
	resave: false,
	saveUninitialiazed: false,
	cookie: { maxAge: 180 * 60 * 1000} //in milliseconds
}));

app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());
app.use(cookieParser());
app.use(express.static(__dirname + '/public'));
app.engine('ejs', engine);
app.set('view engine', 'ejs');

var userloginServer = "http://localhost:5000/"

var isLoggedIn = false;

app.get('/signin', function(request, response) {
	/*if (request.session.userid) {
		return response.redirect("/");
	}*/
	response.render('user/login',  {login: isLoggedIn, cartQuantity: 0});
});

app.get('/signup', function(request, response) {
  response.render('user/signup', {login: isLoggedIn, cartQuantity: 0});
});

app.post('/signin', function(request, response) {

});

app.post('/signup', function(request, response) {
	console.log("Call is in POST signup");
	console.log(JSON.stringify(request.body));

	var xmlhttp = new XMLHttpRequest();
	xmlhttp.open("POST", userloginServer + "users");
	xmlhttp.setRequestHeader("Content-Type", "application/json");
	var temp_userId = request.body.name;
	var jsonToSend = {
		"Name": request.body.name,
		"Email":  request.body.email,
		"Password": request.body.password
	};

	xmlhttp.send(JSON.stringify(jsonToSend));

	xmlhttp.onreadystatechange = function() {
		if (this.readyState === 4 && this.status === 200) {
			request.session.userid = request.body.email;
			response.redirect("/")
		}
		else if (this.readyState === 4 && this.status !== 200) {
			console.log("Cannot post to user database");
			response.redirect("/");
		}
	}
});

app.get('/', function(request, response){
	if (request.session.userid) {
		isLoggedIn = true;
	}
	else {
		isLoggedIn = false;
	}

	response.render('user/login', {login:isLoggedIn, cartQuantity: 0});
});

app.listen(secret.port, function (err) {
    if (err) throw err;
    console.log('Server is listening on port ' + secret.port + '!');
});
