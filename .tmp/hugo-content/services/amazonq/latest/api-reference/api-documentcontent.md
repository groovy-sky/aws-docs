# DocumentContent

The contents of a document.

###### Note

Documents have size limitations. The maximum file size for a document is 50 MB.
The maximum amount of text that can be extracted from a single document is 5 MB. For
more information, see [Supported document formats in Amazon Q Business](../qbusiness-ug/doc-types.md).

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**blob**

The contents of the document. Documents passed to the `blob` parameter must
be base64 encoded. Your code might not need to encode the document file bytes if you're
using an AWS SDK to call Amazon Q Business APIs. If you are calling the
Amazon Q Business endpoint directly using REST, you must base64 encode the contents before
sending.

Type: Base64-encoded binary data object

Required: No

**s3**

The path to the document in an Amazon S3 bucket.

Type: [S3](api-s3.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/documentcontent.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/documentcontent.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/documentcontent.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentAttributeValue

DocumentDetails

All content copied from https://docs.aws.amazon.com/.
