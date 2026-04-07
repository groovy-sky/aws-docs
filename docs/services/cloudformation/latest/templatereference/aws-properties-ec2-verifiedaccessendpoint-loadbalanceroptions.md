This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VerifiedAccessEndpoint LoadBalancerOptions

Describes the load balancer options when creating an AWS Verified Access endpoint using the
`load-balancer` type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LoadBalancerArn" : String,
  "Port" : Integer,
  "PortRanges" : [ PortRange, ... ],
  "Protocol" : String,
  "SubnetIds" : [ String, ... ]
}

```

### YAML

```yaml

  LoadBalancerArn: String
  Port: Integer
  PortRanges:
    - PortRange
  Protocol: String
  SubnetIds:
    - String

```

## Properties

`LoadBalancerArn`

The ARN of the load balancer.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Port`

The IP port number.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PortRanges`

The port ranges.

_Required_: No

_Type_: Array of [PortRange](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-verifiedaccessendpoint-portrange.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The IP protocol.

_Required_: No

_Type_: String

_Allowed values_: `http | https | tcp`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetIds`

The IDs of the subnets. You can specify only one subnet per Availability Zone.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CidrOptions

NetworkInterfaceOptions
