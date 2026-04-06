# RdsDbSnapshotAttributeValue

The name and values of a manual Amazon RDS DB snapshot attribute. Manual DB snapshot
attributes are used to authorize other AWS accounts to restore a manual DB
snapshot.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**accountIds**

The AWS account IDs that have access to the manual Amazon RDS DB snapshot. If the value
`all` is specified, then the Amazon RDS DB snapshot is public and can be copied or
restored by all AWS accounts.

- If the configuration is for an existing Amazon RDS DB snapshot and you do not specify
the `accountIds` in `RdsDbSnapshotAttributeValue`, then the
access preview uses the existing shared `accountIds` for the
snapshot.

- If the access preview is for a new resource and you do not specify the specify the
`accountIds` in `RdsDbSnapshotAttributeValue`, then the
access preview considers the snapshot without any attributes.

- To propose deletion of an existing shared `accountIds`, you can specify
an empty list for `accountIds` in the
`RdsDbSnapshotAttributeValue`.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/rdsdbsnapshotattributevalue.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/rdsdbsnapshotattributevalue.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/rdsdbsnapshotattributevalue.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RdsDbClusterSnapshotConfiguration

RdsDbSnapshotConfiguration
