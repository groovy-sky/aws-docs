# ListMountTargetsDescription

Contains information about a mount target returned in list operations.

## Contents

**mountTargetId**

The ID of the mount target.

Type: String

Length Constraints: Minimum length of 22. Maximum length of 45.

Pattern: `fsmt-[0-9a-f]{17,40}`

Required: Yes

**ownerId**

The AWS account ID of the mount target owner.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 12.

Pattern: `(\d{12})|(\d{4}-{4}-\d{4})`

Required: Yes

**subnetId**

The ID of the subnet where the mount target is located.

Type: String

Length Constraints: Minimum length of 15. Maximum length of 47.

Pattern: `subnet-[0-9a-f]{8,40}`

Required: Yes

**availabilityZoneId**

The Availability Zone ID where the mount target is located.

Type: String

Required: No

**fileSystemId**

The ID of the S3 File System.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 128.

Pattern: `(arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}|fs-[0-9a-f]{17,40})`

Required: No

**ipv4Address**

The IPv4 address of the mount target.

Type: String

Length Constraints: Minimum length of 7. Maximum length of 15.

Pattern: `[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}`

Required: No

**ipv6Address**

The IPv6 address of the mount target.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 39.

Required: No

**networkInterfaceId**

The ID of the network interface associated with the mount target.

Type: String

Required: No

**status**

The current status of the mount target.

Type: String

Valid Values: `available | creating | deleting | deleted | error | updating`

Required: No

**statusMessage**

Additional information about the mount target status.

Type: String

Required: No

**vpcId**

The ID of the VPC where the mount target is located.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3files-2025-05-05/listmounttargetsdescription.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3files-2025-05-05/listmounttargetsdescription.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3files-2025-05-05/listmounttargetsdescription.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListFileSystemsDescription

PosixUser

All content copied from https://docs.aws.amazon.com/.
