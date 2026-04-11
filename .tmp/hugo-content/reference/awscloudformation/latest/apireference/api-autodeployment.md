# AutoDeployment

Describes whether StackSets automatically deploys to AWS Organizations accounts that are
added to a target organization or organizational unit (OU). For more information, see [Enable or\
disable automatic deployments for StackSets in AWS Organizations](../../../../services/cloudformation/latest/userguide/stacksets-orgs-manage-auto-deployment.md) in the
_AWS CloudFormation User Guide_.

## Contents

**DependsOn.member.N**

A list of StackSet ARNs that this StackSet depends on for auto-deployment operations.
When auto-deployment is triggered, operations will be sequenced to ensure all dependencies
complete successfully before this StackSet's operation begins.

Type: Array of strings

Required: No

**Enabled**

If set to `true`, StackSets automatically deploys additional stack instances to
AWS Organizations accounts that are added to a target organization or organizational unit
(OU) in the specified Regions. If an account is removed from a target organization or OU,
StackSets deletes stack instances from the account in the specified Regions.

Type: Boolean

Required: No

**RetainStacksOnAccountRemoval**

If set to `true`, stack resources are retained when an account is removed from a
target organization or OU. If set to `false`, stack resources are deleted. Specify
only if `Enabled` is set to `True`.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/autodeployment.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/autodeployment.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/autodeployment.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Annotation

BatchDescribeTypeConfigurationsError

All content copied from https://docs.aws.amazon.com/.
