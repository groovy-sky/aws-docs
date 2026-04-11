This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Transformer X12ElementLengthValidationRule

Defines a validation rule that specifies custom length constraints for a specific X12
element. This rule allows you to override the standard minimum and maximum length
requirements for an element, enabling validation of trading partner-specific length
requirements that may differ from the X12 specification. Both minimum and maximum length
values must be specified.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ElementId" : String,
  "MaxLength" : Number,
  "MinLength" : Number
}

```

### YAML

```yaml

  ElementId: String
  MaxLength: Number
  MinLength: Number

```

## Properties

`ElementId`

Specifies the four-digit element ID to which the length constraints will be applied.
This identifies which X12 element will have its length requirements modified.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9]{4}$`

_Minimum_: `4`

_Maximum_: `4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxLength`

Specifies the maximum allowed length for the identified element. This value defines the
upper limit for the element's content length.

_Required_: Yes

_Type_: Number

_Minimum_: `1`

_Maximum_: `1000000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinLength`

Specifies the minimum required length for the identified element. This value defines the
lower limit for the element's content length.

_Required_: Yes

_Type_: Number

_Minimum_: `1`

_Maximum_: `1000000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

X12Details

X12ElementRequirementValidationRule

All content copied from https://docs.aws.amazon.com/.
