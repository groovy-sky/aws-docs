This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Billing::BillingView DataFilterExpression

See [Expression](../../../../reference/aws-cost-management/latest/apireference/api-billing-expression.md). Billing view only supports `LINKED_ACCOUNT` and `Tags`.

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

_Type_: [Dimensions](aws-properties-billing-billingview-dimensions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The specific `Tag` to use for `Expression`.

_Required_: No

_Type_: [Tags](aws-properties-billing-billingview-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeRange`

Property description not available.

_Required_: No

_Type_: [TimeRange](aws-properties-billing-billingview-timerange.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Billing::BillingView

Dimensions

All content copied from https://docs.aws.amazon.com/.
