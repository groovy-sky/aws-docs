# Document

A document in an Amazon Q Business application.

## Contents

**id**

The identifier of the document.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1825.

Pattern: `\P{C}*`

Required: Yes

**accessConfiguration**

Configuration information for access permission to a document.

Type: [AccessConfiguration](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_AccessConfiguration.html) object

Required: No

**attributes**

Custom attributes to apply to the document for refining Amazon Q Business web experience
responses.

Type: Array of [DocumentAttribute](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_DocumentAttribute.html) objects

Array Members: Minimum number of 1 item. Maximum number of 500 items.

Required: No

**content**

The contents of the document.

Type: [DocumentContent](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_DocumentContent.html) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

**contentType**

The file type of the document in the Blob field.

If you want to index snippets or subsets of HTML documents instead of the entirety of
the HTML documents, you add the `HTML` start and closing tags
( `<HTML>content</HTML>`) around the content.

Type: String

Valid Values: `PDF | HTML | MS_WORD | PLAIN_TEXT | PPT | RTF | XML | XSLT | MS_EXCEL | CSV | JSON | MD`

Required: No

**documentEnrichmentConfiguration**

The configuration information for altering document metadata and content during the
document ingestion process.

Type: [DocumentEnrichmentConfiguration](api-documentenrichmentconfiguration.md) object

Required: No

**mediaExtractionConfiguration**

The configuration for extracting information from media in the document.

Type: [MediaExtractionConfiguration](api-mediaextractionconfiguration.md) object

Required: No

**title**

The title of the document.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/Document)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/Document)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/Document)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DeleteDocument

DocumentAcl
