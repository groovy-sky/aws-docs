This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackage InferenceSpecification

Defines how to perform inference generation after a training job is run.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Containers" : [ ModelPackageContainerDefinition, ... ],
  "SupportedContentTypes" : [ String, ... ],
  "SupportedRealtimeInferenceInstanceTypes" : [ String, ... ],
  "SupportedResponseMIMETypes" : [ String, ... ],
  "SupportedTransformInstanceTypes" : [ String, ... ]
}

```

### YAML

```yaml

  Containers:
    - ModelPackageContainerDefinition
  SupportedContentTypes:
    - String
  SupportedRealtimeInferenceInstanceTypes:
    - String
  SupportedResponseMIMETypes:
    - String
  SupportedTransformInstanceTypes:
    - String

```

## Properties

`Containers`

The Amazon ECR registry path of the Docker image that contains the inference code.

_Required_: Yes

_Type_: Array of [ModelPackageContainerDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelpackage-modelpackagecontainerdefinition.html)

_Minimum_: `1`

_Maximum_: `15`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SupportedContentTypes`

The supported MIME types for the input data.

_Required_: Yes

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SupportedRealtimeInferenceInstanceTypes`

A list of the instance types that are used to generate inferences in real-time.

This parameter is required for unversioned models, and optional for versioned
models.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SupportedResponseMIMETypes`

The supported MIME types for the output data.

_Required_: Yes

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SupportedTransformInstanceTypes`

A list of the instance types on which a transformation job can be run or on which an
endpoint can be deployed.

This parameter is required for unversioned models, and optional for versioned
models.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FileSource

MetadataProperties
