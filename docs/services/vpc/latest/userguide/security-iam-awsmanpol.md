---
title: "AWS managed policies for Amazon Virtual Private Cloud"
---

# AWS managed policies for Amazon Virtual Private Cloud

An AWS managed policy is a standalone policy that is created and administered by AWS. AWS managed policies are designed
to provide permissions for many common use cases so that you can start assigning permissions to users, groups, and roles.

Keep in mind that AWS managed policies might not grant least-privilege permissions for your specific use cases because
they're available for all AWS customers to use. We recommend that you reduce permissions further by defining
[customer managed policies](../../../iam/latest/userguide/access-policies-managed-vs-inline.md#customer-managed-policies) that are specific to your use cases.

You cannot change the permissions defined in AWS managed policies. If AWS updates the permissions defined in an AWS
managed policy, the update affects all principal identities (users, groups, and roles) that the policy is attached to. AWS is
most likely to update an AWS managed policy when a new AWS service is launched or new API operations become available for
existing services.

For more information, see [AWS managed policies](../../../iam/latest/userguide/access-policies-managed-vs-inline.md#aws-managed-policies) in the
_IAM User Guide_.

## AWS managed policy: AmazonVPCFullAccess

You can attach the `AmazonVPCFullAccess` policy to your IAM identities.
This policy grants permissions that allow full access to Amazon VPC.

To view the permissions for this policy, see [AmazonVPCFullAccess](../../../aws-managed-policy/latest/reference/amazonvpcfullaccess.md)
in the _AWS Managed Policy Reference_.

## AWS managed policy: AmazonVPCReadOnlyAccess

You can attach the `AmazonVPCReadOnlyAccess` policy to your IAM identities.
This policy grants permissions that allow read-only access to Amazon VPC.

To view the permissions for this policy, see [AmazonVPCReadOnlyAccess](../../../aws-managed-policy/latest/reference/amazonvpcreadonlyaccess.md)
in the _AWS Managed Policy Reference_.

## AWS managed policy: AmazonVPCCrossAccountNetworkInterfaceOperations

You can attach the `AmazonVPCCrossAccountNetworkInterfaceOperations` policy to your IAM identities.
This policy grants permissions that allow the identity to create network interfaces and attach them
to cross-account resources.

To view the permissions for this policy, see [AmazonVPCCrossAccountNetworkInterfaceOperations](../../../aws-managed-policy/latest/reference/amazonvpccrossaccountnetworkinterfaceoperations.md)
in the _AWS Managed Policy Reference_.

## AWS managed policy: AWSServiceRoleForNATGateway

You can attach the `AWSServiceRoleForNATGateway` policy to your IAM identities.
This policy grants permissions that allow the identity to work on your behalf to automatically scale regional NAT Gateways.

To view the permissions for this policy, see [AWSServiceRoleForNATGateway](../../../aws-managed-policy/latest/reference/awsservicerolefornatgateway.md)
in the _AWS Managed Policy Reference_.

## Amazon VPC updates to AWS managed policies

View details about updates to AWS managed policies for Amazon VPC since this service
began tracking these changes in March 2021.

ChangeDescriptionDate[AWS managed policy: AmazonVPCFullAccess](#security-iam-awsmanpol-AmazonVPCFullAccess) –
Update to an existing policyActions added to the AWSIPAMServiceRolePolicy managed policy (ec2:ModifyManagedPrefixList, ec2:DescribeManagedPrefixLists, and ec2:GetManagedPrefixListEntries) to enable IPAM modify and read managed prefix lists.October 31, 2025[AWS managed policy: AWSServiceRoleForNATGateway](#security-iam-awsmanpol-AWSServiceRoleForNATGateway) –
New policyNew AWSServiceRoleForNATGateway policy allows the identity to automatically scale regional NAT Gateways.November 19, 2025[AWS managed policy: AmazonVPCFullAccess](#security-iam-awsmanpol-AmazonVPCFullAccess) –
Update to an existing policyAdded the AssociateSecurityGroupVpc,
DescribeSecurityGroupVpcAssociations, and
DisassociateSecurityGroupVpc actions, which allow you
to associate, disassociate, and view security group associations with VPCs.December 9, 2024[AWS managed policy: AmazonVPCReadOnlyAccess](#security-iam-awsmanpol-AmazonVPCReadOnlyAccess) –
Update to an existing policyAdded the DescribeSecurityGroupVpcAssociations action, which allows you to view security group associations with VPCs.December 9, 2024[AWS managed policy: AmazonVPCFullAccess](#security-iam-awsmanpol-AmazonVPCFullAccess) –
Update to an existing policyAdded the GetSecurityGroupsForVpc action, which allows you
to get security groups that are usable in your VPC.February 8, 2024[AWS managed policy: AmazonVPCReadOnlyAccess](#security-iam-awsmanpol-AmazonVPCReadOnlyAccess) –
Update to an existing policyAdded the GetSecurityGroupsForVpc action, which allows you
to get security groups that are usable in your VPC.February 8, 2024[AWS managed policy: AmazonVPCCrossAccountNetworkInterfaceOperations](#security-iam-awsmanpol-AmazonVPCCrossAccountNetworkInterfaceOperations) –
Update to an existing policyAdded the AssignIpv6Addresses and UnassignIpv6Addresses
actions, which allow you to manage the IPv6 addresses associated with network interfaces.September 25, 2023[AWS managed policy: AmazonVPCReadOnlyAccess](#security-iam-awsmanpol-AmazonVPCReadOnlyAccess) –
Update to an existing policyAdded the DescribeSecurityGroupRules action, which allows you
to view [security group rules](security-group-rules.md).August 2, 2021[AWS managed policy: AmazonVPCFullAccess](#security-iam-awsmanpol-AmazonVPCFullAccess) –
Update to an existing policyAdded the DescribeSecurityGroupRules and ModifySecurityGroupRules
actions, which allow you to view and modify [security group rules](security-group-rules.md).August 2, 2021[AWS managed policy: AmazonVPCFullAccess](#security-iam-awsmanpol-AmazonVPCFullAccess) –
Update to an existing policyAdded actions for carrier gateways, IPv6 pools, local gateways, and local
gateway route tables.June 23, 2021[AWS managed policy: AmazonVPCReadOnlyAccess](#security-iam-awsmanpol-AmazonVPCReadOnlyAccess) –
Update to an existing policyAdded actions for carrier gateways, IPv6 pools, local gateways, and local
gateway route tables.June 23, 2021

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Troubleshoot

Using service-linked roles

All content copied from https://docs.aws.amazon.com/.
