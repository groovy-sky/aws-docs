This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InspectorV2::Filter StringFilter

An object that describes the details of a string filter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Comparison" : String,
  "Value" : String
}

```

### YAML

```yaml

  Comparison: String
  Value: String

```

## Properties

`Comparison`

The operator to use when comparing values in the filter.

_Required_: Yes

_Type_: String

_Allowed values_: `EQUALS | PREFIX | NOT_EQUALS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value to filter on.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PortRangeFilter

Next

All content copied from https://docs.aws.amazon.com/.
