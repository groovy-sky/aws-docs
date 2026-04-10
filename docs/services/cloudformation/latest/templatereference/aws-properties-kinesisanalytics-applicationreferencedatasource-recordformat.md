This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalytics::ApplicationReferenceDataSource RecordFormat

Describes the record format and relevant mapping information that should be applied
to schematize the records on the stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MappingParameters" : MappingParameters,
  "RecordFormatType" : String
}

```

### YAML

```yaml

  MappingParameters:
    MappingParameters
  RecordFormatType: String

```

## Properties

`MappingParameters`

When configuring application input at the time of creating or updating an application,
provides additional mapping information specific to the record format (such as JSON,
CSV, or record fields delimited by some delimiter) on the streaming source.

_Required_: No

_Type_: [MappingParameters](aws-properties-kinesisanalytics-applicationreferencedatasource-mappingparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordFormatType`

The type of record format.

_Required_: Yes

_Type_: String

_Allowed values_: `JSON | CSV`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RecordColumn

ReferenceDataSource

All content copied from https://docs.aws.amazon.com/.
