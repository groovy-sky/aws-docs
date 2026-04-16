---
title: "AWS managed policies for Reachability Analyzer"
---

# AWS managed policies for Reachability Analyzer

To add permissions to users, groups, and roles, it is easier to use AWS managed policies
than to write policies yourself. It takes time and expertise to [create IAM customer\
managed policies](../../../iam/latest/userguide/access-policies-create-console.md) that provide your team with only the permissions they need. To get
started quickly, you can use our AWS managed policies. These policies cover common use cases
and are available in your AWS account. For more information about AWS managed policies,
see [AWS managed policies](../../../iam/latest/userguide/access-policies-managed-vs-inline.md#aws-managed-policies) in the _IAM User Guide_.

AWS services maintain and update AWS managed policies. You can't change the
permissions in AWS managed policies. Services occasionally add additional permissions to an
AWS managed policy to support new features. This type of update affects all identities
(users, groups, and roles) where the policy is attached. Services are most likely to update an
AWS managed policy when a new feature is launched or when new operations become available.
Services do not remove permissions from an AWS managed policy, so policy updates won't
break your existing permissions.

Additionally, AWS supports managed policies for job functions that span multiple
services. For example, the **ReadOnlyAccess** AWS managed
policy provides read-only access to all AWS services and resources. When a service launches
a new feature, AWS adds read-only permissions for new operations and resources. For a list
and descriptions of job function policies, see [AWS managed policies for\
job functions](../../../iam/latest/userguide/access-policies-job-functions.md) in the _IAM User Guide_.

## AWS managed policy: AmazonVPCReachabilityAnalyzerFullAccessPolicy

Provides permissions to create, analyze, and delete paths, and to describe path resources, such as EC2 instances,
firewalls, internet gateways, load balancers, NAT gateways, network interfaces, transit gateways, VPC endpoint services,
VPC endpoints, VPC peering connections, and virtual private gateways.

To view the permissions for this policy, see [AmazonVPCReachabilityAnalyzerFullAccessPolicy](../../../aws-managed-policy/latest/reference/amazonvpcreachabilityanalyzerfullaccesspolicy.md) in the _AWS Managed Policy Reference_.

Reachability Analyzer does not support resources from Direct Connect (service prefix: `directconnect`)
or AWS Global Accelerator (service prefix: `globalaccelerator`). If you use this policy as a model
for your own policies, you can omit these actions.

## AWS managed policy: AmazonVPCReachabilityAnalyzerPathComponentReadPolicy

This policy is attached to the role [IAMRoleForReachabilityAnalyzerCrossAccountResourceAccess](cross-account-access-roles.md#IAMRoleForReachabilityAnalyzerCrossAccountResourceAccess). This role is
deployed to the member accounts in an organization when the management account enables
trusted access for Reachability Analyzer using the console. It provides permissions to view resources from
across your organization using the Reachability Analyzer console. For more information, see [Cross-account access\
roles](cross-account-access-roles.md).

To view the permissions for this policy, see [AmazonVPCReachabilityAnalyzerPathComponentReadPolicy](../../../aws-managed-policy/latest/reference/amazonvpcreachabilityanalyzerpathcomponentreadpolicy.md) in the _AWS Managed Policy Reference_.

## AWS managed policy: AWSReachabilityAnalyzerServiceRolePolicy

This policy is attached to a service-linked role that allows Reachability Analyzer to perform actions on
your behalf. For more information, see [Use service-linked roles](using-service-linked-roles.md).

To view the permissions for this policy, see [AWSReachabilityAnalyzerServiceRolePolicy](../../../aws-managed-policy/latest/reference/awsreachabilityanalyzerservicerolepolicy.md) in the _AWS Managed Policy Reference_.

## Reachability Analyzer updates to AWS managed policies

View details about updates to AWS managed policies for Reachability Analyzer since this service began
tracking these changes.

ChangeDescriptionDate[AWSReachabilityAnalyzerServiceRolePolicy](#AWSReachabilityAnalyzerServiceRolePolicy) – Update to an existing policy

Removed actions related to AWS Global Accelerator (service prefix: `globalaccelerator`).

September 10, 2024[AmazonVPCReachabilityAnalyzerFullAccessPolicy](#AmazonVPCReachabilityAnalyzerFullAccessPolicy) – Update to an existing policyAdded the action `elasticloadbalancing:DescribeTargetGroupAttributes`,
which grants permission to describe the attributes of a target group.May 15, 2024[AWSReachabilityAnalyzerServiceRolePolicy](#AWSReachabilityAnalyzerServiceRolePolicy) – Update to an existing policyAdded the action `elasticloadbalancing:DescribeTargetGroupAttributes`,
which grants permission to describe the attributes of a target group.May 15, 2024

[AmazonVPCReachabilityAnalyzerFullAccessPolicy](#AmazonVPCReachabilityAnalyzerFullAccessPolicy) – Update to an existing policy

Removed resource ID prefixes from the resource ARNs used to allow tagging
Reachability Analyzer resources on create.November 3, 2023

[AmazonVPCReachabilityAnalyzerFullAccessPolicy](#AmazonVPCReachabilityAnalyzerFullAccessPolicy) – New policy

Added a policy that provides full access to Reachability Analyzer for single account
use.June 14, 2023

[AmazonVPCReachabilityAnalyzerPathComponentReadPolicy](#AmazonVPCReachabilityAnalyzerPathComponentReadPolicy) – New policy

Added a policy that grants member accounts permission to view resources
from across your organization. The policy is attached to a role that is
deployed to member accounts when the management account enables trusted access
for Reachability Analyzer using the console.May 1, 2023[AWSReachabilityAnalyzerServiceRolePolicy](#AWSReachabilityAnalyzerServiceRolePolicy) – New policyAdded a policy that is attached to a service-linked role that
allows it to access AWS resources and integrate with AWS Organizations on your
behalf.November, 23, 2022Reachability Analyzer started tracking changesReachability Analyzer started tracking changes for its AWS managed policies.March 1, 2021

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Use service-linked roles

Cross-account access
roles

All content copied from https://docs.aws.amazon.com/.
