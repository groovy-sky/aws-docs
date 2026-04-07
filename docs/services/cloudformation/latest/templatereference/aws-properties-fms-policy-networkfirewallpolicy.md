This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FMS::Policy NetworkFirewallPolicy

Configures the firewall policy deployment model of AWS Network Firewall. For information about
Network Firewall deployment models, see [AWS Network Firewall example\
architectures with routing](https://docs.aws.amazon.com/network-firewall/latest/developerguide/architectures.html) in the _Network Firewall Developer_
_Guide_.

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

Defines the deployment model to use for the firewall policy. To use a distributed model,
set [FirewallDeploymentModel](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-thirdpartyfirewallpolicy.html) to
`DISTRIBUTED`.

_Required_: Yes

_Type_: String

_Allowed values_: `DISTRIBUTED | CENTRALIZED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NetworkAclEntrySet

PolicyOption
