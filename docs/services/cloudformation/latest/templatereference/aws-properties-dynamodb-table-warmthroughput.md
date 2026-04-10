This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::Table WarmThroughput

Provides visibility into the number of read and write operations your table or
secondary index can instantaneously support. The settings can be modified using the
`UpdateTable` operation to meet the throughput requirements of an
upcoming peak event.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ReadUnitsPerSecond" : Integer,
  "WriteUnitsPerSecond" : Integer
}

```

### YAML

```yaml

  ReadUnitsPerSecond: Integer
  WriteUnitsPerSecond: Integer

```

## Properties

`ReadUnitsPerSecond`

Represents the number of read operations your base table can instantaneously
support.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WriteUnitsPerSecond`

Represents the number of write operations your base table can instantaneously
support.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TimeToLiveSpecification

Next

All content copied from https://docs.aws.amazon.com/.
