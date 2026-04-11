This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Comprehend::DocumentClassifier DocumentClassifierDocuments

The location of the training documents. This parameter is required in a request to create a semi-structured document classification model.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Uri" : String,
  "TestS3Uri" : String
}

```

### YAML

```yaml

  S3Uri: String
  TestS3Uri: String

```

## Properties

`S3Uri`

The S3 URI location of the training documents specified in the S3Uri CSV file.

_Required_: Yes

_Type_: String

_Pattern_: `s3://[a-z0-9][\.\-a-z0-9]{1,61}[a-z0-9](/.*)?`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TestS3Uri`

The S3 URI location of the test documents included in the TestS3Uri CSV file.
This field is not required if you do not specify a test CSV file.

_Required_: No

_Type_: String

_Pattern_: `s3://[a-z0-9][\.\-a-z0-9]{1,61}[a-z0-9](/.*)?`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AugmentedManifestsListItem

DocumentClassifierInputDataConfig

All content copied from https://docs.aws.amazon.com/.
