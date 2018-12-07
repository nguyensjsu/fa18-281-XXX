## Kong API vs Amazon API gateway:


1. Kong API Gateway:
Kong is a scalable, open source API Gateway. Kong runs in front of any RESTful API and is extended through Plugins, which provide extra functionalities beyond the core platform.

1.1 Kong Architecture:

Kong is on top of NGINX built using the openResty framework.

Kong has two types of users. They are:

Admin users.
Consumers(may be a user or an application which raises requests).
In Kong, everything is under the control of Admin. Adding consumers, APIs, Services, plugins are done by Admin only. Admin has access to perform CRUD operations.

1.2. Ports:

Kong makes use of 4 ports.

1.2.1. For consumers

8000 — for incoming HTTP traffic
8443- for incoming HTTPS traffic
1.2.2. For Admin

8001 — for incoming HTTP traffic
8444- for incoming HTTPS traffic
1.3. Kong Editions:

Kong supports two types of versions. They are

1.3.1. Community Edition:

Kong Community Edition provides all the functionalities like authentication, authorization, traffic control etc for developing and maintaining applications easily. But the disadvantage of this Edition is that it does not has an Admin GUI, everything like adding APIs, services, routes, plugins are done through CURL commands. Kong Community Edition is free to use. Kong does not provide any support for community edition users.

1.3.2. Enterprise Edition:

Kong Enterprise edition is a paid version of kong which provides various advanced features to users such as,

Admin GUI
Developer Portals
API Analytics
More scalable and secure when compared to community edition.
24/7 support from kong.
1.3.3. Plugins:

One of the core principals of Kong is its extensibility through plugins. Plugins allow you to easily add new features to your API or make your API easier to manage. Plugins are written in LUA language. We can use the available plugins or create our own plugins based on our requirement. Kong has many built in plugins.

Example: Rate-Limiting is a plugin which allows you to set rate limiting to your end api.

2. Amazon API Gateway:
Amazon API Gateway is an AWS service that enables developers to create, publish, maintain, monitor, and secure APIs at any scale. You can create APIs that access AWS or other web services, as well as data stored in the AWS cloud.

                                               Figure3: Amazon API Gateway Work Flow
3. Features Provided by Kong and Amazon API Gateway
Both gateways acts as a layer between the end user and upstream API or services.The common features provided by both Kong and Amazon API Gateway are:

3.1. Authentication and Authorization:

3.1.1. Kong Authentication:

kong allows the user to use his own authentication mechanism, or to use plugins provided by kong. Kong has various plugins for Authentication such as:

1. Basic Authentication:

Includes usage of a username and password for accessing Upstream APIs.

2. Key-Authentication(Token-based Authentication):

If this plugin is used for a service or an API, tan a unique id is created for that service and users who are having the secret key(password) can only be authorized to access that service or API.

3.OAUTH-2:

Validate access tokens sent by developers using a third-party OAuth 2.0 Authorization Server, by leveraging its Introspection Endpoint (RFC 7662). This plugin assumes that the Consumer already has an access token that will be validated against a third-party OAuth 2.0 server.

4.JWT:

Kong does not create JWT tokens, it only verifies the generated JWT tokens. But if you want to create a JWT token by kong, tan their is a plugin called JWT-CRAFTER which can be used to generate JWT Tokens. A consumer with the authorized token can only be able to access upstream service.

5.OpenId connect:

OpenID Connect plugin allows the integration with a 3rd party identity provider (IdP) or Kong Oauth 2.0 in a standardized way. This plugin can be used to implement Kong as a (proxying) Oauth 2.0 resource server (RS) and/or as an OpenID Connect relying party (RP) between the client and the upstream service.

6.LDAP:

Add LDAP Bind Authentication to a Route (or the deprecated API entity) with username and password protection.The plugin will check for valid credentials in the Proxy-Authorization  and authorization header.

7.HMAC: 

uses message Authentication protocol for providing the user an access to upstream APIs

3.1.1.1. Security Plugins in Kong:

Kong provides security to our upstream APIs and services using the following plugins:

ACL
CORS
Dynamic SSL
IP Restriction
BOT Detection
Pros of Kong in Authentication:

1. Verifying credentials of the user on obtaining a sign-in request, if credentials are valid, a token is generated (if JWT plugin is used) tan access to upstream service is provided. The generated JWT is stored in the cache for future reference.

2.Provides Role Based Access Control (only for enterprise edition)

3. Supports Groups concept, users can be grouped together and Kong allows admin to impose access limits, restrictions to those groups

4. Allows login using different modes like email/phone/username.

Cons of Kong in Authentication:

1. kong does not has any password policy.

2. Kong does not provide any mechanism for verifying user email Id.

