This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::Fleet ProxyConfiguration

Information about the proxy configurations that apply network access control to your reserved capacity instances.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultBehavior" : String,
  "OrderedProxyRules" : [ FleetProxyRule, ... ]
}

```

### YAML

```yaml

  DefaultBehavior: String
  OrderedProxyRules:
    - FleetProxyRule

```

## Properties

`DefaultBehavior`

The default behavior of outgoing traffic.

_Required_: No

_Type_: String

_Allowed values_: `ALLOW_ALL | DENY_ALL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrderedProxyRules`

An array of `FleetProxyRule` objects that represent the specified destination domains or IPs to allow or deny network access control to.

_Required_: No

_Type_: Array of [FleetProxyRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-codebuild-fleet-fleetproxyrule.html)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FleetProxyRule

ScalingConfigurationInput
