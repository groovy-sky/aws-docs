# NetworkInterfaceAssociation

Describes association information for an Elastic IP address (IPv4 only), or a Carrier
IP address (for a network interface which resides in a subnet in a Wavelength
Zone).

## Contents

**allocationId**

The allocation ID.

Type: String

Required: No

**associationId**

The association ID.

Type: String

Required: No

**carrierIp**

The carrier IP address associated with the network interface.

This option is only available when the network interface is in a subnet which is
associated with a Wavelength Zone.

Type: String

Required: No

**customerOwnedIp**

The customer-owned IP address associated with the network interface.

Type: String

Required: No

**ipOwnerId**

The ID of the Elastic IP address owner.

Type: String

Required: No

**publicDnsName**

The public DNS name.

Type: String

Required: No

**publicIp**

The address of the Elastic IP address bound to the network interface.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/networkinterfaceassociation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/networkinterfaceassociation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/networkinterfaceassociation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

NetworkInterface

NetworkInterfaceAttachment
