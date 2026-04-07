This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::EndpointConfig DataCaptureConfig

Specifies the configuration of your endpoint for model monitor data capture.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CaptureContentTypeHeader" : CaptureContentTypeHeader,
  "CaptureOptions" : [ CaptureOption, ... ],
  "DestinationS3Uri" : String,
  "EnableCapture" : Boolean,
  "InitialSamplingPercentage" : Integer,
  "KmsKeyId" : String
}

```

### YAML

```yaml

  CaptureContentTypeHeader:
    CaptureContentTypeHeader
  CaptureOptions:
    - CaptureOption
  DestinationS3Uri: String
  EnableCapture: Boolean
  InitialSamplingPercentage: Integer
  KmsKeyId: String

```

## Properties

`CaptureContentTypeHeader`

A list of the JSON and CSV content type that the endpoint captures.

_Required_: No

_Type_: [CaptureContentTypeHeader](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-endpointconfig-capturecontenttypeheader.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CaptureOptions`

Specifies whether the endpoint captures input data to your model, output data from your model, or both.

_Required_: Yes

_Type_: Array of [CaptureOption](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-endpointconfig-captureoption.html)

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DestinationS3Uri`

The S3 bucket where model monitor stores captured data.

_Required_: Yes

_Type_: String

_Pattern_: `(https|s3)://([^/])/?(.*)`

_Minimum_: `0`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnableCapture`

Set to `True` to enable data capture.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InitialSamplingPercentage`

The percentage of data to capture.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyId`

The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to
encrypt the captured data at rest using Amazon S3 server-side encryption. The KmsKeyId can be any of the following
formats: Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab Key ARN:
arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab Alias name: alias/ExampleAlias Alias name
ARN: arn:aws:kms:us-west-2:111122223333:alias/ExampleAlias If you don't provide a KMS key ID, Amazon SageMaker uses
the default KMS key for Amazon S3 for your role's account. For more information, see KMS-Managed Encryption Keys
(https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html) in the Amazon Simple Storage Service
Developer Guide. The KMS key policy must grant permission to the IAM role that you specify in your CreateModel
(https://docs.aws.amazon.com/sagemaker/latest/APIReference/API\_CreateModel.html) request. For more information, see
Using Key Policies in AWS KMS
(http://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html) in the AWS Key Management
Service Developer Guide.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9:/_-]*`

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ClarifyTextConfig

ExplainerConfig
