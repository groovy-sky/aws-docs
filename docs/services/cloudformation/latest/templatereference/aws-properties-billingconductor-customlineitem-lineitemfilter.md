This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BillingConductor::CustomLineItem LineItemFilter

A representation of the line item filter for your custom line item. You can use line item filters to include or exclude specific resource values from the billing group's total cost.
For example, if you create a custom line item and you want to filter out a value, such as
Savings Plans discounts, you can update `LineItemFilter` to exclude it.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Attribute" : String,
  "AttributeValues" : [ String, ... ],
  "MatchOption" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Attribute: String
  AttributeValues:
    - String
  MatchOption: String
  Values:
    - String

```

## Properties

`Attribute`

The attribute of the line item filter. This specifies what attribute that you can filter
on.

_Required_: Yes

_Type_: String

_Allowed values_: `LINE_ITEM_TYPE | SERVICE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AttributeValues`

The values of the line item filter. This specifies the values to filter on.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchOption`

The match criteria of the line item filter. This parameter specifies whether not to include the resource value from the billing group total cost.

_Required_: Yes

_Type_: String

_Allowed values_: `NOT_EQUAL | EQUAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The values of the line item filter. This specifies the values to filter on. Currently, you can only exclude Savings Plans discounts.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomLineItemPercentageChargeDetails

PresentationDetails

All content copied from https://docs.aws.amazon.com/.
