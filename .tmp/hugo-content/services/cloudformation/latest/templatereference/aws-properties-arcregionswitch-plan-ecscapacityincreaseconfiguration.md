This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan EcsCapacityIncreaseConfiguration

The configuration for an AWS ECS capacity increase.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CapacityMonitoringApproach" : ,
  "Services" : [ Service, ... ],
  "TargetPercent" : Number,
  "TimeoutMinutes" : Number,
  "Ungraceful" : EcsUngraceful
}

```

### YAML

```yaml

  CapacityMonitoringApproach:

  Services:
    - Service
  TargetPercent: Number
  TimeoutMinutes: Number
  Ungraceful:
    EcsUngraceful

```

## Properties

`CapacityMonitoringApproach`

The monitoring approach specified for the configuration, for example, `Most_Recent`.

_Required_: No

_Type_:

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Services`

The services specified for the configuration.

_Required_: Yes

_Type_: Array of [Service](aws-properties-arcregionswitch-plan-service.md)

_Minimum_: `2`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetPercent`

The target percentage specified for the configuration. The default is 100.

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

_Type_: [EcsUngraceful](aws-properties-arcregionswitch-plan-ecsungraceful.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Ec2Ungraceful

EcsUngraceful

All content copied from https://docs.aws.amazon.com/.
