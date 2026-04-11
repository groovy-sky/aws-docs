# SecurityGroupVpcAssociation

A security group association with a VPC that you made with [AssociateSecurityGroupVpc](api-associatesecuritygroupvpc.md).

## Contents

**groupId**

The association's security group ID.

Type: String

Required: No

**groupOwnerId**

The AWS account ID of the owner of the security group.

Type: String

Required: No

**state**

The association's state.

Type: String

Valid Values: `associating | associated | association-failed | disassociating | disassociated | disassociation-failed`

Required: No

**stateReason**

The association's state reason.

Type: String

Required: No

**vpcId**

The association's VPC ID.

Type: String

Required: No

**vpcOwnerId**

The AWS account ID of the owner of the VPC.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/securitygroupvpcassociation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/securitygroupvpcassociation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/securitygroupvpcassociation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SecurityGroupRuleUpdate

ServiceConfiguration
