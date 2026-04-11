This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Transformer CopyValueEntry

This object defines one value to be copied with the [copyValue](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-copyValue) processor.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OverwriteIfExists" : Boolean,
  "Source" : String,
  "Target" : String
}

```

### YAML

```yaml

  OverwriteIfExists: Boolean
  Source: String
  Target: String

```

## Properties

`OverwriteIfExists`

Specifies whether to overwrite the value if the destination key already exists. If you
omit this, the default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The key to copy.

_Required_: Yes

_Type_: String

_Pattern_: `^.*[a-zA-Z0-9]+.*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Target`

The key of the field to copy the value to.

_Required_: Yes

_Type_: String

_Pattern_: `^.*[a-zA-Z0-9]+.*$`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CopyValue

Csv

All content copied from https://docs.aws.amazon.com/.
