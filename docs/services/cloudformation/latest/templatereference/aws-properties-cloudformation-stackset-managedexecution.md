This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::StackSet ManagedExecution

Describes whether StackSets performs non-conflicting operations concurrently and queues
conflicting operations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Active" : Boolean
}

```

### YAML

```yaml

  Active: Boolean

```

## Properties

`Active`

When `true`, CloudFormation performs non-conflicting operations concurrently and
queues conflicting operations. After conflicting operations finish, CloudFormation starts queued
operations in request order.

###### Note

If there are already running or queued operations, CloudFormation queues all incoming
operations even if they are non-conflicting.

You can't modify your StackSet's execution configuration while there are running or queued
operations for that StackSet.

When `false` (default), StackSets performs one operation at a time in request
order.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeploymentTargets

OperationPreferences

All content copied from https://docs.aws.amazon.com/.
