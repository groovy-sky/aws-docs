# ListAccessPointsDescription

Contains information about an S3 File System Access Point returned in list operations.

## Contents

**accessPointArn**

The Amazon Resource Name (ARN) of the access point.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 256.

Pattern: `arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}/access-point/fsap-[0-9a-f]{17,40}`

Required: Yes

**accessPointId**

The ID of the access point.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 256.

Pattern: `(arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}/access-point/fsap-[0-9a-f]{17,40}|fsap-[0-9a-f]{17,40})`

Required: Yes

**fileSystemId**

The ID of the S3 File System.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 128.

Pattern: `(arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}|fs-[0-9a-f]{17,40})`

Required: Yes

**ownerId**

The AWS account ID of the access point owner.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 12.

Pattern: `(\d{12})|(\d{4}-{4}-\d{4})`

Required: Yes

**status**

The current status of the access point.

Type: String

Valid Values: `available | creating | deleting | deleted | error | updating`

Required: Yes

**name**

The name of the access point.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 256.

Pattern: `([\p{L}\p{Z}\p{N}_.:/=+\-@]*)`

Required: No

**posixUser**

The POSIX identity configured for this access point.

Type: [PosixUser](api-s3files-posixuser.md) object

Required: No

**rootDirectory**

The root directory configuration for this access point.

Type: [RootDirectory](api-s3files-rootdirectory.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3files-2025-05-05/listaccesspointsdescription.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3files-2025-05-05/listaccesspointsdescription.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3files-2025-05-05/listaccesspointsdescription.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImportDataRule

ListFileSystemsDescription

All content copied from https://docs.aws.amazon.com/.
