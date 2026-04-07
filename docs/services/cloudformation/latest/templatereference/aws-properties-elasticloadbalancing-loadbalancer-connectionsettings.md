This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancing::LoadBalancer ConnectionSettings

Specifies the idle timeout value for your Classic Load Balancer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IdleTimeout" : Integer
}

```

### YAML

```yaml

  IdleTimeout: Integer

```

## Properties

`IdleTimeout`

The time, in seconds, that the connection is allowed to be idle (no data has been sent over the connection) before it is closed by the load balancer.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `3600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ModifyLoadBalancerAttributes](https://docs.aws.amazon.com/elasticloadbalancing/2012-06-01/APIReference/API_ModifyLoadBalancerAttributes.html)
in the _Elastic Load Balancing API Reference (version 2012-06-01)_

- [Idle Connection Timeout](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/config-idle-timeout.html)
in the _User Guide for Classic Load Balancers_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ConnectionDrainingPolicy

HealthCheck
