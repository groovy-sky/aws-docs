This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::Table KinesisStreamSpecification

The Kinesis Data Streams configuration for the specified table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApproximateCreationDateTimePrecision" : String,
  "StreamArn" : String
}

```

### YAML

```yaml

  ApproximateCreationDateTimePrecision: String
  StreamArn: String

```

## Properties

`ApproximateCreationDateTimePrecision`

The precision for the time and date that the stream was created.

_Required_: No

_Type_: String

_Allowed values_: `MICROSECOND | MILLISECOND`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamArn`

The ARN for a specific Kinesis data stream.

Length Constraints: Minimum length of 37. Maximum length of 1024.

_Required_: Yes

_Type_: String

_Minimum_: `37`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Change Data Capture for Kinesis Data Streams with DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/kds.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

KeySchema

LocalSecondaryIndex
