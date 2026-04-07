This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BillingConductor::CustomLineItem CustomLineItemChargeDetails

The charge details of a custom line item. It should contain only one of `Flat` or `Percentage`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Flat" : CustomLineItemFlatChargeDetails,
  "LineItemFilters" : [ LineItemFilter, ... ],
  "Percentage" : CustomLineItemPercentageChargeDetails,
  "Type" : String
}

```

### YAML

```yaml

  Flat:
    CustomLineItemFlatChargeDetails
  LineItemFilters:
    - LineItemFilter
  Percentage:
    CustomLineItemPercentageChargeDetails
  Type: String

```

## Properties

`Flat`

A `CustomLineItemFlatChargeDetails` that describes the charge details of a flat custom line item.

_Required_: No

_Type_: [CustomLineItemFlatChargeDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-billingconductor-customlineitem-customlineitemflatchargedetails.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LineItemFilters`

A representation of the line item filter.

_Required_: No

_Type_: Array of [LineItemFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-billingconductor-customlineitem-lineitemfilter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Percentage`

A `CustomLineItemPercentageChargeDetails` that describes the charge details of a percentage custom line item.

_Required_: No

_Type_: [CustomLineItemPercentageChargeDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-billingconductor-customlineitem-customlineitempercentagechargedetails.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the custom line item that indicates whether the charge is a fee or credit.

_Required_: Yes

_Type_: String

_Allowed values_: `FEE | CREDIT`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

BillingPeriodRange

CustomLineItemFlatChargeDetails
