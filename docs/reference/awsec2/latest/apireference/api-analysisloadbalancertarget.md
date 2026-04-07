# AnalysisLoadBalancerTarget

Describes a load balancer target.

## Contents

**address**

The IP address.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 15.

Pattern: `^([0-9]{1,3}.){3}[0-9]{1,3}$`

Required: No

**availabilityZone**

The Availability Zone.

Type: String

Required: No

**availabilityZoneId**

The ID of the Availability Zone.

Type: String

Required: No

**instance**

Information about the instance.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**port**

The port on which the target is listening.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 65535.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/analysisloadbalancertarget.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/analysisloadbalancertarget.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/analysisloadbalancertarget.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AnalysisLoadBalancerListener

AnalysisPacketHeader
