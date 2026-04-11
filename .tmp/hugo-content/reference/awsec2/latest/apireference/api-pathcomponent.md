# PathComponent

Describes a path component.

## Contents

**aclRule**

The network ACL rule.

Type: [AnalysisAclRule](api-analysisaclrule.md) object

Required: No

**AdditionalDetailSet.N**

The additional details.

Type: Array of [AdditionalDetail](api-additionaldetail.md) objects

Required: No

**attachedTo**

The resource to which the path component is attached.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**component**

The component.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**destinationVpc**

The destination VPC.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**elasticLoadBalancerListener**

The load balancer listener.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**ExplanationSet.N**

The explanation codes.

Type: Array of [Explanation](api-explanation.md) objects

Required: No

**firewallStatefulRule**

The Network Firewall stateful rule.

Type: [FirewallStatefulRule](api-firewallstatefulrule.md) object

Required: No

**firewallStatelessRule**

The Network Firewall stateless rule.

Type: [FirewallStatelessRule](api-firewallstatelessrule.md) object

Required: No

**inboundHeader**

The inbound header.

Type: [AnalysisPacketHeader](api-analysispacketheader.md) object

Required: No

**outboundHeader**

The outbound header.

Type: [AnalysisPacketHeader](api-analysispacketheader.md) object

Required: No

**routeTableRoute**

The route table route.

Type: [AnalysisRouteTableRoute](api-analysisroutetableroute.md) object

Required: No

**securityGroupRule**

The security group rule.

Type: [AnalysisSecurityGroupRule](api-analysissecuritygrouprule.md) object

Required: No

**sequenceNumber**

The sequence number.

Type: Integer

Required: No

**serviceName**

The name of the VPC endpoint service.

Type: String

Required: No

**sourceVpc**

The source VPC.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**subnet**

The subnet.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**transitGateway**

The transit gateway.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**transitGatewayRouteTableRoute**

The route in a transit gateway route table.

Type: [TransitGatewayRouteTableRoute](api-transitgatewayroutetableroute.md) object

Required: No

**vpc**

The component VPC.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/pathcomponent.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/pathcomponent.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/pathcomponent.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

PacketHeaderStatementRequest

PathFilter
