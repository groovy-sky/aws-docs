# Using AWS AppConfig Agent with AWS Lambda

An AWS Lambda extension is a companion process that augments the capabilities of a Lambda
function. An extension can start before a function is invoked, run in parallel with a
function, and continue to run after a function invocation is processed. In essence, a Lambda
extension is like a client that runs in parallel to a Lambda invocation. This parallel client
can interface with your function at any point during its lifecycle.

If you use AWS AppConfig feature flags or other dynamic configuration data in a Lambda function,
then we recommend that you add the AWS AppConfig Agent Lambda extension as a layer to your Lambda
function. This makes calling feature flags simpler, and the extension itself includes best
practices that simplify using AWS AppConfig while reducing costs. Reduced costs result from fewer
API calls to the AWS AppConfig service and shorter Lambda function processing times. For more
information about Lambda extensions, see [Lambda extensions](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-extensions-api.html) in the
_AWS Lambda Developer Guide_.

###### Note

AWS AppConfig is a capability of AWS Systems Manager. AWS AppConfig [pricing](https://aws.amazon.com/systems-manager/pricing) is based on the number of
times a configuration is called and received. Your costs increase if your Lambda performs
multiple cold starts and retrieves new configuration data frequently.

###### Topics

- [Understanding how the AWS AppConfig Agent Lambda extension works](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-integration-lambda-extensions-how-it-works.html)

- [Adding the AWS AppConfig Agent Lambda extension](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-integration-lambda-extensions-add.html)

- [Configuring the AWS AppConfig Agent Lambda extension](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-integration-lambda-extensions-config.html)

- [Understanding available versions of the AWS AppConfig Agent Lambda extension](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-integration-lambda-extensions-versions.html)

- [Using AWS AppConfig Agent with Amazon EC2 and on-premises machines](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-integration-ec2.html)

- [Using AWS AppConfig Agent with Amazon ECS and Amazon EKS](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-integration-containers-agent.html)

- [Retrieving basic and multi-variant feature flags](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-integration-retrieving-feature-flags.html)

- [Using a manifest to enable additional retrieval features](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-agent-how-to-use-additional-features.html)

- [Generating a client using the OpenAPI specification](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-integration-OpenAPI.html)

- [Working with AWS AppConfig Agent local development mode](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-agent-how-to-use-local-development.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

How to use AWS AppConfig Agent to retrieve configuration data

Understanding how the AWS AppConfig Agent Lambda extension works
