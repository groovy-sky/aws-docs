This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel VpcOutputSettings

Settings to enable VPC mode in the channel, so that the endpoints for all outputs are in your VPC.

This entity is at the top level in the channel.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PublicAddressAllocationIds" : [ String, ... ],
  "SecurityGroupIds" : [ String, ... ],
  "SubnetIds" : [ String, ... ]
}

```

### YAML

```yaml

  PublicAddressAllocationIds:
    - String
  SecurityGroupIds:
    - String
  SubnetIds:
    - String

```

## Properties

`PublicAddressAllocationIds`

List of public address allocation IDs to associate with ENIs that will be created in Output VPC. Must specify
one for SINGLE\_PIPELINE, two for STANDARD channels

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityGroupIds`

A list of up to 5 EC2 VPC security group IDs to attach to the Output VPC network interfaces.
If none are specified then the VPC default security group will be used

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetIds`

A list of VPC subnet IDs from the same VPC.
If STANDARD channel, subnet IDs must be mapped to two unique availability zones (AZ).

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VideoSelectorSettings

WavSettings
