This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream ProcessorParameter

The `ProcessorParameter` property specifies a processor parameter in a data
processor for an Amazon Kinesis Data Firehose delivery stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ParameterName" : String,
  "ParameterValue" : String
}

```

### YAML

```yaml

  ParameterName: String
  ParameterValue: String

```

## Properties

`ParameterName`

The name of the parameter. Currently the following default values are supported: 3
for `NumberOfRetries` and 60 for the `BufferIntervalInSeconds`. The
`BufferSizeInMBs` ranges between 0.2 MB and up to 3MB. The default buffering
hint is 1MB for all destinations, except Splunk. For Splunk, the default buffering hint is
256 KB.

_Required_: Yes

_Type_: String

_Allowed values_: `LambdaArn | NumberOfRetries | MetadataExtractionQuery | JsonParsingEngine | RoleArn | BufferSizeInMBs | BufferIntervalInSeconds | SubRecordType | Delimiter | CompressionFormat | DataMessageExtraction`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterValue`

The parameter value.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!\s*$).+`

_Minimum_: `1`

_Maximum_: `5120`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Processor

RedshiftDestinationConfiguration

All content copied from https://docs.aws.amazon.com/.
