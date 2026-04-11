This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ProcessingJob ProcessingOutputConfig

Configuration for uploading output from the processing container.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsKeyId" : String,
  "Outputs" : [ ProcessingOutputsObject, ... ]
}

```

### YAML

```yaml

  KmsKeyId: String
  Outputs:
    - ProcessingOutputsObject

```

## Properties

`KmsKeyId`

The AWS Key Management Service (AWS KMS) key that Amazon SageMaker
uses to encrypt the processing job output. `KmsKeyId` can be an ID of a KMS
key, ARN of a KMS key, or alias of a KMS key. The `KmsKeyId` is applied to
all outputs.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9:/_-]*`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Outputs`

An array of outputs configuring the data to upload from the processing
container.

_Required_: Yes

_Type_: Array of [ProcessingOutputsObject](aws-properties-sagemaker-processingjob-processingoutputsobject.md)

_Maximum_: `10`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProcessingInputsObject

ProcessingOutputsObject

All content copied from https://docs.aws.amazon.com/.
