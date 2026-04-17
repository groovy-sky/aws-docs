---
title: "GlobalClusterMember"
---

# GlobalClusterMember

A data structure with information about any primary and
secondary clusters associated with a global cluster (Aurora global database).

## Contents

###### Note

In the following list, the required parameters are described first.

**DBClusterArn**

The Amazon Resource Name (ARN) for each Aurora DB cluster in the global cluster.

Type: String

Required: No

**GlobalWriteForwardingStatus**

The status of write forwarding for a secondary cluster in the global cluster.

Type: String

Valid Values: `enabled | disabled | enabling | disabling | unknown`

Required: No

**IsWriter**

Indicates whether the Aurora DB cluster is the primary cluster
(that is, has read-write capability) for the global
cluster with which it is associated.

Type: Boolean

Required: No

**Readers.member.N**

The Amazon Resource Name (ARN) for each read-only secondary cluster
associated with the global cluster.

Type: Array of strings

Required: No

**SynchronizationStatus**

The status of synchronization of each Aurora DB cluster in the global cluster.

Type: String

Valid Values: `connected | pending-resync`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/rds-2014-10-31/GlobalClusterMember)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/rds-2014-10-31/GlobalClusterMember)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/rds-2014-10-31/GlobalClusterMember)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GlobalCluster

Integration

All content copied from https://docs.aws.amazon.com/.
