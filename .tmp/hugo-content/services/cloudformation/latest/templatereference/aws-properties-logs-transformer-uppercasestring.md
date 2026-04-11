This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Transformer UpperCaseString

This processor converts a string field to uppercase.

For more information about this processor including examples, see [upperCaseString](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-upperCaseString) in the _CloudWatch Logs User_
_Guide_.

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

The array of containing the keys of the field to convert to uppercase.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TypeConverterEntry

Next

All content copied from https://docs.aws.amazon.com/.
