# UserAlias

Aliases attached to a user id within an Amazon Q Business application.

## Contents

**userId**

The identifier of the user id associated with the user aliases.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: Yes

**dataSourceId**

The identifier of the data source that the user aliases are associated with.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: No

**indexId**

The identifier of the index that the user aliases are associated with.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/UserAlias)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/UserAlias)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/UserAlias)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TopicConfiguration

UsersAndGroups
