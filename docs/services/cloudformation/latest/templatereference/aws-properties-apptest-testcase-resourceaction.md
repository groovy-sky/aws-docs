This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppTest::TestCase ResourceAction

Specifies a resource action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudFormationAction" : CloudFormationAction,
  "M2ManagedApplicationAction" : M2ManagedApplicationAction,
  "M2NonManagedApplicationAction" : M2NonManagedApplicationAction
}

```

### YAML

```yaml

  CloudFormationAction:
    CloudFormationAction
  M2ManagedApplicationAction:
    M2ManagedApplicationAction
  M2NonManagedApplicationAction:
    M2NonManagedApplicationAction

```

## Properties

`CloudFormationAction`

The CloudFormation action of the resource action.

_Required_: No

_Type_: [CloudFormationAction](aws-properties-apptest-testcase-cloudformationaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`M2ManagedApplicationAction`

The AWS Mainframe Modernization managed application action of the resource action.

_Required_: No

_Type_: [M2ManagedApplicationAction](aws-properties-apptest-testcase-m2managedapplicationaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`M2NonManagedApplicationAction`

The AWS Mainframe Modernization non-managed application action of the resource action.

_Required_: No

_Type_: [M2NonManagedApplicationAction](aws-properties-apptest-testcase-m2nonmanagedapplicationaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OutputFile

Script

All content copied from https://docs.aws.amazon.com/.
