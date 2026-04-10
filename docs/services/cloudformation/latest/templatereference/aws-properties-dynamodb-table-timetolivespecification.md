This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::Table TimeToLiveSpecification

Represents the settings used to enable or disable Time to Live (TTL) for the specified
table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttributeName" : String,
  "Enabled" : Boolean
}

```

### YAML

```yaml

  AttributeName: String
  Enabled: Boolean

```

## Properties

`AttributeName`

The name of the TTL attribute used to store the expiration time for items in the
table.

###### Note

- The `AttributeName` property is required when enabling the TTL, or
when TTL is already enabled.

- To update this property, you must first disable TTL and then enable TTL with
the new attribute name.

_Required_: Conditional

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Indicates whether TTL is to be enabled (true) or disabled (false) on the table.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

WarmThroughput

All content copied from https://docs.aws.amazon.com/.
