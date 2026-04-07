This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::Table ProvisionedThroughput

Throughput for the specified table, which consists of values for
`ReadCapacityUnits` and `WriteCapacityUnits`. For more information
about the contents of a provisioned throughput structure, see [Amazon DynamoDB Table ProvisionedThroughput](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ProvisionedThroughput.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ReadCapacityUnits" : Integer,
  "WriteCapacityUnits" : Integer
}

```

### YAML

```yaml

  ReadCapacityUnits: Integer
  WriteCapacityUnits: Integer

```

## Properties

`ReadCapacityUnits`

The maximum number of strongly consistent reads consumed per second before DynamoDB
returns a `ThrottlingException`. For more information, see [Specifying\
Read and Write Requirements](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughput.html) in the _Amazon DynamoDB Developer_
_Guide_.

If read/write capacity mode is `PAY_PER_REQUEST` the value is set to
0.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WriteCapacityUnits`

The maximum number of writes consumed per second before DynamoDB returns a
`ThrottlingException`. For more information, see [Specifying\
Read and Write Requirements](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughput.html) in the _Amazon DynamoDB Developer_
_Guide_.

If read/write capacity mode is `PAY_PER_REQUEST` the value is set to
0.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Projection

ResourcePolicy
