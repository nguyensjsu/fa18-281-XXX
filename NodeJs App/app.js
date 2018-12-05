var express = require('express');
var axios = require('axios');
var path = require('path');
var XMLHttpRequest = require("xmlhttprequest").XMLHttpRequest;
var bodyParser = require('body-parser');
var csurf = require('csurf');
var session = require('express-session');
var cookieParser = require('cookie-parser');
var secret = require('./config/secret');
var ejs = require('ejs');
var engine = require('ejs-mate');
var flash = require('express-flash');

var app = express();

app.use(session({
	secret: 'guessTheSecret',
	resave: false,
	saveUninitialiazed: false,
	cookie: { maxAge: 180 * 60 * 1000}
}));

app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());
app.use(cookieParser());
app.use(express.static(__dirname + '/public'));
app.engine('ejs', engine);
app.set('view engine', 'ejs');

var userloginServer = "http://teamprojectuserlogin-291444453.us-east-2.elb.amazonaws.com/";
var productCatalogueServer = "http://Productelb-1814563480.us-east-2.elb.amazonaws.com:3000/";

var isLoggedIn = false;
var cartQuantity = 0;
var cart = Object();

cart.items = [];
cart.total = 0;

app.get('/signin', function(request, response) {
	response.render('user/login',  {login: isLoggedIn, cartQuantity: 0});
});

app.get('/signup', function(request, response) {
  response.render('user/signup', {login: isLoggedIn, cartQuantity: 0});
});

app.post('/signin', function(request, response) {
	var emailID = request.body.email;
	var xmlhttp = new XMLHttpRequest();
	xmlhttp.open("GET", userloginServer+ "users/" +emailID);
	xmlhttp.setRequestHeader("Content-Type", "application/json");
	xmlhttp.send();

	xmlhttp.onreadystatechange = function()
	{
		if (this.readyState === 4 && this.status === 200) {
			var responseText = JSON.parse(this.responseText);

			if(responseText.Email == emailID) {
				isLoggedIn = true;
				response.redirect("/products");
			}
			else {
				response.redirect("/signin");
			}
		}
	}
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
		"Address": request.body.address,
		"Password": request.body.password
	};

	xmlhttp.send(JSON.stringify(jsonToSend));

	xmlhttp.onreadystatechange = function()
	{
		if (this.readyState === 4 && this.status === 200) {
			//request.session.userid = request.body.email;
			response.redirect("/signin")
		}
		else if (this.readyState === 4 && this.status !== 200) {
			console.log("Cannot post to user database");
			response.redirect("/signup");
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

	var xmlhttp = new XMLHttpRequest();
	xmlhttp.open("GET", productCatalogueServer+ "products");
	xmlhttp.setRequestHeader("Content-Type", "application/json");
	xmlhttp.send();
	xmlhttp.onreadystatechange = function() {
		if (this.readyState === 4 && this.status === 200) {
			var products_array = JSON.parse(this.responseText);
			console.log(products_array);
			response.render('./main/catalog', {products: products_array, login: isLoggedIn, cartQuantity: 0});
		}
	}
});

app.get('/products', function(request, response) {
	var xmlhttp1 = new XMLHttpRequest();
	xmlhttp1.open("GET", productCatalogueServer+ "products");
	xmlhttp1.setRequestHeader("Content-Type", "application/json");
	xmlhttp1.send();

	xmlhttp1.onreadystatechange = function()
	{
			if (this.readyState === 4 && this.status === 200) {
				var products_array = JSON.parse(this.responseText);
				response.render('./main/catalog', {products: products_array, login: isLoggedIn, cartQuantity: 0});
			}
	}
});

app.get('/products/:id', function(request, response) {
	var productId = request.params["id"];
	var xmlhttp1 = new XMLHttpRequest();
	xmlhttp1.open("GET", productCatalogueServer+ "products/" + productId);
	xmlhttp1.setRequestHeader("Content-Type", "application/json");
	xmlhttp1.send();

	xmlhttp1.onreadystatechange = function()
	{
		if (this.readyState === 4 && this.status === 200) {
			var product = JSON.parse(this.responseText);
			response.render('./main/product', {product: product, login: isLoggedIn, cartQuantity: 0});
		}
	}
});

app.get('/cart', function(request, response) {

	if(isLoggedIn) {
		response.render('./main/cart', {foundCart: cart, login: isLoggedIn, cartQuantity: 0});
	}
	else {
		response.redirect("/signin");
	}
});

app.get('/logout', function(request, response) {

});

app.post('/remove', function(request, response) {
	var productNameToRemove = request.body.item;

	var dummyCart = cart;
	for (var i=0; i<dummyCart.Products.length;i++) {
		var product = dummyCart.Products[i];

		if(product.ProductName === productNameToRemove) {
			delete dummyCart.Products[i];
		}
	}

	var procutArray = dummyCart.Products;

	var filtered = procutArray.filter(function (element) {
  	return element != null;
	});

	cart = dummyCart;
	cart.Products = filtered;
	cartQuantity = cart.Products.length;
	updateTheCart(cart, ()=> {

	});
	response.redirect('/cart');
});

function caluclateTotal(cart) {

	var totalAmount = 0;

	for (var i=0; i<cart.Products.length; i++) {
		var productPrice = parseInt(cart.Products[i].Price);
		var productQuantity = parseInt(cart.Products[i].Quantity);

		totalAmount+=productPrice*productQuantity;
	}

	cart.Total = totalAmount.toString();
}

function updateTheCart(cart, callback) {
	var xmlhttp = new XMLHttpRequest();
	xmlhttp.open("PUT", cartServer+ "carts");
	xmlhttp.setRequestHeader("Content-Type", "application/json");
	xmlhttp.send(JSON.stringify(cart));

	xmlhttp.onreadystatechange = function()
	{
			if (this.readyState === 4 && this.status === 200) {
				cart = JSON.parse(this.responseText);
				cartQuantity = cart.Products.length;
				callback();
			}
	}
}

app.get('/logout', function(request, response) {
	updateTheCart(cart, ()=>{
		isLoggedIn = false;
		cart = null;
		cartQuantity = 0;
		userID = null;
		response.redirect('/products');
	});
});

app.listen(secret.port, function (err) {
    if (err) throw err;
    console.log('Server is listening on port ' + secret.port + '!');
});
