This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::Scraper EksConfiguration

The `EksConfiguration` structure describes the connection to the Amazon EKS
cluster from which a scraper collects metrics.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClusterArn" : String,
  "SecurityGroupIds" : [ String, ... ],
  "SubnetIds" : [ String, ... ]
}

```

### YAML

```yaml

  ClusterArn: String
  SecurityGroupIds:
    - String
  SubnetIds:
    - String

```

## Properties

`ClusterArn`

ARN of the Amazon EKS cluster.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z]*:eks:[-a-z0-9]+:[0-9]{12}:cluster/.+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityGroupIds`

A list of the security group IDs for the Amazon EKS cluster VPC configuration.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetIds`

A list of subnet IDs for the Amazon EKS cluster VPC configuration.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Destination

RoleConfiguration

All content copied from https://docs.aws.amazon.com/.
