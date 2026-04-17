---
title: "RootDirectory"
---

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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3files-2025-05-05/RootDirectory)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3files-2025-05-05/RootDirectory)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3files-2025-05-05/RootDirectory)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PosixUser

Tag

All content copied from https://docs.aws.amazon.com/.
