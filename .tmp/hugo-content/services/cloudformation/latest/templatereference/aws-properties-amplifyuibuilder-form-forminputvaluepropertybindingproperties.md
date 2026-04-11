This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Form FormInputValuePropertyBindingProperties

Associates a form property to a binding property. This enables exposed properties on the
top level form to propagate data to the form's property values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Field" : String,
  "Property" : String
}

```

### YAML

```yaml

  Field: String
  Property: String

```

## Properties

`Field`

The data field to bind the property to.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Property`

The form property to bind to the data field.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FormInputValueProperty

FormStyle

All content copied from https://docs.aws.amazon.com/.
