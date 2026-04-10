This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::Fleet

The `AWS::CodeBuild::Fleet` resource configures a compute fleet, a set of dedicated instances for your build environment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CodeBuild::Fleet",
  "Properties" : {
      "BaseCapacity" : Integer,
      "ComputeConfiguration" : ComputeConfiguration,
      "ComputeType" : String,
      "EnvironmentType" : String,
      "FleetProxyConfiguration" : ProxyConfiguration,
      "FleetServiceRole" : String,
      "FleetVpcConfig" : VpcConfig,
      "ImageId" : String,
      "Name" : String,
      "OverflowBehavior" : String,
      "ScalingConfiguration" : ScalingConfigurationInput,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CodeBuild::Fleet
Properties:
  BaseCapacity: Integer
  ComputeConfiguration:
    ComputeConfiguration
  ComputeType: String
  EnvironmentType: String
  FleetProxyConfiguration:
    ProxyConfiguration
  FleetServiceRole: String
  FleetVpcConfig:
    VpcConfig
  ImageId: String
  Name: String
  OverflowBehavior: String
  ScalingConfiguration:
    ScalingConfigurationInput
  Tags:
    - Tag

```

## Properties

`BaseCapacity`

The initial number of machines allocated to the compute ﬂeet, which deﬁnes the number of builds that can run in parallel.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComputeConfiguration`

The compute configuration of the compute fleet. This is only required if `computeType` is set to `ATTRIBUTE_BASED_COMPUTE` or `CUSTOM_INSTANCE_TYPE`.

_Required_: No

_Type_: [ComputeConfiguration](aws-properties-codebuild-fleet-computeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComputeType`

Information about the compute resources the compute fleet uses. Available values
include:

- `ATTRIBUTE_BASED_COMPUTE`: Specify the amount of vCPUs, memory, disk space, and the type of machine.

###### Note

If you use `ATTRIBUTE_BASED_COMPUTE`, you must define your attributes by using `computeConfiguration`. AWS CodeBuild
will select the cheapest instance that satisfies your specified attributes. For more information, see [Reserved capacity environment \
types](../../../codebuild/latest/userguide/build-env-ref-compute-types.md#environment-reserved-capacity.types) in the _AWS CodeBuild User Guide_.

- `BUILD_GENERAL1_SMALL`: Use up to 4 GiB memory and 2 vCPUs for
builds.

- `BUILD_GENERAL1_MEDIUM`: Use up to 8 GiB memory and 4 vCPUs for
builds.

- `BUILD_GENERAL1_LARGE`: Use up to 16 GiB memory and 8 vCPUs for
builds, depending on your environment type.

- `BUILD_GENERAL1_XLARGE`: Use up to 72 GiB memory and 36 vCPUs for
builds, depending on your environment type.

- `BUILD_GENERAL1_2XLARGE`: Use up to 144 GiB memory, 72 vCPUs, and
824 GB of SSD storage for builds. This compute type supports Docker images up to
100 GB uncompressed.

- `BUILD_LAMBDA_1GB`: Use up to 1 GiB memory for
builds. Only available for environment type `LINUX_LAMBDA_CONTAINER` and `ARM_LAMBDA_CONTAINER`.

- `BUILD_LAMBDA_2GB`: Use up to 2 GiB memory for
builds. Only available for environment type `LINUX_LAMBDA_CONTAINER` and `ARM_LAMBDA_CONTAINER`.

- `BUILD_LAMBDA_4GB`: Use up to 4 GiB memory for
builds. Only available for environment type `LINUX_LAMBDA_CONTAINER` and `ARM_LAMBDA_CONTAINER`.

- `BUILD_LAMBDA_8GB`: Use up to 8 GiB memory for
builds. Only available for environment type `LINUX_LAMBDA_CONTAINER` and `ARM_LAMBDA_CONTAINER`.

- `BUILD_LAMBDA_10GB`: Use up to 10 GiB memory for
builds. Only available for environment type `LINUX_LAMBDA_CONTAINER` and `ARM_LAMBDA_CONTAINER`.

If you use `BUILD_GENERAL1_SMALL`:

- For environment type `LINUX_CONTAINER`, you can use up to 4 GiB
memory and 2 vCPUs for builds.

- For environment type `LINUX_GPU_CONTAINER`, you can use up to 16
GiB memory, 4 vCPUs, and 1 NVIDIA A10G Tensor Core GPU for builds.

- For environment type `ARM_CONTAINER`, you can use up to 4 GiB
memory and 2 vCPUs on ARM-based processors for builds.

If you use `BUILD_GENERAL1_LARGE`:

- For environment type `LINUX_CONTAINER`, you can use up to 16 GiB
memory and 8 vCPUs for builds.

- For environment type `LINUX_GPU_CONTAINER`, you can use up to 255
GiB memory, 32 vCPUs, and 4 NVIDIA Tesla V100 GPUs for builds.

- For environment type `ARM_CONTAINER`, you can use up to 16 GiB
memory and 8 vCPUs on ARM-based processors for builds.

For more information, see [On-demand environment types](../../../codebuild/latest/userguide/build-env-ref-compute-types.md#environment.types)
in the _AWS CodeBuild User Guide._

_Required_: No

_Type_: String

_Allowed values_: `BUILD_GENERAL1_SMALL | BUILD_GENERAL1_MEDIUM | BUILD_GENERAL1_LARGE | BUILD_GENERAL1_XLARGE | BUILD_GENERAL1_2XLARGE | ATTRIBUTE_BASED_COMPUTE | CUSTOM_INSTANCE_TYPE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentType`

The environment type of the compute fleet.

- The environment type `ARM_CONTAINER` is available only in regions
US East (N. Virginia), US East (Ohio), US West (Oregon), EU (Ireland), Asia Pacific (Mumbai),
Asia Pacific (Tokyo), Asia Pacific (Singapore), Asia Pacific (Sydney), EU (Frankfurt),
and South America (São Paulo).

- The environment type `ARM_EC2` is available only in regions
US East (N. Virginia), US East (Ohio), US West (Oregon), EU (Ireland), EU (Frankfurt),
Asia Pacific (Tokyo), Asia Pacific (Singapore), Asia Pacific (Sydney), South America (São Paulo),
and Asia Pacific (Mumbai).

- The environment type `LINUX_CONTAINER` is available only in regions
US East (N. Virginia), US East (Ohio), US West (Oregon), EU (Ireland), EU (Frankfurt),
Asia Pacific (Tokyo), Asia Pacific (Singapore), Asia Pacific (Sydney), South America (São Paulo),
and Asia Pacific (Mumbai).

- The environment type `LINUX_EC2` is available only in regions
US East (N. Virginia), US East (Ohio), US West (Oregon), EU (Ireland), EU (Frankfurt),
Asia Pacific (Tokyo), Asia Pacific (Singapore), Asia Pacific (Sydney), South America (São Paulo),
and Asia Pacific (Mumbai).

- The environment type `LINUX_GPU_CONTAINER` is available only in
regions US East (N. Virginia), US East (Ohio), US West (Oregon), EU (Ireland),
EU (Frankfurt), Asia Pacific (Tokyo), and Asia Pacific (Sydney).

- The environment type `MAC_ARM` is available only in
regions US East (Ohio), US East (N. Virginia), US West (Oregon), Europe (Frankfurt),
and Asia Pacific (Sydney).

- The environment type `WINDOWS_EC2` is available only in regions
US East (N. Virginia), US East (Ohio), US West (Oregon), EU (Ireland), EU (Frankfurt),
Asia Pacific (Tokyo), Asia Pacific (Singapore), Asia Pacific (Sydney), South America (São Paulo),
and Asia Pacific (Mumbai).

- The environment type `WINDOWS_SERVER_2019_CONTAINER` is available only in regions
US East (N. Virginia), US East (Ohio), US West (Oregon), Asia Pacific (Sydney), Asia Pacific (Tokyo),
Asia Pacific (Mumbai) and EU (Ireland).

- The environment type `WINDOWS_SERVER_2022_CONTAINER` is available only in regions
US East (N. Virginia), US East (Ohio), US West (Oregon), EU (Ireland), EU (Frankfurt), Asia Pacific (Sydney),
Asia Pacific (Singapore), Asia Pacific (Tokyo), South America (São Paulo) and Asia Pacific (Mumbai).

For more information, see [Build environment compute types](../../../codebuild/latest/userguide/build-env-ref-compute-types.md) in the _AWS CodeBuild_
_user guide_.

_Required_: No

_Type_: String

_Allowed values_: `WINDOWS_SERVER_2019_CONTAINER | WINDOWS_SERVER_2022_CONTAINER | LINUX_CONTAINER | LINUX_GPU_CONTAINER | ARM_CONTAINER | MAC_ARM | LINUX_EC2 | ARM_EC2 | WINDOWS_EC2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FleetProxyConfiguration`

Information about the proxy configurations that apply network access control to your reserved capacity instances.

_Required_: No

_Type_: [ProxyConfiguration](aws-properties-codebuild-fleet-proxyconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FleetServiceRole`

The service role associated with the compute fleet. For more information, see [Allow a user to add a permission policy for a fleet service role](../../../codebuild/latest/userguide/auth-and-access-control-iam-identity-based-access-control.md#customer-managed-policies-example-permission-policy-fleet-service-role.html) in the _AWS CodeBuild User Guide_.

_Required_: No

_Type_: String

_Pattern_: `^(?:arn:)[a-zA-Z+-=,._:/@]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FleetVpcConfig`

Information about the VPC configuration that AWS CodeBuild accesses.

_Required_: No

_Type_: [VpcConfig](aws-properties-codebuild-fleet-vpcconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageId`

The Amazon Machine Image (AMI) of the compute fleet.

_Required_: No

_Type_: String

_Pattern_: `^((aws/codebuild/([A-Za-z0-9._-]+|ami/[A-Za-z0-9._-]+):[A-Za-z0-9._-]+)|ami-[a-z0-9]{1,1020})$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the compute fleet.

_Required_: No

_Type_: String

_Minimum_: `2`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OverflowBehavior`

The compute fleet overflow behavior.

- For overflow behavior `QUEUE`, your overflow builds need to wait on
the existing fleet instance to become available.

- For overflow behavior `ON_DEMAND`, your overflow builds run on CodeBuild on-demand.

###### Note

If you choose to set your overflow behavior to on-demand while creating a VPC-connected
fleet, make sure that you add the required VPC permissions to your project service role. For more
information, see [Example \
policy statement to allow CodeBuild access to AWS services required to create a VPC network interface](../../../codebuild/latest/userguide/auth-and-access-control-iam-identity-based-access-control.md#customer-managed-policies-example-create-vpc-network-interface).

_Required_: No

_Type_: String

_Allowed values_: `QUEUE | ON_DEMAND`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScalingConfiguration`

The scaling configuration of the compute fleet.

_Required_: No

_Type_: [ScalingConfigurationInput](aws-properties-codebuild-fleet-scalingconfigurationinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of tag key and value pairs associated with this compute fleet.

These tags are available for use by AWS services that support AWS CodeBuild compute fleet tags.

_Required_: No

_Type_: Array of [Tag](aws-properties-codebuild-fleet-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the AWS CodeBuild compute fleet.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the compute fleet.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS CodeBuild

ComputeConfiguration

All content copied from https://docs.aws.amazon.com/.
