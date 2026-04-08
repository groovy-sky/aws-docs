# DBClusterSnapshotAttributesResult

Contains the results of a successful call to the `DescribeDBClusterSnapshotAttributes`
API action.

Manual DB cluster snapshot attributes are used to authorize other AWS accounts
to copy or restore a manual DB cluster snapshot. For more information, see the `ModifyDBClusterSnapshotAttribute`
API action.

## Contents

###### Note

In the following list, the required parameters are described first.

**DBClusterSnapshotAttributes.DBClusterSnapshotAttribute.N**

The list of attributes and values for the manual DB cluster snapshot.

Type: Array of [DBClusterSnapshotAttribute](api-dbclustersnapshotattribute.md) objects

Required: No

**DBClusterSnapshotIdentifier**

The identifier of the manual DB cluster snapshot that the attributes apply to.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/dbclustersnapshotattributesresult.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/dbclustersnapshotattributesresult.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/dbclustersnapshotattributesresult.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DBClusterSnapshotAttribute

DBClusterStatusInfo

All content copied from https://docs.aws.amazon.com/.
