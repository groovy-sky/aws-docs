This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Budgets::Budget HistoricalOptions

The parameters that define or describe the historical data that your auto-adjusting budget is based on.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BudgetAdjustmentPeriod" : Integer
}

```

### YAML

```yaml

  BudgetAdjustmentPeriod: Integer

```

## Properties

`BudgetAdjustmentPeriod`

The number of budget periods included in the moving-average calculation that determines your auto-adjusted budget amount. The maximum value depends on the `TimeUnit` granularity of the budget:

- For the `DAILY` granularity, the maximum value is `60`.

- For the `MONTHLY` granularity, the maximum value is `12`.

- For the `QUARTERLY` granularity, the maximum value is `4`.

- For the `ANNUALLY` granularity, the maximum value is `1`.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `60`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExpressionDimensionValues

Notification

All content copied from https://docs.aws.amazon.com/.
