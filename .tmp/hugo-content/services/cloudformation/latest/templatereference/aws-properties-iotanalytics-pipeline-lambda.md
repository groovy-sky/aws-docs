This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Pipeline Lambda

An activity that runs a Lambda function to modify the message.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BatchSize" : Integer,
  "LambdaName" : String,
  "Name" : String,
  "Next" : String
}

```

### YAML

```yaml

  BatchSize: Integer
  LambdaName: String
  Name: String
  Next: String

```

## Properties

`BatchSize`

The number of messages passed to the Lambda function for processing.

The AWS Lambda function must be able to process all of these messages within
five minutes, which is the maximum timeout duration for Lambda functions.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaName`

The name of the Lambda function that is run on the message.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9_-]+`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the 'lambda' activity.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Next`

The next activity in the pipeline.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Filter

Math

All content copied from https://docs.aws.amazon.com/.
