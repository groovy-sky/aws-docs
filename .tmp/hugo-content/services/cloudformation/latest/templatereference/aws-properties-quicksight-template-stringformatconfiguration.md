This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template StringFormatConfiguration

Formatting configuration for string fields.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NullValueFormatConfiguration" : NullValueFormatConfiguration,
  "NumericFormatConfiguration" : NumericFormatConfiguration
}

```

### YAML

```yaml

  NullValueFormatConfiguration:
    NullValueFormatConfiguration
  NumericFormatConfiguration:
    NumericFormatConfiguration

```

## Properties

`NullValueFormatConfiguration`

The options that determine the null value format configuration.

_Required_: No

_Type_: [NullValueFormatConfiguration](aws-properties-quicksight-template-nullvalueformatconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumericFormatConfiguration`

The formatting configuration for numeric strings.

_Required_: No

_Type_: [NumericFormatConfiguration](aws-properties-quicksight-template-numericformatconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StringDefaultValues

StringParameterDeclaration

All content copied from https://docs.aws.amazon.com/.
