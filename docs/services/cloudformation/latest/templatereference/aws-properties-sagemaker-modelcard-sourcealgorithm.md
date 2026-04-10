This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelCard SourceAlgorithm

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

_Maximum_: `170`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelDataUrl`

The Amazon S3 path where the model artifacts, which result from model training, are stored.
This path must point to a single `gzip` compressed tar archive
( `.tar.gz` suffix).

###### Note

The model artifacts must be in an S3 bucket that is in the same AWS
region as the algorithm.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SimpleMetric

Tag

All content copied from https://docs.aws.amazon.com/.
