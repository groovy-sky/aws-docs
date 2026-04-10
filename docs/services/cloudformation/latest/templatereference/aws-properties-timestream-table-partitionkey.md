This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Timestream::Table PartitionKey

An attribute used in partitioning data in a table. A dimension key partitions data using the values of the
dimension specified by the dimension-name as partition key, while a measure key partitions data using measure names
(values of the 'measure\_name' column).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnforcementInRecord" : String,
  "Name" : String,
  "Type" : String
}

```

### YAML

```yaml

  EnforcementInRecord: String
  Name: String
  Type: String

```

## Properties

`EnforcementInRecord`

The level of enforcement for the specification of a dimension key in ingested records. Options are REQUIRED
(dimension key must be specified) and OPTIONAL (dimension key does not have to be specified).

_Required_: No

_Type_: String

_Allowed values_: `REQUIRED | OPTIONAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the attribute used for a dimension key.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the partition key. Options are DIMENSION (dimension key) and MEASURE (measure key).

_Required_: Yes

_Type_: String

_Allowed values_: `DIMENSION | MEASURE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MagneticStoreWriteProperties

RetentionProperties

All content copied from https://docs.aws.amazon.com/.
