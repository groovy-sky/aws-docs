This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::FirewallPolicy FlowTimeouts

Describes the amount of time that can pass without any traffic sent through the firewall before the firewall determines that the connection is idle and Network Firewall removes the flow entry from its flow table.
Existing connections and flows are not impacted when you update this value. Only new connections after you update this value are impacted.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TcpIdleTimeoutSeconds" : Integer
}

```

### YAML

```yaml

  TcpIdleTimeoutSeconds: Integer

```

## Properties

`TcpIdleTimeoutSeconds`

The number of seconds that can pass without any TCP traffic sent through the firewall before the firewall determines that the connection is idle.
After the idle timeout passes, data packets are dropped, however, the next TCP SYN packet is considered a new flow and is processed by the firewall.
Clients or targets can use TCP keepalive packets to reset the idle timeout.

You can define the `TcpIdleTimeoutSeconds` value to be between 60 and 6000 seconds. If no value is provided, it defaults to 350 seconds.

_Required_: No

_Type_: Integer

_Minimum_: `60`

_Maximum_: `6000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FirewallPolicy

IPSet

All content copied from https://docs.aws.amazon.com/.
