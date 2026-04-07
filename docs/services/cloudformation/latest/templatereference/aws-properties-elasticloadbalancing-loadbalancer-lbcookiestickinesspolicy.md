This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancing::LoadBalancer LBCookieStickinessPolicy

Specifies a policy for duration-based session stickiness for your Classic Load Balancer.

To associate a policy with a listener, use the [PolicyNames](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-policynames)
property for the listener.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CookieExpirationPeriod" : String,
  "PolicyName" : String
}

```

### YAML

```yaml

  CookieExpirationPeriod: String
  PolicyName: String

```

## Properties

`CookieExpirationPeriod`

The time period, in seconds, after which the cookie should be considered stale. If this parameter is not specified, the stickiness session lasts for the duration of the browser session.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyName`

The name of the policy. This name must be unique within the set of policies for this load balancer.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [CreateLBCookieStickinessPolicy](https://docs.aws.amazon.com/elasticloadbalancing/2012-06-01/APIReference/API_CreateLBCookieStickinessPolicy.html)
in the _Elastic Load Balancing API Reference (version 2012-06-01)_

- [Sticky Sessions](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-sticky-sessions.html)
in the _User Guide for Classic Load Balancers_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

HealthCheck

Listeners
