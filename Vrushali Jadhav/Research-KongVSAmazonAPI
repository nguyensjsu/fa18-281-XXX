
Introduction:
Kong is an open-source API gateway and microservice management layer.


Key concepts:
1] API Object:
Wraps properties of any HTTP(s) endpoint that accomplishes a specific task or delivers some service. Configurations include HTTP methods, endpoint URIs, upstream URL which points to our API servers and will be used for proxying requests, maximum retires, rate limits, timeouts, etc.
Consumer Object – wraps properties of anyone using our API endpoints. It will be used for tracking, access control and more
Upstream Object – describes how incoming requests will be proxied or load balanced, represented by a virtual hostname
Target Object – represents the services are implemented and served, identified by a hostname (or an IP address) and a port. Note that targets of every upstream can only be added or disabled. A history of target changes is maintained by the upstream
Plugin Object – pluggable features to enrich functionalities of our application during the request and response lifecycle. For example, API authentication and rate limiting features can be added by enabling relevant plugins. Kong provides very powerful plugins in its plugins gallery
Admin API – RESTful API endpoints used to manage Kong configurations, endpoints, consumers, plugins, and so on


