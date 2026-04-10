This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::CapacityProvider CapacityProviderVpcConfig

VPC configuration that specifies the network settings for compute instances managed by the capacity provider.

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

A list of security group IDs that control network access for compute instances managed by the capacity provider.

_Required_: Yes

_Type_: Array of String

_Minimum_: `0 | 0`

_Maximum_: `1024 | 5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetIds`

A list of subnet IDs where the capacity provider launches compute instances.

_Required_: Yes

_Type_: Array of String

_Minimum_: `0 | 1`

_Maximum_: `1024 | 16`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CapacityProviderScalingConfig

InstanceRequirements

All content copied from https://docs.aws.amazon.com/.
