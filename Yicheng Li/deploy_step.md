# Steps to deploy nodeJS:
## Create web an app environment follow instructions under links         https://docs.aws.amazon.com/quickstarts/latest/webapp/welcome.html?icmpid=docs_eb_console_new
    Define name, description, url page name; pick nodejs as platform.
## Launch Ubuntu Service 16.04 LTS an instance for pull nodeJS code from github.
    1. AMI:             Ubuntu Server 16.04 LTS (HVM)
    2. Instance Type:   t2.micro
    3. VPC:             ohio
    4. Network:         public subnet
    5. Auto Public IP:  yes
    6. Security Group:  default
    7. SG Open Ports:   22
    8. Key Pair:        cmpe281-ohio

## SSH into the instance
    ssh -i <key>.pem ubuntu@<public ip>

## Install git and nodeJS command line tools
    sudo apt update
    sudo apt install python-pip
    pip install --upgrade --user awsebcli
    sudo apt install nodejs

## Deploy using eb
    eb init()
    select region your web environment locate
    eb status()
    eb deploy

## Go to AES web environment and click the link and direct to your app.
