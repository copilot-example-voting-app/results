# results
Frontend service to visualize the results of the votes.  
![results](https://user-images.githubusercontent.com/879348/95273545-3789e580-07f8-11eb-82e6-e38d5ac4000d.png)

The "results" service illustrates how you can have multiple services share the same load balancer.
The ["vote"](https://github.com/copilot-example-voting-app/vote) service listens on "/" whereas 
"results" listens on "/results". AWS Copilot handles creating multiple listener rules for you 
and assigning them a priority.

## How to create this service?
1. Install the AWS Copilot CLI [https://aws.github.io/copilot-cli/](https://aws.github.io/copilot-cli/)
2. Run
   ```bash
   $ copilot init
   ```
3. Enter "voting-app" for the name of your application.
4. Select "Load Balanced Web Service" for the service type.
5. Enter "results" for the name of the service.
6. Say "Y" to deploying to a "test" environment 🚀

Once deployed, your service will be accessible at an HTTP endpoint provided by the CLI like: http://votin-publi-anelun2kxbrl-XXXXXXX.YYYYY.elb.amazonaws.com/results

## What does it do?
AWS Copilot uses AWS CloudFormation under the hood to provision your infrastructure resources.
You should be able to see a `voting-app-test-results` stack that yours ECS service along with all the peripheral resources
needed for logging, service discovery, and more...


## How does it work?
Copilot stores the infrastructure-as-code for your service under the `copilot/` directory.
```
copilot
└── results
    └── manifest.yml
```
The `manifest.yml` file under `results/` holds the common configuration for a "load balanced web service" pattern.
The difference between [the manifest in the "vote" service](https://github.com/copilot-example-voting-app/vote/blob/main/copilot/vote/manifest.yml) and 
"results" is the `http.path` field.

## Deleting the service
If you'd like to delete only the service from the "voting-app" application. 
```bash
$ copilot svc delete
```
If you'd like to delete the entire application including other services and deployment environments:
```bash
$ copilot app delete
```