This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Transformer X12ValidationRule

Represents a single validation rule that can be applied during X12 EDI processing. This
is a union type that can contain one of several specific validation rule types: code list
validation rules for modifying allowed element codes, element length validation rules for
enforcing custom length constraints, or element requirement validation rules for changing
mandatory/optional status. Each validation rule targets specific aspects of EDI document
validation to ensure compliance with trading partner requirements and business
rules.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CodeListValidationRule" : X12CodeListValidationRule,
  "ElementLengthValidationRule" : X12ElementLengthValidationRule,
  "ElementRequirementValidationRule" : X12ElementRequirementValidationRule
}

```

### YAML

```yaml

  CodeListValidationRule:
    X12CodeListValidationRule
  ElementLengthValidationRule:
    X12ElementLengthValidationRule
  ElementRequirementValidationRule:
    X12ElementRequirementValidationRule

```

## Properties

`CodeListValidationRule`

Specifies a code list validation rule that modifies the allowed code values for a
specific X12 element. This rule enables you to customize which codes are considered valid
for an element, allowing for trading partner-specific code requirements.

_Required_: No

_Type_: [X12CodeListValidationRule](aws-properties-b2bi-transformer-x12codelistvalidationrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ElementLengthValidationRule`

Specifies an element length validation rule that defines custom length constraints for a
specific X12 element. This rule allows you to enforce minimum and maximum length
requirements that may differ from the standard X12 specification.

_Required_: No

_Type_: [X12ElementLengthValidationRule](aws-properties-b2bi-transformer-x12elementlengthvalidationrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ElementRequirementValidationRule`

Specifies an element requirement validation rule that modifies whether a specific X12
element is required or optional within a segment. This rule provides flexibility to
accommodate different trading partner requirements for element presence.

_Required_: No

_Type_: [X12ElementRequirementValidationRule](aws-properties-b2bi-transformer-x12elementrequirementvalidationrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

X12ValidationOptions

Next

All content copied from https://docs.aws.amazon.com/.
