This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::Fleet PlayerGatewayConfiguration

Configuration settings for player gateway. Use these settings to specify advanced options for how player gateway handles connections.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GameServerIpProtocolSupported" : String
}

```

### YAML

```yaml

  GameServerIpProtocolSupported: String

```

## Properties

`GameServerIpProtocolSupported`

The IP protocol that your game servers support for player connections through player gateway. If the value is set to `IPv4`, GameLift will install and execute a lightweight IP translation software on fleet instances to receive and transform incoming IPv6 traffic to IPv4. If the value is set to `DUAL_STACK`, the lightweight IP translation software will not be installed on fleet instances. `DUAL_STACK` provides slightly better performance than `IPv4`.

_Required_: No

_Type_: String

_Allowed values_: `IPv4 | DUAL_STACK`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ManagedCapacityConfiguration

ResourceCreationLimitPolicy
