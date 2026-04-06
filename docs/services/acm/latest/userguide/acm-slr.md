# Use a service-linked role (SLR) with ACM

AWS Certificate Manager uses an AWS Identity and Access Management (IAM) [service-linked role](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_terms-and-concepts.html#iam-term-service-linked-role) to enable enable automatic renewals of private certificates issued from a private CA for another account shared by AWS Resource Access Manager. A service-linked role (SLR) is an IAM role that is
linked directly to the ACM service. SLRs are predefined by ACM
and include all the permissions that the service requires to call other AWS
services on your behalf.

The SLR makes setting up ACM easier because you don’t have to manually
add the necessary permissions for unattended certificate signing. ACM
defines the permissions of its SLR, and unless defined otherwise, only ACM
can assume the role. The defined permissions include the trust policy and the
permissions policy, and that permissions policy cannot be attached to any other
IAM entity.

For information about other services that support SLRs, see [AWS\
Services That Work with IAM](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-services-that-work-with-iam.html) and look for the services that have
**Yes** in the **Service-Linked**
**Role** column. Choose a **Yes** with a
link to view the SLR documentation for that service.

## SLR permissions for ACM

ACM uses an SLR named Amazon Certificate Manager Service Role Policy.

The AWSServiceRoleForCertificateManager SLR trusts the following services to assume the role:

- `acm.amazonaws.com`

The role permissions policy allows ACM to complete the following
actions on the specified resources:

- Actions: `acm-pca:IssueCertificate`,
`acm-pca:GetCertificate` on "\*"

You must configure permissions to allow an IAM entity (such as a user,
group, or role) to create, edit, or delete an SLR. For more information, see
[Service-Linked Role Permissions](https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html#service-linked-role-permissions) in the
_IAM User Guide_.

###### Important

ACM might alert you that it cannot determine whether an SLR exists on your account. If the
required `iam:GetRole` permission has already been granted to the ACM SLR for your account,
then the alert will not recur after the SLR is created. If it does recur,
then you or your account administrator might need to grant the `iam:GetRole`
permission to ACM, or associate your account with the ACM-managed policy
`AWSCertificateManagerFullAccess`.

## Creating the SLR for ACM

You don't need to manually create the SLR that ACM uses. When you issue an
ACM certificate using the AWS Management Console, the AWS CLI, or the AWS API,
ACM creates the SLR for you the first time that you a private
CA for another account shared by AWS RAM to sign your certificate.

If you encounter messages stating that ACM cannot determine whether an SLR
exists on your account, it may mean that your account has not granted a read
permission that AWS Private CA requires. This will not prevent the SLR from being
installed, and you can still issue certificates, but ACM will be unable to
renew the certificates automatically until you resolve the problem. For more
information, see [Problems with the ACM service-linked role (SLR)](https://docs.aws.amazon.com/acm/latest/userguide/slr-problems.html).

###### Important

This SLR can appear in your account if you completed an action in another
service that uses the features supported by this role. Also, if you were
using the ACM service before January 1, 2017, when it began
supporting SLRs, then ACM created the AWSServiceRoleForCertificateManager role in your
account. To learn more, see [A New Role Appeared in My IAM Account](https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_roles.html#troubleshoot_roles_new-role-appeared).

If you delete this SLR, and then need to create it again, you can use either
of these methods:

- In the IAM console, choose **Role**,
**Create role**, **Certificate**
**Manager** to create a new role with the
**CertificateManagerServiceRolePolicy** use case.

- Using the IAM API [CreateServiceLinkedRole](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateServiceLinkedRole.html) or the corresponding AWS CLI command
[create-service-linked-role](https://docs.aws.amazon.com/cli/latest/reference/iam/create-service-linked-role.html), create an SLR with the
`acm.amazonaws.com` service name.

For more information, see [Creating a Service-Linked Role](https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html#create-service-linked-role) in the
_IAM User Guide_.

## Editing the SLR for ACM

ACM does not allow you to edit the AWSServiceRoleForCertificateManager service-linked role.
After you create an SLR, you cannot change the name of the role because various
entities might reference the role. However, you can edit the description of the
role using IAM. For more information, see [Editing a Service-Linked Role](https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html#edit-service-linked-role) in the
_IAM User Guide_.

## Deleting the SLR for ACM

You typically don't need to delete the AWSServiceRoleForCertificateManager SLR. However, you can delete
the role manually using the IAM console, the AWS CLI or the AWS API. For more
information, see [Deleting a Service-Linked Role](https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html#delete-service-linked-role) in the
_IAM User Guide_.

## Supported Regions for ACM SLRs

ACM supports using SLRs in all of the regions where both
ACM and AWS Private CA are available. For more information, see
[AWS Regions and Endpoints](https://docs.aws.amazon.com/general/latest/gr/rande.html).

Region nameRegion identitySupport in ACMUS East (N. Virginia)us-east-1YesUS East (Ohio)us-east-2YesUS West (N. California)us-west-1YesUS West (Oregon)us-west-2YesAsia Pacific (Mumbai)ap-south-1YesAsia Pacific (Osaka)ap-northeast-3YesAsia Pacific (Seoul)ap-northeast-2YesAsia Pacific (Singapore)ap-southeast-1YesAsia Pacific (Sydney)ap-southeast-2YesAsia Pacific (Tokyo)ap-northeast-1YesCanada (Central)ca-central-1YesEurope (Frankfurt)eu-central-1YesEurope (Zurich)eu-central-2YesEurope (Ireland)eu-west-1YesEurope (London)eu-west-2YesEurope (Paris)eu-west-3YesSouth America (São Paulo)sa-east-1YesAWS GovCloud (US-West)us-gov-west-1YesAWS GovCloud (US-East)
Eastus-gov-east-1Yes

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Use condition keys

Troubleshooting
