This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Transformer TrimString

Use this processor to remove leading and trailing whitespace.

For more information about this processor including examples, see [trimString](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-trimString) in the _CloudWatch Logs User Guide_.

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

The array containing the keys of the fields to trim.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SubstituteStringEntry

TypeConverter

All content copied from https://docs.aws.amazon.com/.
