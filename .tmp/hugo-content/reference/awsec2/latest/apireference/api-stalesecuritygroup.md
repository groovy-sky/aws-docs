# StaleSecurityGroup

Describes a stale security group (a security group that contains stale rules).

## Contents

**description**

The description of the security group.

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

**StaleIpPermissions.N**

Information about the stale inbound rules in the security group.

Type: Array of [StaleIpPermission](api-staleippermission.md) objects

Required: No

**StaleIpPermissionsEgress.N**

Information about the stale outbound rules in the security group.

Type: Array of [StaleIpPermission](api-staleippermission.md) objects

Required: No

**vpcId**

The ID of the VPC for the security group.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/stalesecuritygroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/stalesecuritygroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/stalesecuritygroup.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

StaleIpPermission

StateReason
