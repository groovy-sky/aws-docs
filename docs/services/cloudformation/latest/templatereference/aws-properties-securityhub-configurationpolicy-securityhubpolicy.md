This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::ConfigurationPolicy SecurityHubPolicy

An object that defines how AWS Security Hub CSPM is configured. The configuration policy includes whether
Security Hub CSPM is enabled or disabled, a list of enabled security standards, a list of enabled or
disabled security controls, and a list of custom parameter values for specified controls. If you provide a list of security controls that are enabled in the configuration
policy, Security Hub CSPM disables all other controls (including newly released controls). If you provide a
list of security controls that are disabled in the configuration policy, Security Hub CSPM enables all other
controls (including newly released controls).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnabledStandardIdentifiers" : [ String, ... ],
  "SecurityControlsConfiguration" : SecurityControlsConfiguration,
  "ServiceEnabled" : Boolean
}

```

### YAML

```yaml

  EnabledStandardIdentifiers:
    - String
  SecurityControlsConfiguration:
    SecurityControlsConfiguration
  ServiceEnabled: Boolean

```

## Properties

`EnabledStandardIdentifiers`

A list that defines which security standards are enabled in the configuration policy.

This property is required only if `ServiceEnabled` is set to `true` in your configuration policy.

_Required_: Conditional

_Type_: Array of String

_Maximum_: `2048 | 1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityControlsConfiguration`

An object that defines which security controls are enabled in the configuration policy. The enablement status
of a control is aligned across all of the enabled standards in an account.

This property is required only if `ServiceEnabled` is set to true in your configuration policy.

_Required_: Conditional

_Type_: [SecurityControlsConfiguration](aws-properties-securityhub-configurationpolicy-securitycontrolsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceEnabled`

Indicates whether Security Hub CSPM is enabled in the policy.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SecurityControlsConfiguration

AWS::SecurityHub::ConnectorV2

All content copied from https://docs.aws.amazon.com/.
