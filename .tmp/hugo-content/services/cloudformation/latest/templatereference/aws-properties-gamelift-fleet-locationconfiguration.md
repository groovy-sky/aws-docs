This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::Fleet LocationConfiguration

A remote location where a multi-location fleet can deploy game servers for game
hosting.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Location" : String,
  "LocationCapacity" : LocationCapacity,
  "PlayerGatewayStatus" : String
}

```

### YAML

```yaml

  Location: String
  LocationCapacity:
    LocationCapacity
  PlayerGatewayStatus: String

```

## Properties

`Location`

An AWS Region code, such as `us-west-2`. For a list of supported Regions
and Local Zones, see [Amazon GameLift Servers service\
locations](../../../gamelift/latest/developerguide/gamelift-regions.md) for managed hosting.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9\-]+`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocationCapacity`

Current resource capacity settings for managed EC2 fleets and managed container fleets. For
multi-location fleets, location values might refer to a fleet's remote location or its
home Region.

**Returned by:** [DescribeFleetCapacity](../../../../reference/gamelift/latest/apireference/api-describefleetcapacity.md), [DescribeFleetLocationCapacity](../../../../reference/gamelift/latest/apireference/api-describefleetlocationcapacity.md), [UpdateFleetCapacity](../../../../reference/gamelift/latest/apireference/api-updatefleetcapacity.md)

_Required_: No

_Type_: [LocationCapacity](aws-properties-gamelift-fleet-locationcapacity.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PlayerGatewayStatus`

The current status of player gateway in this location for this fleet. Note, even if a fleet has PlayerGatewayMode configured as `ENABLED`, player gateway might not be available in a specific location. For more information about locations where player gateway is supported, see [Amazon GameLift Servers service locations](../../../gameliftservers/latest/developerguide/gamelift-regions.md).

Possible values include:

- `ENABLED` \-\- Player gateway is available for this fleet location.

- `DISABLED` \-\- Player gateway is not available for this fleet location.

_Required_: No

_Type_: String

_Allowed values_: `DISABLED | ENABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LocationCapacity

ManagedCapacityConfiguration

All content copied from https://docs.aws.amazon.com/.
