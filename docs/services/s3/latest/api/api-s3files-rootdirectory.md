# RootDirectory

Specifies the root directory path and optional creation permissions for newly created directories.

## Contents

**creationPermissions**

The permissions to set on newly created directories.

Type: [CreationPermissions](api-s3files-creationpermissions.md) object

Required: No

**path**

The path to use as the root directory for the access point.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 100.

Pattern: ``(\/|(\/(?!\.)+[^$#<>;;`|&?{}^*/\n]+){1,4})``

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3files-2025-05-05/rootdirectory.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3files-2025-05-05/rootdirectory.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3files-2025-05-05/rootdirectory.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PosixUser

Tag

All content copied from https://docs.aws.amazon.com/.
