This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BillingConductor::CustomLineItem CustomLineItemFlatChargeDetails

The charge details of a custom line item. It should contain only one of `Flat` or `Percentage`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChargeValue" : Number
}

```

### YAML

```yaml

  ChargeValue: Number

```

## Properties

`ChargeValue`

The custom line item's fixed charge value in USD.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `1000000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomLineItemChargeDetails

CustomLineItemPercentageChargeDetails

All content copied from https://docs.aws.amazon.com/.
