This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application InputLambdaProcessor

An object that contains the Amazon Resource Name (ARN) of the Amazon Lambda function that is
used to preprocess records in the stream in a SQL-based Kinesis Data Analytics application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ResourceARN" : String
}

```

### YAML

```yaml

  ResourceARN: String

```

## Properties

`ResourceARN`

The ARN of the Amazon Lambda function that operates on records in the stream.

###### Note

To specify an earlier version of the Lambda function than the latest, include the
Lambda function version in the Lambda function ARN. For more information about
Lambda ARNs, see [Example\
ARNs: Amazon Lambda](../../../../general/latest/gr/aws-arns-and-namespaces.md#arn-syntax-lambda)

_Required_: Yes

_Type_: String

_Pattern_: `^arn:.*$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [InputLambdaProcessor](../../../managed-flink/latest/apiv2/api-inputlambdaprocessor.md) in the _Amazon Kinesis Data_
_Analytics API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Input

InputParallelism

All content copied from https://docs.aws.amazon.com/.
