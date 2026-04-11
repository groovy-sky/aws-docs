This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ReceiptRule LambdaAction

When included in a receipt rule, this action calls an AWS Lambda function and,
optionally, publishes a notification to Amazon Simple Notification Service (Amazon SNS).

To enable Amazon SES to call your AWS Lambda function or to publish to an Amazon SNS topic of
another account, Amazon SES must have permission to access those resources. For information
about giving permissions, see the [Amazon SES Developer\
Guide](../../../ses/latest/dg/receiving-email-permissions.md).

For information about using AWS Lambda actions in receipt rules, see the [Amazon SES\
Developer Guide](../../../ses/latest/dg/receiving-email-action-lambda.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FunctionArn" : String,
  "InvocationType" : String,
  "TopicArn" : String
}

```

### YAML

```yaml

  FunctionArn: String
  InvocationType: String
  TopicArn: String

```

## Properties

`FunctionArn`

The Amazon Resource Name (ARN) of the AWS Lambda function. An example of an AWS Lambda
function ARN is `arn:aws:lambda:us-west-2:account-id:function:MyFunction`.
For more information about AWS Lambda, see the [AWS Lambda Developer Guide](../../../lambda/latest/dg/welcome.md).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InvocationType`

The invocation type of the AWS Lambda function. An invocation type of
`RequestResponse` means that the execution of the function immediately
results in a response, and a value of `Event` means that the function is
invoked asynchronously. The default value is `Event`. For information about
AWS Lambda invocation types, see the [AWS Lambda Developer Guide](../../../lambda/latest/dg/api-invoke.md).

###### Important

There is a 30-second timeout on `RequestResponse` invocations. You
should use `Event` invocation in most cases. Use
`RequestResponse` only to make a mail flow decision, such as whether
to stop the receipt rule or the receipt rule set.

_Required_: No

_Type_: String

_Allowed values_: `Event | RequestResponse`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopicArn`

The Amazon Resource Name (ARN) of the Amazon SNS topic to notify when the Lambda action is
executed. You can find the ARN of a topic by using the [ListTopics](../../../sns/latest/api/api-listtopics.md) operation in
Amazon SNS.

For more information about Amazon SNS topics, see the [Amazon SNS Developer Guide](../../../sns/latest/dg/createtopic.md).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectAction

Rule

All content copied from https://docs.aws.amazon.com/.
