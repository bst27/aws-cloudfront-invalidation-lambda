# About
This application can be used as (local) CLI or AWS Lambda application to invalidate
AWS Cloudfront distributions. When running in AWS Lambda you may setup
an AWS API Gateway to invoke the application via HTTPS (e.g. in a webhook).

The secret token has to be defined as an environment variable. It is checked
before an invalidation is created. It cannot be empty.

When running on AWS Lambda make sure you give permissions to create
invalidation requests to the corresponding execution role. When running
as a CLI application make sure  you have setup an AWS credential provider.

# Build
Execute `script.sh` to build the project. You will receive two files in the `dist` directory:
* **main** is the app executable
* **lambda-function.zip** can be used to upload to AWS Lambda  

# Config
Permission statement to allow creating AWS Cloudfront invalidations. Give this
permission to the AWS Lambda execution role:
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "AllowCloudfrontCreateInvalidation",
            "Effect": "Allow",
            "Action": "cloudfront:CreateInvalidation",
            "Resource": [
                "arn:aws:cloudfront::<myAwsAccountId>:distribution/<myDistributionId1>>",
                "arn:aws:cloudfront::<myAwsAccountId>:distribution/<myDistributionId2>>",
            ]
        }
    ]
}
```

# Examples
Curl request example:
```
curl \
--header "Content-Type: application/json" \
--request POST \
--data '{"SecretToken":"123","myDistributionId":"ABC","InvalidationPath":"/*"}' \
https://a1b2c3d4.execute-api.eu-central-1.amazonaws.com/default/cloudfront-invalidator
```

CLI invocation example:
```
./main cli <secretToken> <distributionId> "<invalidationPath>"
```