# ReplicationGroupUpdate

Represents one of the following:

- A new replica to be added to an existing regional table or global table. This
request invokes the `CreateTableReplica` action in the destination
Region.

- New parameters for an existing replica. This request invokes the
`UpdateTable` action in the destination Region.

- An existing replica to be deleted. The request invokes the
`DeleteTableReplica` action in the destination Region, deleting
the replica and all if its items in the destination Region.

###### Note

When you manually remove a table or global table replica, you do not automatically
remove any associated scalable targets, scaling policies, or CloudWatch
alarms.

## Contents

###### Note

In the following list, the required parameters are described first.

**Create**

The parameters required for creating a replica for the table.

Type: [CreateReplicationGroupMemberAction](api-createreplicationgroupmemberaction.md) object

Required: No

**Delete**

The parameters required for deleting a replica for the table.

Type: [DeleteReplicationGroupMemberAction](api-deletereplicationgroupmemberaction.md) object

Required: No

**Update**

The parameters required for updating a replica for the table.

Type: [UpdateReplicationGroupMemberAction](api-updatereplicationgroupmemberaction.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/replicationgroupupdate.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/replicationgroupupdate.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/replicationgroupupdate.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicaSettingsUpdate

ReplicaUpdate

All content copied from https://docs.aws.amazon.com/.