3. Credentials for a consumer are set and maintained by Admin and consumer cannot change them.

4. No concept of multi-factor Authentication

5.No risk analysis for sign in operation

6. Kong functions in a better way if we integrate our own authentication server and pass the generated tokens to kong gateway for validating them.

3.1.2 Amazon Gateway Authentication

we can implement all the above-mentioned features in Amazon API Gateway by the use of Cognito AWS Service as an Authorizer.

1.Cognito:

Amazon Cognito provides authentication, authorization, and user management for your web and mobile apps. Your users can sign in directly with a username and password, or through a third party such as Facebook, Amazon, or Google.

Types of Authentication Mechanisms in Cognito:

User Pool
Social Identity Providers such as Amazon, Facebook etc.
SAML
Components of Cognito:

i. User Pool:

A user pool is a user directory in Amazon Cognito. With a user pool, your users can sign in to your web or mobile app through Amazon Cognito. Your users can also sign in through social identity providers like Facebook or Amazon, and through SAML identity providers.

User pools provide:

Sign-up and sign-in services.
A built-in, customizable web UI to sign in users.
Social sign-in with Facebook, Google, and Login with Amazon, as well as sign-in with SAML identity providers from your user pool.
User directory management and user profiles.
Security features such as multi-factor authentication (MFA), checks for compromised credentials, account takeover protection, and phone and email verification.
Customized workflows and user migration through AWS Lambda triggers.
ii. Identity Pool:

With an identity pool, your users can obtain temporary AWS credentials to access AWS services, such as Amazon S3 and DynamoDB. Identity pools support anonymous guest users, as well as the following identity providers that you can use to authenticate users for identity pools

To save user profile information, your identity pool needs to be integrated with a user pool.

                                    
        Figure4: Cognito Authentication Using User Pool for accessing backend Resources
                      
Figure5: Cognito Authentication Using User Pool, Identity Pool for accessing AWS services
Creating User Pool:

We can create a user pool in cognito console. While creating User pool, Cognito asks users to set following features to user pool:

1.Select how the end users need to sing-In either through username or through email/phone number

2. Select which attributes are to be displayed for sign up.

3. Password Policy

4. Select who can create users, either admin can only create user or users can sign up themselves.

5. select How quickly should user accounts created by administrators to expire if not used? (set maximum days)

6. Specify if you want Multi-Factor Authentication or not.

7. Email and phone number verification(if required)

Adding users to user pool:

1. Admin can add the user.

2. The user can signup themselves.

3. we can impost a CSV file of users into user pool.

The developer can use AWS Lambda triggers with Amazon Cognito user pools to customize workflows at various stages in the life cycle of a user account

Security Management in Cognito:

Cognito uses Multi-Factor Authentication (MFA) for advanced security management. It includes the following methods:

i. SMS Text Message MFA:

When a user signs in with MFA turned on, he or she first enters and submits his or her username and password. The client application will receive a response indicating where the authorization code was sent. The client application should indicate to the user where to look for the code (such as which phone number the code was sent to), provide a form for entering the code, and tan submit the code to complete the sign-in process. The destination is masked (only the last 4 digits of the phone number are displayed). If an application is using the Amazon Cognito hosted UI, it shows a page for the user to enter the MFA code. The SMS text message authorization code is valid for 3 minutes.

ii. TOTP Software Token MFA:

A user is challenged to complete authentication using a time-based one-time (TOTP) password after their username and password has been verified when TOTP software token MFA is enabled. If yourapplication is using the Amazon Cognito hosted UI to sign in users, the UI will show the second page for user to enter the TOTP password after they has submitted their username and password.

Advance Settings in Cognito:

i. Compromised Credentials Protection:

users sometimes may reuse the same credentials (me.e., username and password) for multiple websites and applications. If those reused credentials are stolen through website breaches and malware, they can become available on the internet, and criminals could try them at other locations. Cognito protections detect if a user’s credentials has been compromised elsewhere and block their use in Amazon Cognito User Pools. Users will be asked to choose another password if they try to use compromised credentials.

ii.Adaptive Authentication:

Adaptive authentication increases the security of user sign-in with Amazon Cognito User Pools without adding unnecessary friction for users. For each sign-in attempt, Amazon Cognito calculates a risk score for whether the attempt is from an attacker. This risk score is based on many factors, including whether the device is unrecognized, the user location is new, the IP address is new, etc. You can configure your user pool to block sign-ins or require second factors at different risk levels.

User History:

In the Users and Groups tab of the Amazon Cognito console, you can select a user to see that user’s recent sign-in events. Each sign-in event has an event ID, context data such as location, device details, and risk detection results associated with it.

Other Authorization and Access mechanisms in Amazon API Gateway:

