# Lambda-backed custom resources

When you associate a Lambda function with a custom resource, the function is invoked
whenever the custom resource is created, updated, or deleted.

CloudFormation calls a Lambda API to invoke the function and to pass all the request data (such
as the request type and resource properties) to the function. The power and customizability
of Lambda functions in combination with CloudFormation enable a wide range of scenarios, such as
dynamically looking up AMI IDs during stack creation, or implementing and using utility
functions, such as string reversal functions.

For an introduction to custom resources and how they work, see [How custom resources work](template-custom-resources.md#how-custom-resources-work).

###### Topics

- [Walkthrough: Create a delay mechanism with a Lambda-backed custom resource](walkthrough-lambda-backed-custom-resources.md)

- [cfn-response module](cfn-lambda-function-code-cfnresponsemodule.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon SNS-backed custom resources

Walkthrough: Create a delay mechanism with a Lambda-backed custom
resource

All content copied from https://docs.aws.amazon.com/.
