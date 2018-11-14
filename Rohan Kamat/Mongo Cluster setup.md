# Mongo CLuster Setup for Cart API




### Launch Instance
1. AMI:             Ubuntu Server 16.04 LTS (HVM)
2. Instance Type:   t2.micro
3. VPC:             cmpe281
4. Network:         public subnet
5. Auto Public IP:  no
6. Security Group:  mongodb-cluster 
7. SG Open Ports:   22, 27017
8. Key Pair:        cmpe281-us-west-1
 

###Elastic IP associated to each instance 
```sh
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 0C49F3730359A14518585931BC711F9BA15703C6
echo "deb [ arch=amd64 ] http://repo.mongodb.org/apt/ubuntu trusty/mongodb-org/3.4 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-3.4.list
sudo apt-get update
sudo apt-get install -y mongodb-org
```
###Configure Key
```sh
openssl rand -base64 741 > keyFile
sudo mkdir -p /opt/mongodb
sudo cp keyFile /opt/mongodb
sudo chown mongodb:mongodb /opt/mongodb/keyFile
sudo chmod 0600 /opt/mongodb/keyFile
```
### Configure Mongo Cluster
```sh
sudo vi /etc/mongod.conf
```
1.  remove or comment out bindIp: 127.0.0.1
replace with bindIp: 0.0.0.0 (binds on all ips) 
network interfaces
net:
    port: 27017
    bindIp: 0.0.0.0

2. Uncomment security section & add key file

security:
    keyFile: /opt/mongodb/keyFile

3. Uncomment Replication section. Name Replica Set = cmpe281

replication:
   replSetName: cmpe281     

4. Create mongod.service
```sh
sudo vi /etc/systemd/system/mongod.service
```
    [Unit]
        Description=High-performance, schema-free document-oriented database
        After=network.target

    [Service]
        User=mongodb
        ExecStart=/usr/bin/mongod --quiet --config /etc/mongod.conf

    [Install]
        WantedBy=multi-user.target
```sh
sudo systemctl enable mongod.service
sudo service mongod restart
sudo service mongod status
```
### Create AMI
Launch Similar 2 Instances
### On Primary Instance 
```sh
mongo
 rs.initiate( {
       _id : "cmpe281",
       members: [
          { _id: 0, host: "primary:27017" },
          { _id: 1, host: "secondary1:27017" },
          { _id: 2, host: "secondary2:27017" }
       ]
    })
rs.status()

```
### Create Admin Account
### Use Robo Mongo to configure Database
