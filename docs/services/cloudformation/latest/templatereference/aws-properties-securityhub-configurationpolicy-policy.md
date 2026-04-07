This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::ConfigurationPolicy Policy

An object that defines how AWS Security Hub CSPM is configured. It includes whether Security Hub CSPM is
enabled or disabled, a list of enabled security standards, a list of enabled or disabled security controls, and a list of custom parameter values for specified controls.
If you provide a list of security controls that are enabled in the configuration policy, Security Hub CSPM
disables all other controls (including newly released controls). If you provide a list of security controls that
are disabled in the configuration policy, Security Hub CSPM enables all other controls (including newly
released controls).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecurityHub" : SecurityHubPolicy
}

```

### YAML

```yaml

  SecurityHub:
    SecurityHubPolicy

```

## Properties

`SecurityHub`

The AWS service that the configuration policy applies to.

_Required_: No

_Type_: [SecurityHubPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-configurationpolicy-securityhubpolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ParameterValue

SecurityControlCustomParameter
