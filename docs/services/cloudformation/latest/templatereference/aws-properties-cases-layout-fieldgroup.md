This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cases::Layout FieldGroup

Object for a group of fields and associated properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Fields" : [ FieldItem, ... ],
  "Name" : String
}

```

### YAML

```yaml

  Fields:
    - FieldItem
  Name: String

```

## Properties

`Fields`

Represents an ordered list containing field related information.

_Required_: Yes

_Type_: Array of [FieldItem](aws-properties-cases-layout-fielditem.md)

_Maximum_: `220`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Name of the field group.

_Required_: No

_Type_: String

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BasicLayout

FieldItem

All content copied from https://docs.aws.amazon.com/.
