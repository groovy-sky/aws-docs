# InstanceSecondaryInterfaceSpecificationRequest

Describes a secondary interface specification for launching an instance.

## Contents

**DeleteOnTermination**

Indicates whether the secondary interface is deleted when the instance is terminated.

The only supported value for this field is `true`.

Type: Boolean

Required: No

**DeviceIndex**

The device index for the secondary interface attachment.

Type: Integer

Required: No

**InterfaceType**

The type of secondary interface.

Type: String

Valid Values: `secondary`

Required: No

**NetworkCardIndex**

The index of the network card. The network card must support secondary interfaces.

Type: Integer

Required: No

**PrivateIpAddress.N**

The private IPv4 addresses to assign to the secondary interface.

Type: Array of [InstanceSecondaryInterfacePrivateIpAddressRequest](api-instancesecondaryinterfaceprivateipaddressrequest.md) objects

Required: No

**PrivateIpAddressCount**

The number of private IPv4 addresses to assign to the secondary interface.

Type: Integer

Required: No

**SecondarySubnetId**

The ID of the secondary subnet.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/instancesecondaryinterfacespecificationrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/instancesecondaryinterfacespecificationrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/instancesecondaryinterfacespecificationrequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

InstanceSecondaryInterfacePrivateIpAddressRequest

InstanceSpecification
