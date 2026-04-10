This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Space ResourceSpec

Specifies the ARN's of a SageMaker AI image and SageMaker AI image version, and the instance type that
the version runs on.

###### Note

When both `SageMakerImageVersionArn` and `SageMakerImageArn` are passed, `SageMakerImageVersionArn` is used. Any updates to `SageMakerImageArn` will not take effect if `SageMakerImageVersionArn` already exists in the `ResourceSpec` because `SageMakerImageVersionArn` always takes precedence. To clear the value set for `SageMakerImageVersionArn`, pass `None` as the value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InstanceType" : String,
  "LifecycleConfigArn" : String,
  "SageMakerImageArn" : String,
  "SageMakerImageVersionArn" : String
}

```

### YAML

```yaml

  InstanceType: String
  LifecycleConfigArn: String
  SageMakerImageArn: String
  SageMakerImageVersionArn: String

```

## Properties

`InstanceType`

The instance type that the image version runs on.

###### Note

**JupyterServer apps** only support the `system` value.

For **KernelGateway apps**, the `system`
value is translated to `ml.t3.medium`. KernelGateway apps also support all other values for available
instance types.

_Required_: No

_Type_: String

_Allowed values_: `system | ml.t3.micro | ml.t3.small | ml.t3.medium | ml.t3.large | ml.t3.xlarge | ml.t3.2xlarge | ml.m5.large | ml.m5.xlarge | ml.m5.2xlarge | ml.m5.4xlarge | ml.m5.8xlarge | ml.m5.12xlarge | ml.m5.16xlarge | ml.m5.24xlarge | ml.c5.large | ml.c5.xlarge | ml.c5.2xlarge | ml.c5.4xlarge | ml.c5.9xlarge | ml.c5.12xlarge | ml.c5.18xlarge | ml.c5.24xlarge | ml.p3.2xlarge | ml.p3.8xlarge | ml.p3.16xlarge | ml.g4dn.xlarge | ml.g4dn.2xlarge | ml.g4dn.4xlarge | ml.g4dn.8xlarge | ml.g4dn.12xlarge | ml.g4dn.16xlarge | ml.r5.large | ml.r5.xlarge | ml.r5.2xlarge | ml.r5.4xlarge | ml.r5.8xlarge | ml.r5.12xlarge | ml.r5.16xlarge | ml.r5.24xlarge | ml.p3dn.24xlarge | ml.m5d.large | ml.m5d.xlarge | ml.m5d.2xlarge | ml.m5d.4xlarge | ml.m5d.8xlarge | ml.m5d.12xlarge | ml.m5d.16xlarge | ml.m5d.24xlarge | ml.g5.xlarge | ml.g5.2xlarge | ml.g5.4xlarge | ml.g5.8xlarge | ml.g5.12xlarge | ml.g5.16xlarge | ml.g5.24xlarge | ml.g5.48xlarge | ml.p4d.24xlarge | ml.p4de.24xlarge | ml.geospatial.interactive | ml.trn1.2xlarge | ml.trn1.32xlarge | ml.trn1n.32xlarge`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LifecycleConfigArn`

The Amazon Resource Name (ARN) of the Lifecycle Configuration attached to the
Resource.

_Required_: No

_Type_: String

_Pattern_: `arn:aws[a-z\-]*:sagemaker:[a-z0-9\-]*:[0-9]{12}:studio-lifecycle-config/.*`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SageMakerImageArn`

The ARN of the SageMaker AI image that the image version belongs to.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(-[\w]+)*:sagemaker:.+:[0-9]{12}:image/[a-z0-9]([-.]?[a-z0-9])*$`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SageMakerImageVersionArn`

The ARN of the image version created on the instance. To clear the value set for `SageMakerImageVersionArn`, pass `None` as the value.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(-[\w]+)*:sagemaker:.+:[0-9]{12}:image-version/[a-z0-9]([-.]?[a-z0-9])*/[0-9]+$`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OwnershipSettings

S3FileSystem

All content copied from https://docs.aws.amazon.com/.
