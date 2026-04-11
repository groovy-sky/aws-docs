This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::InferenceExperiment CaptureContentTypeHeader

Configuration specifying how to treat different headers. If no headers are specified
Amazon SageMaker AI will by default base64 encode when capturing the data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CsvContentTypes" : [ String, ... ],
  "JsonContentTypes" : [ String, ... ]
}

```

### YAML

```yaml

  CsvContentTypes:
    - String
  JsonContentTypes:
    - String

```

## Properties

`CsvContentTypes`

The list of all content type headers that Amazon SageMaker AI will treat as CSV and
capture accordingly.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `256 | 10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JsonContentTypes`

The list of all content type headers that SageMaker AI will treat as JSON and
capture accordingly.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `256 | 10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SageMaker::InferenceExperiment

DataStorageConfig

All content copied from https://docs.aws.amazon.com/.
