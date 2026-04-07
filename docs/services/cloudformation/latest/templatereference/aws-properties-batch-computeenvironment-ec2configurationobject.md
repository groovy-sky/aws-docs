This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::ComputeEnvironment Ec2ConfigurationObject

Provides information used to select Amazon Machine Images (AMIs) for instances in the
compute environment. If `Ec2Configuration` isn't specified, the default is
`ECS_AL2023` ( [Amazon ECS-optimized Amazon Linux 2023](../../../amazonecs/latest/developerguide/ecs-optimized-ami.md)) for EC2 (ECS) compute environments and `EKS_AL2023` ( [Amazon EKS-optimized Amazon Linux 2023\
AMI](https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html)) for EKS compute environments.

###### Note

This object isn't applicable to jobs that are running on Fargate resources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ImageIdOverride" : String,
  "ImageKubernetesVersion" : String,
  "ImageType" : String
}

```

### YAML

```yaml

  ImageIdOverride: String
  ImageKubernetesVersion: String
  ImageType: String

```

## Properties

`ImageIdOverride`

The AMI ID used for instances launched in the compute environment that match the image type.
This setting overrides the `imageId` set in the `computeResource`
object.

###### Note

The AMI that you choose for a compute environment must match the architecture of the instance types that
you intend to use for that compute environment. For example, if your compute environment uses A1 instance types,
the compute resource AMI that you choose must support ARM instances. Amazon ECS vends both x86 and ARM versions of the
Amazon ECS-optimized Amazon Linux 2023 AMI. For more information, see [Amazon ECS-optimized\
Amazon Linux 2023 AMI](../../../amazonecs/latest/developerguide/ecs-optimized-ami.md#ecs-optimized-ami-linux-variants.html)
in the _Amazon Elastic Container Service Developer Guide_.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`ImageKubernetesVersion`

The Kubernetes version for the compute environment. If you don't specify a value, the latest
version that AWS Batch supports is used.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`ImageType`

The image type to match with the instance type to select an AMI. The supported values are
different for `ECS` and `EKS` resources.

ECS

If the `imageIdOverride` parameter isn't specified, then a recent [Amazon ECS-optimized Amazon Linux 2023 AMI](../../../amazonecs/latest/developerguide/ecs-optimized-ami.md) ( `ECS_AL2023`) is used. If a new image type is
specified in an update, but neither an `imageId` nor a `imageIdOverride`
parameter is specified, then the latest Amazon ECS optimized AMI for that image type that's
supported by AWS Batch is used.

###### Important

AWS is ending support for Amazon ECS Amazon Linux 2-optimized and accelerated AMIs on June 30,
2026\. On January 12, 2026, AWS Batch changed the default AMI for new Amazon ECS compute environments
from Amazon Linux 2 to Amazon Linux 2023. Effective June 30, 2026, AWS Batch will block creation
of new Amazon ECS compute environments using Batch-provided Amazon Linux 2 AMIs. We strongly recommend
migrating your existing AWS Batch Amazon ECS compute environments to Amazon Linux 2023 prior to June 30,
2026\. For more information on upgrading from AL2 to AL2023, see [How to migrate\
from ECS AL2 to ECS AL2023](https://docs.aws.amazon.com/batch/latest/userguide/ecs-migration-2023.html) in the _AWS Batch User Guide_.

ECS\_AL2

[Amazon Linux\
2](../../../amazonecs/latest/developerguide/ecs-optimized-ami.md): Used for non-GPU instance families.

ECS\_AL2\_NVIDIA

[Amazon Linux 2\
(GPU)](../../../amazonecs/latest/developerguide/ecs-optimized-ami.md#gpuami): Used for GPU instance families (for example `P4` and
`G4`) and non AWS Graviton-based instance types.

ECS\_AL2023

[Amazon Linux 2023](../../../amazonecs/latest/developerguide/ecs-optimized-ami.md): Default
for all non-GPU instance families.

###### Note

Amazon Linux 2023 does not support `A1` instances.

ECS\_AL2023\_NVIDIA

[Amazon Linux 2023\
(GPU)](../../../amazonecs/latest/developerguide/ecs-optimized-ami.md#gpuami): Default for all GPU instance families and can be used for all non AWS Graviton-based instance types.

###### Note

ECS\_AL2023\_NVIDIA doesn't support `p3` and `g3` instance types.

EKS

If the `imageIdOverride` parameter isn't specified, then a recent [Amazon EKS-optimized Amazon Linux 2023\
AMI](https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html) ( `EKS_AL2023`) is used. If a new image type is specified in an update,
but neither an `imageId` nor a `imageIdOverride` parameter is specified,
then the latest Amazon EKS optimized AMI for that image type that AWS Batch supports is used.

###### Important

Amazon Linux 2023 AMIs are the
default on AWS Batch for Amazon EKS.

AWS ended support for Amazon EKS AL2-optimized and AL2-accelerated AMIs on November 26,
2025\. AWS Batch Amazon EKS compute environments using Amazon Linux 2 will no longer receive software
updates, security patches, or bug fixes from AWS. We recommend migrating to Amazon Linux
2023\. For more information on upgrading from AL2 to AL2023, see [How to upgrade from EKS AL2 to EKS AL2023](https://docs.aws.amazon.com/batch/latest/userguide/eks-migration-2023.html) in the _AWS Batch User Guide_.

EKS\_AL2

[Amazon\
Linux 2](https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html): Used for non-GPU instance families.

EKS\_AL2\_NVIDIA

[Amazon\
Linux 2 (accelerated)](https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html): Used for GPU instance families (for example,
`P4` and `G4`) and can be used for all non AWS Graviton-based
instance types.

EKS\_AL2023

[Amazon\
Linux 2023](https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html): Default for non-GPU instance families.

###### Note

Amazon Linux 2023 does not support `A1` instances.

EKS\_AL2023\_NVIDIA

[Amazon\
Linux 2023 (accelerated)](https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html): Default for GPU instance families and can be used for all non AWS
Graviton-based instance types.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ComputeScalingPolicy

EksConfiguration
