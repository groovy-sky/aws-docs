# Index

Summary information for your Amazon Q Business index.

## Contents

**createdAt**

The Unix timestamp when the index was created.

Type: Timestamp

Required: No

**displayName**

The name of the index.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1000.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9_-]*`

Required: No

**indexId**

The identifier for the index.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: No

**status**

The current status of the index. When the status is `ACTIVE`, the index is
ready.

Type: String

Valid Values: `CREATING | ACTIVE | DELETING | FAILED | UPDATING`

Required: No

**updatedAt**

The Unix timestamp when the index was last updated.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/Index)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/Index)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/Index)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ImageSourceDetails

IndexCapacityConfiguration