me. User Authentication mechanism:

User can use his own authentication mechanism, but in this case, the gateway does not verify your request is valid or not, it simply sends the headers to backend service and that service should verify headers.

ii. Use the API Gateway Resource policy:

Amazon API Gateway resource policies are JSON policy documents that you attach to an API to control whether a specified principal (typically an IAM user or role) can invoke the API. You can use API Gateway resource policies to enable users from a different AWS account to securely access your API or to allow the API to be invoked only from specified source IP address ranges or CIDR blocks.

iii. Use IAM Permissions:

You control access to Amazon API Gateway with IAM permission by controlling access to the following two API Gateway component processes:

1. To create, deploy, and manage an API in API Gateway, you must grant the API developer permissions to perform the required actions supported by the API management component of API Gateway.

2. To call a deployed API or to refresh the API caching, you must grant the API caller permissions to perform required IAM actions supported by the API execution component of API Gateway.

iv.Using client-side SSL certificates:

You can use API Gateway to generate an SSL certificate and use its public key in the backend to verify that HTTP requests to your backend system are from API Gateway. This allows your HTTP backend to control and accept only requests originating from Amazon API Gateway, even if the backend is publicly accessible.

3.2. Traffic Control :

3.2.1 Kong API Gateway

Kong provides the feature to limit the access to the UpStream API by imposing a rate-limit plugin. This plugin allows the admin to specify the number of times the consumer can hit the app per second or per minute or per day etc. If a user tries to call the API even after the limit is exceeded tan “API limit exceeded ” error will be thrown. This rate-limit can be imposed to complete application, or per API or sometimes per customer( in the case of the malicious user who tries to access upstream API for fraudulent activity).

3.2.2 Amazon API Gateway

The feature of imposing rate-limits to the Upstream API is referred to as Throttling in Amazon API Gateway, We can specify no of requests can be made at per second basis. Each time a request is made, a throttling check happens, if rate-limit is not exceeded, backend API is called and the response is generated. Else If rate-limit is exceeded, error code 429 is thrown.

3.3. Serverless:

Both Kong and Amazon API Gateway support Server less Mechanism using AWS lamda , openwhisk(only kong supports).

3.4. Latency:

Both kong and Amazon API Gateway reduces the latency time by storing the responses in the cache. As a result, if the same request is made again, tan instead of making a database call again, the gateway finds the response of that request from the cache and provides it to end user.

3.5. Request/Response monitoring:

3.5.1. Kong Gateway:

Kong transforms the request into the format supported by the upstream API and tan forwards the request to the upstream API. The same mechanism is used for transforming the upstream API response to the format required by an end user.

3.5.2. Amazon API Gateway

Amazon API Gateway uses Apache velocity for transforming data which includes transforming payload or JSON of request to the format needed by backend API and also transforming the response payload or JSON of response generated by backend API to the format needed by an end user.

3.6. Logging:

3.6.1 Kong Gateway

Kong API Gateway provides various plugins for storing logs, some of them are

Syslog
UDP Log
TCP Log
FILE Log
3.6.2 Amazon API Gateway

Amazon API Gateway allows users to use 0.5GB to 237GB of cache. Every request and response details are stored in the cache. We can configure a cache key and TIME TO LIVE (TTL) . The cache key gets removed as soon as TTL expires.

3.7 Proxy Requests

Both Kong and Amazon API Gateway supports proxy requests.

4. Advantages of Kong API Gateway
Kong has several advantages compared to other API management platforms and tools. Community favorites include:

Open-Source: No black box. For enterprise or free usage, Kong is entirely open-source, always.
Based on Nginx: Kong is embedded in Nginx and benefits from its amazing proxying performances.
Customizable: Write plugins to cover all your architecture use-cases.
Data Ownership: Kong and its underlying datastore run on your servers.
Easy to scale: All Kong nodes are stateless. Spawning new nodes in your cluster is very easy.
Integrations: Many plugins integrate with popular third-party services in the microservices world.
5. Advantages of Amazon API Gateway:
Supports Versioning.
AWS Authentication using sig4.
Global distribution of APIs using integration of cloud front.
provides DDOS protection for malicious requests.
Apache Velocity is used for request/response mapping.
We can easily integrate AWS services with Amazon API Gateway.
Conclusion:

Both Kong and Amazon API Gateways are providing various features to simply our Application Development and Maintenance. We can select an API Gateway based on our requirements,

If their is a need to use Nginx and Lua in our Application, tan we can go with Kong API Gateway. We can add and remove plugins easily
If we need proper Autantication and Authorization mechanism, easy integration with all the aws services, proper versioning, easy maintenance, proper caching mechanism and if we need to use Cognito for Autantication tan we can go with Amazon aws API Gateway.
