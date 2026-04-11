# PosixUser

Specifies the POSIX identity with uid, gid, and secondary group IDs for user enforcement.

## Contents

**gid**

The POSIX group ID.

Type: Long

Valid Range: Minimum value of 0. Maximum value of 4294967295.

Required: Yes

**uid**

The POSIX user ID.

Type: Long

Valid Range: Minimum value of 0. Maximum value of 4294967295.

Required: Yes

**secondaryGids**

An array of secondary POSIX group IDs.

Type: Array of longs

Valid Range: Minimum value of 0. Maximum value of 4294967295.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3files-2025-05-05/posixuser.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3files-2025-05-05/posixuser.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3files-2025-05-05/posixuser.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListMountTargetsDescription

RootDirectory

All content copied from https://docs.aws.amazon.com/.
