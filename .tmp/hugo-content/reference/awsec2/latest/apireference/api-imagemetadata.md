# ImageMetadata

Information about the AMI.

## Contents

**creationDate**

The date and time the AMI was created.

Type: String

Required: No

**deprecationTime**

The deprecation date and time of the AMI, in UTC, in the following format:
_YYYY_- _MM_- _DD_ T _HH_: _MM_: _SS_ Z.

Type: String

Required: No

**imageAllowed**

If `true`, the AMI satisfies the criteria for Allowed AMIs and can be
discovered and used in the account. If `false`, the AMI can't be discovered or used
in the account.

For more information, see [Control the discovery and use of AMIs in\
Amazon EC2 with Allowed AMIs](../../../../services/ec2/latest/userguide/ec2-allowed-amis.md) in
_Amazon EC2 User Guide_.

Type: Boolean

Required: No

**imageId**

The ID of the AMI.

Type: String

Required: No

**imageOwnerAlias**

The alias of the AMI owner.

Valid values: `amazon` \| `aws-backup-vault` \|
`aws-marketplace`

Type: String

Required: No

**imageOwnerId**

The ID of the AWS account that owns the AMI.

Type: String

Required: No

**imageState**

The current state of the AMI. If the state is `available`, the AMI is
successfully registered and can be used to launch an instance.

Type: String

Valid Values: `pending | available | invalid | deregistered | transient | failed | error | disabled`

Required: No

**isPublic**

Indicates whether the AMI has public launch permissions. A value of `true`
means this AMI has public launch permissions, while `false` means it has only
implicit (AMI owner) or explicit (shared with your account) launch permissions.

Type: Boolean

Required: No

**name**

The name of the AMI.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/imagemetadata.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/imagemetadata.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/imagemetadata.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ImageDiskContainer

ImageRecycleBinInfo
