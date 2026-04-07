# AvailabilityZoneAddress

For regional NAT gateways only: The configuration specifying which Elastic IP address (EIP) to use for handling outbound NAT traffic from a specific Availability Zone.

A regional NAT gateway is a single NAT Gateway that works across multiple availability zones (AZs) in your VPC, providing redundancy, scalability and availability across all the AZs in a Region.

For more information, see [Regional NAT gateways for automatic multi-AZ expansion](../../../../services/vpc/latest/userguide/nat-gateways-regional.md) in the _Amazon VPC User Guide_.

## Contents

**AllocationId.N**

The allocation IDs of the Elastic IP addresses (EIPs) to be used for handling outbound NAT traffic in this specific Availability Zone.

Type: Array of strings

Required: No

**AvailabilityZone**

For regional NAT gateways only: The Availability Zone where this specific NAT gateway configuration will be active. Each AZ in a regional NAT gateway has its own configuration to handle outbound NAT traffic from that AZ.

A regional NAT gateway is a single NAT Gateway that works across multiple availability zones (AZs) in your VPC, providing redundancy, scalability and availability across all the AZs in a Region.

Type: String

Required: No

**AvailabilityZoneId**

For regional NAT gateways only: The ID of the Availability Zone where this specific NAT gateway configuration will be active. Each AZ in a regional NAT gateway has its own configuration to handle outbound NAT traffic from that AZ. Use this instead of AvailabilityZone for consistent identification of AZs across AWS Regions.

A regional NAT gateway is a single NAT Gateway that works across multiple availability zones (AZs) in your VPC, providing redundancy, scalability and availability across all the AZs in a Region.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/availabilityzoneaddress.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/availabilityzoneaddress.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/availabilityzoneaddress.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AvailabilityZone

AvailabilityZoneGeography
