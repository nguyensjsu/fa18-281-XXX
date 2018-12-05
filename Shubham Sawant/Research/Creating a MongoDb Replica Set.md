### Task: Creating a MongoDb Image so that it can be used later for creating instances.

#### Subtask 1: Launch an Ubuntu type instance in public subnet. The configuration has given below:

1. Launch 'Ubuntu Server 16.04 LTS (HVM)' type instance.
2. In security group, open ports 22 and 27017.
3. Create this instance in public subnet.

#### Subtask 2: SSH into the primary node.

1. SSH into the created node using keyValue pair.

#### Subtask 3: Install MongoDb.

1. sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 9DA31620334BD75D9DCB49F368818C72E52529D4
2. echo "deb [ arch=amd64 ] https://repo.mongodb.org/apt/ubuntu trusty/mongodb-org/4.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-4.0.list
3. sudo apt-get update
4. sudo apt-get install -y mongodb-org

#### Subtask 4: Installation Verification.
  * sudo systemctl enable mongod
  * sudo systemctl start mongod
  * sudo systemctl stop mongod
  * sudo systemctl restart mongod
  * mongod --version

#### Subtask 5: Create a key file that will be used to secure authentication between the members of your replica set.

1. Issue this command to generate your key file:
  * openssl rand -base64 741 > keyFile
2. Create the /opt/mongo directory to store your key file:
  * sudo mkdir -p /opt/mongodb
3. Move the keyfile to /opt/mongo, and assign it the correct permissions:
  * sudo cp keyFile /opt/mongodb
4. Update the ownership of the file.
  * sudo chown mongodb:mongodb /opt/mongodb/keyFile
  * sudo chmod 0600 /opt/mongodb/keyFile

#### Subtask 6: Config mongod.conf

* sudo nano /etc/mongod.conf

1.  remove or comment out bindIp: 127.0.0.1
    replace with bindIp: 0.0.0.0 (binds on all ips)

    (network interfaces)
    net:
        port: 27017
        bindIp: 0.0.0.0

2. Uncomment security section & add key file

    security:
      keyFile: /opt/mongodb/keyFile

3. Uncomment Replication section. Name Replica Set = cmpe281

    replication:
        replSetName: cmpe281

#### Subtask 7: Create mongod.service

  * sudo nano /etc/systemd/system/mongod.service

    [Unit]
        Description=High-performance, schema-free document-oriented database
        After=network.target

    [Service]
        User=mongodb
        ExecStart=/usr/bin/mongod --quiet --config /etc/mongod.conf

    [Install]
        WantedBy=multi-user.target

5. Enable Mongo Service

    * sudo systemctl enable mongod.service

6. Restart MongoDB to apply our changes

    * sudo service mongod restart
    * sudo service mongod status

#### Subtask 8: Create admin account.
    1. Create an admin user to access the database.

      mongo

    2. Select admin database.

      use admin

    3. Create admin account.

      db.createUser( {
          user: "admin",
          pwd: "*****",
          roles: [{ role: "root", db: "admin" }]
      });
      db.createUser( {
          user: "admin",
          pwd: "cmpe281",
          roles: [{ role: "root", db: "admin" }]
      });

