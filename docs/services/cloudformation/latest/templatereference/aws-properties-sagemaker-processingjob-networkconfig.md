This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ProcessingJob NetworkConfig

Networking options for a job, such as network traffic encryption between containers,
whether to allow inbound and outbound network calls to and from containers, and the VPC
subnets and security groups to use for VPC-enabled jobs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnableInterContainerTrafficEncryption" : Boolean,
  "EnableNetworkIsolation" : Boolean,
  "VpcConfig" : VpcConfig
}

```

### YAML

```yaml

  EnableInterContainerTrafficEncryption: Boolean
  EnableNetworkIsolation: Boolean
  VpcConfig:
    VpcConfig

```

## Properties

`EnableInterContainerTrafficEncryption`

Whether to encrypt all communications between distributed processing jobs. Choose
`True` to encrypt communications. Encryption provides greater security
for distributed processing jobs, but the processing might take longer.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnableNetworkIsolation`

Whether to allow inbound and outbound network calls to and from the containers used for
the processing job.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcConfig`

Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources
have access to. You can control access to and from your resources by configuring a VPC. For more information, see
[Give SageMaker Access to\
Resources in your Amazon VPC](../../../sagemaker/latest/dg/infrastructure-give-access.md).

_Required_: No

_Type_: [VpcConfig](aws-properties-sagemaker-processingjob-vpcconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FeatureStoreOutput

ProcessingInputsObject

All content copied from https://docs.aws.amazon.com/.
