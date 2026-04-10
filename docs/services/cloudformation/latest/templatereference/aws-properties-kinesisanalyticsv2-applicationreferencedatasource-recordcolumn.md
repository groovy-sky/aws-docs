This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource RecordColumn

For a SQL-based Kinesis Data Analytics application, describes the mapping of each
data element in the streaming source to the corresponding column in the in-application
stream.

Also used to describe the format of the reference data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Mapping" : String,
  "Name" : String,
  "SqlType" : String
}

```

### YAML

```yaml

  Mapping: String
  Name: String
  SqlType: String

```

## Properties

`Mapping`

A reference to the data element in the streaming input or the reference data
source.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the column that is created in the in-application input stream or reference
table.

_Required_: Yes

_Type_: String

_Pattern_: `[^-\s<>&]*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SqlType`

The type of column created in the in-application input stream or reference table.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [RecordColumn](../../../managed-flink/latest/apiv2/api-recordcolumn.md) in the _Amazon Kinesis Data Analytics API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MappingParameters

RecordFormat

All content copied from https://docs.aws.amazon.com/.
