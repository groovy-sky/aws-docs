This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::EndpointConfig AsyncInferenceOutputConfig

Specifies the configuration for asynchronous inference invocation outputs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsKeyId" : String,
  "NotificationConfig" : AsyncInferenceNotificationConfig,
  "S3FailurePath" : String,
  "S3OutputPath" : String
}

```

### YAML

```yaml

  KmsKeyId: String
  NotificationConfig:
    AsyncInferenceNotificationConfig
  S3FailurePath: String
  S3OutputPath: String

```

## Properties

`KmsKeyId`

The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt
the asynchronous inference output in Amazon S3.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NotificationConfig`

Specifies the configuration for notifications of inference results for asynchronous inference.

_Required_: No

_Type_: [AsyncInferenceNotificationConfig](aws-properties-sagemaker-endpointconfig-asyncinferencenotificationconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3FailurePath`

The Amazon S3 location to upload failure inference responses to.

_Required_: No

_Type_: String

_Pattern_: `(https|s3)://([^/])/?(.*)`

_Minimum_: `0`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3OutputPath`

The Amazon S3 location to upload inference responses to.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AsyncInferenceNotificationConfig

CaptureContentTypeHeader

All content copied from https://docs.aws.amazon.com/.
