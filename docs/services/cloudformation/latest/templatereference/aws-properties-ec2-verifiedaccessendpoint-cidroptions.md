This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VerifiedAccessEndpoint CidrOptions

Describes the CIDR options for a Verified Access endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Cidr" : String,
  "PortRanges" : [ PortRange, ... ],
  "Protocol" : String,
  "SubnetIds" : [ String, ... ]
}

```

### YAML

```yaml

  Cidr: String
  PortRanges:
    - PortRange
  Protocol: String
  SubnetIds:
    - String

```

## Properties

`Cidr`

The CIDR.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PortRanges`

The port ranges.

_Required_: No

_Type_: Array of [PortRange](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-verifiedaccessendpoint-portrange.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The protocol.

_Required_: No

_Type_: String

_Allowed values_: `http | https | tcp`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetIds`

The IDs of the subnets.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::VerifiedAccessEndpoint

LoadBalancerOptions
