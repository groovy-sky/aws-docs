# RdsDbClusterSnapshotConfiguration

The proposed access control configuration for an Amazon RDS DB cluster snapshot. You can
propose a configuration for a new Amazon RDS DB cluster snapshot or an Amazon RDS DB cluster snapshot
that you own by specifying the `RdsDbClusterSnapshotAttributeValue` and optional
AWS KMS encryption key. For more information, see [ModifyDBClusterSnapshotAttribute](../../../amazonrds/latest/apireference/api-modifydbclustersnapshotattribute.md).

## Contents

**attributes**

The names and values of manual DB cluster snapshot attributes. Manual DB cluster
snapshot attributes are used to authorize other AWS accounts to restore a manual DB
cluster snapshot. The only valid value for `AttributeName` for the attribute map
is `restore`

Type: String to [RdsDbClusterSnapshotAttributeValue](api-rdsdbclustersnapshotattributevalue.md) object map

Required: No

**kmsKeyId**

The KMS key identifier for an encrypted Amazon RDS DB cluster snapshot. The KMS key
identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.

- If the configuration is for an existing Amazon RDS DB cluster snapshot and you do not
specify the `kmsKeyId`, or you specify an empty string, then the access
preview uses the existing `kmsKeyId` of the snapshot.

- If the access preview is for a new resource and you do not specify the specify the
`kmsKeyId`, then the access preview considers the snapshot as
unencrypted.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/rdsdbclustersnapshotconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/rdsdbclustersnapshotconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/rdsdbclustersnapshotconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RdsDbClusterSnapshotAttributeValue

RdsDbSnapshotAttributeValue
