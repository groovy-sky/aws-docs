This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource MappingParameters

When you configure a SQL-based Kinesis Data Analytics application's input at the
time of creating or updating an application, provides additional mapping information specific
to the record format (such as JSON, CSV, or record fields delimited by some delimiter) on the
streaming source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CSVMappingParameters" : CSVMappingParameters,
  "JSONMappingParameters" : JSONMappingParameters
}

```

### YAML

```yaml

  CSVMappingParameters:
    CSVMappingParameters
  JSONMappingParameters:
    JSONMappingParameters

```

## Properties

`CSVMappingParameters`

Provides additional mapping information when the record format uses delimiters
(for example, CSV).

_Required_: No

_Type_: [CSVMappingParameters](aws-properties-kinesisanalyticsv2-applicationreferencedatasource-csvmappingparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JSONMappingParameters`

Provides additional mapping information when JSON is the record format on the streaming source.

_Required_: No

_Type_: [JSONMappingParameters](aws-properties-kinesisanalyticsv2-applicationreferencedatasource-jsonmappingparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [MappingParameters](../../../managed-flink/latest/apiv2/api-mappingparameters.md) in the _Amazon Kinesis Data_
_Analytics API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JSONMappingParameters

RecordColumn

All content copied from https://docs.aws.amazon.com/.
