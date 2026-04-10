This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SNS::Topic LoggingConfig

The `LoggingConfig` property type specifies the `Delivery` status
logging configuration for an [`AWS::SNS::Topic`](../userguide/aws-resource-sns-topic.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FailureFeedbackRoleArn" : String,
  "Protocol" : String,
  "SuccessFeedbackRoleArn" : String,
  "SuccessFeedbackSampleRate" : String
}

```

### YAML

```yaml

  FailureFeedbackRoleArn: String
  Protocol: String
  SuccessFeedbackRoleArn: String
  SuccessFeedbackSampleRate: String

```

## Properties

`FailureFeedbackRoleArn`

The IAM role ARN to be used when logging failed message deliveries in Amazon
CloudWatch.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

Indicates one of the supported protocols for the Amazon SNS topic.

###### Note

At least one of the other three `LoggingConfig` properties is recommend
along with `Protocol`.

_Required_: Yes

_Type_: String

_Allowed values_: `http/s | sqs | lambda | firehose | application`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SuccessFeedbackRoleArn`

The IAM role ARN to be used when logging successful message deliveries in Amazon
CloudWatch.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SuccessFeedbackSampleRate`

The percentage of successful message deliveries to be logged in Amazon CloudWatch. Valid
percentage values range from 0 to 100.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SNS::Topic

Subscription

All content copied from https://docs.aws.amazon.com/.
