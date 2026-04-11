This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackage SourceAlgorithm

Specifies an algorithm that was used to create the model package. The algorithm must
be either an algorithm resource in your SageMaker account or an algorithm in AWS Marketplace that you are subscribed to.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlgorithmName" : String,
  "ModelDataUrl" : String
}

```

### YAML

```yaml

  AlgorithmName: String
  ModelDataUrl: String

```

## Properties

`AlgorithmName`

The name of an algorithm that was used to create the model package. The algorithm must
be either an algorithm resource in your SageMaker account or an algorithm in AWS Marketplace that you are subscribed to.

_Required_: Yes

_Type_: String

_Pattern_: `(arn:aws[a-z\-]*:sagemaker:[a-z0-9\-]*:[0-9]{12}:[a-z\-]*\/)?([a-zA-Z0-9]([a-zA-Z0-9-]){0,62})(?<!-)$`

_Minimum_: `1`

_Maximum_: `170`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelDataUrl`

The Amazon S3 path where the model artifacts, which result from model training, are stored.
This path must point to a single `gzip` compressed tar archive
( `.tar.gz` suffix).

###### Note

The model artifacts must be in an S3 bucket that is in the same AWS
region as the algorithm.

_Required_: No

_Type_: String

_Pattern_: `^(https|s3)://([^/]+)/?(.*)$`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SecurityConfig

SourceAlgorithmSpecification

All content copied from https://docs.aws.amazon.com/.
