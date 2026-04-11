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

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/webexperience.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/webexperience.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/webexperience.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VideoSourceDetails

WebExperienceAuthConfiguration

All content copied from https://docs.aws.amazon.com/.
