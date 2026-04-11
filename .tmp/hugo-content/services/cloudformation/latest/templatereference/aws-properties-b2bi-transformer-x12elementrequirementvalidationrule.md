This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Transformer X12ElementRequirementValidationRule

Defines a validation rule that modifies the requirement status of a specific X12 element
within a segment. This rule allows you to make optional elements mandatory or mandatory
elements optional, providing flexibility to accommodate different trading partner
requirements and business rules. The rule targets a specific element position within a
segment and sets its requirement status to either OPTIONAL or MANDATORY.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ElementPosition" : String,
  "Requirement" : String
}

```

### YAML

```yaml

  ElementPosition: String
  Requirement: String

```

## Properties

`ElementPosition`

Specifies the position of the element within an X12 segment for which the requirement
status will be modified. The format follows the pattern of segment identifier followed by
element position (e.g., "ST-01" for the first element of the ST segment).

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9]+(?:-\d{2})(?:-\d{2})?$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Requirement`

Specifies the requirement status for the element at the specified position. Valid values
are OPTIONAL (the element may be omitted) or MANDATORY (the element must be
present).

_Required_: Yes

_Type_: String

_Allowed values_: `OPTIONAL | MANDATORY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

X12ElementLengthValidationRule

X12SplitOptions

All content copied from https://docs.aws.amazon.com/.
