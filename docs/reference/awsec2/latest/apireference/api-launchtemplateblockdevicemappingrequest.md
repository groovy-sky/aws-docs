# LaunchTemplateBlockDeviceMappingRequest

Describes a block device mapping.

## Contents

**DeviceName**

The device name (for example, /dev/sdh or xvdh).

Type: String

Required: No

**Ebs**

Parameters used to automatically set up EBS volumes when the instance is
launched.

Type: [LaunchTemplateEbsBlockDeviceRequest](api-launchtemplateebsblockdevicerequest.md) object

Required: No

**NoDevice**

To omit the device from the block device mapping, specify an empty string.

Type: String

Required: No

**VirtualName**

The virtual device name (ephemeralN). Instance store volumes are numbered starting
from 0. An instance type with 2 available instance store volumes can specify mappings
for ephemeral0 and ephemeral1. The number of available instance store volumes depends on
the instance type. After you connect to the instance, you must mount the volume.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/launchtemplateblockdevicemappingrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/launchtemplateblockdevicemappingrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/launchtemplateblockdevicemappingrequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LaunchTemplateBlockDeviceMapping

LaunchTemplateCapacityReservationSpecificationRequest
