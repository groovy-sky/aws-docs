This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::InferenceExperiment DataStorageConfig

The Amazon S3 location and configuration for storing inference request and response data.

This is an optional parameter that you can use for data capture. For more information, see [Capture data](../../../sagemaker/latest/dg/model-monitor-data-capture.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContentType" : CaptureContentTypeHeader,
  "Destination" : String,
  "KmsKey" : String
}

```

### YAML

```yaml

  ContentType:
    CaptureContentTypeHeader
  Destination: String
  KmsKey: String

```

## Properties

`ContentType`

Configuration specifying how to treat different headers. If no headers are specified SageMaker will by default
base64 encode when capturing the data.

_Required_: No

_Type_: [CaptureContentTypeHeader](aws-properties-sagemaker-inferenceexperiment-capturecontenttypeheader.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Destination`

The Amazon S3 bucket where the inference request and response data is stored.

_Required_: Yes

_Type_: String

_Pattern_: `^(https|s3)://([^/])/?(.*)$`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKey`

The AWS Key Management Service key that Amazon SageMaker uses to encrypt captured data at rest using Amazon S3
server-side encryption.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CaptureContentTypeHeader

EndpointMetadata

All content copied from https://docs.aws.amazon.com/.
