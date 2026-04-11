# Explanation

Describes an explanation code for an unreachable path. For more information, see [Reachability Analyzer explanation codes](../../../../services/vpc/latest/reachability/explanation-codes.md).

## Contents

**acl**

The network ACL.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**aclRule**

The network ACL rule.

Type: [AnalysisAclRule](api-analysisaclrule.md) object

Required: No

**address**

The IPv4 address, in CIDR notation.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 15.

Pattern: `^([0-9]{1,3}.){3}[0-9]{1,3}$`

Required: No

**AddressSet.N**

The IPv4 addresses, in CIDR notation.

Type: Array of strings

Length Constraints: Minimum length of 0. Maximum length of 15.

Pattern: `^([0-9]{1,3}.){3}[0-9]{1,3}$`

Required: No

**attachedTo**

The resource to which the component is attached.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**AvailabilityZoneIdSet.N**

The IDs of the Availability Zones.

Type: Array of strings

Required: No

**AvailabilityZoneSet.N**

The Availability Zones.

Type: Array of strings

Required: No

**CidrSet.N**

The CIDR ranges.

Type: Array of strings

Required: No

**classicLoadBalancerListener**

The listener for a Classic Load Balancer.

Type: [AnalysisLoadBalancerListener](api-analysisloadbalancerlistener.md) object

Required: No

**component**

The component.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**componentAccount**

The AWS account for the component.

Type: String

Pattern: `\d{12}`

Required: No

**componentRegion**

The Region for the component.

Type: String

Pattern: `[a-z]{2}-[a-z]+-[1-9]+`

Required: No

**customerGateway**

The customer gateway.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**destination**

The destination.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**destinationVpc**

The destination VPC.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**direction**

The direction. The following are the possible values:

- egress

- ingress

Type: String

Required: No

**elasticLoadBalancerListener**

The load balancer listener.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**explanationCode**

The explanation code.

Type: String

Required: No

**firewallStatefulRule**

The Network Firewall stateful rule.

Type: [FirewallStatefulRule](api-firewallstatefulrule.md) object

Required: No

**firewallStatelessRule**

The Network Firewall stateless rule.

Type: [FirewallStatelessRule](api-firewallstatelessrule.md) object

Required: No

**ingressRouteTable**

The route table.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**internetGateway**

The internet gateway.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**loadBalancerArn**

The Amazon Resource Name (ARN) of the load balancer.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1283.

Required: No

**loadBalancerListenerPort**

The listener port of the load balancer.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 65535.

Required: No

**loadBalancerTarget**

The target.

Type: [AnalysisLoadBalancerTarget](api-analysisloadbalancertarget.md) object

Required: No

**loadBalancerTargetGroup**

The target group.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**LoadBalancerTargetGroupSet.N**

The target groups.

Type: Array of [AnalysisComponent](api-analysiscomponent.md) objects

Required: No

**loadBalancerTargetPort**

The target port.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 65535.

Required: No

**missingComponent**

The missing component.

Type: String

Required: No

**natGateway**

The NAT gateway.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**networkInterface**

The network interface.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**packetField**

The packet field.

Type: String

Required: No

**port**

The port.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 65535.

Required: No

**PortRangeSet.N**

The port ranges.

Type: Array of [PortRange](api-portrange.md) objects

Required: No

**prefixList**

The prefix list.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**ProtocolSet.N**

The protocols.

Type: Array of strings

Required: No

**routeTable**

The route table.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**routeTableRoute**

The route table route.

Type: [AnalysisRouteTableRoute](api-analysisroutetableroute.md) object

Required: No

**securityGroup**

The security group.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**securityGroupRule**

The security group rule.

Type: [AnalysisSecurityGroupRule](api-analysissecuritygrouprule.md) object

Required: No

**SecurityGroupSet.N**

The security groups.

Type: Array of [AnalysisComponent](api-analysiscomponent.md) objects

Required: No

**sourceVpc**

The source VPC.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**state**

The state.

Type: String

Required: No

**subnet**

The subnet.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**subnetRouteTable**

The route table for the subnet.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**transitGateway**

The transit gateway.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**transitGatewayAttachment**

The transit gateway attachment.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**transitGatewayRouteTable**

The transit gateway route table.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**transitGatewayRouteTableRoute**

The transit gateway route table route.

Type: [TransitGatewayRouteTableRoute](api-transitgatewayroutetableroute.md) object

Required: No

**vpc**

The component VPC.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**vpcEndpoint**

The VPC endpoint.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**vpcPeeringConnection**

The VPC peering connection.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**vpnConnection**

The VPN connection.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**vpnGateway**

The VPN gateway.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/explanation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/explanation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/explanation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

EventInformation

ExportImageTask
