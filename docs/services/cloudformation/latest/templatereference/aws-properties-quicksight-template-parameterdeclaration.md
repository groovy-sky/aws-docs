This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template ParameterDeclaration

The declaration definition of a parameter.

For more information, see [Parameters in Amazon Quick Sight](../../../quicksight/latest/user/parameters-in-quicksight.md) in the _Amazon Quick Suite User Guide_.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DateTimeParameterDeclaration" : DateTimeParameterDeclaration,
  "DecimalParameterDeclaration" : DecimalParameterDeclaration,
  "IntegerParameterDeclaration" : IntegerParameterDeclaration,
  "StringParameterDeclaration" : StringParameterDeclaration
}

```

### YAML

```yaml

  DateTimeParameterDeclaration:
    DateTimeParameterDeclaration
  DecimalParameterDeclaration:
    DecimalParameterDeclaration
  IntegerParameterDeclaration:
    IntegerParameterDeclaration
  StringParameterDeclaration:
    StringParameterDeclaration

```

## Properties

`DateTimeParameterDeclaration`

A parameter declaration for the `DateTime` data type.

_Required_: No

_Type_: [DateTimeParameterDeclaration](aws-properties-quicksight-template-datetimeparameterdeclaration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DecimalParameterDeclaration`

A parameter declaration for the `Decimal` data type.

_Required_: No

_Type_: [DecimalParameterDeclaration](aws-properties-quicksight-template-decimalparameterdeclaration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntegerParameterDeclaration`

A parameter declaration for the `Integer` data type.

_Required_: No

_Type_: [IntegerParameterDeclaration](aws-properties-quicksight-template-integerparameterdeclaration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StringParameterDeclaration`

A parameter declaration for the `String` data type.

_Required_: No

_Type_: [StringParameterDeclaration](aws-properties-quicksight-template-stringparameterdeclaration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParameterDateTimePickerControl

ParameterDropDownControl

All content copied from https://docs.aws.amazon.com/.
