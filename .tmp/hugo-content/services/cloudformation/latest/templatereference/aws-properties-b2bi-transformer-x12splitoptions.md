This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Transformer X12SplitOptions

Contains options for splitting X12 EDI files into smaller units. This is useful for
processing large EDI files more efficiently.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SplitBy" : String
}

```

### YAML

```yaml

  SplitBy: String

```

## Properties

`SplitBy`

Specifies the method used to split X12 EDI files. Valid values include
`TRANSACTION` (split by individual transaction sets), or `NONE`
(no splitting).

_Required_: No

_Type_: String

_Allowed values_: `NONE | TRANSACTION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

X12ElementRequirementValidationRule

X12ValidationOptions

All content copied from https://docs.aws.amazon.com/.
