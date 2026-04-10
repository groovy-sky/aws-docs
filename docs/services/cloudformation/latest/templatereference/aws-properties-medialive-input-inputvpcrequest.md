This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Input InputVpcRequest

Settings that apply only if the input is an push input where the source is on Amazon VPC.

The parent of this entity is Input.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecurityGroupIds" : [ String, ... ],
  "SubnetIds" : [ String, ... ]
}

```

### YAML

```yaml

  SecurityGroupIds:
    - String
  SubnetIds:
    - String

```

## Properties

`SecurityGroupIds`

The list of up to five VPC security group IDs to attach to the input VPC network interfaces. The security groups
require subnet IDs. If none are specified, MediaLive uses the VPC default security group.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetIds`

The list of two VPC subnet IDs from the same VPC. You must associate subnet IDs to two unique Availability
Zones.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InputSourceRequest

MediaConnectFlowRequest

All content copied from https://docs.aws.amazon.com/.
