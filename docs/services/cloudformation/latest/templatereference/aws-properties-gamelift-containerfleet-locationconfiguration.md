This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::ContainerFleet LocationConfiguration

A remote location where a multi-location fleet can deploy game servers for game
hosting.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Location" : String,
  "LocationCapacity" : LocationCapacity,
  "PlayerGatewayStatus" : String,
  "StoppedActions" : [ String, ... ]
}

```

### YAML

```yaml

  Location: String
  LocationCapacity:
    LocationCapacity
  PlayerGatewayStatus: String
  StoppedActions:
    - String

```

## Properties

`Location`

An AWS Region code, such as `us-west-2`. For a list of supported Regions
and Local Zones, see [Amazon GameLift Servers service\
locations](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-regions.html) for managed hosting.

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

**Returned by:** [DescribeFleetCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetCapacity.html), [DescribeFleetLocationCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetLocationCapacity.html), [UpdateFleetCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_UpdateFleetCapacity.html)

_Required_: No

_Type_: [LocationCapacity](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-containerfleet-locationcapacity.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PlayerGatewayStatus`

The current status of player gateway in this location for this container fleet. Note, even if a container fleet has PlayerGatewayMode configured as `ENABLED`, player gateway might not be available in a specific location. For more information about locations where player gateway is supported, see [Amazon GameLift Servers service locations](https://docs.aws.amazon.com/gameliftservers/latest/developerguide/gamelift-regions.html).

Possible values include:

- `ENABLED` \-\- Player gateway is available for this container fleet location.

- `DISABLED` \-\- Player gateway is not available for this container fleet location.

_Required_: No

_Type_: String

_Allowed values_: `DISABLED | ENABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StoppedActions`

A list of fleet actions that have been suspended in the fleet location.

_Required_: No

_Type_: Array of String

_Allowed values_: `AUTO_SCALING`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LocationCapacity

LogConfiguration
