# BlockDeviceMappingResponse

Describes a block device mapping, which defines the EBS volumes and instance store
volumes to attach to an instance at launch.

## Contents

**deviceName**

The device name (for example, `/dev/sdh` or `xvdh`).

Type: String

Required: No

**ebs**

Parameters used to automatically set up EBS volumes when the instance is
launched.

Type: [EbsBlockDeviceResponse](api-ebsblockdeviceresponse.md) object

Required: No

**noDevice**

Suppresses the specified device included in the block device mapping.

Type: String

Required: No

**virtualName**

The virtual device name.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/blockdevicemappingresponse.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/blockdevicemappingresponse.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/blockdevicemappingresponse.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

BlockDeviceMapping

BlockPublicAccessStates
