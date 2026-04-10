This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::ConfigurationPolicy SecurityControlCustomParameter

A list of security controls and control parameter values that are included in a configuration policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Parameters" : {Key: Value, ...},
  "SecurityControlId" : String
}

```

### YAML

```yaml

  Parameters:
    Key: Value
  SecurityControlId: String

```

## Properties

`Parameters`

An object that specifies parameter values for a control in a configuration policy.

_Required_: No

_Type_: Object of [ParameterConfiguration](aws-properties-securityhub-configurationpolicy-parameterconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityControlId`

The ID of the security control.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Policy

SecurityControlsConfiguration

All content copied from https://docs.aws.amazon.com/.
