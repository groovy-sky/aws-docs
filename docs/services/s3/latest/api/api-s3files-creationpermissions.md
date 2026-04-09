# CreationPermissions

Specifies the permissions to set on newly created directories within the file system.

## Contents

**ownerGid**

The POSIX group ID to assign to newly created directories.

Type: Long

Valid Range: Minimum value of 0. Maximum value of 4294967295.

Required: Yes

**ownerUid**

The POSIX user ID to assign to newly created directories.

Type: Long

Valid Range: Minimum value of 0. Maximum value of 4294967295.

Required: Yes

**permissions**

The octal permissions to assign to newly created directories.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 4.

Pattern: `[0-7]{3,4}`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3files-2025-05-05/creationpermissions.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3files-2025-05-05/creationpermissions.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3files-2025-05-05/creationpermissions.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon S3 Files

ExpirationDataRule

All content copied from https://docs.aws.amazon.com/.
