# WebExperience

Provides information for an Amazon Q Business web experience.

## Contents

**createdAt**

The Unix timestamp when the Amazon Q Business application was last updated.

Type: Timestamp

Required: No

**defaultEndpoint**

The endpoint URLs for your Amazon Q Business web experience. The URLs are unique and fully
hosted by AWS.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `(https?|ftp|file)://([^\s]*)`

Required: No

**status**

The status of your Amazon Q Business web experience.

Type: String

Valid Values: `CREATING | ACTIVE | DELETING | FAILED | PENDING_AUTH_CONFIG`

Required: No

**updatedAt**

The Unix timestamp when your Amazon Q Business web experience was updated.

Type: Timestamp

Required: No

**webExperienceId**

The identifier of your Amazon Q Business web experience.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/WebExperience)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/WebExperience)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/WebExperience)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VideoSourceDetails

WebExperienceAuthConfiguration
