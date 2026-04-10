This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Flow VpcInterface

The details of a VPC interface.

###### Note

When configuring VPC interfaces for NDI outputs, keep in mind the following:

- VPC interfaces must be defined as nested attributes within the
`AWS::MediaConnect::Flow` resource, and not within the top-level
`AWS::MediaConnect::FlowVpcInterface` resource.

- There's a maximum limit of three VPC interfaces for each flow. If you've already reached
this limit, you can't update the flow to use a different
VPC interface without first removing an existing one.

To update your VPC interfaces in this scenario, you must first remove the VPC
interface that’s not being used. Next, add the new VPC interfaces. Lastly, update
the `VpcInterfaceAdapter` in the `NDIConfig` property. These
changes must be performed as separate manual operations and cannot be done through
a single template update.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "NetworkInterfaceIds" : [ String, ... ],
  "NetworkInterfaceType" : String,
  "RoleArn" : String,
  "SecurityGroupIds" : [ String, ... ],
  "SubnetId" : String,
  "Tags" : [ Tag, ... ]
}

```

### YAML

```yaml

  Name: String
  NetworkInterfaceIds:
    - String
  NetworkInterfaceType: String
  RoleArn: String
  SecurityGroupIds:
    - String
  SubnetId: String
  Tags:
    - Tag

```

## Properties

`Name`

Immutable and has to be a unique against other VpcInterfaces in this Flow.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkInterfaceIds`

IDs of the network interfaces created in customer's account by MediaConnect.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkInterfaceType`

The type of network interface.

_Required_: No

_Type_: String

_Allowed values_: `ena | efa`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

A role Arn MediaConnect can assume to create ENIs in your account.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*):iam::[0-9]{12}:role/[a-zA-Z0-9_+=,.@-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroupIds`

Security Group IDs to be used on ENI.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetId`

Subnet must be in the AZ of the Flow.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-mediaconnect-flow-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VideoMonitoringSetting

VpcInterfaceAttachment

All content copied from https://docs.aws.amazon.com/.
