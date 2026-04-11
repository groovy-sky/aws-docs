This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LaunchTemplate MaintenanceOptions

The maintenance options of your instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoRecovery" : String
}

```

### YAML

```yaml

  AutoRecovery: String

```

## Properties

`AutoRecovery`

Disables the automatic recovery behavior of your instance or sets it to
default.

_Required_: No

_Type_: String

_Allowed values_: `default | disabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LicenseSpecification

MemoryGiBPerVCpu

All content copied from https://docs.aws.amazon.com/.
