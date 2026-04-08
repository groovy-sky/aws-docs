# DBSnapshotAttributesResult

Contains the results of a successful call to the `DescribeDBSnapshotAttributes`
API action.

Manual DB snapshot attributes are used to authorize other AWS accounts
to copy or restore a manual DB snapshot. For more information, see the `ModifyDBSnapshotAttribute`
API action.

## Contents

###### Note

In the following list, the required parameters are described first.

**DBSnapshotAttributes.DBSnapshotAttribute.N**

The list of attributes and values for the manual DB snapshot.

Type: Array of [DBSnapshotAttribute](api-dbsnapshotattribute.md) objects

Required: No

**DBSnapshotIdentifier**

The identifier of the manual DB snapshot that the attributes apply to.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/dbsnapshotattributesresult.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/dbsnapshotattributesresult.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/dbsnapshotattributesresult.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DBSnapshotAttribute

DBSnapshotTenantDatabase

All content copied from https://docs.aws.amazon.com/.
