This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Endpoint FailoverConfig

The failover configuration for an endpoint. This includes what triggers failover and what
happens when it's triggered.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Primary" : Primary,
  "Secondary" : Secondary
}

```

### YAML

```yaml

  Primary:
    Primary
  Secondary:
    Secondary

```

## Properties

`Primary`

The main Region of the endpoint.

_Required_: Yes

_Type_: [Primary](aws-properties-events-endpoint-primary.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Secondary`

The Region that events are routed to when failover is triggered or event replication is
enabled.

_Required_: Yes

_Type_: [Secondary](aws-properties-events-endpoint-secondary.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EndpointEventBus

Primary

All content copied from https://docs.aws.amazon.com/.
