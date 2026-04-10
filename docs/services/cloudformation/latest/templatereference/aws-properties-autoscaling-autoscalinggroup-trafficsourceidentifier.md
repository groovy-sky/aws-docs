This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::AutoScalingGroup TrafficSourceIdentifier

Identifying information for a traffic source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Identifier" : String,
  "Type" : String
}

```

### YAML

```yaml

  Identifier: String
  Type: String

```

## Properties

`Identifier`

Identifies the traffic source.

For Application Load Balancers, Gateway Load Balancers, Network Load Balancers, and VPC Lattice, this will be the Amazon Resource Name
(ARN) for a target group in this account and Region. For Classic Load Balancers, this will be the name
of the Classic Load Balancer in this account and Region.

For example:

- Application Load Balancer ARN:
`arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/my-targets/1234567890123456`

- Classic Load Balancer name: `my-classic-load-balancer`

- VPC Lattice ARN:
`arn:aws:vpc-lattice:us-west-2:123456789012:targetgroup/tg-1234567890123456`

To get the ARN of a target group for a Application Load Balancer, Gateway Load Balancer, or Network Load Balancer, or the name of a
Classic Load Balancer, use the Elastic Load Balancing [DescribeTargetGroups](../../../../reference/elasticloadbalancing/latest/apireference/api-describetargetgroups.md) and [DescribeLoadBalancers](../../../../reference/elasticloadbalancing/latest/apireference/api-describeloadbalancers.md) API operations.

To get the ARN of a target group for VPC Lattice, use the VPC Lattice [GetTargetGroup](../../../../reference/vpc-lattice/latest/apireference/api-gettargetgroup.md) API operation.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `1`

_Maximum_: `511`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Provides additional context for the value of `Identifier`.

The following lists the valid values:

- `elb` if `Identifier` is the name of a Classic Load Balancer.

- `elbv2` if `Identifier` is the ARN of an Application Load Balancer, Gateway Load Balancer,
or Network Load Balancer target group.

- `vpc-lattice` if `Identifier` is the ARN of a VPC Lattice
target group.

Required if the identifier is the name of a Classic Load Balancer.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `1`

_Maximum_: `511`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TotalLocalStorageGBRequest

VCpuCountRequest

All content copied from https://docs.aws.amazon.com/.
