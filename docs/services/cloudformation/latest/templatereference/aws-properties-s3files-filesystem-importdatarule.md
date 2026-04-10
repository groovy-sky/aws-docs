This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Files::FileSystem ImportDataRule

Specifies a rule that controls how data is imported from S3 into the file
system.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Prefix" : String,
  "SizeLessThan" : Integer,
  "Trigger" : String
}

```

### YAML

```yaml

  Prefix: String
  SizeLessThan: Integer
  Trigger: String

```

## Properties

`Prefix`

The S3 key prefix that scopes this import rule. Only objects with keys beginning with
this prefix are subject to the rule.

_Required_: Yes

_Type_: String

_Pattern_: `^(|.*/)$`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SizeLessThan`

The upper size limit in bytes for this import rule. Only objects with a size strictly
less than this value will have data imported into the file system.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `52673613135872`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Trigger`

The event that triggers data import. Valid values are
`ON_DIRECTORY_FIRST_ACCESS` (import when a directory is first accessed) and
`ON_FILE_ACCESS` (import when a file is accessed).

_Required_: Yes

_Type_: String

_Allowed values_: `ON_DIRECTORY_FIRST_ACCESS | ON_FILE_ACCESS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExpirationDataRule

SynchronizationConfiguration

All content copied from https://docs.aws.amazon.com/.
