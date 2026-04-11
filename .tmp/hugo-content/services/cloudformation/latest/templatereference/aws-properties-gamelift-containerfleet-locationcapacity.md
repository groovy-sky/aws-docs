This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::ContainerFleet LocationCapacity

Current resource capacity settings for managed EC2 fleets and managed container fleets. For
multi-location fleets, location values might refer to a fleet's remote location or its
home Region.

**Returned by:** [DescribeFleetCapacity](../../../../reference/gamelift/latest/apireference/api-describefleetcapacity.md), [DescribeFleetLocationCapacity](../../../../reference/gamelift/latest/apireference/api-describefleetlocationcapacity.md), [UpdateFleetCapacity](../../../../reference/gamelift/latest/apireference/api-updatefleetcapacity.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DesiredEC2Instances" : Integer,
  "ManagedCapacityConfiguration" : ManagedCapacityConfiguration,
  "MaxSize" : Integer,
  "MinSize" : Integer
}

```

### YAML

```yaml

  DesiredEC2Instances: Integer
  ManagedCapacityConfiguration:
    ManagedCapacityConfiguration
  MaxSize: Integer
  MinSize: Integer

```

## Properties

`DesiredEC2Instances`

The number of Amazon EC2 instances you want to maintain in the specified fleet location.
This value must fall between the minimum and maximum size limits. Changes in desired
instance value can take up to 1 minute to be reflected when viewing the fleet's capacity
settings.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManagedCapacityConfiguration`

Use ManagedCapacityConfiguration with the "SCALE\_TO\_AND\_FROM\_ZERO" ZeroCapacityStrategy to enable Amazon
GameLift Servers to fully manage the MinSize value, switching between 0 and 1 based on game session
activity. This is ideal for eliminating compute costs during periods of no game activity.
It is particularly beneficial during development when you're away from your desk, iterating on builds
for extended periods, in production environments serving low-traffic locations, or for games with long,
predictable downtime windows. By automatically managing capacity between 0 and 1 instances, you avoid paying
for idle instances while maintaining the ability to serve game sessions when demand arrives. Note that while
scale-out is triggered immediately upon receiving a game session request, actual game session availability
depends on your server process startup time, so this approach works best with multi-location Fleets where
cold-start latency is tolerable. With a "MANUAL" ZeroCapacityStrategy Amazon GameLift Servers will not
modify Fleet MinSize values automatically and will not scale out from zero instances in response to game
sessions.

_Required_: No

_Type_: [ManagedCapacityConfiguration](aws-properties-gamelift-containerfleet-managedcapacityconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxSize`

The maximum number of instances that are allowed in the specified fleet location. If
this parameter is not set, the default is 1.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinSize`

The minimum number of instances that are allowed in the specified fleet location. If
this parameter is not set, the default is 0. This parameter's value will be ignored when using a
ManagedCapacityConfiguration where ZeroCapacityStrategy has a value of SCALE\_TO\_AND\_FROM\_ZERO.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IpPermission

LocationConfiguration

All content copied from https://docs.aws.amazon.com/.
