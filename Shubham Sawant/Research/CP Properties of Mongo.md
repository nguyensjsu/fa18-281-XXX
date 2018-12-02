### CP Properties of MongoDB

* MongoDB is strongly consistent by default - if you do a write and then do a read, assuming the write was successful you will always be able to read the result of the write you just read. This is because MongoDB is a single-master system and all reads go to the primary by default. If you optionally enable reading from the secondaries then MongoDB becomes eventually consistent where it's possible to read out-of-date results.

* MongoDB also gets high-availability through automatic failover in replica sets: http://www.mongodb.org/display/DOCS/Replica+Sets
