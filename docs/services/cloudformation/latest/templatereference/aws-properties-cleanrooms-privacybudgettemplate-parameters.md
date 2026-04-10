This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::PrivacyBudgetTemplate Parameters

Specifies the epsilon and noise parameters for the privacy budget template.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BudgetParameters" : [ BudgetParameter, ... ],
  "Epsilon" : Integer,
  "ResourceArn" : String,
  "UsersNoisePerQuery" : Integer
}

```

### YAML

```yaml

  BudgetParameters:
    - BudgetParameter
  Epsilon: Integer
  ResourceArn: String
  UsersNoisePerQuery: Integer

```

## Properties

`BudgetParameters`

Property description not available.

_Required_: No

_Type_: Array of [BudgetParameter](aws-properties-cleanrooms-privacybudgettemplate-budgetparameter.md)

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Epsilon`

The epsilon value that you want to use.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceArn`

Property description not available.

_Required_: No

_Type_: String

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UsersNoisePerQuery`

Noise added per query is measured in terms of the number of users whose contributions you want to obscure. This value governs the rate at which the privacy budget is depleted.

_Required_: No

_Type_: Integer

_Minimum_: `10`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BudgetParameter

Tag

All content copied from https://docs.aws.amazon.com/.
