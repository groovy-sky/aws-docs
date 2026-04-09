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

Type: [AccessConfiguration](api-accessconfiguration.md) object

Required: No

**attributes**

Custom attributes to apply to the document for refining Amazon Q Business web experience
responses.

Type: Array of [DocumentAttribute](api-documentattribute.md) objects

Array Members: Minimum number of 1 item. Maximum number of 500 items.

Required: No

**content**

The contents of the document.

Type: [DocumentContent](api-documentcontent.md) object

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

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/document.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/document.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/document.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteDocument

DocumentAcl

All content copied from https://docs.aws.amazon.com/.
