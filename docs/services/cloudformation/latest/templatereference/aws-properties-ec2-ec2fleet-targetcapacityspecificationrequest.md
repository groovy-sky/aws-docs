This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::EC2Fleet TargetCapacitySpecificationRequest

Specifies the number of units to request for an EC2 Fleet. You can choose to set the
target capacity in terms of instances or a performance characteristic that is important to
your application workload, such as vCPUs, memory, or I/O. If the request type is
`maintain`, you can specify a target capacity of `0` and add
capacity later.

`TargetCapacitySpecificationRequest` is a property of the [AWS::EC2::EC2Fleet](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultTargetCapacityType" : String,
  "OnDemandTargetCapacity" : Integer,
  "SpotTargetCapacity" : Integer,
  "TargetCapacityUnitType" : String,
  "TotalTargetCapacity" : Integer
}

```

### YAML

```yaml

  DefaultTargetCapacityType: String
  OnDemandTargetCapacity: Integer
  SpotTargetCapacity: Integer
  TargetCapacityUnitType: String
  TotalTargetCapacity: Integer

```

## Properties

`DefaultTargetCapacityType`

The default target capacity type.

_Required_: No

_Type_: String

_Allowed values_: `on-demand | spot | reserved-capacity`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OnDemandTargetCapacity`

The number of On-Demand units to request.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpotTargetCapacity`

The number of Spot units to request.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetCapacityUnitType`

The unit for the target capacity. You can specify this parameter only when using
attributed-based instance type selection.

Default: `units` (the number of instances)

_Required_: No

_Type_: String

_Allowed values_: `vcpu | memory-mib | units`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TotalTargetCapacity`

The number of units to request, filled using the default target capacity type.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [TargetCapacitySpecificationRequest](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TargetCapacitySpecificationRequest.html) in the _Amazon EC2 API_
_Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TagSpecification

TotalLocalStorageGBRequest
