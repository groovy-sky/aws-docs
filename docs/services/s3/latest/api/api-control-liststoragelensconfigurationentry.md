# ListStorageLensConfigurationEntry

Part of `ListStorageLensConfigurationResult`. Each entry includes the
description of the S3 Storage Lens configuration, its home Region, whether it is enabled, its
Amazon Resource Name (ARN), and config ID.

## Contents

**HomeRegion**

A container for the S3 Storage Lens home Region. Your metrics data is stored and retained in
your designated S3 Storage Lens home Region.

Type: String

Length Constraints: Minimum length of 5. Maximum length of 30.

Pattern: `[a-z0-9\-]+`

Required: Yes

**Id**

A container for the S3 Storage Lens configuration ID.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Pattern: `[a-zA-Z0-9\-\_\.]+`

Required: Yes

**StorageLensArn**

The ARN of the S3 Storage Lens configuration. This property is read-only.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Pattern: `arn:[a-z\-]+:s3:[a-z0-9\-]+:\d{12}:storage\-lens\/.*`

Required: Yes

**IsEnabled**

A container for whether the S3 Storage Lens configuration is enabled. This property is
required.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/liststoragelensconfigurationentry.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/liststoragelensconfigurationentry.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/liststoragelensconfigurationentry.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListCallerAccessGrantsEntry

ListStorageLensGroupEntry

All content copied from https://docs.aws.amazon.com/.
