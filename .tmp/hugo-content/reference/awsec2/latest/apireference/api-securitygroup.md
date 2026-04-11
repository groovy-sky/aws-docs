# SecurityGroup

Describes a security group.

## Contents

**groupDescription**

A description of the security group.

Type: String

Required: No

**groupId**

The ID of the security group.

Type: String

Required: No

**groupName**

The name of the security group.

Type: String

Required: No

**IpPermissions.N**

The inbound rules associated with the security group.

Type: Array of [IpPermission](api-ippermission.md) objects

Required: No

**IpPermissionsEgress.N**

The outbound rules associated with the security group.

Type: Array of [IpPermission](api-ippermission.md) objects

Required: No

**ownerId**

The AWS account ID of the owner of the security group.

Type: String

Required: No

**securityGroupArn**

The ARN of the security group.

Type: String

Required: No

**TagSet.N**

Any tags assigned to the security group.

Type: Array of [Tag](api-tag.md) objects

Required: No

**vpcId**

The ID of the VPC for the security group.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/securitygroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/securitygroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/securitygroup.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SecondarySubnetIpv4CidrBlockAssociation

SecurityGroupForVpc
