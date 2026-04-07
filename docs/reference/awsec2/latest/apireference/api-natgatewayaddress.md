# NatGatewayAddress

Describes the IP addresses and network interface associated with a NAT gateway.

## Contents

**allocationId**

\[Public NAT gateway only\] The allocation ID of the Elastic IP address that's associated with the NAT gateway.

Type: String

Required: No

**associationId**

\[Public NAT gateway only\] The association ID of the Elastic IP address that's associated with the NAT gateway.

Type: String

Required: No

**availabilityZone**

The Availability Zone where this Elastic IP address (EIP) is being used to handle outbound NAT traffic.

Type: String

Required: No

**availabilityZoneId**

The ID of the Availability Zone where this Elastic IP address (EIP) is being used to handle outbound NAT traffic. Use this instead of AvailabilityZone for consistent identification of AZs across AWS Regions.

Type: String

Required: No

**failureMessage**

The address failure message.

Type: String

Required: No

**isPrimary**

Defines if the IP address is the primary address.

Type: Boolean

Required: No

**networkInterfaceId**

The ID of the network interface associated with the NAT gateway.

Type: String

Required: No

**privateIp**

The private IP address associated with the NAT gateway.

Type: String

Required: No

**publicIp**

\[Public NAT gateway only\] The Elastic IP address associated with the NAT gateway.

Type: String

Required: No

**status**

The address status.

Type: String

Valid Values: `assigning | unassigning | associating | disassociating | succeeded | failed`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/natgatewayaddress.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/natgatewayaddress.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/natgatewayaddress.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

NatGateway

NatGatewayAttachedAppliance
