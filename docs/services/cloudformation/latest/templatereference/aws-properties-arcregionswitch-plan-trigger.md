This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan Trigger

Defines a condition that can automatically trigger the execution of a Region switch plan.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String,
  "Conditions" : [ TriggerCondition, ... ],
  "Description" : String,
  "MinDelayMinutesBetweenExecutions" : Number,
  "TargetRegion" : String
}

```

### YAML

```yaml

  Action: String
  Conditions:
    - TriggerCondition
  Description: String
  MinDelayMinutesBetweenExecutions: Number
  TargetRegion: String

```

## Properties

`Action`

The action to perform when the trigger fires. Valid values include `activate` and `deactivate`.

_Required_: Yes

_Type_: String

_Allowed values_: `activate | deactivate | postRecovery`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Conditions`

The conditions that must be met for the trigger to fire.

_Required_: Yes

_Type_: Array of [TriggerCondition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-arcregionswitch-plan-triggercondition.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description for a trigger.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinDelayMinutesBetweenExecutions`

The minimum time, in minutes, that must elapse between automatic executions of the plan.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetRegion`

The AWS Region for a trigger.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z]{2}-[a-z-]+-\d+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Step

TriggerCondition
