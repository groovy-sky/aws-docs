This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application InputParallelism

For a SQL-based Kinesis Data Analytics application, describes the number of
in-application streams to create for a given streaming source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Count" : Integer
}

```

### YAML

```yaml

  Count: Integer

```

## Properties

`Count`

The number of in-application streams to create.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [InputParallelism](../../../managed-flink/latest/apiv2/api-inputparallelism.md) in the _Amazon Kinesis Data Analytics_
_API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InputLambdaProcessor

InputProcessingConfiguration

All content copied from https://docs.aws.amazon.com/.
