This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FMS::Policy ThirdPartyFirewallPolicy

Configures the deployment model for the third-party firewall.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FirewallDeploymentModel" : String
}

```

### YAML

```yaml

  FirewallDeploymentModel: String

```

## Properties

`FirewallDeploymentModel`

Defines the deployment model to use for the third-party firewall policy.

_Required_: Yes

_Type_: String

_Allowed values_: `DISTRIBUTED | CENTRALIZED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SecurityServicePolicyData

AWS::FMS::ResourceSet

All content copied from https://docs.aws.amazon.com/.
