# InstanceTopology

Information about the instance topology.

## Contents

**availabilityZone**

The name of the Availability Zone or Local Zone that the instance is in.

Type: String

Required: No

**capacityBlockId**

The ID of the Capacity Block. This parameter is only supported for UltraServer
instances and identifies instances within the UltraServer domain.

Type: String

Required: No

**groupName**

The name of the placement group that the instance is in.

Type: String

Required: No

**instanceId**

The instance ID.

Type: String

Required: No

**instanceType**

The instance type.

Type: String

Required: No

**NetworkNodeSet.N**

The network nodes. The nodes are hashed based on your account. Instances from
different accounts running under the same server will return a different hashed list of
strings.

The value is `null` or empty if:

- The instance type is not supported.

- The instance is in a state other than `running`.

Type: Array of strings

Required: No

**zoneId**

The ID of the Availability Zone or Local Zone that the instance is in.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/instancetopology.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/instancetopology.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/instancetopology.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

InstanceTagNotificationAttribute

InstanceTypeInfo
