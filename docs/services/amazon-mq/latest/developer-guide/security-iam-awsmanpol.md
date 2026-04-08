# AWS managed policies for Amazon MQ

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

Amazon MQ supports the following AWS managed policies:

- [AmazonMQApiFullAccess](../../../aws-managed-policy/latest/reference/amazonmqapifullaccess.md)

- [AmazonMQApiReadOnlyAccess](../../../aws-managed-policy/latest/reference/amazonmqapireadonlyaccess.md)

- [AmazonMQFullAccess](../../../aws-managed-policy/latest/reference/amazonmqfullaccess.md)

- [AmazonMQReadOnlyAccess](../../../aws-managed-policy/latest/reference/amazonmqreadonlyaccess.md)

- [AmazonMQServiceRolePolicy](../../../aws-managed-policy/latest/reference/amazonmqservicerolepolicy.md)

## AWS managed policy: AmazonMQServiceRolePolicy

You can't attach `AmazonMQServiceRolePolicy` to your IAM entities. This
policy is attached to a service-linked role that allows Amazon MQ to perform actions on your
behalf. For more information about this permission policy and the actions it allows Amazon MQ
to perform, see [Service-linked role permissions for Amazon MQ](using-service-linked-roles.md#slr-permissions).

## Amazon MQ updates to AWS managed policies

View details about updates to AWS managed policies for Amazon MQ since this service began
tracking these changes. For automatic alerts about changes to this page, subscribe to the
RSS feed on the Amazon MQ [Document history](amazon-mq-release-notes.md)
page.

ChangeDescriptionDate

Amazon MQ started tracking changes

Amazon MQ started tracking changes for its AWS managed policies.

May 5, 2021

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Broker authentication and authorization

Using service-linked roles

All content copied from https://docs.aws.amazon.com/.
