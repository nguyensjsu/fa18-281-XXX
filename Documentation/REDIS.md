# REDIS BASIC COMMANDS
```
SET KEY “VALUE”
GET KEY>>VALUE	
INCR KEY>>VALUE+1
DEL KEY >>deletes the key/value
SETNX(SETS only if key does not exists)
EXPIRE KEY <TIME in seconds>
TTL KEY>> Time left(-2 means expired,-1 means no expiry )
Keys can have a list:
RPUSH  NAME >>adds at end of list
LPUSH NAME>> Adds at beginning of list
LRANGE KEY <param,param> 
LLEN KEY>>Length of list
LPOP KEY>> POPS first element
Sets
SADD KEY “VALUE”
SREM KEY “VALUE”
SISMEMBER KEY “VALUE”
1>>exits
0>>non
SUNION KEY KEY …
HASH
HSET KEY {KEY:VALUE}
HGET KEY {KEY}>>VALUE
```
______________________________________________________________________________
# REDIS WITH GOLANG
Basic connection
```golang

go get -u github.com/go-redis/redis
func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
```
______________________________________________________________________________
# REFERENCES
*[redis with goLang]

*[working with redis]

*[redis github code example]

*[redis official documentation]

*[goLang documentaion]


[redis with goLang]:https://golangme.com/blog/how-to-use-redis-with-golang/

[working with redis]:https://www.alexedwards.net/blog/working-with-redis

[redis github code example]:https://github.com/go-redis/redis

[redis official documentation]:http://try.redis.io/

[goLang documentaion]:https://godoc.org/gopkg.in/redis.v3






 












 







  
  
   
