---
title: "User"
---

# User

## Contents

###### Note

In the following list, the required parameters are described first.

**AccessString**

Access permissions string used for this user.

Type: String

Required: No

**ARN**

The Amazon Resource Name (ARN) of the user.

Type: String

Required: No

**Authentication**

Denotes whether the user requires a password to authenticate.

Type: [Authentication](api-authentication.md) object

Required: No

**Engine**

The options are valkey or redis.

Type: String

Pattern: `[a-zA-Z]*`

Required: No

**MinimumEngineVersion**

The minimum engine version required, which is Redis OSS 6.0

Type: String

Required: No

**Status**

Indicates the user status. Can be "active", "modifying" or "deleting".

Type: String

Required: No

**UserGroupIds.member.N**

Returns a list of the user group IDs the user belongs to.

Type: Array of strings

Length Constraints: Minimum length of 1.

Pattern: `[a-zA-Z][a-zA-Z0-9\-]*`

Required: No

**UserId**

The ID of the user.

Type: String

Required: No

**UserName**

The username of the user.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/elasticache-2015-02-02/User)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/elasticache-2015-02-02/User)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/elasticache-2015-02-02/User)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateAction

UserGroup

All content copied from https://docs.aws.amazon.com/.
