AWS App Runner will no longer be open to new customers starting April 30, 2026. If you would like to use
App Runner, sign up prior to that date. Existing customers can continue to use the service as normal. For more information, see
[AWS App Runner availability\
change](apprunner-availability-change.md).

# AWS managed policies for AWS App Runner

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

## App Runner updates to AWS managed policies

View details about updates to AWS managed policies for App Runner since this service began tracking these changes. For automatic alerts about
changes to this page, subscribe to the RSS feed on the App Runner Document history page.

ChangeDescriptionDate

[AWSAppRunnerReadOnlyAccess](security-iam-service-with-iam.md#security_iam_service-with-iam-users) – New policy

App Runner added a new policy to allow users to list and view details about App Runner resources.

Feb 24, 2022

[AWSAppRunnerFullAccess](security-iam-service-with-iam.md#security_iam_service-with-iam-users) – Update to an existing
policy

App Runner updated the resource list for the `iam:CreateServiceLinkedRole` action to allow creation of
`AWSServiceRoleForAppRunnerNetworking` service-linked role.

Feb 8, 2022

[AppRunnerNetworkingServiceRolePolicy](using-service-linked-roles-networking.md) – New policy

App Runner added a new policy to allow App Runner to make calls to Amazon Virtual Private Cloud to attach a VPC to your App Runner service and manage network
interfaces on behalf of App Runner services. The policy is used in the `AWSServiceRoleForAppRunnerNetworking` service-linked role.

Feb 8, 2022

[AWSAppRunnerFullAccess](security-iam-service-with-iam.md#security_iam_service-with-iam-users) – New policy

App Runner added a new policy to allow users to perform all App Runner actions.

Jan 10, 2022

[AppRunnerServiceRolePolicy](using-service-linked-roles-management.md) – New policy

App Runner added a new policy to allow App Runner to make calls to Amazon CloudWatch Logs and Amazon CloudWatch Events on behalf of App Runner services. The policy is used in
the `AWSServiceRoleForAppRunner` service-linked role.

Mar 1, 2021

[AWSAppRunnerServicePolicyForECRAccess](security-iam-service-with-iam.md#security_iam_service-with-iam-roles-service.access) – New
policy

App Runner added a new policy to allow App Runner to access Amazon Elastic Container Registry (Amazon ECR) images in your account.

Mar 1, 2021

App Runner started tracking changes

App Runner started tracking changes for its AWS managed policies.

Mar 1, 2021

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Networking role

Troubleshooting

All content copied from https://docs.aws.amazon.com/.
