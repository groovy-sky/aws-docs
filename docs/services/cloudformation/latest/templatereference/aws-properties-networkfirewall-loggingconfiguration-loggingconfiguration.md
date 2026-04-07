This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::LoggingConfiguration LoggingConfiguration

Defines how AWS Network Firewall performs logging for a firewall.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogDestinationConfigs" : [ LogDestinationConfig, ... ]
}

```

### YAML

```yaml

  LogDestinationConfigs:
    - LogDestinationConfig

```

## Properties

`LogDestinationConfigs`

Defines the logging destinations for the logs for a firewall. Network Firewall generates
logs for stateful rule groups.

_Required_: Yes

_Type_: Array of [LogDestinationConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-networkfirewall-loggingconfiguration-logdestinationconfig.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LogDestinationConfig

AWS::NetworkFirewall::RuleGroup
