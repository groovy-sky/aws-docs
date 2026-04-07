This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::GlobalTable AttributeDefinition

Represents an attribute for describing the schema for the table and indexes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttributeName" : String,
  "AttributeType" : String
}

```

### YAML

```yaml

  AttributeName: String
  AttributeType: String

```

## Properties

`AttributeName`

A name for the attribute.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AttributeType`

The data type for the attribute, where:

- `S` \- the attribute is of type String

- `N` \- the attribute is of type Number

- `B` \- the attribute is of type Binary

_Required_: Yes

_Type_: String

_Allowed values_: `S | N | B`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::DynamoDB::GlobalTable

CapacityAutoScalingSettings
