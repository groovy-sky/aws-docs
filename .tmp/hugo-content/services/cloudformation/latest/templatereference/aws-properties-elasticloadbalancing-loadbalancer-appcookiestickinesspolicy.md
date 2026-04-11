This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancing::LoadBalancer AppCookieStickinessPolicy

Specifies a policy for application-controlled session stickiness for your Classic Load Balancer.

To associate a policy with a listener, use the [PolicyNames](../userguide/aws-properties-ec2-elb-listener.md#cfn-ec2-elb-listener-policynames)
property for the listener.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CookieName" : String,
  "PolicyName" : String
}

```

### YAML

```yaml

  CookieName: String
  PolicyName: String

```

## Properties

`CookieName`

The name of the application cookie used for stickiness.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyName`

The mnemonic name for the policy being created. The name must be unique within a set of policies for this load balancer.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [CreateAppCookieStickinessPolicy](../../../../reference/elasticloadbalancing/2012-06-01/apireference/api-createappcookiestickinesspolicy.md)
in the _Elastic Load Balancing API Reference (version 2012-06-01)_

- [Sticky Sessions](../../../elasticloadbalancing/latest/classic/elb-sticky-sessions.md)
in the _User Guide for Classic Load Balancers_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccessLoggingPolicy

ConnectionDrainingPolicy

All content copied from https://docs.aws.amazon.com/.
