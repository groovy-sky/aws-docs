---
title: "Event"
---

# Event

Represents a single occurrence of something interesting within the system. Some
examples of events are creating a cluster, adding or removing a cache node, or rebooting
a node.

## Contents

###### Note

In the following list, the required parameters are described first.

**Date**

The date and time when the event occurred.

Type: Timestamp

Required: No

**Message**

The text of the event.

Type: String

Required: No

**SourceIdentifier**

The identifier for the source of the event. For example, if the event occurred at the
cluster level, the identifier would be the name of the cluster.

Type: String

Required: No

**SourceType**

Specifies the origin of this event - a cluster, a parameter group, a security group,
etc.

Type: String

Valid Values: `cache-cluster | cache-parameter-group | cache-security-group | cache-subnet-group | replication-group | serverless-cache | serverless-cache-snapshot | user | user-group`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/elasticache-2015-02-02/Event)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/elasticache-2015-02-02/Event)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/elasticache-2015-02-02/Event)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EngineDefaults

Filter

All content copied from https://docs.aws.amazon.com/.
