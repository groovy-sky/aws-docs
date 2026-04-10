This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Transformer AddKeyEntry

This object defines one key that will be added with the [addKeys](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-addKey) processor.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "OverwriteIfExists" : Boolean,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  OverwriteIfExists: Boolean
  Value: String

```

## Properties

`Key`

The key of the new entry to be added to the log event

_Required_: Yes

_Type_: String

_Pattern_: `^.*[a-zA-Z0-9]+.*$`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OverwriteIfExists`

Specifies whether to overwrite the value if the key already exists in the log event.
If you omit this, the default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of the new entry to be added to the log event

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Logs::Transformer

AddKeys

All content copied from https://docs.aws.amazon.com/.
