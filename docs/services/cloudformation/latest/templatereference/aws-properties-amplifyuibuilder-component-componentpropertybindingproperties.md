This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Component ComponentPropertyBindingProperties

The `ComponentPropertyBindingProperties` property specifies a component
property to associate with a binding property. This enables exposed properties on the top
level component to propagate data to the component's property values.

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

The component property to bind to the data field.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComponentProperty

ComponentVariant

All content copied from https://docs.aws.amazon.com/.
