# SecondaryInterface

Describes a secondary interface.

## Contents

**attachment**

The attachment information for the secondary interface.

Type: [SecondaryInterfaceAttachment](api-secondaryinterfaceattachment.md) object

Required: No

**availabilityZone**

The Availability Zone of the secondary interface.

Type: String

Required: No

**availabilityZoneId**

The ID of the Availability Zone of the secondary interface.

Type: String

Required: No

**macAddress**

The MAC address of the secondary interface.

Type: String

Required: No

**ownerId**

The ID of the AWS account that owns the secondary interface.

Type: String

Required: No

**PrivateIpv4AddressSet.N**

The private IPv4 addresses associated with the secondary interface.

Type: Array of [SecondaryInterfaceIpv4Address](api-secondaryinterfaceipv4address.md) objects

Required: No

**secondaryInterfaceArn**

The Amazon Resource Name (ARN) of the secondary interface.

Type: String

Required: No

**secondaryInterfaceId**

The ID of the secondary interface.

Type: String

Required: No

**secondaryInterfaceType**

The type of secondary interface.

Type: String

Valid Values: `secondary`

Required: No

**secondaryNetworkId**

The ID of the secondary network.

Type: String

Required: No

**secondaryNetworkType**

The type of the secondary network.

Type: String

Valid Values: `rdma`

Required: No

**secondarySubnetId**

The ID of the secondary subnet.

Type: String

Required: No

**sourceDestCheck**

Indicates whether source/destination checking is enabled.

Type: Boolean

Required: No

**status**

The status of the secondary interface.

Type: String

Valid Values: `available | in-use`

Required: No

**TagSet.N**

The tags assigned to the secondary interface.

Type: Array of [Tag](api-tag.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/secondaryinterface.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/secondaryinterface.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/secondaryinterface.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ScheduledInstancesPrivateIpAddressConfig

SecondaryInterfaceAttachment