![create_admin_user](https://github.com/nguyensjsu/cmpe281-ShubhamSawantSJSU/blob/master/Test%20Cases%20%26%20Conclusions/MongoDB%20ScreenShots/create_admin_user.png)

#### Subtask 10: Create an image of this instance so that we can use this to create other 5 instances.

* Right click on the instance and say create image under image. This is shown as below.

![create_image](https://github.com/nguyensjsu/cmpe281-ShubhamSawantSJSU/blob/master/Test%20Cases%20%26%20Conclusions/MongoDB%20ScreenShots/create_image.png)

### Task: VPC peering between 2 VPCs created in 2 different regions.

#### Subtask 1: Creating a different VPC in North California Region.

* Create a new VPC in north california region with one public and one private subnet and setup the key-pair. So now we have 2 different VPCs in 2 regions (Oregon and North California).

#### Subtask 2: Connecting 2 VPC using peering connections.

* Go to 'Peering Connections' under VPC dashboard.

![peering_connections](https://github.com/nguyensjsu/cmpe281-ShubhamSawantSJSU/blob/master/Test%20Cases%20%26%20Conclusions/MongoDB%20ScreenShots/peering_connections.png)

* Create a new peering connection. Give a name tag to the VPC, add current VPC in requester and the target VPC in accepter.

![create_peering_connection](https://github.com/nguyensjsu/cmpe281-ShubhamSawantSJSU/blob/master/Test%20Cases%20%26%20Conclusions/MongoDB%20ScreenShots/create_peering_connection.png)

* Now it is the time to accept the peering request. Go to peering connection of accepter, and accept the request.

![accept_peering_request](https://github.com/nguyensjsu/cmpe281-ShubhamSawantSJSU/blob/master/Test%20Cases%20%26%20Conclusions/MongoDB%20ScreenShots/accept_peering_request.png)

* My VPCs are connected now. The instances created in these VPCs can communicate with each other.

* Last step is to modify the route table of the private subnets so that they can forward the requests to each other. I have added screenshot for that:

![add_route_in_route_table](https://github.com/nguyensjsu/cmpe281-ShubhamSawantSJSU/blob/master/Test%20Cases%20%26%20Conclusions/MongoDB%20ScreenShots/add_route_in_route_table.png)

### Task: Launch 2 EC2 instances in first VPC (North California) and 3 EC2 instances in second VPC (Oregon).

#### Subtask 1: Launch 2 EC2 instances using already created image (AMI), as shown in the screenshot below.

![launch_using_image](https://github.com/nguyensjsu/cmpe281-ShubhamSawantSJSU/blob/master/Test%20Cases%20%26%20Conclusions/MongoDB%20ScreenShots/launch_using_image.png)

![2_ec2_instances](https://github.com/nguyensjsu/cmpe281-ShubhamSawantSJSU/blob/master/Test%20Cases%20%26%20Conclusions/MongoDB%20ScreenShots/2_ec2_instances.png)

![3_ec2_instances](https://github.com/nguyensjsu/cmpe281-ShubhamSawantSJSU/blob/master/Test%20Cases%20%26%20Conclusions/MongoDB%20ScreenShots/3_ec2_instances.png)

### Task: Create a Jumpbox and ssh into the primary node.

* Create a normal EC2 instance in public instance and name it as Jumpbox.

![Jumpbox](https://github.com/nguyensjsu/cmpe281-ShubhamSawantSJSU/blob/master/Test%20Cases%20%26%20Conclusions/MongoDB%20ScreenShots/Jumpbox.png)

* SSH into Jumpbox and from their ssh into private instance. The command for that is given below:

    1. sudo ssh -i northCaliforniaKeyPair.pem ec2-user@10.0.3.175

* Move the pem file to server using WinScp.

![Move_the_pem_file](https://github.com/nguyensjsu/cmpe281-ShubhamSawantSJSU/blob/master/Test%20Cases%20%26%20Conclusions/MongoDB%20ScreenShots/Move_the_pem_file.png)

* Now SSH into the private instance.

![ssh_into_primary](https://github.com/nguyensjsu/cmpe281-ShubhamSawantSJSU/blob/master/Test%20Cases%20%26%20Conclusions/MongoDB%20ScreenShots/ssh_into_primary.png)

### Task: Initiate the replicaset and add nodes to the replicaset.

* rs.initiate()

![rs.initiate()](https://github.com/nguyensjsu/cmpe281-ShubhamSawantSJSU/blob/master/Test%20Cases%20%26%20Conclusions/MongoDB%20ScreenShots/rs.initiate().png)

* Add other nodes using rs.add() command.

![rs.add()](https://github.com/nguyensjsu/cmpe281-ShubhamSawantSJSU/blob/master/Test%20Cases%20%26%20Conclusions/MongoDB%20ScreenShots/rs.add().png)

* Check status of the cluster using rs.status() command.

![rs.status()](https://github.com/nguyensjsu/cmpe281-ShubhamSawantSJSU/blob/master/Test%20Cases%20%26%20Conclusions/MongoDB%20ScreenShots/rs.status().png)

* Now our replicaSet is ready and we can proceed with the testing part.
