This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Transformer DeleteKeys

This processor deletes entries from a log event. These entries are key-value pairs.

For more information about this processor including examples, see [deleteKeys](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-deleteKeys) in the _CloudWatch Logs User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "WithKeys" : [ String, ... ]
}

```

### YAML

```yaml

  WithKeys:
    - String

```

## Properties

`WithKeys`

The list of keys to delete.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DateTimeConverter

Grok

All content copied from https://docs.aws.amazon.com/.
