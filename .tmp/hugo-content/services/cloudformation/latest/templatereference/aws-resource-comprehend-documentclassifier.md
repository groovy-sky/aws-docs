This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Comprehend::DocumentClassifier

This resource creates and trains a document classifier to categorize documents. You provide a set of training documents
that are labeled with the categories that you want to identify.
After the classifier is trained you can use it to categorize a set of labeled
documents into the categories. For more information, see
[Document Classification](../../../comprehend/latest/dg/how-document-classification.md)
in the Comprehend Developer Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Comprehend::DocumentClassifier",
  "Properties" : {
      "DataAccessRoleArn" : String,
      "DocumentClassifierName" : String,
      "InputDataConfig" : DocumentClassifierInputDataConfig,
      "LanguageCode" : String,
      "Mode" : String,
      "ModelKmsKeyId" : String,
      "ModelPolicy" : String,
      "OutputDataConfig" : DocumentClassifierOutputDataConfig,
      "Tags" : [ Tag, ... ],
      "VersionName" : String,
      "VolumeKmsKeyId" : String,
      "VpcConfig" : VpcConfig
    }
}

```

### YAML

```yaml

Type: AWS::Comprehend::DocumentClassifier
Properties:
  DataAccessRoleArn: String
  DocumentClassifierName: String
  InputDataConfig:
    DocumentClassifierInputDataConfig
  LanguageCode: String
  Mode: String
  ModelKmsKeyId: String
  ModelPolicy: String
  OutputDataConfig:
    DocumentClassifierOutputDataConfig
  Tags:
    - Tag
  VersionName: String
  VolumeKmsKeyId: String
  VpcConfig:
    VpcConfig

```

## Properties

`DataAccessRoleArn`

The Amazon Resource Name (ARN) of the IAM role that
grants Amazon Comprehend read access to your input data.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws(-[^:]+)?:iam::[0-9]{12}:role/.+`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DocumentClassifierName`

The name of the document classifier.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InputDataConfig`

Specifies the format and location of the input data for the job.

_Required_: Yes

_Type_: [DocumentClassifierInputDataConfig](aws-properties-comprehend-documentclassifier-documentclassifierinputdataconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LanguageCode`

The language of the input documents. You can specify any of the languages
supported by Amazon Comprehend. All documents must be in the same language.

_Required_: Yes

_Type_: String

_Allowed values_: `en | es | fr | it | de | pt`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Mode`

Indicates the mode in which the classifier will be trained. The classifier can be trained
in multi-class (single-label) mode or multi-label mode.
Multi-class mode identifies a single class label for each document and
multi-label mode identifies one or more class labels for each document. Multiple
labels for an individual document are separated by a delimiter. The default delimiter between
labels is a pipe (\|).

_Required_: No

_Type_: String

_Allowed values_: `MULTI_CLASS | MULTI_LABEL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelKmsKeyId`

ID for the AWS KMS key that Amazon Comprehend uses to encrypt
trained custom models. The ModelKmsKeyId can be either of the following formats:

- KMS Key ID: `"1234abcd-12ab-34cd-56ef-1234567890ab"`

- Amazon Resource Name (ARN) of a KMS Key:
`"arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"`

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelPolicy`

The resource-based policy to attach to your custom document classifier model. You can use
this policy to allow another AWS account to import your custom model.

Provide your policy as a JSON body that you enter as a UTF-8 encoded string without line
breaks. To provide valid JSON, enclose the attribute names and values in double quotes. If the
JSON body is also enclosed in double quotes, then you must escape the double quotes that are
inside the policy:

`"{\"attribute\": \"value\", \"attribute\": [\"value\"]}"`

To avoid escaping quotes, you can use single quotes to enclose the policy and double
quotes to enclose the JSON names and values:

`'{"attribute": "value", "attribute": ["value"]}'`

_Required_: No

_Type_: String

_Pattern_: `[\u0009\u000A\u000D\u0020-\u00FF]+`

_Minimum_: `1`

_Maximum_: `20000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputDataConfig`

Provides output results configuration parameters for custom classifier jobs.

_Required_: No

_Type_: [DocumentClassifierOutputDataConfig](aws-properties-comprehend-documentclassifier-documentclassifieroutputdataconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Tags to associate with the document classifier. A tag is a key-value
pair that adds as a metadata to a resource used by Amazon Comprehend. For example, a tag with
"Sales" as the key might be added to a resource to indicate its use by the sales department.

_Required_: No

_Type_: Array of [Tag](aws-properties-comprehend-documentclassifier-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VersionName`

The version name given to the newly created classifier. Version names can have a maximum
of 256 characters. Alphanumeric characters, hyphens (-) and underscores (\_) are allowed. The
version name must be unique among all models with the same classifier name in the AWS account/AWS Region.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VolumeKmsKeyId`

ID for the AWS Key Management Service (KMS) key that Amazon Comprehend uses to encrypt
data on the storage volume attached to the ML compute instance(s) that process the analysis
job. The VolumeKmsKeyId can be either of the following formats:

- KMS Key ID: `"1234abcd-12ab-34cd-56ef-1234567890ab"`

- Amazon Resource Name (ARN) of a KMS Key:
`"arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"`

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcConfig`

Configuration parameters for a private Virtual Private Cloud (VPC) containing the
resources you are using for your custom classifier. For more information, see [Amazon\
VPC](../../../vpc/latest/userguide/what-is-amazon-vpc.md).

_Required_: No

_Type_: [VpcConfig](aws-properties-comprehend-documentclassifier-vpcconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the document classifier.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the document classifier.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Comprehend

AugmentedManifestsListItem

All content copied from https://docs.aws.amazon.com/.
