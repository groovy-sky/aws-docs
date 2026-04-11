This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::ConfigurationPolicy SecurityControlsConfiguration

An object that defines which security controls are enabled in an AWS Security Hub CSPM configuration policy.
The enablement status of a control is aligned across all of the enabled standards in an account.

This property is required only if `ServiceEnabled` is set to `true` in your configuration policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DisabledSecurityControlIdentifiers" : [ String, ... ],
  "EnabledSecurityControlIdentifiers" : [ String, ... ],
  "SecurityControlCustomParameters" : [ SecurityControlCustomParameter, ... ]
}

```

### YAML

```yaml

  DisabledSecurityControlIdentifiers:
    - String
  EnabledSecurityControlIdentifiers:
    - String
  SecurityControlCustomParameters:
    - SecurityControlCustomParameter

```

## Properties

`DisabledSecurityControlIdentifiers`

A list of security controls that are disabled in the configuration policy.

Provide only one of `EnabledSecurityControlIdentifiers` or
`DisabledSecurityControlIdentifiers`.

If you provide `DisabledSecurityControlIdentifiers`, Security Hub CSPM enables all other controls not in
the list, and enables
[AutoEnableControls](../../../../reference/securityhub/1-0/apireference/api-updatesecurityhubconfiguration.md#securityhub-UpdateSecurityHubConfiguration-request-AutoEnableControls).

_Required_: No

_Type_: Array of String

_Maximum_: `2048 | 1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnabledSecurityControlIdentifiers`

A list of security controls that are enabled in the configuration policy.

Provide only one of `EnabledSecurityControlIdentifiers` or
`DisabledSecurityControlIdentifiers`.

If you provide `EnabledSecurityControlIdentifiers`, Security Hub CSPM disables all other controls not in
the list, and disables
[AutoEnableControls](../../../../reference/securityhub/1-0/apireference/api-updatesecurityhubconfiguration.md#securityhub-UpdateSecurityHubConfiguration-request-AutoEnableControls).

_Required_: No

_Type_: Array of String

_Maximum_: `2048 | 1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityControlCustomParameters`

A list of security controls and control parameter values that are included in a configuration policy.

_Required_: No

_Type_: Array of [SecurityControlCustomParameter](aws-properties-securityhub-configurationpolicy-securitycontrolcustomparameter.md)

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SecurityControlCustomParameter

SecurityHubPolicy

All content copied from https://docs.aws.amazon.com/.
