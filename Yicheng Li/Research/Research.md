# AWS Lambda
    AWS Lambda is an event-driven, serverless computing platform. It belongs to a computing service in
    Amazon. The use of Lambda is to build smaller, on-demand applications that are responsive to events
    and new information. Benefits using Lambda are we don't need to worry about server config, pay only
    what you use. An example using AWS Lambda is to resize and update thumbnails.

    Question:
    need more time to investigate how it can apply to our project

# ElasticBeanStalk
   AWS Elastic Beanstalk is an easy-to-use service for deploying and scaling web applications and services developed with Java, .NET, PHP, Node.js, Python, Ruby, Go, and Docker on familiar servers such as Apache, Nginx, Passenger, and IIS.

   Advantage:
   Auto-scale,
   load balancing for requests,
   no additional charge, only pay for the AWS resources needed
   to store and run our application,
   have a log monitor for latency, CPU usage and ect.

   Use case:
   Deploy our nodeJS application for public access

   Steps to deploy nodeJS:
   1. Create web an app environment follow instructions under links https://docs.aws.amazon.com/quickstarts/latest/webapp/welcome.html?icmpid=docs_eb_console_new
   2. Create an instance to pull nodeJS code inside the instance.
   3. Under the folder you nodejs code is in, setup eb
      .eb init()
        .select the region for your elasticBeanStalk environment, and pick the web environment where the nodeJS application deploy on.
      .eb status()
      .eb deploy
    4. Go to AES web environment and click the link and direct to your app.
