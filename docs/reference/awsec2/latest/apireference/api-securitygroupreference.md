# SecurityGroupReference

Describes a VPC with a security group that references your security group.

## Contents

**groupId**

The ID of your security group.

Type: String

Required: No

**referencingVpcId**

The ID of the VPC with the referencing security group.

Type: String

Required: No

**transitGatewayId**

The ID of the transit gateway (if applicable).

Type: String

Required: No

**vpcPeeringConnectionId**

The ID of the VPC peering connection (if applicable). For more information about security group referencing for peering connections, see
[Update your security groups to reference peer security groups](../../../../services/vpc/latest/peering/vpc-peering-security-groups.md)
in the _VPC Peering Guide_.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/securitygroupreference.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/securitygroupreference.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/securitygroupreference.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SecurityGroupIdentifier

SecurityGroupRule
