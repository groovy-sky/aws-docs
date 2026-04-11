This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KendraRanking::ExecutionPlan Tag

A key-value pair that identifies or categorizes a rescore
execution plan. A rescore execution plan is an
Amazon Kendra Intelligent Ranking resource used for
provisioning the `Rescore` API. You can also use
a tag to help control access to a rescore execution plan.
A tag key and value can consist of Unicode letters, digits,
white space, and any of the following symbols: \_ . : / = + - @.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

The key for the tag. Keys are not case sensitive and must
be unique.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value associated with the tag. The value can be an
empty string but it can't be null.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CapacityUnitsConfiguration

Next

All content copied from https://docs.aws.amazon.com/.
