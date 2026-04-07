This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Billing::BillingView DataFilterExpression

See [Expression](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_billing_Expression.html). Billing view only supports `LINKED_ACCOUNT` and `Tags`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Dimensions" : Dimensions,
  "Tags" : Tags,
  "TimeRange" : TimeRange
}

```

### YAML

```yaml

  Dimensions:
    Dimensions
  Tags:
    Tags
  TimeRange:
    TimeRange

```

## Properties

`Dimensions`

The specific `Dimension` to use for `Expression`.

_Required_: No

_Type_: [Dimensions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-billing-billingview-dimensions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The specific `Tag` to use for `Expression`.

_Required_: No

_Type_: [Tags](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-billing-billingview-tags.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeRange`

Property description not available.

_Required_: No

_Type_: [TimeRange](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-billing-billingview-timerange.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Billing::BillingView

Dimensions
