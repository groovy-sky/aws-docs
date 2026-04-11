This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Replicator KafkaClusterClientVpcConfig

Details of an Amazon VPC which has network connectivity to the Apache Kafka cluster.

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

The security groups to attach to the ENIs for the broker nodes.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `16`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetIds`

The list of subnets in the client VPC to connect to.

_Required_: Yes

_Type_: Array of String

_Minimum_: `2`

_Maximum_: `3`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KafkaCluster

ReplicationInfo

All content copied from https://docs.aws.amazon.com/.
