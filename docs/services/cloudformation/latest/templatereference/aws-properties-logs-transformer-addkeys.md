This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Transformer AddKeys

This processor adds new key-value pairs to the log event.

For more information about this processor including examples, see [addKeys](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-addKeys) in the _CloudWatch Logs User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Entries" : [ AddKeyEntry, ... ]
}

```

### YAML

```yaml

  Entries:
    - AddKeyEntry

```

## Properties

`Entries`

An array of objects, where each object contains the information about one key to add
to the log event.

_Required_: Yes

_Type_: Array of [AddKeyEntry](aws-properties-logs-transformer-addkeyentry.md)

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AddKeyEntry

CopyValue

All content copied from https://docs.aws.amazon.com/.
