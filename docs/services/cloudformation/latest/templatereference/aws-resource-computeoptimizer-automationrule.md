This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ComputeOptimizer::AutomationRule

The `AWS::ComputeOptimizer::AutomationRule` resource specifies an automation rule that automatically implements recommended actions based on your defined criteria and schedule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ComputeOptimizer::AutomationRule",
  "Properties" : {
      "Criteria" : Criteria,
      "Description" : String,
      "Name" : String,
      "OrganizationConfiguration" : OrganizationConfiguration,
      "Priority" : String,
      "RecommendedActionTypes" : [ String, ... ],
      "RuleType" : String,
      "Schedule" : Schedule,
      "Status" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ComputeOptimizer::AutomationRule
Properties:
  Criteria:
    Criteria
  Description: String
  Name: String
  OrganizationConfiguration:
    OrganizationConfiguration
  Priority: String
  RecommendedActionTypes:
    - String
  RuleType: String
  Schedule:
    Schedule
  Status: String
  Tags:
    - Tag

```

## Properties

`Criteria`

A set of conditions that specify which recommended action qualify for implementation. When a rule is active and a recommended action matches these criteria, Compute Optimizer implements the action at the scheduled run time. You can specify up to 20 conditions per filter criteria and 20 values per condition.

_Required_: No

_Type_: [Criteria](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-computeoptimizer-automationrule-criteria.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the automation rule.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-\s@\.]*$`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the automation rule.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]*$`

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrganizationConfiguration`

Configuration settings for organization-wide rules.

_Required_: No

_Type_: [OrganizationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-computeoptimizer-automationrule-organizationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Priority`

A string representation of a decimal number between 0 and 1 (having up to 30 digits after the decimal point) that determines the priority of the rule. When multiple rules match the same recommended action, Compute Optimizer assigns the action to the rule with the lowest priority value (highest priority), even if that rule is scheduled to run later than other matching rules.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecommendedActionTypes`

List of recommended action types that this rule can execute.

_Required_: Yes

_Type_: Array of String

_Allowed values_: `SnapshotAndDeleteUnattachedEbsVolume | UpgradeEbsVolumeType`

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleType`

The type of automation rule (OrganizationRule or AccountRule).

_Required_: Yes

_Type_: String

_Allowed values_: `AccountRule | OrganizationRule`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Schedule`

The schedule configuration for when the automation rule should execute.

_Required_: Yes

_Type_: [Schedule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-computeoptimizer-automationrule-schedule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The current status of the automation rule (Active or Inactive).

_Required_: Yes

_Type_: String

_Allowed values_: `Active | Inactive`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to associate with the rule.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-computeoptimizer-automationrule-tag.html)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the automation rule.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AccountId`

The 12-digit AWS account ID that owns this automation rule.

`CreatedTimestamp`

The timestamp when the automation rule was created.

`LastUpdatedTimestamp`

The timestamp when the automation rule was last updated.

`RuleArn`

The Amazon Resource Name (ARN) of the automation rule.

`RuleId`

The unique identifier of the automation rule.

`RuleRevision`

The revision number of the automation rule.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Compute Optimizer

Criteria
