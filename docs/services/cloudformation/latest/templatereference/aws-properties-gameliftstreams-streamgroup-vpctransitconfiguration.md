This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLiftStreams::StreamGroup VpcTransitConfiguration

Configuration for connecting a stream group location to resources in your Amazon VPC using
a Transit Gateway. When you specify a VPC transit configuration, Amazon GameLift Streams creates a Transit Gateway and
shares it with your account using Resource Access Manager. After the stream group is active, you must
complete the setup by accepting the resource share, creating a VPC attachment, and
configuring routing.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Ipv4CidrBlocks" : [ String, ... ],
  "VpcId" : String
}

```

### YAML

```yaml

  Ipv4CidrBlocks:
    - String
  VpcId: String

```

## Properties

`Ipv4CidrBlocks`

A list of IPv4 CIDR blocks in your VPC that you want the stream group to be able to
access. You can specify up to 5 CIDR blocks. The CIDR blocks must be valid subsets of
the VPC's CIDR blocks and cannot overlap with the service VPC CIDR block.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

The ID of the Amazon VPC that you want to connect to the stream group. The VPC must be in
the same AWS account as the stream group. This value cannot be changed after the stream group is created.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LocationConfiguration

Next

All content copied from https://docs.aws.amazon.com/.
