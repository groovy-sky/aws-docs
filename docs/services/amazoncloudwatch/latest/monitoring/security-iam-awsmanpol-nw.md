# AWS managed policies for Network Synthetic Monitor

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
Services do not remove permissions from an AWS managed policy, so policy updates won't break
your existing permissions.

Additionally, AWS supports managed policies for job functions that span multiple
services. For example, the `ReadOnlyAccess` AWS managed policy provides read-only
access to all AWS services and resources. When a service launches a new feature, AWS adds
read-only permissions for new operations and resources. For a list and descriptions of job
function policies, see [AWS managed policies for\
job functions](../../../iam/latest/userguide/access-policies-job-functions.md) in the _IAM User Guide_.

## AWS managed policy: CloudWatchNetworkMonitorServiceRolePolicy

The `CloudWatchNetworkMonitorServiceRolePolicy` is attached to a service-linked
role that allows the service to perform actions on your behalf and access resources associated
with CloudWatch Network Synthetic Monitor. You cannot attach this policy to your IAM identities. For more
information, see [Using a service-linked role for Network Synthetic Monitor](monitoring-using-service-linked-roles-nw.md).

## Network Synthetic Monitor updates to AWS managed policies

To view details about updates to AWS managed policies for Network Synthetic Monitor since this service
began tracking these changes,
see [CloudWatch updates to AWS managed policies](managed-policies-cloudwatch.md#security-iam-awsmanpol-updates).
For automatic alerts about managed policy changes in CloudWatch, subscribe to the RSS feed on the CloudWatch
[Document history](documenthistory.md) page.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Troubleshooting

IAM permissions

All content copied from https://docs.aws.amazon.com/.
