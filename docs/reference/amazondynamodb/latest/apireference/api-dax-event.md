# Event

Represents a single occurrence of something interesting within the system. Some
examples of events are creating a DAX cluster, adding or removing a node, or rebooting a
node.

## Contents

###### Note

In the following list, the required parameters are described first.

**Date**

The date and time when the event occurred.

Type: Timestamp

Required: No

**Message**

A user-defined message associated with the event.

Type: String

Required: No

**SourceName**

The source of the event. For example, if the event occurred at the node level, the
source would be the node ID.

Type: String

Required: No

**SourceType**

Specifies the origin of this event - a cluster, a parameter group, a node ID,
etc.

Type: String

Valid Values: `CLUSTER | PARAMETER_GROUP | SUBNET_GROUP`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dax-2017-04-19/event.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dax-2017-04-19/event.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dax-2017-04-19/event.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Endpoint

Node

All content copied from https://docs.aws.amazon.com/.
