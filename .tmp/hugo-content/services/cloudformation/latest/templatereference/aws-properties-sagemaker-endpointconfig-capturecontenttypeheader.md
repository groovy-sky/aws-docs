This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::EndpointConfig CaptureContentTypeHeader

Specifies the JSON and CSV content types of the data that the endpoint captures.

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

A list of the CSV content types of the data that the endpoint captures. For the endpoint to capture the data,
you must also specify the content type when you invoke the endpoint.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`JsonContentTypes`

A list of the JSON content types of the data that the endpoint captures. For the endpoint to capture the data,
you must also specify the content type when you invoke the endpoint.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AsyncInferenceOutputConfig

CaptureOption

All content copied from https://docs.aws.amazon.com/.
