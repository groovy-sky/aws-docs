# BlockPublicAccessStates

The state of VPC Block Public Access (BPA).

## Contents

**internetGatewayBlockMode**

The mode of VPC BPA.

- `off`: VPC BPA is not enabled and traffic is allowed to and from internet gateways and egress-only internet gateways in this Region.

- `block-bidirectional`: Block all traffic to and from internet gateways and egress-only internet gateways in this Region (except for excluded VPCs and subnets).

- `block-ingress`: Block all internet traffic to the VPCs in this Region (except for VPCs or subnets which are excluded). Only traffic to and from NAT gateways and egress-only internet gateways is allowed because these gateways only allow outbound connections to be established.

Type: String

Valid Values: `off | block-bidirectional | block-ingress`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/blockpublicaccessstates.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/blockpublicaccessstates.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/blockpublicaccessstates.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

BlockDeviceMappingResponse

BundleTask
