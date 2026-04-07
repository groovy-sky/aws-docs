This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Budgets::BudgetsAction

The `AWS::Budgets::BudgetsAction` resource enables you to take predefined actions that are initiated when a budget threshold has been exceeded.
For more information, see [Managing Your Costs with Budgets](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/budgets-managing-costs.html)
in the _AWS Billing and Cost Management User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Budgets::BudgetsAction",
  "Properties" : {
      "ActionThreshold" : ActionThreshold,
      "ActionType" : String,
      "ApprovalModel" : String,
      "BudgetName" : String,
      "Definition" : Definition,
      "ExecutionRoleArn" : String,
      "NotificationType" : String,
      "ResourceTags" : [ ResourceTag, ... ],
      "Subscribers" : [ Subscriber, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Budgets::BudgetsAction
Properties:
  ActionThreshold:
    ActionThreshold
  ActionType: String
  ApprovalModel: String
  BudgetName: String
  Definition:
    Definition
  ExecutionRoleArn: String
  NotificationType: String
  ResourceTags:
    - ResourceTag
  Subscribers:
    - Subscriber

```

## Properties

`ActionThreshold`

The trigger threshold of the action.

_Required_: Yes

_Type_: [ActionThreshold](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-budgets-budgetsaction-actionthreshold.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ActionType`

The type of action. This defines the type of tasks that can be carried out by this action. This field also determines the format for definition.

_Required_: Yes

_Type_: String

_Allowed values_: `APPLY_IAM_POLICY | APPLY_SCP_POLICY | RUN_SSM_DOCUMENTS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ApprovalModel`

This specifies if the action needs manual or automatic approval.

_Required_: No

_Type_: String

_Allowed values_: `AUTOMATIC | MANUAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BudgetName`

A string that represents the budget name. ":" and "\\" characters aren't allowed.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Definition`

Specifies all of the type-specific parameters.

_Required_: Yes

_Type_: [Definition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-budgets-budgetsaction-definition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionRoleArn`

The role passed for action execution and reversion. Roles and actions must be in the same account.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-eusc|-cn|-us-gov|-iso|-iso-[a-z]{1})?:iam::\d{12}:role(\u002F[\u0021-\u007F]+\u002F|\u002F)[\w+=,.@-]+$`

_Minimum_: `32`

_Maximum_: `618`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationType`

The type of a notification.

_Required_: Yes

_Type_: String

_Allowed values_: `ACTUAL | FORECASTED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceTags`

An optional list of tags to associate with the specified budget action. Each tag
consists of a key and a value, and each key must be unique for the resource.

_Required_: No

_Type_: Array of [ResourceTag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-budgets-budgetsaction-resourcetag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subscribers`

A list of subscribers.

_Required_: Yes

_Type_: Array of [Subscriber](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-budgets-budgetsaction-subscriber.html)

_Minimum_: `1`

_Maximum_: `11`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

`ActionId`

A system-generated universally unique identifier (UUID) for the action.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TimePeriod

ActionThreshold
