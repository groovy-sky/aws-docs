This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource JSONMappingParameters

For a SQL-based Kinesis Data Analytics application, provides additional mapping
information when JSON is the record format on the streaming source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RecordRowPath" : String
}

```

### YAML

```yaml

  RecordRowPath: String

```

## Properties

`RecordRowPath`

The path to the top-level parent that contains the records.

_Required_: Yes

_Type_: String

_Pattern_: `^(?=^\$)(?=^\S+$).*$`

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [JSONMappingParameters](../../../managed-flink/latest/apiv2/api-jsonmappingparameters.md) in the _Amazon Kinesis Data_
_Analytics API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CSVMappingParameters

MappingParameters

All content copied from https://docs.aws.amazon.com/.
