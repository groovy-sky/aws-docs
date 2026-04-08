# DBSnapshotAttribute

Contains the name and values of a manual DB snapshot attribute

Manual DB snapshot attributes are used to authorize other AWS accounts
to restore a manual DB snapshot. For more information, see the `ModifyDBSnapshotAttribute`
API.

## Contents

###### Note

In the following list, the required parameters are described first.

**AttributeName**

The name of the manual DB snapshot attribute.

The attribute named `restore` refers to the list of AWS accounts that
have permission to copy or restore the manual DB cluster snapshot. For more information,
see the `ModifyDBSnapshotAttribute`
API action.

Type: String

Required: No

**AttributeValues.AttributeValue.N**

The value or values for the manual DB snapshot attribute.

If the `AttributeName` field is set to `restore`, then this element
returns a list of IDs of the AWS accounts that are authorized to copy or restore the manual
DB snapshot. If a value of `all` is in the list, then the manual DB snapshot
is public and available for any AWS account to copy or restore.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/dbsnapshotattribute.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/dbsnapshotattribute.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/dbsnapshotattribute.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DBSnapshot

DBSnapshotAttributesResult

All content copied from https://docs.aws.amazon.com/.
