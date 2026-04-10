This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BCMDataExports::Export S3Destination

Describes the destination Amazon Simple Storage Service (Amazon S3) bucket name and object
keys of a data exports file.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Bucket" : String,
  "S3BucketOwner" : String,
  "S3OutputConfigurations" : S3OutputConfigurations,
  "S3Prefix" : String,
  "S3Region" : String
}

```

### YAML

```yaml

  S3Bucket: String
  S3BucketOwner: String
  S3OutputConfigurations:
    S3OutputConfigurations
  S3Prefix: String
  S3Region: String

```

## Properties

`S3Bucket`

The name of the Amazon S3 bucket used as the destination of a data export file.

_Required_: Yes

_Type_: String

_Pattern_: `^[\S\s]*$`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BucketOwner`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3OutputConfigurations`

The output configuration for the data export.

_Required_: Yes

_Type_: [S3OutputConfigurations](aws-properties-bcmdataexports-export-s3outputconfigurations.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Prefix`

The S3 path prefix you want prepended to the name of your data export.

_Required_: Yes

_Type_: String

_Pattern_: `^[\S\s]*$`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Region`

The S3 bucket Region.

_Required_: Yes

_Type_: String

_Pattern_: `^[\S\s]*$`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceTag

S3OutputConfigurations

All content copied from https://docs.aws.amazon.com/.
