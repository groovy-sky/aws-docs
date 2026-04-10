This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis StringDefaultValues

The default values of the `StringParameterDeclaration`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DynamicValue" : DynamicDefaultValue,
  "StaticValues" : [ String, ... ]
}

```

### YAML

```yaml

  DynamicValue:
    DynamicDefaultValue
  StaticValues:
    - String

```

## Properties

`DynamicValue`

The dynamic value of the `StringDefaultValues`. Different defaults displayed according to users, groups, and values mapping.

_Required_: No

_Type_: [DynamicDefaultValue](aws-properties-quicksight-analysis-dynamicdefaultvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StaticValues`

The static values of the `DecimalDefaultValues`.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `50000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StaticFileUrlSourceOptions

StringFormatConfiguration

All content copied from https://docs.aws.amazon.com/.
