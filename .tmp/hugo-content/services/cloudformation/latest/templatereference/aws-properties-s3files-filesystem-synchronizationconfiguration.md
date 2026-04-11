This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Files::FileSystem SynchronizationConfiguration

The synchronization configuration for the file system.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExpirationDataRules" : [ ExpirationDataRule, ... ],
  "ImportDataRules" : [ ImportDataRule, ... ],
  "LatestVersionNumber" : Integer
}

```

### YAML

```yaml

  ExpirationDataRules:
    - ExpirationDataRule
  ImportDataRules:
    - ImportDataRule
  LatestVersionNumber: Integer

```

## Properties

`ExpirationDataRules`

The rules that control when cached data expires from the file system.

_Required_: Yes

_Type_: Array of [ExpirationDataRule](aws-properties-s3files-filesystem-expirationdatarule.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImportDataRules`

The rules that control how data is imported from S3 into the file system.

_Required_: Yes

_Type_: Array of [ImportDataRule](aws-properties-s3files-filesystem-importdatarule.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LatestVersionNumber`

The latest version number of the synchronization configuration.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImportDataRule

Tag

All content copied from https://docs.aws.amazon.com/.
