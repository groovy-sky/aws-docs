This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Transformer ListToMap

This processor takes a list of objects that contain key fields, and converts them into
a map of target keys.

For more information about this processor including examples, see [listToMap](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-listToMap) in the _CloudWatch Logs User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Flatten" : Boolean,
  "FlattenedElement" : String,
  "Key" : String,
  "Source" : String,
  "Target" : String,
  "ValueKey" : String
}

```

### YAML

```yaml

  Flatten: Boolean
  FlattenedElement: String
  Key: String
  Source: String
  Target: String
  ValueKey: String

```

## Properties

`Flatten`

A Boolean value to indicate whether the list will be flattened into single items.
Specify `true` to flatten the list. The default is `false`

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FlattenedElement`

If you set `flatten` to `true`, use
`flattenedElement` to specify which element, `first` or
`last`, to keep.

You must specify this parameter if `flatten` is `true`

_Required_: No

_Type_: String

_Allowed values_: `first | last`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Key`

The key of the field to be extracted as keys in the generated map

_Required_: Yes

_Type_: String

_Pattern_: `^.*[a-zA-Z0-9]+.*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The key in the log event that has a list of objects that will be converted to a
map.

_Required_: Yes

_Type_: String

_Pattern_: `^.*[a-zA-Z0-9]+.*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Target`

The key of the field that will hold the generated map

_Required_: No

_Type_: String

_Pattern_: `^.*[a-zA-Z0-9]+.*$`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueKey`

If this is specified, the values that you specify in this parameter will be extracted
from the `source` objects and put into the values of the generated map.
Otherwise, original objects in the source list will be put into the values of the
generated map.

_Required_: No

_Type_: String

_Pattern_: `^.*[a-zA-Z0-9]+.*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Grok

LowerCaseString

All content copied from https://docs.aws.amazon.com/.
