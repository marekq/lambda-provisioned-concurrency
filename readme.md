lambda-provisioned-concurrency
==============================

This demo application helps to test the new Provisioned Concurrency (PC) feature of AWS Lambda. Once you deploy the stack, two Go Lambda functions will be created. The first Lambda will generate SQS messages with a random integer and put them on the included SQS queue. The second function receives messages in batches of one from the SQS queue and prints them. You can also invoke the backend Lambda function over an HTTP endpoint and generate load that way. The backend Lambda is configured with one PC which can autoscale up to a maximum amount you can set in the SAM template. 


Setup
-----

Prior to deploying the app, you need to do the following;

* Run 'bash deploy.sh' in order to deploy the stack. You will be asked to fill in the stacks name, AWS region and whether to create an SQS and HTTP endpoint. 


Contact
-------

In case you have any suggestions, questions or remarks, please raise an issue or reach out to @marekq.
