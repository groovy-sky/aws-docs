# Using service-linked roles for AWS Artifact

AWS Artifact uses AWS Identity and Access Management (IAM)
[service-linked roles](../../../iam/latest/userguide/id-roles-terms-and-concepts.md#iam-term-service-linked-role). A service-linked role is a unique type of IAM role that is
linked directly to AWS Artifact. Service-linked roles are predefined by AWS Artifact and
include all the permissions that the service requires to call other AWS services on your
behalf.

A service-linked role makes setting up AWS Artifact easier because you don't have to
manually add the necessary permissions. AWS Artifact defines the permissions of its
service-linked roles, and unless defined otherwise, only AWS Artifact can assume its roles. The
defined permissions include the trust policy and the permissions policy, and that permissions
policy cannot be attached to any other IAM entity.

You can delete a service-linked role only after first deleting its related resources. This
protects your AWS Artifact resources because you can't inadvertently remove permission to
access the resources.

For information about other services that support service-linked roles, see
[AWS services that work with IAM](../../../iam/latest/userguide/reference-aws-services-that-work-with-iam.md) and look for the services that have
**Yes** in the
**Service-linked roles** column.
Choose a
**Yes** with a link to view the service-linked role
documentation for that service.

## Service-linked role permissions for AWS Artifact

AWS Artifact uses the service-linked role named
**AWSServiceRoleForArtifact** –
Allows AWS Artifact to gather information about an organization via AWS Organizations.

The AWSServiceRoleForArtifact service-linked role trusts the following services to assume the
role:

- `artifact.amazonaws.com`

The role permissions policy named AWSArtifactServiceRolePolicy allows AWS Artifact to complete the
following actions on the
`organizations` resource.

- `DescribeOrganization`

- `DescribeAccount`

- `ListAccounts`

- `ListAWSServiceAccessForOrganization`

## Creating a service-linked role for AWS Artifact

You don't need to manually create a service-linked role. When you
go to the **Organization agreements** tab in an organization management account and choose the **Get started** link in the AWS Management Console, AWS Artifact creates
the service-linked role for you.

If you delete this service-linked role, and then need to create it again, you can use the
same process to recreate the role in your account. When you go to the **Organization agreements** tab in an organization management account and choose the **Get started** link,
AWS Artifact creates the service-linked role for you again.

## Editing a service-linked role for AWS Artifact

AWS Artifact does not allow you to edit the AWSServiceRoleForArtifact service-linked role. After you
create a service-linked role, you cannot change the name of the role because various entities
might reference the role. However, you can edit the description of the role using IAM. For
more information, see
[Editing\
a service-linked role](../../../iam/latest/userguide/using-service-linked-roles.md#edit-service-linked-role) in the
_IAM User Guide_.

## Deleting a service-linked role for AWS Artifact

If you no longer need to use a feature or service that requires a service-linked role, we
recommend that you delete that role. That way you don't have an unused entity that is not
actively monitored or maintained. However, you must clean up the resources for your
service-linked role before you can manually delete it.

###### Note

If the AWS Artifact service is using the role when you try to delete the resources,
then the deletion might fail. If that happens, wait for a few minutes and try the operation
again.

###### To delete AWS Artifact resources used by the AWSServiceRoleForArtifact

1. Visit the 'Organization Agreements' table in the AWS Artifact console

2. Terminate any active Organization agreements

**To manually delete the service-linked role using**
**IAM**

Use the IAM console, the AWS CLI, or the AWS API to delete the AWSServiceRoleForArtifact service-linked
role. For more information, see
[Deleting a service-linked role](../../../iam/latest/userguide/using-service-linked-roles.md#delete-service-linked-role) in the
_IAM User Guide_.

## Supported Regions for AWS Artifact service-linked roles

AWS Artifact does not support using service-linked roles in every Region where the
service is available. You can use the AWSServiceRoleForArtifact role in the following Regions.

Region nameRegion identitySupport in AWS ArtifactUS East (N. Virginia)us-east-1YesUS East (Ohio)us-east-2NoUS West (N. California)us-west-1NoUS West (Oregon)us-west-2YesAfrica (Cape Town)af-south-1NoAsia Pacific (Hong Kong)ap-east-1NoAsia Pacific (Jakarta)ap-southeast-3NoAsia Pacific (Mumbai)ap-south-1NoAsia Pacific (Osaka)ap-northeast-3NoAsia Pacific (Seoul)ap-northeast-2NoAsia Pacific (Singapore)ap-southeast-1NoAsia Pacific (Sydney)ap-southeast-2NoAsia Pacific (Tokyo)ap-northeast-1NoCanada (Central)ca-central-1NoEurope (Frankfurt)eu-central-1NoEurope (Ireland)eu-west-1NoEurope (London)eu-west-2NoEurope (Milan)eu-south-1NoEurope (Paris)eu-west-3NoEurope (Stockholm)eu-north-1NoMiddle East (Bahrain)me-south-1NoMiddle East (UAE)me-central-1NoSouth America (São Paulo)sa-east-1NoAWS GovCloud (US-East)

us-gov-east-1NoAWS GovCloud (US-West)

us-gov-west-1Yes

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Using AWS managed policies

Using IAM condition keys

All content copied from https://docs.aws.amazon.com/.
