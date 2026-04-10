This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Transformer MoveKeys

This processor moves a key from one field to another. The original key is
deleted.

For more information about this processor including examples, see [moveKeys](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-moveKeys) in the _CloudWatch Logs User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Entries" : [ MoveKeyEntry, ... ]
}

```

### YAML

```yaml

  Entries:
    - MoveKeyEntry

```

## Properties

`Entries`

An array of objects, where each object contains the information about one key to move.

_Required_: Yes

_Type_: Array of [MoveKeyEntry](aws-properties-logs-transformer-movekeyentry.md)

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MoveKeyEntry

ParseCloudfront

All content copied from https://docs.aws.amazon.com/.
