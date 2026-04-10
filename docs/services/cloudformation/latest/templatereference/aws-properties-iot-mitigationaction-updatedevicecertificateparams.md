This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::MitigationAction UpdateDeviceCertificateParams

Parameters to define a mitigation action that changes the state of the device certificate to inactive.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String
}

```

### YAML

```yaml

  Action: String

```

## Properties

`Action`

The action that you want to apply to the device certificate. The only supported value is `DEACTIVATE`.

_Required_: Yes

_Type_: String

_Allowed values_: `DEACTIVATE | UNSET_VALUE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateCACertificateParams

AWS::IoT::Policy

All content copied from https://docs.aws.amazon.com/.
