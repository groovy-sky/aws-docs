This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::FlowVpcInterface

The `AWS::MediaConnect::FlowVpcInterface` resource is a connection between your AWS Elemental MediaConnect flow and a virtual private cloud (VPC) that you
created using the Amazon Virtual Private Cloud service.

To avoid streaming your content over the public internet, you can add up to two VPC
interfaces to your flow and use those connections to transfer content between your VPC
and MediaConnect.

You can update an existing flow to add a VPC interface. If you haven’t created the
flow yet, you must create the flow with a temporary standard source by doing the
following:

1. Use CloudFormation to create a flow with a standard source that uses to the
    flow’s public IP address.

2. Use CloudFormation to create a VPC interface to add to this flow. This can
    also be done as part of the previous step.

3. After CloudFormation has created the flow and the VPC interface, update the
    source to point to the VPC interface that you created.

###### Important

The previous steps must be undone before the CloudFormation stack can be deleted. Because the source is manually updated in step 3, CloudFormation is not aware of this change. The source must be returned to a standard source before CloudFormation stack deletion.

###### Note

When configuring NDI outputs for your flow, define the VPC interface as a nested attribute within the `AWS::MediaConnect::Flow` resource. Do not use the top-level `AWS::MediaConnect::FlowVpcInterface` resource type to specify NDI configurations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaConnect::FlowVpcInterface",
  "Properties" : {
      "FlowArn" : String,
      "Name" : String,
      "RoleArn" : String,
      "SecurityGroupIds" : [ String, ... ],
      "SubnetId" : String
    }
}

```

### YAML

```yaml

Type: AWS::MediaConnect::FlowVpcInterface
Properties:
  FlowArn: String
  Name: String
  RoleArn: String
  SecurityGroupIds:
    - String
  SubnetId: String

```

## Properties

`FlowArn`

The Amazon Resource Name (ARN) of the flow.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*):mediaconnect:[a-z0-9-]+:[0-9]{12}:flow:[a-zA-Z0-9-]+:[a-zA-Z0-9_-]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name for the VPC interface. This name must be unique within the flow.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The Amazon Resource Name (ARN) of the role that you created when you set up MediaConnect as a trusted service.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*):iam::[0-9]{12}:role/[a-zA-Z0-9_+=,.@-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroupIds`

A virtual firewall to control inbound and outbound traffic.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetId`

The subnet IDs that you want to use for your VPC interface. A range of IP addresses in your VPC. When you create your VPC, you specify a range of IPv4 addresses for the VPC in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16. This is the primary CIDR block for your VPC. When you create a subnet for your VPC, you specify the CIDR block for the subnet, which is a subset of the VPC CIDR block. The subnets that you use across all VPC interfaces on the flow must be in the same Availability Zone as the flow.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the flow ARN and the name of the VPC interface. For
example:

`{ "Ref":
            "arn:aws:mediaconnect:us-east-1:111122223333:flow:1-23aBC45dEF67hiJ8-12AbC34DE5fG:BasketballGame|MyVPCInterface"
            }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`NetworkInterfaceIds`

The IDs of the network interfaces that MediaConnect created in your account.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcInterfaceAttachment

AWS::MediaConnect::Gateway

All content copied from https://docs.aws.amazon.com/.
