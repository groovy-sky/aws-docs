This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::EC2Fleet

Specifies the configuration information to launch a fleet--or group--of instances. An
EC2 Fleet can launch multiple instance types across multiple Availability Zones, using the
On-Demand Instance, Reserved Instance, and Spot Instance purchasing models together. Using
EC2 Fleet, you can define separate On-Demand and Spot capacity targets, specify the
instance types that work best for your applications, and specify how Amazon EC2 should
distribute your fleet capacity within each purchasing model. For more information, see
[Launching an\
EC2 Fleet](../../../ec2/latest/userguide/ec2-fleet.md) in the _Amazon EC2 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::EC2Fleet",
  "Properties" : {
      "Context" : String,
      "ExcessCapacityTerminationPolicy" : String,
      "LaunchTemplateConfigs" : [ FleetLaunchTemplateConfigRequest, ... ],
      "OnDemandOptions" : OnDemandOptionsRequest,
      "ReplaceUnhealthyInstances" : Boolean,
      "ReservedCapacityOptions" : ReservedCapacityOptionsRequest,
      "SpotOptions" : SpotOptionsRequest,
      "TagSpecifications" : [ TagSpecification, ... ],
      "TargetCapacitySpecification" : TargetCapacitySpecificationRequest,
      "TerminateInstancesWithExpiration" : Boolean,
      "Type" : String,
      "ValidFrom" : String,
      "ValidUntil" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::EC2Fleet
Properties:
  Context: String
  ExcessCapacityTerminationPolicy: String
  LaunchTemplateConfigs:
    - FleetLaunchTemplateConfigRequest
  OnDemandOptions:
    OnDemandOptionsRequest
  ReplaceUnhealthyInstances: Boolean
  ReservedCapacityOptions:
    ReservedCapacityOptionsRequest
  SpotOptions:
    SpotOptionsRequest
  TagSpecifications:
    - TagSpecification
  TargetCapacitySpecification:
    TargetCapacitySpecificationRequest
  TerminateInstancesWithExpiration: Boolean
  Type: String
  ValidFrom: String
  ValidUntil: String

```

## Properties

`Context`

Reserved.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcessCapacityTerminationPolicy`

Indicates whether running instances should be terminated if the total target capacity of
the EC2 Fleet is decreased below the current size of the EC2 Fleet.

Supported only for fleets of type `maintain`.

_Required_: No

_Type_: String

_Allowed values_: `termination | no-termination`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LaunchTemplateConfigs`

The configuration for the EC2 Fleet.

_Required_: Yes

_Type_: Array of [FleetLaunchTemplateConfigRequest](aws-properties-ec2-ec2fleet-fleetlaunchtemplateconfigrequest.md)

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OnDemandOptions`

Describes the configuration of On-Demand Instances in an EC2 Fleet.

_Required_: No

_Type_: [OnDemandOptionsRequest](aws-properties-ec2-ec2fleet-ondemandoptionsrequest.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReplaceUnhealthyInstances`

Indicates whether EC2 Fleet should replace unhealthy Spot Instances. Supported only for
fleets of type `maintain`. For more information, see [EC2 Fleet\
health checks](../../../ec2/latest/userguide/manage-ec2-fleet.md#ec2-fleet-health-checks) in the _Amazon EC2 User Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReservedCapacityOptions`

Defines EC2 Fleet preferences for utilizing reserved capacity when DefaultTargetCapacityType is set to `reserved-capacity`.

_Required_: No

_Type_: [ReservedCapacityOptionsRequest](aws-properties-ec2-ec2fleet-reservedcapacityoptionsrequest.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SpotOptions`

Describes the configuration of Spot Instances in an EC2 Fleet.

_Required_: No

_Type_: [SpotOptionsRequest](aws-properties-ec2-ec2fleet-spotoptionsrequest.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TagSpecifications`

The key-value pair for tagging the EC2 Fleet request on creation. For more information, see
[Tag your resources](../../../ec2/latest/userguide/using-tags.md#tag-resources).

If the fleet type is `instant`, specify a resource type of `fleet`
to tag the fleet or `instance` to tag the instances at launch.

If the fleet type is `maintain` or `request`, specify a resource
type of `fleet` to tag the fleet. You cannot specify a resource type of
`instance`. To tag instances at launch, specify the tags in a [launch template](../../../ec2/latest/userguide/ec2-launch-templates.md#create-launch-template).

_Required_: No

_Type_: Array of [TagSpecification](aws-properties-ec2-ec2fleet-tagspecification.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TargetCapacitySpecification`

The number of units to request.

_Required_: Yes

_Type_: [TargetCapacitySpecificationRequest](aws-properties-ec2-ec2fleet-targetcapacityspecificationrequest.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TerminateInstancesWithExpiration`

Indicates whether running instances should be terminated when the EC2 Fleet expires.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The fleet type. The default value is `maintain`.

- `maintain` \- The EC2 Fleet places an asynchronous request for your desired
capacity, and continues to maintain your desired Spot capacity by replenishing
interrupted Spot Instances.

- `request` \- The EC2 Fleet places an asynchronous one-time request for your
desired capacity, but does submit Spot requests in alternative capacity pools if Spot
capacity is unavailable, and does not maintain Spot capacity if Spot Instances are
interrupted.

- `instant` \- The EC2 Fleet places a synchronous one-time request for your
desired capacity, and returns errors for any instances that could not be
launched.

For more information, see [EC2 Fleet\
request types](../../../ec2/latest/userguide/ec2-fleet-request-type.md) in the _Amazon EC2 User Guide_.

_Required_: No

_Type_: String

_Allowed values_: `maintain | request | instant`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ValidFrom`

The start date and time of the request, in UTC format (for example,
_YYYY_- _MM_- _DD_ T _HH_: _MM_: _SS_ Z).
The default is to start fulfilling the request immediately.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ValidUntil`

The end date and time of the request, in UTC format (for example,
_YYYY_- _MM_- _DD_ T _HH_: _MM_: _SS_ Z).
At this point, no new EC2 Fleet requests are placed or able to fulfill the request. If no value is specified, the request remains until you cancel it.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the fleet ID, such as
`fleet-1fe24079-d272-4023-8e7c-70e10784cb0e`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`FleetId`

The ID of the EC2 Fleet.

## See also

- [CreateFleet](../../../../reference/awsec2/latest/apireference/api-createfleet.md) in the _Amazon EC2 API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AcceleratorCountRequest

All content copied from https://docs.aws.amazon.com/.
