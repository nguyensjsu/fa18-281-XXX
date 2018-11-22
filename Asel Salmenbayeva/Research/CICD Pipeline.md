# CI/CD Pipeline
Continuous Integration/Continuous Deployment pipeline is a good way to save time and effort on repetitive process of safe deployment on each code update. The pipeline generally automates building, testing and deploying the application and ensures reliability of the new release. Since we are required to deploy our project on AWS and Heroku, the use of pipelines provided by these services is reasonable.

## AWS Pipeline
The AWS CodePipeline service offers following stages:
* Source
* Build
* Staging
* Production
_Source_ is where it integrates with GitHub repository. Changes are built at the _Build_ stage. Then, code is deployed and tested at _Staging_ and will be deployed to public servers at _Production_ only after developer's manual approval.
There's a [guide](https://docs.aws.amazon.com/lambda/latest/dg/build-pipeline.html) to set up pipeline for AWS Lambda.

## Heroku Pipeline
The stages of Heroku differ only by their names:
* Development
* Review
* Staging
* Production
As soon as there's a change in code hosted on Github, Heroku creates _Review_ app to test. Then if everything's ok, it's automatically deployed to _Staging_ app for further testing. When the change is ready, developer <b>promotes</b> the staging application to _Production_, making it available to the app’s end users. Promoting is a function of manual approving on Heroku so the concept's basically the same as on AWS.
Also Heroku provides a clear [guide](https://devcenter.heroku.com/articles/pipelines) with some visual examples.