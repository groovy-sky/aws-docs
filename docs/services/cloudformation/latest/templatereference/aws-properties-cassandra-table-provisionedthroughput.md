This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cassandra::Table ProvisionedThroughput

The provisioned throughput for the table, which consists of
`ReadCapacityUnits` and `WriteCapacityUnits`.

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

The amount of read capacity that's provisioned for the table. For more information,
see [Read/write capacity\
mode](../../../keyspaces/latest/devguide/readwritecapacitymode.md) in the _Amazon Keyspaces Developer Guide_.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WriteCapacityUnits`

The amount of write capacity that's provisioned for the table. For more information,
see [Read/write capacity\
mode](../../../keyspaces/latest/devguide/readwritecapacitymode.md) in the _Amazon Keyspaces Developer Guide_.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionSpecification

ReplicaSpecification

All content copied from https://docs.aws.amazon.com/.
