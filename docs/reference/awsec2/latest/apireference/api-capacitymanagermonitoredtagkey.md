---
title: "CapacityManagerMonitoredTagKey"
---

# CapacityManagerMonitoredTagKey

Describes a tag key that is being monitored by Capacity Manager, including its activation status and the earliest available data point.

## Contents

**capacityManagerProvided**

Indicates whether this tag key is provided by Capacity Manager by default, rather than being user-activated.

Type: Boolean

Required: No

**earliestDatapointTimestamp**

The earliest timestamp from which tag data is available for queries, in UTC ISO 8601 format.

Type: Timestamp

Required: No

**status**

The current status of the monitored tag key. Valid values are `activating`, `activated`, `deactivating`, and `suspended`.

Type: String

Valid Values: `activating | activated | deactivating | suspended`

Required: No

**statusMessage**

A message providing additional details about the current status of the monitored tag key.

Type: String

Required: No

**tagKey**

The tag key being monitored.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/capacitymanagermonitoredtagkey.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/capacitymanagermonitoredtagkey.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/capacitymanagermonitoredtagkey.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CapacityManagerDimension

CapacityManagerTagDimension

All content copied from https://docs.aws.amazon.com/.
