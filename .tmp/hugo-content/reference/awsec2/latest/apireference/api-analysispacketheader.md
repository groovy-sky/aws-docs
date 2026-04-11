# AnalysisPacketHeader

Describes a header. Reflects any changes made by a component as traffic passes through.
The fields of an inbound header are null except for the first component of a path.

## Contents

**DestinationAddressSet.N**

The destination addresses.

Type: Array of strings

Length Constraints: Minimum length of 0. Maximum length of 15.

Pattern: `^([0-9]{1,3}.){3}[0-9]{1,3}$`

Required: No

**DestinationPortRangeSet.N**

The destination port ranges.

Type: Array of [PortRange](api-portrange.md) objects

Required: No

**protocol**

The protocol.

Type: String

Required: No

**SourceAddressSet.N**

The source addresses.

Type: Array of strings

Length Constraints: Minimum length of 0. Maximum length of 15.

Pattern: `^([0-9]{1,3}.){3}[0-9]{1,3}$`

Required: No

**SourcePortRangeSet.N**

The source port ranges.

Type: Array of [PortRange](api-portrange.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/analysispacketheader.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/analysispacketheader.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/analysispacketheader.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AnalysisLoadBalancerTarget

AnalysisRouteTableRoute
