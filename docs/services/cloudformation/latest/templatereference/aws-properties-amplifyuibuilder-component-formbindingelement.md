This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Component FormBindingElement

Describes how to bind a component property to form data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Element" : String,
  "Property" : String
}

```

### YAML

```yaml

  Element: String
  Property: String

```

## Properties

`Element`

The name of the component to retrieve a value from.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Property`

The property to retrieve a value from.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComponentVariant

MutationActionSetStateParameter

All content copied from https://docs.aws.amazon.com/.
