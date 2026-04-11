This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LaunchTemplate Placement

Specifies the placement of an instance.

`Placement` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](../userguide/aws-properties-ec2-launchtemplate-launchtemplatedata.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Affinity" : String,
  "AvailabilityZone" : String,
  "GroupId" : String,
  "GroupName" : String,
  "HostId" : String,
  "HostResourceGroupArn" : String,
  "PartitionNumber" : Integer,
  "SpreadDomain" : String,
  "Tenancy" : String
}

```

### YAML

```yaml

  Affinity: String
  AvailabilityZone: String
  GroupId: String
  GroupName: String
  HostId: String
  HostResourceGroupArn: String
  PartitionNumber: Integer
  SpreadDomain: String
  Tenancy: String

```

## Properties

`Affinity`

The affinity setting for an instance on a Dedicated Host.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailabilityZone`

The Availability Zone for the instance.

Either `AvailabilityZone` or `AvailabilityZoneId` can be specified, but not both

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupId`

The Group Id of a placement group. You must specify the Placement Group **Group Id** to launch an instance in a shared placement
group.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupName`

The name of the placement group for the instance.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HostId`

The ID of the Dedicated Host for the instance.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HostResourceGroupArn`

The ARN of the host resource group in which to launch the instances. If you specify a
host resource group ARN, omit the **Tenancy** parameter or
set it to `host`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PartitionNumber`

The number of the partition the instance should launch in. Valid only if the placement
group strategy is set to `partition`.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpreadDomain`

Reserved for future use.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tenancy`

The tenancy of the instance. An instance with a tenancy of dedicated runs on
single-tenant hardware.

_Required_: No

_Type_: String

_Allowed values_: `default | dedicated | host`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [LaunchTemplatePlacementRequest](../../../../reference/awsec2/latest/apireference/api-launchtemplateplacementrequest.md) in the _Amazon EC2 API_
_Reference_

- [Create a launch template using advanced settings](../../../autoscaling/ec2/userguide/advanced-settings-for-your-launch-template.md) in the _Amazon EC2 Auto Scaling User Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NetworkPerformanceOptions

PrivateDnsNameOptions

All content copied from https://docs.aws.amazon.com/.
