This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Budgets::Budget AutoAdjustData

Determine the budget amount for an auto-adjusting budget.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoAdjustType" : String,
  "HistoricalOptions" : HistoricalOptions
}

```

### YAML

```yaml

  AutoAdjustType: String
  HistoricalOptions:
    HistoricalOptions

```

## Properties

`AutoAdjustType`

The string that defines whether your budget auto-adjusts based on historical or forecasted data.

_Required_: Yes

_Type_: String

_Allowed values_: `HISTORICAL | FORECAST`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HistoricalOptions`

The parameters that define or describe the historical data that your auto-adjusting budget is based on.

_Required_: No

_Type_: [HistoricalOptions](aws-properties-budgets-budget-historicaloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Creating an auto-adjusting AWS budget based on historical data

The following example creates an auto-adjusting AWS budget based on
historical data with a 6 months adjustment period. The maximum value of the budget
adjustment period depends on the `TimeUnit` granularity of the budget. For
example, if you set `TimeUnit` to use the `MONTHLY` value, then the
maximum value of `BudgetAdjustmentPeriod` is 12. For more information, see
[BudgetAdjustmentPeriod](../userguide/aws-properties-budgets-budget-historicaloptions.md#cfn-budgets-budget-historicaloptions-budgetadjustmentperiod) in `HistoricalOptions`.

#### JSON

```json

{
    "Resources": {
        "BudgetExample": {
            "Type": "AWS::Budgets::Budget",
            "Properties": {
                "Budget": {
                    "BudgetType": "COST",
                    "TimeUnit": "MONTHLY",
                    "AutoAdjustData": {
                        "AutoAdjustType": "HISTORICAL",
                        "HistoricalOptions": {
                            "BudgetAdjustmentPeriod": 6
                        }
                    }
                }
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  BudgetExample:
    Type: "AWS::Budgets::Budget"
    Properties:
      Budget:
        BudgetType: COST
        TimeUnit: MONTHLY
        AutoAdjustData:
          AutoAdjustType: HISTORICAL
          HistoricalOptions:
              BudgetAdjustmentPeriod: 6
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Budgets::Budget

BudgetData

All content copied from https://docs.aws.amazon.com/.
