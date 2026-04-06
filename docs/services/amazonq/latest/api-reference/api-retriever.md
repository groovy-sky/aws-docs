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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/Retriever)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/Retriever)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/Retriever)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ResponseConfiguration

RetrieverConfiguration
