# ListFileSystemsDescription

Contains information about an S3 File System returned in list operations.

## Contents

**bucket**

The Amazon Resource Name (ARN) of the S3 bucket.

Type: String

Pattern: `(arn:aws[a-zA-Z0-9-]*:s3:::.+)`

Required: Yes

**creationTime**

The time when the file system was created.

Type: Timestamp

Required: Yes

**fileSystemArn**

The Amazon Resource Name (ARN) of the file system.

Type: String

Pattern: `(arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40})`

Required: Yes

**fileSystemId**

The ID of the file system.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 128.

Pattern: `(arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}|fs-[0-9a-f]{17,40})`

Required: Yes

**ownerId**

The AWS account ID of the file system owner.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 12.

Pattern: `(\d{12})|(\d{4}-{4}-\d{4})`

Required: Yes

**roleArn**

The Amazon Resource Name (ARN) of the IAM role used for S3 access.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 2048.

Pattern: `arn:(aws[a-zA-Z-]*)?:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+`

Required: Yes

**status**

The current status of the file system.

Type: String

Valid Values: `available | creating | deleting | deleted | error | updating`

Required: Yes

**name**

The name of the file system.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 256.

Pattern: `([\p{L}\p{Z}\p{N}_.:/=+\-@]*)`

Required: No

**statusMessage**

Additional information about the file system status.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3files-2025-05-05/listfilesystemsdescription.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3files-2025-05-05/listfilesystemsdescription.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3files-2025-05-05/listfilesystemsdescription.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListAccessPointsDescription

ListMountTargetsDescription

All content copied from https://docs.aws.amazon.com/.
