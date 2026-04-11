# UserAppItem

An Amazon Q App associated with a user, either owned by the user or favorited.

## Contents

**appArn**

The Amazon Resource Name (ARN) of the Q App.

Type: String

Required: Yes

**appId**

The unique identifier of the Q App.

Type: String

Pattern: `[\da-f]{8}-[\da-f]{4}-[45][\da-f]{3}-[89ABab][\da-f]{3}-[\da-f]{12}`

Required: Yes

**createdAt**

The date and time the user's association with the Q App was created.

Type: Timestamp

Required: Yes

**title**

The title of the Q App.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 100.

Pattern: `[^{}\\"<>]+`

Required: Yes

**canEdit**

A flag indicating whether the user can edit the Q App.

Type: Boolean

Required: No

**description**

The description of the Q App.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 500.

Required: No

**isVerified**

Indicates whether the Q App has been verified.

Type: Boolean

Required: No

**status**

The status of the user's association with the Q App.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qapps-2023-11-27/userappitem.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qapps-2023-11-27/userappitem.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qapps-2023-11-27/userappitem.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

User

Common Parameters

All content copied from https://docs.aws.amazon.com/.
