This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan ExecutionApprovalConfiguration

Configuration for approval steps in a Region switch plan execution. Approval steps require manual intervention before the execution can proceed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApprovalRole" : String,
  "TimeoutMinutes" : Number
}

```

### YAML

```yaml

  ApprovalRole: String
  TimeoutMinutes: Number

```

## Properties

`ApprovalRole`

The IAM approval role for the configuration.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutMinutes`

The timeout value specified for the configuration.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EksResourceScalingUngraceful

ExecutionBlockConfiguration

All content copied from https://docs.aws.amazon.com/.
