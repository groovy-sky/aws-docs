This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Component SortProperty

The `SortProperty` property specifies how to sort the data that you bind to a
component.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Direction" : String,
  "Field" : String
}

```

### YAML

```yaml

  Direction: String
  Field: String

```

## Properties

`Direction`

The direction of the sort, either ascending or descending.

_Required_: Yes

_Type_: String

_Allowed values_: `ASC | DESC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Field`

The field to perform the sort on.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Predicate

AWS::AmplifyUIBuilder::Form

All content copied from https://docs.aws.amazon.com/.
