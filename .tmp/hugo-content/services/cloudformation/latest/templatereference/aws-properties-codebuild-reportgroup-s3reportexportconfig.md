This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::ReportGroup S3ReportExportConfig

Information about the S3 bucket where the raw data of a report are exported.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bucket" : String,
  "BucketOwner" : String,
  "EncryptionDisabled" : Boolean,
  "EncryptionKey" : String,
  "Packaging" : String,
  "Path" : String
}

```

### YAML

```yaml

  Bucket: String
  BucketOwner: String
  EncryptionDisabled: Boolean
  EncryptionKey: String
  Packaging: String
  Path: String

```

## Properties

`Bucket`

The name of the S3 bucket where the raw data of a report are exported.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketOwner`

The AWS account identifier of the owner of the Amazon S3 bucket. This allows report data to be exported to an Amazon S3 bucket
that is owned by an account other than the account running the build.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionDisabled`

A boolean value that specifies if the results of a report are encrypted.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionKey`

The encryption key for the report's encrypted raw data.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Packaging`

The type of build output artifact to create. Valid values include:

- `NONE`: CodeBuild creates the raw data in the output bucket. This
is the default if packaging is not specified.

- `ZIP`: CodeBuild creates a ZIP file with the raw data in the
output bucket.

_Required_: No

_Type_: String

_Allowed values_: `ZIP | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Path`

The path to the exported report's raw data results.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReportExportConfig

Tag

All content copied from https://docs.aws.amazon.com/.
