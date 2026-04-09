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

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/useralias.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/useralias.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/useralias.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TopicConfiguration

UsersAndGroups

All content copied from https://docs.aws.amazon.com/.
