This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Cluster BrokerNodeGroupInfo

Describes the setup to be used for the broker nodes in the cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BrokerAZDistribution" : String,
  "ClientSubnets" : [ String, ... ],
  "ConnectivityInfo" : ConnectivityInfo,
  "InstanceType" : String,
  "SecurityGroups" : [ String, ... ],
  "StorageInfo" : StorageInfo
}

```

### YAML

```yaml

  BrokerAZDistribution: String
  ClientSubnets:
    - String
  ConnectivityInfo:
    ConnectivityInfo
  InstanceType: String
  SecurityGroups:
    - String
  StorageInfo:
    StorageInfo

```

## Properties

`BrokerAZDistribution`

This parameter is currently not in use.

_Required_: No

_Type_: String

_Minimum_: `6`

_Maximum_: `9`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClientSubnets`

The list of subnets to connect to in the client virtual private cloud (VPC). Amazon creates elastic network interfaces (ENIs) inside these subnets. Client applications use ENIs to produce and consume data.

If you use the US West (N. California) Region, specify exactly two subnets. For other Regions where Amazon MSK is available, you can specify either two or three subnets. The subnets that you specify must be in distinct Availability Zones. When you create a cluster, Amazon MSK distributes the broker nodes evenly across the subnets that you specify.

Client subnets can't occupy the Availability Zone with ID `use1-az3`.

_Required_: Yes

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConnectivityInfo`

Information about the cluster's connectivity setting.

_Required_: No

_Type_: [ConnectivityInfo](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-msk-cluster-connectivityinfo.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceType`

The type of Amazon EC2 instances to use for brokers. Depending on the [broker type](https://docs.aws.amazon.com/msk/latest/developerguide/broker-instance-types.html), Amazon MSK supports the following broker sizes:

**Standard broker sizes**

- kafka.t3.small

###### Note

You can't select the kafka.t3.small instance type when the metadata mode is KRaft.

- kafka.m5.large, kafka.m5.xlarge, kafka.m5.2xlarge, kafka.m5.4xlarge, kafka.m5.8xlarge, kafka.m5.12xlarge, kafka.m5.16xlarge, kafka.m5.24xlarge

- kafka.m7g.large, kafka.m7g.xlarge, kafka.m7g.2xlarge, kafka.m7g.4xlarge, kafka.m7g.8xlarge, kafka.m7g.12xlarge, kafka.m7g.16xlarge

**Express broker sizes**

- express.m7g.large, express.m7g.xlarge, express.m7g.2xlarge, express.m7g.4xlarge, express.m7g.8xlarge, express.m7g.12xlarge, express.m7g.16xlarge

###### Note

Some broker sizes might not be available in certian AWS Regions. See the updated [Pricing tools](https://aws.amazon.com/msk/pricing) section on the Amazon MSK pricing page for the latest list of available instances by Region.

_Required_: Yes

_Type_: String

_Minimum_: `5`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroups`

The security groups to associate with the ENIs in order to specify who can connect to and communicate with the Amazon MSK cluster. If you don't specify a security group, Amazon MSK uses the default security group associated with the VPC. If you specify security groups that were shared with you, you must ensure that you have permissions to them. Specifically, you need the `ec2:DescribeSecurityGroups` permission.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageInfo`

Contains information about storage volumes attached to Amazon MSK broker nodes.

_Required_: No

_Type_: [StorageInfo](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-msk-cluster-storageinfo.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

BrokerLogs

ClientAuthentication
