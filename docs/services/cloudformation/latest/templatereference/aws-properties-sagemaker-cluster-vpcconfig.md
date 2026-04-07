This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Cluster VpcConfig

Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources
have access to. You can control access to and from your resources by configuring a VPC.
For more information, see [Give SageMaker Access to\
Resources in your Amazon VPC](https://docs.aws.amazon.com/sagemaker/latest/dg/infrastructure-give-access.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecurityGroupIds" : [ String, ... ],
  "Subnets" : [ String, ... ]
}

```

### YAML

```yaml

  SecurityGroupIds:
    - String
  Subnets:
    - String

```

## Properties

`SecurityGroupIds`

The VPC security group IDs, in the form `sg-xxxxxxxx`. Specify the security
groups for the VPC that is specified in the `Subnets` field.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `32 | 5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Subnets`

The ID of the subnets in the VPC to which you want to connect your training job or
model. For information about the availability of specific instance types, see [Supported\
Instance Types and Availability Zones](https://docs.aws.amazon.com/sagemaker/latest/dg/instance-types-az.html).

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `32 | 16`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TieredStorageConfig

AWS::SageMaker::CodeRepository
