# Проект OPYS
For going to OPYS app u must go to OPYS/cmd/miniOpys/main.go or main.exe 

You must have DB in postgres named postgres, password admin and tables
/*
create table if not exists users
(
	id uuid not null,
	username varchar(256) not null,
	normalizedusername varchar(256) not null,
	email varchar(256) not null,
	emailconfirmed boolean not null,
	normalizedemail text,
	passwordhash text not null,
	securitystamp text,
	concurrencystamp text,
	lockoutend date,
	lockoutenabled boolean,
	accessfailedcount integer
)
;

create unique index if not exists users_email_uindex
	on users (email)
;

create table if not exists userclaims
(
	id uuid not null,
	userid uuid not null,
	claimtype text,
	claimvalue text
)
;
*/

There are routes
    "localhost/signup/" wich gets JSON like
    	{
	 		"Email":"myemail@gmail.com",
	  		"Password":"123456789",
	  		"ConfirmPassword":"123456789",
	  		"Claims": [
	    	{
	      		"type":"firstname",
		      	"value":"Jane"
	    	},
	    	{
	      		"type":"lastname",
	      		"value":"Doe"
	    	},
	    	{
	      		"type":"birthdate",
	      		"value":"1999.12.31"
	    	}]
		}
    "localhost/signin/" gets JSON
        {
		"Email":"myemail@example.com",
	  	"Password":"123456789"
	    }