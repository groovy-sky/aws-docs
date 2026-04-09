# DeleteGlobalTableWitnessGroupMemberAction

Specifies the action to remove a witness Region from a MRSC global table. You cannot
delete a single witness from a MRSC global table - you must delete both a replica and
the witness together. The deletion of both a witness and replica converts the remaining
replica to a single-Region DynamoDB table.

## Contents

###### Note

In the following list, the required parameters are described first.

**RegionName**

The witness Region name to be removed from the MRSC global table.

Type: String

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/deleteglobaltablewitnessgroupmemberaction.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/deleteglobaltablewitnessgroupmemberaction.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/deleteglobaltablewitnessgroupmemberaction.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteGlobalSecondaryIndexAction

DeleteReplicaAction

All content copied from https://docs.aws.amazon.com/.
