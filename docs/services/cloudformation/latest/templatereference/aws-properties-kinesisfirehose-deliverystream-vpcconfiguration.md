This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream VpcConfiguration

The details of the VPC of the Amazon ES destination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RoleARN" : String,
  "SecurityGroupIds" : [ String, ... ],
  "SubnetIds" : [ String, ... ]
}

```

### YAML

```yaml

  RoleARN: String
  SecurityGroupIds:
    - String
  SubnetIds:
    - String

```

## Properties

`RoleARN`

The ARN of the IAM role that you want the delivery stream to use to create endpoints
in the destination VPC. You can use your existing Kinesis Data Firehose delivery role or
you can specify a new role. In either case, make sure that the role trusts the Kinesis Data
Firehose service principal and that it grants the following permissions:

- `ec2:DescribeVpcs`

- `ec2:DescribeVpcAttribute`

- `ec2:DescribeSubnets`

- `ec2:DescribeSecurityGroups`

- `ec2:DescribeNetworkInterfaces`

- `ec2:CreateNetworkInterface`

- `ec2:CreateNetworkInterfacePermission`

- `ec2:DeleteNetworkInterface`

If you revoke these permissions after you create the delivery stream, Kinesis Data
Firehose can't scale out by creating more ENIs when necessary. You might therefore see a
degradation in performance.

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityGroupIds`

The IDs of the security groups that you want Kinesis Data Firehose to use when it
creates ENIs in the VPC of the Amazon ES destination. You can use the same security group
that the Amazon ES domain uses or different ones. If you specify different security groups
here, ensure that they allow outbound HTTPS traffic to the Amazon ES domain's security
group. Also ensure that the Amazon ES domain's security group allows HTTPS traffic from the
security groups specified here. If you use the same security group for both your delivery
stream and the Amazon ES domain, make sure the security group inbound rule allows HTTPS
traffic.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `1024 | 5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetIds`

The IDs of the subnets that Kinesis Data Firehose uses to create ENIs in the VPC of
the Amazon ES destination. Make sure that the routing tables and inbound and outbound rules
allow traffic to flow from the subnets whose IDs are specified here to the subnets that
have the destination Amazon ES endpoints. Kinesis Data Firehose creates at least one ENI in
each of the subnets that are specified here. Do not delete or modify these ENIs.

The number of ENIs that Kinesis Data Firehose creates in the subnets specified here
scales up and down automatically based on throughput. To enable Kinesis Data Firehose to
scale up the number of ENIs to match throughput, ensure that you have sufficient quota. To
help you calculate the quota you need, assume that Kinesis Data Firehose can create up to
three ENIs for this delivery stream for each of the subnets specified here.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `1024 | 16`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Next

All content copied from https://docs.aws.amazon.com/.
