This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Billing::BillingView Dimensions

The specific `Dimension` to use for `Expression`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Key: String
  Values:
    - String

```

## Properties

`Key`

The key that's associated with the tag.

_Required_: No

_Type_: String

_Allowed values_: `LINKED_ACCOUNT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The metadata that you can use to filter and group your results.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1024 | 200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataFilterExpression

Tag

All content copied from https://docs.aws.amazon.com/.
