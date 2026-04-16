---
title: "AWS managed policies for transit gateways in AWS Transit Gateway"
---

# AWS managed policies for transit gateways in AWS Transit Gateway

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

To work with a transit gateway, one of the following AWS managed policies might meet your needs:

- [AmazonEC2FullAccess](../../../aws-managed-policy/latest/reference/amazonec2fullaccess.md)

- [AmazonEC2ReadOnlyAccess](../../../aws-managed-policy/latest/reference/amazonec2readonlyaccess.md)

- [PowerUserAccess](../../../aws-managed-policy/latest/reference/poweruseraccess.md)

- [ReadOnlyAccess](../../../aws-managed-policy/latest/reference/readonlyaccess.md)

## AWS managed policy: AWSVPCTransitGatewayServiceRolePolicy

This policy is attached to the role [AWSServiceRoleForVPCTransitGateway](service-linked-roles.md).
This allows Amazon VPC to create and manage resources for your transit gateway attachments.

To view the permissions for this policy, see [AWSVPCTransitGatewayServiceRolePolicy](../../../aws-managed-policy/latest/reference/awsvpctransitgatewayservicerolepolicy.md)
in the _AWS Managed Policy Reference_.

## Transit gateway updates to AWS managed policies

View details about updates to AWS managed policies for transit gateways since Amazon VPC
began tracking these changes in March 2021.

ChangeDescriptionDateAmazon VPC started tracking changesAmazon VPC started tracking changes to its AWS managed policies.March 1, 2021

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Service-linked roles

Network ACLs

All content copied from https://docs.aws.amazon.com/.
