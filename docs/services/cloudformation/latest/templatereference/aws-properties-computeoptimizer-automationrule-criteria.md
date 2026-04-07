This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ComputeOptimizer::AutomationRule Criteria

A set of conditions that specify which recommended action qualify for implementation. When a rule is active and a recommended action matches these criteria, Compute Optimizer implements the action at the scheduled run time. You can specify up to 20 conditions per filter criteria and 20 values per condition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EbsVolumeSizeInGib" : [ IntegerCriteriaCondition, ... ],
  "EbsVolumeType" : [ StringCriteriaCondition, ... ],
  "EstimatedMonthlySavings" : [ DoubleCriteriaCondition, ... ],
  "LookBackPeriodInDays" : [ IntegerCriteriaCondition, ... ],
  "Region" : [ StringCriteriaCondition, ... ],
  "ResourceArn" : [ StringCriteriaCondition, ... ],
  "ResourceTag" : [ ResourceTagsCriteriaCondition, ... ],
  "RestartNeeded" : [ StringCriteriaCondition, ... ]
}

```

### YAML

```yaml

  EbsVolumeSizeInGib:
    - IntegerCriteriaCondition
  EbsVolumeType:
    - StringCriteriaCondition
  EstimatedMonthlySavings:
    - DoubleCriteriaCondition
  LookBackPeriodInDays:
    - IntegerCriteriaCondition
  Region:
    - StringCriteriaCondition
  ResourceArn:
    - StringCriteriaCondition
  ResourceTag:
    - ResourceTagsCriteriaCondition
  RestartNeeded:
    - StringCriteriaCondition

```

## Properties

`EbsVolumeSizeInGib`

Filter criteria for EBS volume sizes in gibibytes (GiB).

_Required_: No

_Type_: Array of [IntegerCriteriaCondition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-computeoptimizer-automationrule-integercriteriacondition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EbsVolumeType`

Filter criteria for EBS volume types, such as gp2, gp3, io1, io2, st1, or sc1.

_Required_: No

_Type_: Array of [StringCriteriaCondition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-computeoptimizer-automationrule-stringcriteriacondition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EstimatedMonthlySavings`

Filter criteria for estimated monthly cost savings from the recommended action.

_Required_: No

_Type_: Array of [DoubleCriteriaCondition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-computeoptimizer-automationrule-doublecriteriacondition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LookBackPeriodInDays`

Filter criteria for the lookback period in days used to analyze resource utilization.

_Required_: No

_Type_: Array of [IntegerCriteriaCondition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-computeoptimizer-automationrule-integercriteriacondition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Region`

Filter criteria for AWS regions where resources must be located.

_Required_: No

_Type_: Array of [StringCriteriaCondition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-computeoptimizer-automationrule-stringcriteriacondition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceArn`

Filter criteria for specific resource ARNs to include or exclude.

_Required_: No

_Type_: Array of [StringCriteriaCondition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-computeoptimizer-automationrule-stringcriteriacondition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceTag`

Filter criteria for resource tags, allowing filtering by tag key and value combinations.

_Required_: No

_Type_: Array of [ResourceTagsCriteriaCondition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-computeoptimizer-automationrule-resourcetagscriteriacondition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RestartNeeded`

Filter criteria indicating whether the recommended action requires a resource restart.

_Required_: No

_Type_: Array of [StringCriteriaCondition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-computeoptimizer-automationrule-stringcriteriacondition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ComputeOptimizer::AutomationRule

DoubleCriteriaCondition
