# Using service-linked roles for EC2 Capacity Manager

EC2 Capacity Manager uses AWS Identity and Access Management (IAM) [service-linked roles](../../../iam/latest/userguide/id-roles-terms-and-concepts.md#iam-term-service-linked-role). A service-linked role is a unique type of IAM role that is
linked directly to Capacity Manager. Service-linked roles are predefined by Capacity Manager and
include all the permissions that the service requires to call other AWS services on your
behalf.

A service-linked role makes setting up Capacity Manager easier because you don't have to
manually add the necessary permissions. Capacity Manager defines the permissions of its
service-linked roles, and unless defined otherwise, only Capacity Manager can assume its roles. The
defined permissions include the trust policy and the permissions policy, and that permissions
policy cannot be attached to any other IAM entity.

You can delete a service-linked role only after first deleting their related resources. This
protects your Capacity Manager resources because you can't inadvertently remove permission to
access the resources.

For information about other services that support service-linked roles, see [AWS services that work with IAM](../../../iam/latest/userguide/reference-aws-services-that-work-with-iam.md) and look for the services that have **Yes** in the **Service-linked roles** column.
Choose a **Yes** with a link to view the service-linked role
documentation for that service.

## Service-linked role permissions for Capacity Manager

Capacity Manager uses the service-linked role named **AWSServiceRoleForEC2CapacityManager** to allow
you to manage capacity resources and integrate with AWS Organizations
on your behalf.

The AWSServiceRoleForEC2CapacityManager service-linked role trusts the following services to assume the
role:

- `ec2.capacitymanager.amazonaws.com`

The role permissions policy named AWSEC2CapacityManagerServiceRolePolicy allows Capacity Manager to complete the
following actions:

- `organizations:DescribeOrganization`

- `organizations:ListAccounts`

- `organizations:ListChildren`

- `organizations:ListAWSServiceAccessForOrganization`

- `organizations:ListDelegatedAdministrators`

You must configure permissions to allow your users, groups, or roles to create, edit, or
delete a service-linked role. For more information, see [Service-linked role permissions](../../../iam/latest/userguide/using-service-linked-roles.md#service-linked-role-permissions) in the
_IAM User Guide_.

## Creating a service-linked role for Capacity Manager

You can use the IAM console to create a service-linked role with the
**AWSEC2CapacityManagerServiceRolePolicy** use case. In the AWS CLI or the AWS API, create a
service-linked role with the `ec2.capacitymanager.amazonaws.com` service name. For more
information, see [Creating a service-linked role](../../../iam/latest/userguide/using-service-linked-roles.md#create-service-linked-role) in the _IAM User Guide_. If you
delete this service-linked role, you can use this same process to create the role
again.

## Editing a service-linked role for Capacity Manager

Capacity Manager does not allow you to edit the AWSServiceRoleForEC2CapacityManager service-linked role. After you
create a service-linked role, you cannot change the name of the role because various entities
might reference the role. However, you can edit the description of the role using IAM. For
more information, see [Editing\
a service-linked role](../../../iam/latest/userguide/using-service-linked-roles.md#edit-service-linked-role) in the _IAM User Guide_.

## Deleting a service-linked role for Capacity Manager

If you no longer need to use a feature or service that requires a service-linked role, we
recommend that you delete that role. That way you don't have an unused entity that is not
actively monitored or maintained. However, you must clean up the resources for your
service-linked role before you can manually delete it.

###### Note

If the Capacity Manager service is using the role when you try to delete the resources,
then the deletion might fail. If that happens, wait for a few minutes and try the operation
again.

###### To remove Capacity Manager resources used by the AWSServiceRoleForEC2CapacityManager

1. All delegated administrators must have disabled their Capacity Manager before removing organizations access.

2. You must delete any active data exports before disabling a capacity manager.

**To manually delete the service-linked role using**
**IAM**

Use the IAM console, the AWS CLI, or the AWS API to delete the AWSServiceRoleForEC2CapacityManager service-linked
role. For more information, see [Deleting a service-linked role](../../../iam/latest/userguide/using-service-linked-roles.md#delete-service-linked-role) in the _IAM User Guide_.

## Supported Regions for Capacity Manager service-linked roles

Capacity Manager supports using service-linked roles in all of the Regions where the service
is available. For more information, see [AWS Regions and endpoints](../../../../general/general/latest/gr/rande.md).

Capacity Manager does not support using service-linked roles in every Region where the
service is available. You can use the AWSServiceRoleForEC2CapacityManager role in the following Regions.

Region nameRegion identitySupport in Capacity ManagerUS East (N. Virginia)us-east-1YesUS East (Ohio)us-east-2YesUS West (N. California)us-west-1YesUS West (Oregon)us-west-2YesAfrica (Cape Town)af-south-1NoAsia Pacific (Hong Kong)ap-east-1NoAsia Pacific (Jakarta)ap-southeast-3NoAsia Pacific (Mumbai)ap-south-1YesAsia Pacific (Osaka)ap-northeast-3YesAsia Pacific (Seoul)ap-northeast-2YesAsia Pacific (Singapore)ap-southeast-1YesAsia Pacific (Sydney)ap-southeast-2YesAsia Pacific (Tokyo)ap-northeast-1YesCanada (Central)ca-central-1YesEurope (Frankfurt)eu-central-1YesEurope (Ireland)eu-west-1YesEurope (London)eu-west-2YesEurope (Milan)eu-south-1NoEurope (Paris)eu-west-3YesEurope (Stockholm)eu-north-1YesMiddle East (Bahrain)me-south-1NoMiddle East (UAE)me-central-1NoSouth America (São Paulo)sa-east-1YesAWS GovCloud (US-East)us-gov-east-1NoAWS GovCloud (US-West)us-gov-west-1No

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Registering a delegated administrator

Organizing your data
