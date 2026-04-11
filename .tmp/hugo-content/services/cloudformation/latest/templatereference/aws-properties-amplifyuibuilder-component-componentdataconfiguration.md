This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Component ComponentDataConfiguration

The `ComponentDataConfiguration` property specifies the configuration for
binding a component's properties to data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Identifiers" : [ String, ... ],
  "Model" : String,
  "Predicate" : Predicate,
  "Sort" : [ SortProperty, ... ]
}

```

### YAML

```yaml

  Identifiers:
    - String
  Model: String
  Predicate:
    Predicate
  Sort:
    - SortProperty

```

## Properties

`Identifiers`

A list of IDs to use to bind data to a component. Use this property to bind specifically
chosen data, rather than data retrieved from a query.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Model`

The name of the data model to use to bind data to a component.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Predicate`

Represents the conditional logic to use when binding data to a component. Use this
property to retrieve only a subset of the data in a collection.

_Required_: No

_Type_: [Predicate](aws-properties-amplifyuibuilder-component-predicate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sort`

Describes how to sort the component's properties.

_Required_: No

_Type_: Array of [SortProperty](aws-properties-amplifyuibuilder-component-sortproperty.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComponentConditionProperty

ComponentEvent

All content copied from https://docs.aws.amazon.com/.
