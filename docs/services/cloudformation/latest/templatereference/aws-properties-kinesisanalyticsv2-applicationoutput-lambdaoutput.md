This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::ApplicationOutput LambdaOutput

When you configure a SQL-based Kinesis Data Analytics application's output,
identifies an Amazon Lambda function as the destination. You provide the function Amazon Resource
Name (ARN) of the Lambda function.

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

The Amazon Resource Name (ARN) of the destination Lambda function to write to.

###### Note

To specify an earlier version of the Lambda function than the latest, include the Lambda function version in the Lambda function ARN. For more information about Lambda ARNs, see [Example ARNs: Amazon Lambda](../../../../general/latest/gr/aws-arns-and-namespaces.md#arn-syntax-lambda)

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [LambdaOutput](../../../managed-flink/latest/apiv2/api-lambdaoutput.md) in the _Amazon Kinesis Data Analytics API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KinesisStreamsOutput

Output

All content copied from https://docs.aws.amazon.com/.
