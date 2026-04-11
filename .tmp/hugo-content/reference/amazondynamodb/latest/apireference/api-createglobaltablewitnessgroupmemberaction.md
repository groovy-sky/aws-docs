# CreateGlobalTableWitnessGroupMemberAction

Specifies the action to add a new witness Region to a MRSC global table. A MRSC global
table can be configured with either three replicas, or with two replicas and one
witness.

## Contents

###### Note

In the following list, the required parameters are described first.

**RegionName**

The AWS Region name to be added as a witness Region for the MRSC global
table. The witness must be in a different Region than the replicas and within the same
Region set:

- US Region set: US East (N. Virginia), US East (Ohio), US West (Oregon)

- EU Region set: Europe (Ireland), Europe (London), Europe (Paris), Europe
(Frankfurt)

- AP Region set: Asia Pacific (Tokyo), Asia Pacific (Seoul), Asia Pacific
(Osaka)

Type: String

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/createglobaltablewitnessgroupmemberaction.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/createglobaltablewitnessgroupmemberaction.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/createglobaltablewitnessgroupmemberaction.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateGlobalSecondaryIndexAction

CreateReplicaAction

All content copied from https://docs.aws.amazon.com/.
