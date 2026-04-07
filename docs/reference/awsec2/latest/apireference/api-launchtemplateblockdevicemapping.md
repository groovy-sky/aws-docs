# LaunchTemplateBlockDeviceMapping

Describes a block device mapping.

## Contents

**deviceName**

The device name.

Type: String

Required: No

**ebs**

Information about the block device for an EBS volume.

Type: [LaunchTemplateEbsBlockDevice](api-launchtemplateebsblockdevice.md) object

Required: No

**noDevice**

To omit the device from the block device mapping, specify an empty string.

Type: String

Required: No

**virtualName**

The virtual device name (ephemeralN).

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/launchtemplateblockdevicemapping.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/launchtemplateblockdevicemapping.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/launchtemplateblockdevicemapping.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LaunchTemplateAndOverridesResponse

LaunchTemplateBlockDeviceMappingRequest
