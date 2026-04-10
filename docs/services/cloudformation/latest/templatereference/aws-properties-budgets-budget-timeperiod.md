This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Budgets::Budget TimePeriod

The period of time that is covered by a budget. The period has a start date and an end
date. The start date must come before the end date. There are no restrictions on the end date.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "End" : String,
  "Start" : String
}

```

### YAML

```yaml

  End: String
  Start: String

```

## Properties

`End`

The end date for a budget. If you didn't specify an end date, AWS set
your end date to `06/15/87 00:00 UTC`. The defaults are the same for the
AWS Billing and Cost Management console and the API.

After the end date, AWS deletes the budget and all the associated
notifications and subscribers. You can change your end date with the
`UpdateBudget` operation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Start`

The start date for a budget. If you created your budget and didn't specify a start
date, the start date defaults to the start of the chosen time period (MONTHLY, QUARTERLY, or
ANNUALLY). For example, if you create your budget on January 24, 2019, choose
`MONTHLY`, and don't set a start date, the start date defaults to
`01/01/19 00:00 UTC`. The defaults are the same for the AWS Billing and Cost Management console and the API.

You can change your start date with the `UpdateBudget` operation.

Valid values depend on the value of `BudgetType`:

- If `BudgetType` is `COST` or `USAGE`: Valid values are
`MONTHLY`, `QUARTERLY`, and `ANNUALLY`.

- If `BudgetType` is `RI_UTILIZATION` or `RI_COVERAGE`: Valid values are
`DAILY`, `MONTHLY`, `QUARTERLY`, and `ANNUALLY`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [TimePeriod](../../../../reference/aws-cost-management/latest/apireference/api-budgets-timeperiod.md)
in the _AWS Cost Explorer Service Cost Management APIs_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagValues

AWS::Budgets::BudgetsAction

All content copied from https://docs.aws.amazon.com/.
