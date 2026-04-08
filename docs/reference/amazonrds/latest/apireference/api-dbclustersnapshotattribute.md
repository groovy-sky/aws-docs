# DBClusterSnapshotAttribute

Contains the name and values of a manual DB cluster snapshot attribute.

Manual DB cluster snapshot attributes are used to authorize other AWS accounts
to restore a manual DB cluster snapshot. For more information, see the `ModifyDBClusterSnapshotAttribute`
API action.

## Contents

###### Note

In the following list, the required parameters are described first.

**AttributeName**

The name of the manual DB cluster snapshot attribute.

The attribute named `restore` refers to the list of AWS accounts that
have permission to copy or restore the manual DB cluster snapshot. For more information,
see the `ModifyDBClusterSnapshotAttribute`
API action.

Type: String

Required: No

**AttributeValues.AttributeValue.N**

The value(s) for the manual DB cluster snapshot attribute.

If the `AttributeName` field is set to `restore`, then this element
returns a list of IDs of the AWS accounts that are authorized to copy or restore the manual
DB cluster snapshot. If a value of `all` is in the list, then the manual DB cluster snapshot
is public and available for any AWS account to copy or restore.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/dbclustersnapshotattribute.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/dbclustersnapshotattribute.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/dbclustersnapshotattribute.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DBClusterSnapshot

DBClusterSnapshotAttributesResult

All content copied from https://docs.aws.amazon.com/.
