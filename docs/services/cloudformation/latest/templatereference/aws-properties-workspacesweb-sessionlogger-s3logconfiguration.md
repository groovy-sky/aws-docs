This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesWeb::SessionLogger S3LogConfiguration

The S3 log configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bucket" : String,
  "BucketOwner" : String,
  "FolderStructure" : String,
  "KeyPrefix" : String,
  "LogFileFormat" : String
}

```

### YAML

```yaml

  Bucket: String
  BucketOwner: String
  FolderStructure: String
  KeyPrefix: String
  LogFileFormat: String

```

## Properties

`Bucket`

The S3 bucket name where logs are delivered.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9][\.\-a-z0-9]{1,61}[a-z0-9]$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketOwner`

The expected bucket owner of the target S3 bucket. The caller must have permissions to write to the target bucket.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FolderStructure`

The folder structure that defines the organizational structure for log files in S3.

_Required_: Yes

_Type_: String

_Allowed values_: `Flat | NestedByDate`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyPrefix`

The S3 path prefix that determines where log files are stored.

_Required_: No

_Type_: String

_Pattern_: `^[\d\w\-_/!().*']+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogFileFormat`

The format of the LogFile that is written to S3.

_Required_: Yes

_Type_: String

_Allowed values_: `JSONLines | Json`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
