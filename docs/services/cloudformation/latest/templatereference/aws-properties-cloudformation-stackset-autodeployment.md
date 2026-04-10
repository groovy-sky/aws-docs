This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::StackSet AutoDeployment

Describes whether StackSets automatically deploys to AWS Organizations accounts
that are added to a target organization or organizational unit (OU). For more
information, see [Enable or disable automatic deployments for StackSets in AWS Organizations](../userguide/stacksets-orgs-manage-auto-deployment.md) in the _CloudFormation User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DependsOn" : [ String, ... ],
  "Enabled" : Boolean,
  "RetainStacksOnAccountRemoval" : Boolean
}

```

### YAML

```yaml

  DependsOn:
    - String
  Enabled: Boolean
  RetainStacksOnAccountRemoval: Boolean

```

## Properties

`DependsOn`

A list of StackSet ARNs that this StackSet depends on for auto-deployment operations. When auto-deployment is triggered, operations will be sequenced to ensure all dependencies complete successfully before this StackSet's operation begins.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

If set to `true`, StackSets automatically deploys additional stack
instances to AWS Organizations accounts that are added to a target organization or
organizational unit (OU) in the specified Regions. If an account is removed from a
target organization or OU, StackSets deletes stack instances from the account in the
specified Regions.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetainStacksOnAccountRemoval`

If set to `true`, stack resources are retained when an account is removed
from a target organization or OU. If set to `false`, stack resources are
deleted. Specify only if `Enabled` is set to `True`.

_Required_: Conditional

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CloudFormation::StackSet

DeploymentTargets

All content copied from https://docs.aws.amazon.com/.
