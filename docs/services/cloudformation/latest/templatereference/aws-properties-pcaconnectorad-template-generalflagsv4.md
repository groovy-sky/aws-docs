This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorAD::Template GeneralFlagsV4

General flags for v4 template schema that defines if the template is for a machine or a
user and if the template can be issued using autoenrollment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoEnrollment" : Boolean,
  "MachineType" : Boolean
}

```

### YAML

```yaml

  AutoEnrollment: Boolean
  MachineType: Boolean

```

## Properties

`AutoEnrollment`

Allows certificate issuance using autoenrollment. Set to TRUE to allow
autoenrollment.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MachineType`

Defines if the template is for machines or users. Set to TRUE if the template is for
machines. Set to FALSE if the template is for users

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeneralFlagsV3

KeyUsage

All content copied from https://docs.aws.amazon.com/.
