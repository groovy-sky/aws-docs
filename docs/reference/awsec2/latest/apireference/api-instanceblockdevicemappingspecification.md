# InstanceBlockDeviceMappingSpecification

Describes a block device mapping entry.

## Contents

**DeviceName**

The device name. For available device names, see [Device names for volumes](../../../../services/ec2/latest/userguide/device-naming.md).

Type: String

Required: No

**Ebs**

Parameters used to automatically set up EBS volumes when the instance is
launched.

Type: [EbsInstanceBlockDeviceSpecification](api-ebsinstanceblockdevicespecification.md) object

Required: No

**NoDevice**

Suppresses the specified device included in the block device mapping.

Type: String

Required: No

**VirtualName**

The virtual device name.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/instanceblockdevicemappingspecification.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/instanceblockdevicemappingspecification.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/instanceblockdevicemappingspecification.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

InstanceBlockDeviceMapping

InstanceCapacity
