# LibraryItemMember

A library item is a snapshot of an Amazon Q App that can be published so the users in their
Amazon Q Apps library can discover it, clone it, and run it.

## Contents

**appId**

The unique identifier of the Q App associated with the library item.

Type: String

Pattern: `[\da-f]{8}-[\da-f]{4}-[45][\da-f]{3}-[89ABab][\da-f]{3}-[\da-f]{12}`

Required: Yes

**appVersion**

The version of the Q App associated with the library item.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 2147483647.

Required: Yes

**categories**

The categories associated with the library item.

Type: Array of [Category](api-qapps-category.md) objects

Array Members: Minimum number of 0 items. Maximum number of 3 items.

Required: Yes

**createdAt**

The date and time the library item was created.

Type: Timestamp

Required: Yes

**createdBy**

The user who created the library item.

Type: String

Required: Yes

**libraryItemId**

The unique identifier of the library item.

Type: String

Pattern: `[\da-f]{8}-[\da-f]{4}-[45][\da-f]{3}-[89ABab][\da-f]{3}-[\da-f]{12}`

Required: Yes

**ratingCount**

The number of ratings the library item has received.

Type: Integer

Required: Yes

**status**

The status of the library item.

Type: String

Required: Yes

**isRatedByUser**

Whether the current user has rated the library item.

Type: Boolean

Required: No

**isVerified**

Indicates whether the library item has been verified.

Type: Boolean

Required: No

**updatedAt**

The date and time the library item was last updated.

Type: Timestamp

Required: No

**updatedBy**

The user who last updated the library item.

Type: String

Required: No

**userCount**

The number of users who have the associated Q App.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qapps-2023-11-27/LibraryItemMember)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qapps-2023-11-27/LibraryItemMember)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qapps-2023-11-27/LibraryItemMember)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FormInputCardMetadata

PermissionInput
