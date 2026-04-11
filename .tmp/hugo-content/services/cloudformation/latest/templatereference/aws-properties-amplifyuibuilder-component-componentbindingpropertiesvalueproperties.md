This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Component ComponentBindingPropertiesValueProperties

The `ComponentBindingPropertiesValueProperties` property specifies the data
binding configuration for a specific property using data stored in AWS. For
AWS connected properties, you can bind a property to data stored in an
Amazon S3 bucket, an Amplify DataStore model or an authenticated user
attribute.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bucket" : String,
  "DefaultValue" : String,
  "Field" : String,
  "Key" : String,
  "Model" : String,
  "Predicates" : [ Predicate, ... ],
  "SlotName" : String,
  "UserAttribute" : String
}

```

### YAML

```yaml

  Bucket: String
  DefaultValue: String
  Field: String
  Key: String
  Model: String
  Predicates:
    - Predicate
  SlotName: String
  UserAttribute: String

```

## Properties

`Bucket`

An Amazon S3 bucket.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultValue`

The default value to assign to the property.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Field`

The field to bind the data to.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Key`

The storage key for an Amazon S3 bucket.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Model`

An Amplify DataStore model.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Predicates`

A list of predicates for binding a component's properties to data.

_Required_: No

_Type_: Array of [Predicate](aws-properties-amplifyuibuilder-component-predicate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SlotName`

The name of a component slot.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserAttribute`

An authenticated user attribute.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComponentBindingPropertiesValue

ComponentChild

All content copied from https://docs.aws.amazon.com/.
