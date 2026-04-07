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

Type: [EbsInstanceBlockDeviceSpecification](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_EbsInstanceBlockDeviceSpecification.html) object

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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/InstanceBlockDeviceMappingSpecification)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/InstanceBlockDeviceMappingSpecification)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/InstanceBlockDeviceMappingSpecification)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

InstanceBlockDeviceMapping

InstanceCapacity
