This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Transformer RenameKeyEntry

This object defines one key that will be renamed with the [renameKey](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-renameKey) processor.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "OverwriteIfExists" : Boolean,
  "RenameTo" : String
}

```

### YAML

```yaml

  Key: String
  OverwriteIfExists: Boolean
  RenameTo: String

```

## Properties

`Key`

The key to rename

_Required_: Yes

_Type_: String

_Pattern_: `^.*[a-zA-Z0-9]+.*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OverwriteIfExists`

Specifies whether to overwrite the existing value if the destination key already
exists. The default is `false`

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RenameTo`

The string to use for the new key name

_Required_: Yes

_Type_: String

_Pattern_: `^.*[a-zA-Z0-9]+.*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Processor

RenameKeys

All content copied from https://docs.aws.amazon.com/.
