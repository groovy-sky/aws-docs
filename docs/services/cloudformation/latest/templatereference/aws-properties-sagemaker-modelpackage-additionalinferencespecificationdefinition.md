This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackage AdditionalInferenceSpecificationDefinition

A structure of additional Inference Specification. Additional Inference Specification
specifies details about inference jobs that can be run with models based on
this model package

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Containers" : [ ModelPackageContainerDefinition, ... ],
  "Description" : String,
  "Name" : String,
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
  Description: String
  Name: String
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

_Type_: Array of [ModelPackageContainerDefinition](aws-properties-sagemaker-modelpackage-modelpackagecontainerdefinition.md)

_Minimum_: `1`

_Maximum_: `15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the additional Inference specification

_Required_: No

_Type_: String

_Pattern_: `.*`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A unique name to identify the additional inference specification. The name must
be unique within the list of your additional inference specifications for a
particular model package.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SupportedContentTypes`

The supported MIME types for the input data.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SupportedRealtimeInferenceInstanceTypes`

A list of the instance types that are used to generate inferences in real-time.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SupportedResponseMIMETypes`

The supported MIME types for the output data.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SupportedTransformInstanceTypes`

A list of the instance types on which a transformation job can be run
or on which an endpoint can be deployed.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SageMaker::ModelPackage

Bias

All content copied from https://docs.aws.amazon.com/.
