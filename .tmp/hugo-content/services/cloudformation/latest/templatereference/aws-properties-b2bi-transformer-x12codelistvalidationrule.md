This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Transformer X12CodeListValidationRule

Code list validation rule configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CodesToAdd" : [ String, ... ],
  "CodesToRemove" : [ String, ... ],
  "ElementId" : String
}

```

### YAML

```yaml

  CodesToAdd:
    - String
  CodesToRemove:
    - String
  ElementId: String

```

## Properties

`CodesToAdd`

Specifies a list of code values to add to the element's allowed values. These codes will
be considered valid for the specified element in addition to the standard codes defined by
the X12 specification.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CodesToRemove`

Specifies a list of code values to remove from the element's allowed values. These codes
will be considered invalid for the specified element, even if they are part of the standard
codes defined by the X12 specification.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ElementId`

Specifies the four-digit element ID to which the code list modifications apply. This
identifies which X12 element will have its allowed code values modified.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9]{4}$`

_Minimum_: `4`

_Maximum_: `4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

X12AdvancedOptions

X12Details

All content copied from https://docs.aws.amazon.com/.
