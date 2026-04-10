This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan Ec2AsgCapacityIncreaseConfiguration

Configuration for increasing the capacity of Amazon EC2 Auto Scaling groups during a Region switch.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Asgs" : [ Asg, ... ],
  "CapacityMonitoringApproach" : ,
  "TargetPercent" : Number,
  "TimeoutMinutes" : Number,
  "Ungraceful" : Ec2Ungraceful
}

```

### YAML

```yaml

  Asgs:
    - Asg
  CapacityMonitoringApproach:

  TargetPercent: Number
  TimeoutMinutes: Number
  Ungraceful:
    Ec2Ungraceful

```

## Properties

`Asgs`

The EC2 Auto Scaling groups for the configuration.

_Required_: Yes

_Type_: Array of [Asg](aws-properties-arcregionswitch-plan-asg.md)

_Minimum_: `2`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CapacityMonitoringApproach`

The monitoring approach that you specify EC2 Auto Scaling groups for the configuration.

_Required_: No

_Type_:

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetPercent`

The target percentage that you specify for EC2 Auto Scaling groups. The default is 100.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutMinutes`

The timeout value specified for the configuration.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ungraceful`

The settings for ungraceful execution.

_Required_: No

_Type_: [Ec2Ungraceful](aws-properties-arcregionswitch-plan-ec2ungraceful.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentDbUngraceful

Ec2Ungraceful

All content copied from https://docs.aws.amazon.com/.
