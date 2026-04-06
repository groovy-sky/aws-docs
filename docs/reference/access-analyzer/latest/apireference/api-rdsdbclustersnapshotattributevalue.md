# RdsDbClusterSnapshotAttributeValue

The values for a manual Amazon RDS DB cluster snapshot attribute.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**accountIds**

The AWS account IDs that have access to the manual Amazon RDS DB cluster snapshot. If the
value `all` is specified, then the Amazon RDS DB cluster snapshot is public and can
be copied or restored by all AWS accounts.

- If the configuration is for an existing Amazon RDS DB cluster snapshot and you do not
specify the `accountIds` in
`RdsDbClusterSnapshotAttributeValue`, then the access preview uses the
existing shared `accountIds` for the snapshot.

- If the access preview is for a new resource and you do not specify the specify the
`accountIds` in `RdsDbClusterSnapshotAttributeValue`, then
the access preview considers the snapshot without any attributes.

- To propose deletion of existing shared `accountIds`, you can specify an
empty list for `accountIds` in the
`RdsDbClusterSnapshotAttributeValue`.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/accessanalyzer-2019-11-01/RdsDbClusterSnapshotAttributeValue)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/accessanalyzer-2019-11-01/RdsDbClusterSnapshotAttributeValue)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/accessanalyzer-2019-11-01/RdsDbClusterSnapshotAttributeValue)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Position

RdsDbClusterSnapshotConfiguration
