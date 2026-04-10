# StackSet

A structure that contains information about a StackSet. With StackSets, you can provision
stacks across AWS accounts and Regions from a single CloudFormation template. Each stack is based
on the same CloudFormation template, but you can customize individual stacks using parameters.

## Contents

**AdministrationRoleARN**

The Amazon Resource Name (ARN) of the IAM role used to create or update the stack
set.

Use customized administrator roles to control which users or groups can manage specific
StackSets within the same administrator account. For more information, see [Prerequisites for using CloudFormation StackSets](../../../../services/cloudformation/latest/userguide/stacksets-prereqs.md) in the
_AWS CloudFormation User Guide_.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Required: No

**AutoDeployment**

Describes whether StackSets automatically deploys to AWS Organizations accounts that are
added to a target organization or organizational unit (OU). Valid only if the StackSet uses
service-managed permissions.

Type: [AutoDeployment](api-autodeployment.md) object

Required: No

**Capabilities.member.N**

The capabilities that are allowed in the StackSet. Some StackSet templates might include
resources that can affect permissions in your AWS account—for example, by creating new
AWS Identity and Access Management (IAM) users. For more information, see [Acknowledging IAM resources in CloudFormation templates](../../../../services/cloudformation/latest/userguide/control-access-with-iam.md#using-iam-capabilities).

Type: Array of strings

Valid Values: `CAPABILITY_IAM | CAPABILITY_NAMED_IAM | CAPABILITY_AUTO_EXPAND`

Required: No

**Description**

A description of the StackSet that you specify when the StackSet is created or
updated.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**ExecutionRoleName**

The name of the IAM execution role used to create or update the StackSet.

Use customized execution roles to control which stack resources users and groups can include
in their StackSets.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Pattern: `[a-zA-Z_0-9+=,.@-]+`

Required: No

**ManagedExecution**

Describes whether StackSets performs non-conflicting operations concurrently and queues
conflicting operations.

Type: [ManagedExecution](api-managedexecution.md) object

Required: No

**OrganizationalUnitIds.member.N**

\[Service-managed permissions\] The organization root ID or organizational unit (OU) IDs that
you specified for [DeploymentTargets](api-deploymenttargets.md).

Type: Array of strings

Pattern: `^(ou-[a-z0-9]{4,32}-[a-z0-9]{8,32}|r-[a-z0-9]{4,32})$`

Required: No

**Parameters.member.N**

A list of input parameters for a StackSet.

Type: Array of [Parameter](api-parameter.md) objects

Required: No

**PermissionModel**

Describes how the IAM roles required for StackSet operations are created.

- With `self-managed` permissions, you must create the administrator and
execution roles required to deploy to target accounts. For more information, see [Grant\
self-managed permissions](../../../../services/cloudformation/latest/userguide/stacksets-prereqs-self-managed.md).

- With `service-managed` permissions, StackSets automatically creates the IAM
roles required to deploy to accounts managed by AWS Organizations. For more information,
see [Activate\
trusted access for StackSets with AWS Organizations](../../../../services/cloudformation/latest/userguide/stacksets-orgs-activate-trusted-access.md).

Type: String

Valid Values: `SERVICE_MANAGED | SELF_MANAGED`

Required: No

**Regions.member.N**

Returns a list of all AWS Regions the given StackSet has stack instances deployed in. The
AWS Regions list output is in no particular order.

Type: Array of strings

Pattern: `^[a-zA-Z0-9-]{1,128}$`

Required: No

**StackSetARN**

The Amazon Resource Name (ARN) of the StackSet.

Type: String

Required: No

**StackSetDriftDetectionDetails**

Detailed information about the drift status of the StackSet.

For StackSets, contains information about the last _completed_ drift
operation performed on the StackSet. Information about drift operations currently in progress
isn't included.

Type: [StackSetDriftDetectionDetails](api-stacksetdriftdetectiondetails.md) object

Required: No

**StackSetId**

The ID of the StackSet.

Type: String

Required: No

**StackSetName**

The name that's associated with the StackSet.

Type: String

Required: No

**Status**

The status of the StackSet.

Type: String

Valid Values: `ACTIVE | DELETED`

Required: No

**Tags.member.N**

A list of tags that specify information about the StackSet. A maximum number of 50 tags can
be specified.

Type: Array of [Tag](api-tag.md) objects

Array Members: Maximum number of 50 items.

Required: No

**TemplateBody**

The structure that contains the body of the template that was used to create or update the
StackSet.

Type: String

Length Constraints: Minimum length of 1.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/stackset.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/stackset.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/stackset.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StackResourceSummary

StackSetAutoDeploymentTargetSummary

All content copied from https://docs.aws.amazon.com/.
