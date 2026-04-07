This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Endpoint RoutingConfig

The routing configuration of the endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FailoverConfig" : FailoverConfig
}

```

### YAML

```yaml

  FailoverConfig:
    FailoverConfig

```

## Properties

`FailoverConfig`

The failover configuration for an endpoint. This includes what triggers failover and what
happens when it's triggered.

_Required_: Yes

_Type_: [FailoverConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-events-endpoint-failoverconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ReplicationConfig

Secondary
