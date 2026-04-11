This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Budgets::Budget Spend

The amount of cost or usage that's measured for a budget.

_Cost example:_ A `Spend` for `3 USD` of
costs has the following parameters:

- An `Amount` of `3`

- A `Unit` of `USD`

_Usage example:_ A `Spend` for `3 GB` of S3
usage has the following parameters:

- An `Amount` of `3`

- A `Unit` of `GB`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Amount" : Number,
  "Unit" : String
}

```

### YAML

```yaml

  Amount: Number
  Unit: String

```

## Properties

`Amount`

The cost or usage amount that's associated with a budget forecast, actual spend, or
budget threshold.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unit`

The unit of measurement that's used for the budget forecast, actual spend, or budget
threshold.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Spend](../../../../reference/aws-cost-management/latest/apireference/api-budgets-spend.md)
in the _AWS Cost Explorer Service Cost Management APIs_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceTag

Subscriber

All content copied from https://docs.aws.amazon.com/.
