# FailedDocument

A list of documents that could not be removed from an Amazon Q Business index. Each entry
contains an error message that indicates why the document couldn't be removed from the
index.

## Contents

**dataSourceId**

The identifier of the Amazon Q Business data source connector that contains the failed
document.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: No

**error**

An explanation for why the document couldn't be removed from the index.

Type: [ErrorDetail](api-errordetail.md) object

Required: No

**id**

The identifier of the document that couldn't be removed from the Amazon Q Business
index.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1825.

Pattern: `\P{C}*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/faileddocument.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/faileddocument.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/faileddocument.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FailedAttachmentEvent

GroupMembers

All content copied from https://docs.aws.amazon.com/.
