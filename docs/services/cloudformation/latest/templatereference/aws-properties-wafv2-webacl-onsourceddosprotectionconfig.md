This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL OnSourceDDoSProtectionConfig

Configures the level of DDoS protection that applies to web ACLs associated with Application Load Balancers.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ALBLowReputationMode" : String
}

```

### YAML

```yaml

  ALBLowReputationMode: String

```

## Properties

`ALBLowReputationMode`

The level of DDoS protection that applies to web ACLs associated with Application Load Balancers. `ACTIVE_UNDER_DDOS` protection is enabled by default whenever a web ACL is associated with an Application Load Balancer.
In the event that an Application Load Balancer experiences high-load conditions or suspected DDoS attacks, the `ACTIVE_UNDER_DDOS` protection automatically rate limits traffic from known low reputation sources without disrupting Application Load Balancer availability.
`ALWAYS_ON` protection provides constant, always-on monitoring of known low reputation sources for suspected DDoS attacks. While this provides a higher level of protection, there may be potential impacts on legitimate traffic.

_Required_: Yes

_Type_: String

_Allowed values_: `ACTIVE_UNDER_DDOS | ALWAYS_ON`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NotStatement

OrStatement
