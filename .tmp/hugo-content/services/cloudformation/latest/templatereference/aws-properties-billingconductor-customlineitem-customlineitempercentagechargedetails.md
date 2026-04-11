This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BillingConductor::CustomLineItem CustomLineItemPercentageChargeDetails

A representation of the charge details associated with a percentage custom line item.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChildAssociatedResources" : [ String, ... ],
  "PercentageValue" : Number
}

```

### YAML

```yaml

  ChildAssociatedResources:
    - String
  PercentageValue: Number

```

## Properties

`ChildAssociatedResources`

A list of resource ARNs to associate to the percentage custom line item.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PercentageValue`

The custom line item's percentage value. This will be multiplied against the combined value of its associated resources to determine its charge value.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomLineItemFlatChargeDetails

LineItemFilter

All content copied from https://docs.aws.amazon.com/.
