# Identity-based policy examples for AWS AppFabric

By default, users and roles don't have permission to create or modify AppFabric
resources. To grant users permission to perform actions on the
resources that they need, an IAM administrator can create IAM policies.

To learn how to create an IAM identity-based policy by using these example JSON policy
documents, see [Create IAM policies (console)](../../../iam/latest/userguide/access-policies-create-console.md) in the
_IAM User Guide_.

For details about actions and resource types defined by AppFabric, including the format of the ARNs for each of the resource types, see [Actions, resources, and condition keys for AWS AppFabric](../../../iam/latest/userguide/list-appfabric.md) in the _Service Authorization Reference_.

###### Contents

- [Policy best practices](security-iam-id-based-policy-examples.md#security_iam_service-with-iam-policy-best-practices)

- [Using the AppFabric console](security-iam-id-based-policy-examples.md#security_iam_id-based-policy-examples-console)

- [AppFabric for security IAM policy examples](security-iam-id-based-policy-examples.md#appfabric-for-security-policy-examples)

  - [Allow access to app bundles](security-iam-id-based-policy-examples.md#security_iam_id-based-policy-examples-allow-app-bundle-access)

  - [Restrict access to app bundles](security-iam-id-based-policy-examples.md#security_iam_id-based-policy-examples-restrict-app-bundle-access)

  - [Restrict deleting or stopping ingestions](security-iam-id-based-policy-examples.md#security_iam_id-based-policy-examples-restrict-delete-stop-ingestion)
- [AppFabric for productivity IAM policy examples](security-iam-id-based-policy-examples.md#appfabric-for-productivity-policy-examples)

  - [Allow access read-only access to productivity features](security-iam-id-based-policy-examples.md#security_iam_id-based-policy-examples-productivity-read-only)

  - [Allow full access to productivity features](security-iam-id-based-policy-examples.md#security_iam_id-based-policy-examples-productivity-full-access)

  - [Allow access to create AppClients](security-iam-id-based-policy-examples.md#security_iam_id-based-policy-examples-productivity-create-appclient)

  - [Allow access to get details of AppClients](security-iam-id-based-policy-examples.md#security_iam_id-based-policy-examples-productivity-get-appclient)

  - [Allow access to list AppClients](security-iam-id-based-policy-examples.md#security_iam_id-based-policy-examples-productivity-list-appclient)

  - [Allow access to update AppClients](security-iam-id-based-policy-examples.md#security_iam_id-based-policy-examples-productivity-update-appclient)

  - [Allow access to delete AppClients](security-iam-id-based-policy-examples.md#security_iam_id-based-policy-examples-productivity-delete-appclient)

  - [Allow access to authorize applications](security-iam-id-based-policy-examples.md#security_iam_id-based-policy-examples-productivity-token)
- [Other IAM policy examples](security-iam-id-based-policy-examples.md#other-iam-policy-examples)

  - [Allow users to view their own permissions](security-iam-id-based-policy-examples.md#security_iam_id-based-policy-examples-view-own-permissions)

## Policy best practices

Identity-based policies determine whether someone can create, access, or delete AppFabric resources in your
account. These actions can incur costs for your AWS account. When you create or edit identity-based policies, follow these guidelines and
recommendations:

- **Get started with AWS managed policies and move toward least-privilege permissions**
– To get started granting permissions to your users and workloads, use the _AWS_
_managed policies_ that grant permissions for many common use cases. They are
available in your AWS account. We recommend that you reduce permissions further by
defining AWS customer managed policies that are specific to your use cases. For more information, see
[AWS managed policies](../../../iam/latest/userguide/access-policies-managed-vs-inline.md#aws-managed-policies) or [AWS managed policies for job functions](../../../iam/latest/userguide/access-policies-job-functions.md) in the _IAM User Guide_.

- **Apply least-privilege permissions** –
When you set permissions with IAM policies, grant only the permissions required to
perform a task. You do this by defining the actions that can be taken on specific resources
under specific conditions, also known as _least-privilege permissions_.
For more information about using IAM to apply permissions, see [Policies and permissions in IAM](../../../iam/latest/userguide/access-policies.md) in the _IAM User Guide_.

- **Use conditions in IAM policies to further restrict access**
– You can add a condition to your policies to limit access to actions and resources. For example, you can write a policy condition to specify that all requests must
be sent using SSL. You can also use conditions to grant access to service actions
if they are used through a specific AWS service, such as CloudFormation. For more information, see
[IAM JSON policy elements: Condition](../../../iam/latest/userguide/reference-policies-elements-condition.md) in the _IAM User Guide_.

- **Use IAM Access Analyzer to validate your IAM policies to ensure secure and functional permissions**
– IAM Access Analyzer validates new and existing policies so that the policies adhere to the IAM policy language (JSON) and IAM best practices.
IAM Access Analyzer provides more than 100 policy checks and actionable recommendations to help
you author secure and functional policies. For more information, see [Validate policies with IAM Access Analyzer](../../../iam/latest/userguide/access-analyzer-policy-validation.md) in the _IAM User Guide_.

- **Require multi-factor authentication (MFA)** –
If you have a scenario that requires IAM users or a root user in your AWS account, turn on MFA for additional security. To require
MFA when API operations are called, add MFA conditions to your policies. For
more information, see [Secure API access with MFA](../../../iam/latest/userguide/id-credentials-mfa-configure-api-require.md) in the _IAM User Guide_.

For more information about best practices in IAM, see [Security best practices in IAM](../../../iam/latest/userguide/best-practices.md) in the _IAM User Guide_.

## Using the AppFabric console

Attach the `AWSAppFabricReadOnlyAccess` AWS managed policy to your IAM
identities to grant them read-only permission to the AppFabric service, including the AppFabric
console in the AWS Management Console. Or, you can attach the `AWSAppFabricFullAccess`
AWS managed policy to your IAM identities to grant them full administrative
permission to the AppFabric service. For more information, see [AWS managed policies for AWS AppFabric](security-iam-awsmanpol.md).

## AppFabric for security IAM policy examples

The following policy examples apply to the AppFabric for security features.

### Allow access to app bundles

The following policy example grants access to app bundles in the AppFabric
service.

JSON

```json

{
    "Version":"2012-10-17",

    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "appfabric:StartUserAccessTasks",
                "appfabric:BatchGetUserAccessTasks"
            ],
            "Resource": ["arn:aws:appfabric:*:*:appbundle/*"]
        }
    ]
}

```

### Restrict access to app bundles

The following policy example restricts access to app bundles in the AppFabric
service.

JSON

```json

{
    "Version":"2012-10-17",

    "Statement": [
        {
            "Action": ["appfabric:*"],
            "Effect": "Allow",
            "Resource": "*"
        },
        {
            "Effect": "Deny",
            "Action": [
                "appfabric:StartUserAccessTasks",
                "appfabric:BatchGetUserAccessTasks"
            ],
            "Resource": ["arn:aws:appfabric:*:*:appbundle/*"]
        }
    ]
}

```

### Restrict deleting or stopping ingestions

The following policy example restricts the deletion or stopping of ingestions in
the AppFabric service.

JSON

```json

{
    "Version":"2012-10-17",

    "Statement": [
        {
            "Action": ["appfabric:*"],
            "Effect": "Allow",
            "Resource": "*"
        },
        {
            "Effect": "Deny",
            "Action": [
                "appfabric:StopIngestion",
                "appfabric:DeleteIngestion",
                "appfabric:DeleteIngestionDestination"
            ],
            "Resource": ["arn:aws:appfabric:*:*:appbundle/*"]
        }
    ]
}

```

## AppFabric for productivity IAM policy examples

The AWS AppFabric for productivity feature is in preview and is subject to change.

The following policy examples apply to the AppFabric for productivity features.

### Allow access read-only access to productivity features

The following policy example grants read-only access to the AppFabric for productivity
features.

###### Important

You might see an invalid action error when adding this policy in the JSON
policy editor of the IAM console. This is because the AppFabric for productivity features are
currently in preview. You should ignore the error and proceed to create the
policy.

### Allow full access to productivity features

The following policy example grants full access to the AppFabric for productivity features.

###### Important

You might see an invalid action error when adding this policy in the JSON
policy editor of the IAM console. This is because the AppFabric for productivity features are
currently in preview. You should ignore the error and proceed to create the
policy.

### Allow access to create AppClients

The following policy example grants access to create AppClients. For more
information, see [Create an AppFabric for productivity\
AppClient](getting-started-appdeveloper-productivity.md#create_appclient).

###### Important

You might see an invalid action error when adding this policy in the JSON
policy editor of the IAM console. This is because the AppFabric for productivity features are
currently in preview. You should ignore the error and proceed to create the
policy.

### Allow access to get details of AppClients

The following policy example grants access to get details of AppClients. For more
information, see [Get details of an\
AppClient](manage-appclients.md#get_appclient_details).

###### Important

You might see an invalid action error when adding this policy in the JSON
policy editor of the IAM console. This is because the AppFabric for productivity features are
currently in preview. You should ignore the error and proceed to create the
policy.

### Allow access to list AppClients

The following policy example grants access to list AppClients. For more
information, see [Get details of an\
AppClient](manage-appclients.md#list_appclients).

###### Important

You might see an invalid action error when adding this policy in the JSON
policy editor of the IAM console. This is because the AppFabric for productivity features are
currently in preview. You should ignore the error and proceed to create the
policy.

### Allow access to update AppClients

The following policy example grants access to update AppClients. For more
information, see [Update an AppClient](manage-appclients.md#update_appclient).

###### Important

You might see an invalid action error when adding this policy in the JSON
policy editor of the IAM console. This is because the AppFabric for productivity features are
currently in preview. You should ignore the error and proceed to create the
policy.

### Allow access to delete AppClients

The following policy example grants access to delete AppClients. For more
information, see [Update an AppClient](manage-appclients.md#delete_appclient).

###### Important

You might see an invalid action error when adding this policy in the JSON
policy editor of the IAM console. This is because the AppFabric for productivity features are
currently in preview. You should ignore the error and proceed to create the
policy.

### Allow access to authorize applications

The following policy example grants access to authorize applications using the
Token API. For more information, see [Authenticate and authorize your application](getting-started-appdeveloper-productivity.md#authorize_data_access).

###### Important

You might see an invalid action error when adding this policy in the JSON
policy editor of the IAM console. This is because the AppFabric for productivity features are
currently in preview. You should ignore the error and proceed to create the
policy.

## Other IAM policy examples

### Allow users to view their own permissions

This example shows how you might create a policy that allows IAM users to view the inline and managed policies that are attached to their user
identity. This policy includes permissions to complete this action on the console or programmatically using the AWS CLI or AWS API.

```json

{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "ViewOwnUserInfo",
            "Effect": "Allow",
            "Action": [
                "iam:GetUserPolicy",
                "iam:ListGroupsForUser",
                "iam:ListAttachedUserPolicies",
                "iam:ListUserPolicies",
                "iam:GetUser"
            ],
            "Resource": ["arn:aws:iam::*:user/${aws:username}"]
        },
        {
            "Sid": "NavigateInConsole",
            "Effect": "Allow",
            "Action": [
                "iam:GetGroupPolicy",
                "iam:GetPolicyVersion",
                "iam:GetPolicy",
                "iam:ListAttachedGroupPolicies",
                "iam:ListGroupPolicies",
                "iam:ListPolicyVersions",
                "iam:ListPolicies",
                "iam:ListUsers"
            ],
            "Resource": "*"
        }
    ]
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

How AWS AppFabric works with IAM

Using service-linked roles

All content copied from https://docs.aws.amazon.com/.
