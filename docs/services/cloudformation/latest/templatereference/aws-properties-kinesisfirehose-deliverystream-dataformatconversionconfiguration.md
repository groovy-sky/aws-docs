This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream DataFormatConversionConfiguration

Specifies that you want Kinesis Data Firehose to convert data from the JSON format to
the Parquet or ORC format before writing it to Amazon S3. Kinesis Data Firehose uses the
serializer and deserializer that you specify, in addition to the column information from
the AWS Glue table, to deserialize your input data from JSON and then
serialize it to the Parquet or ORC format. For more information, see [Kinesis\
Data Firehose Record Format Conversion](../../../firehose/latest/dev/record-format-conversion.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "InputFormatConfiguration" : InputFormatConfiguration,
  "OutputFormatConfiguration" : OutputFormatConfiguration,
  "SchemaConfiguration" : SchemaConfiguration
}

```

### YAML

```yaml

  Enabled: Boolean
  InputFormatConfiguration:
    InputFormatConfiguration
  OutputFormatConfiguration:
    OutputFormatConfiguration
  SchemaConfiguration:
    SchemaConfiguration

```

## Properties

`Enabled`

Defaults to `true`. Set it to `false` if you want to disable
format conversion while preserving the configuration details.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputFormatConfiguration`

Specifies the deserializer that you want Firehose to use to convert the
format of your data from JSON. This parameter is required if `Enabled` is set to
true.

_Required_: No

_Type_: [InputFormatConfiguration](aws-properties-kinesisfirehose-deliverystream-inputformatconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputFormatConfiguration`

Specifies the serializer that you want Firehose to use to convert the
format of your data to the Parquet or ORC format. This parameter is required if
`Enabled` is set to true.

_Required_: No

_Type_: [OutputFormatConfiguration](aws-properties-kinesisfirehose-deliverystream-outputformatconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SchemaConfiguration`

Specifies the AWS Glue Data Catalog table that contains the column
information. This parameter is required if `Enabled` is set to true.

_Required_: No

_Type_: [SchemaConfiguration](aws-properties-kinesisfirehose-deliverystream-schemaconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DatabaseTables

DeliveryStreamEncryptionConfigurationInput

All content copied from https://docs.aws.amazon.com/.
