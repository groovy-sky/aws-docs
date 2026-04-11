# Retriever

Summary information for the retriever used for your Amazon Q Business application.

## Contents

**applicationId**

The identifier of the Amazon Q Business application using the retriever.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: No

**displayName**

The name of your retriever.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1000.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9_-]*`

Required: No

**retrieverId**

The identifier of the retriever used by your Amazon Q Business application.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: No

**status**

The status of your retriever.

Type: String

Valid Values: `CREATING | ACTIVE | FAILED`

Required: No

**type**

The type of your retriever.

Type: String

Valid Values: `NATIVE_INDEX | KENDRA_INDEX`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/retriever.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/retriever.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/retriever.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResponseConfiguration

RetrieverConfiguration

All content copied from https://docs.aws.amazon.com/.
