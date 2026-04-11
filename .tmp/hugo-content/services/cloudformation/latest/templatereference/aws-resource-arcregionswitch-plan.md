This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan

Represents a Region switch plan. A plan defines the steps required to shift traffic from one AWS Region to another.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ARCRegionSwitch::Plan",
  "Properties" : {
      "AssociatedAlarms" : {Key: Value, ...},
      "Description" : String,
      "ExecutionRole" : String,
      "Name" : String,
      "PrimaryRegion" : String,
      "RecoveryApproach" : String,
      "RecoveryTimeObjectiveMinutes" : Number,
      "Regions" : [ String, ... ],
      "ReportConfiguration" : ReportConfiguration,
      "Tags" : {Key: Value, ...},
      "Triggers" : [ Trigger, ... ],
      "Workflows" : [ Workflow, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ARCRegionSwitch::Plan
Properties:
  AssociatedAlarms:
    Key: Value
  Description: String
  ExecutionRole: String
  Name: String
  PrimaryRegion: String
  RecoveryApproach: String
  RecoveryTimeObjectiveMinutes: Number
  Regions:
    - String
  ReportConfiguration:
    ReportConfiguration
  Tags:
    Key: Value
  Triggers:
    - Trigger
  Workflows:
    - Workflow

```

## Properties

`AssociatedAlarms`

The associated application health alarms for a plan.

_Required_: No

_Type_: Object of [AssociatedAlarm](aws-properties-arcregionswitch-plan-associatedalarm.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description for a plan.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionRole`

The execution role for a plan.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[a-zA-Z0-9-]*:iam::[0-9]{12}:role/.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name for a plan.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,30}[a-zA-Z0-9])?$`

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PrimaryRegion`

The primary Region for a plan.

_Required_: No

_Type_: String

_Pattern_: `^[a-z]{2}-[a-z-]+-\d+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RecoveryApproach`

The recovery approach for a Region switch plan, which can be active/active (activeActive) or active/passive (activePassive).

_Required_: Yes

_Type_: String

_Allowed values_: `activeActive | activePassive`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RecoveryTimeObjectiveMinutes`

The recovery time objective for a plan.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `10080`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Regions`

The AWS Regions for a plan.

_Required_: Yes

_Type_: Array of String

_Minimum_: `2`

_Maximum_: `2`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReportConfiguration`

The report configuration for a plan.

_Required_: No

_Type_: [ReportConfiguration](aws-properties-arcregionswitch-plan-reportconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Triggers`

The triggers for a plan.

_Required_: No

_Type_: Array of [Trigger](aws-properties-arcregionswitch-plan-trigger.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Workflows`

The workflows for a plan.

_Required_: Yes

_Type_: Array of [Workflow](aws-properties-arcregionswitch-plan-workflow.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the plan.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the plan.

`Owner`

The owner of a plan.

`PlanHealthChecks`

Property description not available.

`Version`

The version for the plan.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ARCRegionSwitch

ArcRoutingControlConfiguration

All content copied from https://docs.aws.amazon.com/.
