# ScheduledInstancesBlockDeviceMapping

Describes a block device mapping for a Scheduled Instance.

## Contents

**DeviceName**

The device name (for example, `/dev/sdh` or `xvdh`).

Type: String

Required: No

**Ebs**

Parameters used to set up EBS volumes automatically when the instance is launched.

Type: [ScheduledInstancesEbs](api-scheduledinstancesebs.md) object

Required: No

**NoDevice**

To omit the device from the block device mapping, specify an empty string.

Type: String

Required: No

**VirtualName**

The virtual device name ( `ephemeral` N). Instance store volumes are numbered
starting from 0. An instance type with two available instance store volumes can specify mappings
for `ephemeral0` and `ephemeral1`. The number of available instance store
volumes depends on the instance type. After you connect to the instance, you must mount the
volume.

Constraints: For M3 instances, you must specify instance store volumes in the block device
mapping for the instance. When you launch an M3 instance, we ignore any instance store volumes
specified in the block device mapping for the AMI.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/scheduledinstancesblockdevicemapping.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/scheduledinstancesblockdevicemapping.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/scheduledinstancesblockdevicemapping.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ScheduledInstanceRecurrenceRequest

ScheduledInstancesEbs
