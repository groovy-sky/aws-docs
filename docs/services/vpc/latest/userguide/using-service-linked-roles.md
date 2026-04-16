---
title: "Using service-linked roles for VPC"
---

# Using service-linked roles for VPC

Amazon VPC uses AWS Identity and Access Management (IAM) [service-linked roles](../../../iam/latest/userguide/id-roles-terms-and-concepts.md#iam-term-service-linked-role). A service-linked role is a unique type of IAM role that is
linked directly to VPC. Service-linked roles are predefined by VPC and
include all the permissions that the service requires to call other AWS services on your
behalf.

A service-linked role makes setting up VPC easier because you don't have to
manually add the necessary permissions. VPC defines the permissions of its
service-linked roles, and unless defined otherwise, only VPC can assume its roles. The
defined permissions include the trust policy and the permissions policy, and that permissions
policy cannot be attached to any other IAM entity.

You can delete a service-linked role only after first deleting their related resources. This
protects your VPC resources because you can't inadvertently remove permission to
access the resources.

For information about other services that support service-linked roles, see [AWS services that work with IAM](../../../iam/latest/userguide/reference-aws-services-that-work-with-iam.md) and look for the services that have **Yes** in the **Service-linked roles** column.
Choose a **Yes** with a link to view the service-linked role
documentation for that service.

## Service-linked role permissions for VPC

VPC uses the service-linked role named **AWSServiceRoleForNATGateway** –
This service-linked role enables Amazon VPC to allocate Elastic IP addresses on your behalf to automatically scale regional NAT Gateways, to associate and disassociate your existing Elastic IPs to regional NAT Gateways at your request, and to describe Network Interfaces to identify your existing infrastructure to automatically expand into new Availability Zones.

The AWSServiceRoleForNATGateway service-linked role trusts the following services to assume the
role:

- `ec2-nat-gateway.amazonaws.com`

The role permissions policy named AWSNATGatewayServiceRolePolicy allows VPC to complete the
following actions on the specified resources:

- Action: `AllocateAddress` on Service Managed EIPs to allocate EIPs on your behalf. Service Managed EIPs handles subsequent tagging with service managed tags and ReleaseAddress automatically.

- Action: `AssociateAddress` on your pre-existing Elastic IP addresses to manually associate them to your regional NAT Gateway at your request.

- Action: `DisassociateAddress` on your pre-existing Elastic IP addresses to remove them from the regional NAT Gateway at your request.

- Action: `DescribeAddresses` to obtain Public IP Address information from customer-provided EIPs on associate.

- Action: `DescribeNetworkInterface` on your existing Network Interfaces to automatically identify the Availability Zones your infrastructure resides in to automatically scale out to new zones.

You must configure permissions to allow your users, groups, or roles to create, edit, or
delete a service-linked role. For more information, see [Service-linked role permissions](../../../iam/latest/userguide/using-service-linked-roles.md#service-linked-role-permissions) in the
_IAM User Guide_.

## Creating a service-linked role for VPC

You don't need to manually create a service-linked role. When you
create a NAT Gateway with a "regional" Availability Mode in the AWS Management Console, the AWS CLI, or the AWS API, VPC creates
the service-linked role for you.

###### Important

This service-linked role can appear in your account if you completed an action in
another service that uses the features supported by this role. Also, if you were using the
VPC service before January 1, 2017, when it began supporting service-linked roles,
then VPC created the AWSServiceRoleForNATGateway role in your account. To learn more, see [A\
new role appeared in my AWS account](../../../iam/latest/userguide/troubleshoot-roles.md#troubleshoot_roles_new-role-appeared).

If you delete this service-linked role, and then need to create it again, you can use the
same process to recreate the role in your account. When you create a NAT Gateway with a "regional" Availability Mode,
VPC creates the service-linked role for you again.

You can also use the IAM console to create a service-linked role with the
**AWSServiceRoleForNATGateway** use case. In the AWS CLI or the AWS API, create a
service-linked role with the `ec2-nat-gateway.amazonaws.com` service name. For more
information, see [Creating a service-linked role](../../../iam/latest/userguide/using-service-linked-roles.md#create-service-linked-role) in the _IAM User Guide_. If you
delete this service-linked role, you can use this same process to create the role
again.

## Editing a service-linked role for VPC

VPC does not allow you to edit the AWSServiceRoleForNATGateway service-linked role. After you
create a service-linked role, you cannot change the name of the role because various entities
might reference the role. However, you can edit the description of the role using IAM. For
more information, see [Editing\
a service-linked role](../../../iam/latest/userguide/using-service-linked-roles.md#edit-service-linked-role) in the _IAM User Guide_.

## Deleting a service-linked role for VPC

If you no longer need to use a feature or service that requires a service-linked role, we
recommend that you delete that role. That way you don't have an unused entity that is not
actively monitored or maintained. However, you must clean up the resources for your
service-linked role before you can manually delete it.

###### Note

If the VPC service is using the role when you try to delete the resources,
then the deletion might fail. If that happens, wait for a few minutes and try the operation
again.

###### To delete VPC resources used by the AWSServiceRoleForNATGateway

- Delete all regional NAT Gateways across all Regions in which they have been deployed.

**To manually delete the service-linked role using**
**IAM**

Use the IAM console, the AWS CLI, or the AWS API to delete the AWSServiceRoleForNATGateway service-linked
role. For more information, see [Deleting a service-linked role](../../../iam/latest/userguide/using-service-linked-roles.md#delete-service-linked-role) in the _IAM User Guide_.

## Supported Regions for VPC service-linked roles

VPC supports using service-linked roles in all of the Regions where the service
is available. For more information, see [AWS Regions and endpoints](../../../../general/latest/gr/rande.md).

VPC does not support using service-linked roles in every Region where the
service is available. You can use the AWSServiceRoleForNATGateway role in the following Regions.

Region nameRegion identitySupport in VPCUS East (N. Virginia)us-east-1YesUS East (Ohio)us-east-2YesUS West (N. California)us-west-1YesUS West (Oregon)us-west-2YesAfrica (Cape Town)af-south-1YesAsia Pacific (Hong Kong)ap-east-1YesAsia Pacific (Taipei)ap-east-2YesAsia Pacific (Jakarta)ap-southeast-3YesAsia Pacific (Mumbai)ap-south-1YesAsia Pacific (Hyderabad)ap-south-2YesAsia Pacific (Osaka)ap-northeast-3YesAsia Pacific (Seoul)ap-northeast-2YesAsia Pacific (Singapore)ap-southeast-1YesAsia Pacific (Sydney)ap-southeast-2YesAsia Pacific (Tokyo)ap-northeast-1YesAsia Pacific (Melbourne)ap-southeast-4YesAsia Pacific (Malaysia)ap-southeast-5YesAsia Pacific (New Zealand)ap-southeast-6YesAsia Pacific (Thailand)ap-southeast-7YesCanada (Central)ca-central-1YesCanada West (Calgary)ca-west-1YesEurope (Frankfurt)eu-central-1YesEurope (Zurich)eu-central-2YesEurope (Ireland)eu-west-1YesEurope (London)eu-west-2YesEurope (Milan)eu-south-1YesEurope (Spain)eu-south-2YesEurope (Paris)eu-west-3YesEurope (Stockholm)eu-north-1YesIsrael (Tel Aviv)il-central-1YesMiddle East (Bahrain)me-south-1YesMiddle East (UAE)me-central-1YesMiddle East (Saudi Arabia)me-west-1YesMexico (Central)mx-central-1YesSouth America (São Paulo)sa-east-1YesAWS GovCloud (US-East)us-gov-east-1NoAWS GovCloud (US-West)us-gov-west-1No

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS managed policies

Infrastructure security

All content copied from https://docs.aws.amazon.com/.
