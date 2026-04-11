This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Comprehend::DocumentClassifier DocumentReaderConfig

Provides configuration parameters to override the default actions for extracting text from PDF documents and image files.

By default, Amazon Comprehend performs the following actions to extract text from files, based on the input file type:

- **Word files** \- Amazon Comprehend parser extracts the text.

- **Digital PDF files** \- Amazon Comprehend parser extracts the text.

- **Image files and scanned PDF files** \- Amazon Comprehend uses the Amazon Textract `DetectDocumentText`
API to extract the text.

`DocumentReaderConfig` does not apply to plain text files or Word files.

For image files and PDF documents, you can override these default actions using the fields listed below.
For more information, see [Setting text extraction options](../../../comprehend/latest/dg/idp-set-textract-options.md) in the Comprehend Developer Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DocumentReadAction" : String,
  "DocumentReadMode" : String,
  "FeatureTypes" : [ String, ... ]
}

```

### YAML

```yaml

  DocumentReadAction: String
  DocumentReadMode: String
  FeatureTypes:
    - String

```

## Properties

`DocumentReadAction`

This field defines the Amazon Textract API operation that Amazon Comprehend uses to extract text from PDF files and image files.
Enter one of the following values:

- `TEXTRACT_DETECT_DOCUMENT_TEXT` \- The Amazon Comprehend service uses the `DetectDocumentText`
API operation.

- `TEXTRACT_ANALYZE_DOCUMENT` \- The Amazon Comprehend service uses the `AnalyzeDocument`
API operation.

_Required_: Yes

_Type_: String

_Allowed values_: `TEXTRACT_DETECT_DOCUMENT_TEXT | TEXTRACT_ANALYZE_DOCUMENT`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DocumentReadMode`

Determines the text extraction actions for PDF files. Enter one of the following values:

- `SERVICE_DEFAULT` \- use the Amazon Comprehend service defaults for PDF files.

- `FORCE_DOCUMENT_READ_ACTION` \- Amazon Comprehend uses the Textract API specified by
DocumentReadAction for all PDF files, including digital PDF files.

_Required_: No

_Type_: String

_Allowed values_: `SERVICE_DEFAULT | FORCE_DOCUMENT_READ_ACTION`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FeatureTypes`

Specifies the type of Amazon Textract features to apply. If you chose `TEXTRACT_ANALYZE_DOCUMENT`
as the read action, you must specify one or both of the following values:

- `TABLES` \- Returns additional information about any tables that are detected in the input document.

- `FORMS` \- Returns additional information about any forms that are detected in the input document.

_Required_: No

_Type_: Array of String

_Allowed values_: `TABLES | FORMS`

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentClassifierOutputDataConfig

Tag

All content copied from https://docs.aws.amazon.com/.
