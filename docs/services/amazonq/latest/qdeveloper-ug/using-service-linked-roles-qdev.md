# Using service-linked roles for Amazon Q Developer

Amazon Q Developer uses AWS Identity and Access Management (IAM) [service-linked roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_terms-and-concepts.html#iam-term-service-linked-role). A service-linked role is a unique type of IAM role that is
linked directly to Amazon Q Developer. Service-linked roles are predefined by Amazon Q Developer and
include all the permissions that the service requires to call other AWS services on your
behalf.

A service-linked role makes setting up Amazon Q Developer easier because you don’t have to
manually add the necessary permissions. Amazon Q Developer defines the permissions of its
service-linked roles, and unless defined otherwise, only Amazon Q Developer can assume its roles.
The defined permissions include the trust policy and the permissions policy, and that
permissions policy cannot be attached to any other IAM entity.

You can delete a service-linked role only after first deleting their related resources.
This protects your Amazon Q Developer resources because you can't inadvertently remove permission
to access the resources.

For information about other services that support service-linked roles, see [AWS services\
that work with IAM](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-services-that-work-with-iam.html) and look for the services that have **Yes** in the **Service-linked roles** column. Choose
a **Yes** with a link to view the service-linked role
documentation for that service.

Learn about [AWS managed policies for Amazon Q Developer](managed-policy.md).

## Service-linked role permissions for Amazon Q Developer

Amazon Q Developer uses the service-linked role named **AWSServiceRoleForAmazonQDeveloper**
– This role grants permissions to Amazon Q to access data in your account to calculate billing, provides access to create and access security reports in Amazon CodeGuru, and emit data to CloudWatch.

The AWSServiceRoleForAmazonQDeveloper service-linked role trusts the following services to assume the
role:

- `q.amazonaws.com`

The role permissions policy named AWSServiceRoleForAmazonQDeveloper allows Amazon Q Developer to complete the
following actions on the specified resources:

- Action: `cloudwatch:PutMetricData` on `AWS/Q CloudWatch
                namespace`

You must configure permissions to allow your users, groups, or roles to create, edit, or
delete a service-linked role. For more information, see [Service-linked role permissions](https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html#service-linked-role-permissions) in the
_IAM User Guide_.

## Creating a service-linked role for Amazon Q Developer

You don't need to manually create a service-linked role. When you
create a profile for Amazon Q in the AWS Management Console, Amazon Q Developer creates the service-linked role
for you.

If you delete this service-linked role, and then need to create it again, you can use
the same process to recreate the role in your account. When you update the settings, Amazon Q
creates the service-linked role for you again.

You can also use the IAM console or AWS CLI to create a service-linked role with the
`q.amazonaws.com` service name. For more information, see [Creating a service-linked role](https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html#create-service-linked-role) in the _IAM User Guide_. If
you delete this service-linked role, you can use this same process to create the role
again.

## Editing a service-linked role for Amazon Q Developer

Amazon Q Developer does not allow you to edit the AWSServiceRoleForAmazonQDeveloper service-linked role. After you
create a service-linked role, you cannot change the name of the role because various
entities might reference the role. However, you can edit the description of the role using
IAM. For more information, see [Editing a service-linked role](https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html#edit-service-linked-role) in the
_IAM User Guide_.

## Deleting a service-linked role for Amazon Q Developer

If you no longer need to use a feature or service that requires a service-linked role,
we recommend that you delete that role. That way you don’t have an unused entity that is not
actively monitored or maintained. However, you must clean up the resources for your
service-linked role before you can manually delete it.

###### Note

If the Amazon Q Developer service is using the role when you try to delete the resources,
then the deletion might fail. If that happens, wait for a few minutes and try the
operation again.

**To manually delete the service-linked role using**
**IAM**

Use the IAM console, the AWS CLI, or the AWS API to delete the AWSServiceRoleForAmazonQDeveloper
service-linked role. For more information, see [Deleting a service-linked role](https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html#delete-service-linked-role) in the
_IAM User Guide_.

## Supported Regions for Amazon Q Developer service-linked roles

Amazon Q Developer does not support using service-linked roles in every Region where the
service is available. You can use the AWSServiceRoleForAmazonQDeveloper role in the following Regions. For more
information, see [AWS Regions and endpoints](https://docs.aws.amazon.com/general/latest/gr/rande.html).

Region nameRegion identitySupport in Amazon Q DeveloperUS East (N. Virginia)us-east-1YesUS East (Ohio)us-east-2NoUS West (N. California)us-west-1NoUS West (Oregon)us-west-2NoAfrica (Cape Town)af-south-1NoAsia Pacific (Hong Kong)ap-east-1NoAsia Pacific (Jakarta)ap-southeast-3NoAsia Pacific (Mumbai)ap-south-1NoAsia Pacific (Osaka)ap-northeast-3NoAsia Pacific (Seoul)ap-northeast-2NoAsia Pacific (Singapore)ap-southeast-1NoAsia Pacific (Sydney)ap-southeast-2NoAsia Pacific (Tokyo)ap-northeast-1NoCanada (Central)ca-central-1NoEurope (Frankfurt)eu-central-1NoEurope (Ireland)eu-west-1NoEurope (London)eu-west-2NoEurope (Milan)eu-south-1NoEurope (Paris)eu-west-3NoEurope (Stockholm)eu-north-1NoMiddle East (Bahrain)me-south-1NoMiddle East (UAE)me-central-1NoSouth America (São Paulo)sa-east-1NoAWS GovCloud (US-East)us-gov-east-1NoAWS GovCloud (US-West)us-gov-west-1No

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Using service-linked roles

For User Subscriptions
