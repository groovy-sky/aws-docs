This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::CapacityProvider ManagedInstancesNetworkConfiguration

The network configuration for Amazon ECS Managed Instances. This specifies the VPC
subnets and security groups that instances use for network connectivity. Amazon ECS
Managed Instances support multiple network modes including `awsvpc`
(instances receive ENIs for task isolation), `host` (instances share network
namespace with tasks), and `none` (no external network connectivity),
ensuring backward compatibility for migrating workloads from Fargate or Amazon
EC2.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecurityGroups" : [ String, ... ],
  "Subnets" : [ String, ... ]
}

```

### YAML

```yaml

  SecurityGroups:
    - String
  Subnets:
    - String

```

## Properties

`SecurityGroups`

The list of security group IDs to apply to Amazon ECS Managed Instances. These
security groups control the network traffic allowed to and from the instances.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subnets`

The list of subnet IDs where Amazon ECS can launch Amazon ECS Managed Instances.
Instances are distributed across the specified subnets for high availability. All
subnets must be in the same VPC.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ManagedInstancesLocalStorageConfiguration

ManagedInstancesProvider

All content copied from https://docs.aws.amazon.com/.
