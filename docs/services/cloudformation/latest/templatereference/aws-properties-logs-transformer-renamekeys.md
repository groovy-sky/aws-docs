This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Transformer RenameKeys

Use this processor to rename keys in a log event.

For more information about this processor including examples, see [renameKeys](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-renameKeys) in the _CloudWatch Logs User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Entries" : [ RenameKeyEntry, ... ]
}

```

### YAML

```yaml

  Entries:
    - RenameKeyEntry

```

## Properties

`Entries`

An array of `RenameKeyEntry` objects, where each object contains the
information about a single key to rename.

_Required_: Yes

_Type_: Array of [RenameKeyEntry](aws-properties-logs-transformer-renamekeyentry.md)

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RenameKeyEntry

SplitString

All content copied from https://docs.aws.amazon.com/.
