# GlobalTableWitnessGroupUpdate

Represents one of the following:

- A new witness to be added to a new global table.

- An existing witness to be removed from an existing global table.

You can configure one witness per MRSC global table.

## Contents

###### Note

In the following list, the required parameters are described first.

**Create**

Specifies a witness Region to be added to a new MRSC global table. The witness must be
added when creating the MRSC global table.

Type: [CreateGlobalTableWitnessGroupMemberAction](api-createglobaltablewitnessgroupmemberaction.md) object

Required: No

**Delete**

Specifies a witness Region to be removed from an existing global table. Must be done
in conjunction with removing a replica. The deletion of both a witness and replica
converts the remaining replica to a single-Region DynamoDB table.

Type: [DeleteGlobalTableWitnessGroupMemberAction](api-deleteglobaltablewitnessgroupmemberaction.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/globaltablewitnessgroupupdate.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/globaltablewitnessgroupupdate.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/globaltablewitnessgroupupdate.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GlobalTableWitnessDescription

ImportSummary

All content copied from https://docs.aws.amazon.com/.
