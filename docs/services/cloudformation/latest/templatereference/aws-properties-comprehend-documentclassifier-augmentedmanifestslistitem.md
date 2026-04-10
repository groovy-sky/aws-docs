This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Comprehend::DocumentClassifier AugmentedManifestsListItem

An augmented manifest file that provides training data for your custom model. An augmented
manifest file is a labeled dataset that is produced by Amazon SageMaker Ground Truth.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttributeNames" : [ String, ... ],
  "S3Uri" : String,
  "Split" : String
}

```

### YAML

```yaml

  AttributeNames:
    - String
  S3Uri: String
  Split: String

```

## Properties

`AttributeNames`

The JSON attribute that contains the annotations for your training documents. The number
of attribute names that you specify depends on whether your augmented manifest file is the
output of a single labeling job or a chained labeling job.

If your file is the output of a single labeling job, specify the LabelAttributeName key
that was used when the job was created in Ground Truth.

If your file is the output of a chained labeling job, specify the LabelAttributeName key
for one or more jobs in the chain. Each LabelAttributeName key provides the annotations from
an individual job.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3Uri`

The Amazon S3 location of the augmented manifest file.

_Required_: Yes

_Type_: String

_Pattern_: `s3://[a-z0-9][\.\-a-z0-9]{1,61}[a-z0-9](/.*)?`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Split`

The purpose of the data you've provided in the augmented manifest. You can either train or
test this data. If you don't specify, the default is train.

TRAIN - all of the documents in the manifest will be used for training. If no test
documents are provided, Amazon Comprehend will automatically reserve a portion of the training
documents for testing.

TEST - all of the documents in the manifest will be used for testing.

_Required_: No

_Type_: String

_Allowed values_: `TRAIN | TEST`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Comprehend::DocumentClassifier

DocumentClassifierDocuments

All content copied from https://docs.aws.amazon.com/.
