This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Comprehend::DocumentClassifier DocumentClassifierInputDataConfig

The input properties for training a document classifier.

For more information on how the input file is formatted, see
[Preparing training data](../../../comprehend/latest/dg/prep-classifier-data.md) in the Comprehend Developer Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AugmentedManifests" : [ AugmentedManifestsListItem, ... ],
  "DataFormat" : String,
  "DocumentReaderConfig" : DocumentReaderConfig,
  "Documents" : DocumentClassifierDocuments,
  "DocumentType" : String,
  "LabelDelimiter" : String,
  "S3Uri" : String,
  "TestS3Uri" : String
}

```

### YAML

```yaml

  AugmentedManifests:
    - AugmentedManifestsListItem
  DataFormat: String
  DocumentReaderConfig:
    DocumentReaderConfig
  Documents:
    DocumentClassifierDocuments
  DocumentType: String
  LabelDelimiter: String
  S3Uri: String
  TestS3Uri: String

```

## Properties

`AugmentedManifests`

A list of augmented manifest files that provide training data for your custom model. An
augmented manifest file is a labeled dataset that is produced by Amazon SageMaker Ground
Truth.

This parameter is required if you set `DataFormat` to
`AUGMENTED_MANIFEST`.

_Required_: No

_Type_: Array of [AugmentedManifestsListItem](aws-properties-comprehend-documentclassifier-augmentedmanifestslistitem.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataFormat`

The format of your training data:

- `COMPREHEND_CSV`: A two-column CSV file, where labels are provided in the
first column, and documents are provided in the second. If you use this value, you must
provide the `S3Uri` parameter in your request.

- `AUGMENTED_MANIFEST`: A labeled dataset that is produced by Amazon
SageMaker Ground Truth. This file is in JSON lines format. Each line is a complete JSON
object that contains a training document and its associated labels.

If you use this value, you must provide the `AugmentedManifests` parameter
in your request.

If you don't specify a value, Amazon Comprehend uses `COMPREHEND_CSV` as the
default.

_Required_: No

_Type_: String

_Allowed values_: `COMPREHEND_CSV | AUGMENTED_MANIFEST`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DocumentReaderConfig`

Property description not available.

_Required_: No

_Type_: [DocumentReaderConfig](aws-properties-comprehend-documentclassifier-documentreaderconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Documents`

The S3 location of the training documents.
This parameter is required in a request to create a native document model.

_Required_: No

_Type_: [DocumentClassifierDocuments](aws-properties-comprehend-documentclassifier-documentclassifierdocuments.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DocumentType`

The type of input documents for training the model. Provide plain-text documents to create a plain-text model, and
provide semi-structured documents to create a native document model.

_Required_: No

_Type_: String

_Allowed values_: `PLAIN_TEXT_DOCUMENT | SEMI_STRUCTURED_DOCUMENT`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LabelDelimiter`

Indicates the delimiter used to separate each label for training a multi-label classifier.
The default delimiter between labels is a pipe (\|). You can use a different character as a
delimiter (if it's an allowed character) by specifying it under Delimiter for labels. If the
training documents use a delimiter other than the default or the delimiter you specify, the
labels on that line will be combined to make a single unique label, such as
LABELLABELLABEL.

_Required_: No

_Type_: String

_Pattern_: `^[ ~!@#$%^*\-_+=|\\:;\t>?/]$`

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3Uri`

The Amazon S3 URI for the input data. The S3 bucket must be in the same Region as the API
endpoint that you are calling. The URI can point to a single input file or it can provide the
prefix for a collection of input files.

For example, if you use the URI `S3://bucketName/prefix`, if the prefix is a
single file, Amazon Comprehend uses that file as input. If more than one file begins with the
prefix, Amazon Comprehend uses all of them as input.

This parameter is required if you set `DataFormat` to
`COMPREHEND_CSV`.

_Required_: No

_Type_: String

_Pattern_: `s3://[a-z0-9][\.\-a-z0-9]{1,61}[a-z0-9](/.*)?`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TestS3Uri`

This specifies the Amazon S3 location that contains the test annotations for the document classifier.
The URI must be in the same AWS Region as the API endpoint that you are calling.

_Required_: No

_Type_: String

_Pattern_: `s3://[a-z0-9][\.\-a-z0-9]{1,61}[a-z0-9](/.*)?`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentClassifierDocuments

DocumentClassifierOutputDataConfig

All content copied from https://docs.aws.amazon.com/.
