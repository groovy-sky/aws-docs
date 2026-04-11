This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Transformer SubstituteString

This processor matches a key’s value against a regular expression and replaces all
matches with a replacement string.

For more information about this processor including examples, see [substituteString](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-substituteString) in the _CloudWatch Logs User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Entries" : [ SubstituteStringEntry, ... ]
}

```

### YAML

```yaml

  Entries:
    - SubstituteStringEntry

```

## Properties

`Entries`

An array of objects, where each object contains the information about one key to match
and replace.

_Required_: Yes

_Type_: Array of [SubstituteStringEntry](aws-properties-logs-transformer-substitutestringentry.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SplitStringEntry

SubstituteStringEntry

All content copied from https://docs.aws.amazon.com/.
