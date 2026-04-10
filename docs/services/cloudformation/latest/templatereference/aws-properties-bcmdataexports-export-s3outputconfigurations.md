This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BCMDataExports::Export S3OutputConfigurations

The compression type, file format, and overwrite preference for the data export.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Compression" : String,
  "Format" : String,
  "OutputType" : String,
  "Overwrite" : String
}

```

### YAML

```yaml

  Compression: String
  Format: String
  OutputType: String
  Overwrite: String

```

## Properties

`Compression`

The compression type for the data export.

_Required_: Yes

_Type_: String

_Allowed values_: `GZIP | PARQUET`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Format`

The file format for the data export.

_Required_: Yes

_Type_: String

_Allowed values_: `TEXT_OR_CSV | PARQUET`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputType`

The output type for the data export.

_Required_: Yes

_Type_: String

_Allowed values_: `CUSTOM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Overwrite`

The rule to follow when generating a version of the data export file. You have the choice
to overwrite the previous version or to be delivered in addition to the previous versions.
Overwriting exports can save on Amazon S3 storage costs. Creating new export versions allows
you to track the changes in cost and usage data over time.

_Required_: Yes

_Type_: String

_Allowed values_: `CREATE_NEW_REPORT | OVERWRITE_REPORT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3Destination

Next

All content copied from https://docs.aws.amazon.com/.
