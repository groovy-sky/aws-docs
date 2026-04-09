# GlobalReplicationGroupMember

A member of a Global datastore. It contains the Replication Group Id, the Amazon
region and the role of the replication group.

## Contents

###### Note

In the following list, the required parameters are described first.

**AutomaticFailover**

Indicates whether automatic failover is enabled for the replication group.

Type: String

Valid Values: `enabled | disabled | enabling | disabling`

Required: No

**ReplicationGroupId**

The replication group id of the Global datastore member.

Type: String

Required: No

**ReplicationGroupRegion**

The Amazon region of the Global datastore member.

Type: String

Required: No

**Role**

Indicates the role of the replication group, primary or secondary.

Type: String

Required: No

**Status**

The status of the membership of the replication group.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/globalreplicationgroupmember.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/globalreplicationgroupmember.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/globalreplicationgroupmember.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GlobalReplicationGroupInfo

KinesisFirehoseDestinationDetails

All content copied from https://docs.aws.amazon.com/.
