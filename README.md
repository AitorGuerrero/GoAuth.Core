# GoAuth
An application to manage authentication of users.

Made with GoLang, and intended to be implemented in a REST API, and Rabbit.

Target specs:
* A user is only composed by an identifier string and a passkey string. This application is not intended to store users 
personal information.
* The manager is the entity who wants to control the user authentication (a web server, a mobile application, a 
company). 
* Namespaces: Each manager will have a namespace, and could have many sub namespaces, granting access to the users
depending on the namespace.
* A manager would be able to create and get information of users only in his namespace.
* A manager will have a secret token to operate with the application. This token could be restored anytime by the 
manager or the application.
* A user could login without the manager, indicating the namespace to the application, and then tell the manager 
the sessions token. The manager should check with the application if the session is correct.