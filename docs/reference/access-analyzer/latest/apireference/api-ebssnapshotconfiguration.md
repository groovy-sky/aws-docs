# EbsSnapshotConfiguration

The proposed access control configuration for an Amazon EBS volume snapshot. You can propose
a configuration for a new Amazon EBS volume snapshot or an Amazon EBS volume snapshot that you own by
specifying the user IDs, groups, and optional AWS KMS encryption key. For more information,
see [ModifySnapshotAttribute](../../../awsec2/latest/apireference/api-modifysnapshotattribute.md).

## Contents

**groups**

The groups that have access to the Amazon EBS volume snapshot. If the value `all`
is specified, then the Amazon EBS volume snapshot is public.

- If the configuration is for an existing Amazon EBS volume snapshot and you do not
specify the `groups`, then the access preview uses the existing shared
`groups` for the snapshot.

- If the access preview is for a new resource and you do not specify the
`groups`, then the access preview considers the snapshot without any
`groups`.

- To propose deletion of existing shared `groups`, you can specify an
empty list for `groups`.

Type: Array of strings

Required: No

**kmsKeyId**

The KMS key identifier for an encrypted Amazon EBS volume snapshot. The KMS key
identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.

- If the configuration is for an existing Amazon EBS volume snapshot and you do not
specify the `kmsKeyId`, or you specify an empty string, then the access
preview uses the existing `kmsKeyId` of the snapshot.

- If the access preview is for a new resource and you do not specify the
`kmsKeyId`, the access preview considers the snapshot as
unencrypted.

Type: String

Required: No

**userIds**

The IDs of the AWS accounts that have access to the Amazon EBS volume snapshot.

- If the configuration is for an existing Amazon EBS volume snapshot and you do not
specify the `userIds`, then the access preview uses the existing shared
`userIds` for the snapshot.

- If the access preview is for a new resource and you do not specify the
`userIds`, then the access preview considers the snapshot without any
`userIds`.

- To propose deletion of existing shared `accountIds`, you can specify an
empty list for `userIds`.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/ebssnapshotconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/ebssnapshotconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/ebssnapshotconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DynamodbTableConfiguration

EcrRepositoryConfiguration
