This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Budgets::Budget CostTypes

The types of cost that are included in a `COST` budget, such as tax and
subscriptions.

`USAGE`, `RI_UTILIZATION`, `RI_COVERAGE`,
`SAVINGS_PLANS_UTILIZATION`, and `SAVINGS_PLANS_COVERAGE`
budgets don't have `CostTypes`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IncludeCredit" : Boolean,
  "IncludeDiscount" : Boolean,
  "IncludeOtherSubscription" : Boolean,
  "IncludeRecurring" : Boolean,
  "IncludeRefund" : Boolean,
  "IncludeSubscription" : Boolean,
  "IncludeSupport" : Boolean,
  "IncludeTax" : Boolean,
  "IncludeUpfront" : Boolean,
  "UseAmortized" : Boolean,
  "UseBlended" : Boolean
}

```

### YAML

```yaml

  IncludeCredit: Boolean
  IncludeDiscount: Boolean
  IncludeOtherSubscription: Boolean
  IncludeRecurring: Boolean
  IncludeRefund: Boolean
  IncludeSubscription: Boolean
  IncludeSupport: Boolean
  IncludeTax: Boolean
  IncludeUpfront: Boolean
  UseAmortized: Boolean
  UseBlended: Boolean

```

## Properties

`IncludeCredit`

Specifies whether a budget includes credits.

The default value is `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeDiscount`

Specifies whether a budget includes discounts.

The default value is `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeOtherSubscription`

Specifies whether a budget includes non-RI subscription costs.

The default value is `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeRecurring`

Specifies whether a budget includes recurring fees such as monthly RI fees.

The default value is `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeRefund`

Specifies whether a budget includes refunds.

The default value is `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeSubscription`

Specifies whether a budget includes subscriptions.

The default value is `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeSupport`

Specifies whether a budget includes support subscription fees.

The default value is `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeTax`

Specifies whether a budget includes taxes.

The default value is `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeUpfront`

Specifies whether a budget includes upfront RI costs.

The default value is `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseAmortized`

Specifies whether a budget uses the amortized rate.

The default value is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseBlended`

Specifies whether a budget uses a blended rate.

The default value is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [CostTypes](../../../../reference/aws-cost-management/latest/apireference/api-budgets-costtypes.md)
in the _AWS Cost Explorer Service Cost Management APIs_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CostCategoryValues

Expression

All content copied from https://docs.aws.amazon.com/.
