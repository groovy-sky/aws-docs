This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLiftStreams::StreamGroup LocationConfiguration

Configuration settings that define a stream group's stream capacity for a location. When configuring a location for the first time, you
must specify a numeric value for at least one of the two capacity types.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlwaysOnCapacity" : Integer,
  "LocationName" : String,
  "MaximumCapacity" : Integer,
  "OnDemandCapacity" : Integer,
  "TargetIdleCapacity" : Integer,
  "VpcTransitConfiguration" : VpcTransitConfiguration
}

```

### YAML

```yaml

  AlwaysOnCapacity: Integer
  LocationName: String
  MaximumCapacity: Integer
  OnDemandCapacity: Integer
  TargetIdleCapacity: Integer
  VpcTransitConfiguration:
    VpcTransitConfiguration

```

## Properties

`AlwaysOnCapacity`

This setting, if non-zero, indicates minimum streaming capacity
which is allocated to you and is never released back to the service. You pay for this
base level of capacity at all times, whether used or idle.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocationName`

A location's name. For example, `us-east-1`. For a complete list of locations that Amazon GameLift Streams supports, refer to [Regions, quotas, and limitations](../../../gameliftstreams/latest/developerguide/regions-quotas.md) in the _Amazon GameLift Streams Developer Guide_.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]+$`

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumCapacity`

This indicates the maximum capacity that the service can allocate
for you. Newly created streams may take a few minutes to start. Capacity is released back
to the service when idle. You pay for capacity that is allocated to you until it is released.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OnDemandCapacity`

This field is deprecated. Use MaximumCapacity instead. This parameter is ignored when MaximumCapacity is specified.

The streaming capacity that Amazon GameLift Streams can allocate in response to stream requests, and then de-allocate when the session has terminated.
This offers a cost control measure at the expense of a greater startup time (typically under 5 minutes). Default is 0 when you create a stream group or add a location.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetIdleCapacity`

This indicates idle capacity which the service pre-allocates and
holds for you in anticipation of future activity. This helps to insulate your users from
capacity-allocation delays. You pay for capacity which is held in this intentional idle
state.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcTransitConfiguration`

Configuration for connecting the stream group to resources in your Amazon VPC using a Transit Gateway. This setting is optional. If specified, Amazon GameLift Streams creates a
Transit Gateway to enable private network connectivity between the service VPC and your VPC. The VPC ID
cannot be changed after the stream group is created, but you can update the CIDR blocks.

_Required_: No

_Type_: [VpcTransitConfiguration](aws-properties-gameliftstreams-streamgroup-vpctransitconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DefaultApplication

VpcTransitConfiguration

All content copied from https://docs.aws.amazon.com/.
