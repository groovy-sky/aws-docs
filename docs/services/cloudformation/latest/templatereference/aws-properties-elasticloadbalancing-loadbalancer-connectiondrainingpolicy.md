This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancing::LoadBalancer ConnectionDrainingPolicy

Specifies the connection draining settings for your Classic Load Balancer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "Timeout" : Integer
}

```

### YAML

```yaml

  Enabled: Boolean
  Timeout: Integer

```

## Properties

`Enabled`

Specifies whether connection draining is enabled for the load balancer.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timeout`

The maximum time, in seconds, to keep the existing connections open before deregistering the instances.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ModifyLoadBalancerAttributes](https://docs.aws.amazon.com/elasticloadbalancing/2012-06-01/APIReference/API_ModifyLoadBalancerAttributes.html)
in the _Elastic Load Balancing API Reference (version 2012-06-01)_

- [Connection Draining](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/config-conn-drain.html)
in the _User Guide for Classic Load Balancers_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AppCookieStickinessPolicy

ConnectionSettings
