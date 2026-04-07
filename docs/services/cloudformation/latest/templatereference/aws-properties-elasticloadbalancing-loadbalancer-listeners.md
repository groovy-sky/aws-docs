This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancing::LoadBalancer Listeners

Specifies a listener for your Classic Load Balancer.

Modifying any property replaces the listener.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InstancePort" : String,
  "InstanceProtocol" : String,
  "LoadBalancerPort" : String,
  "PolicyNames" : [ String, ... ],
  "Protocol" : String,
  "SSLCertificateId" : String
}

```

### YAML

```yaml

  InstancePort: String
  InstanceProtocol: String
  LoadBalancerPort: String
  PolicyNames:
    - String
  Protocol: String
  SSLCertificateId: String

```

## Properties

`InstancePort`

The port on which the instance is listening.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`InstanceProtocol`

The protocol to use for routing traffic to instances: HTTP, HTTPS, TCP, or SSL.

If the front-end protocol is TCP or SSL, the back-end protocol must be TCP or SSL.
If the front-end protocol is HTTP or HTTPS, the back-end protocol must be HTTP or HTTPS.

If there is another listener with the same `InstancePort` whose `InstanceProtocol` is secure,
(HTTPS or SSL), the listener's `InstanceProtocol` must also be secure.

If there is another listener with the same `InstancePort` whose `InstanceProtocol` is HTTP or TCP,
the listener's `InstanceProtocol` must be HTTP or TCP.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`LoadBalancerPort`

The port on which the load balancer is listening. On EC2-VPC, you can specify any port from the range 1-65535. On EC2-Classic, you can specify any port from the following list: 25, 80, 443, 465, 587, 1024-65535.

_Required_: Yes

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`PolicyNames`

The names of the policies to associate with the listener.

_Required_: No

_Type_: Array of String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Protocol`

The load balancer transport protocol to use for routing: HTTP, HTTPS, TCP, or SSL.

_Required_: Yes

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`SSLCertificateId`

The Amazon Resource Name (ARN) of the server certificate.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

## See also

- [CreateLoadBalancerListeners](https://docs.aws.amazon.com/elasticloadbalancing/2012-06-01/APIReference/API_CreateLoadBalancerListeners.html)
in the _Elastic Load Balancing API Reference (version 2012-06-01)_

- [Listeners](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-listener-config.html)
in the _User Guide for Classic Load Balancers_

- [HTTPS Listeners](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-https-load-balancers.html)
in the _User Guide for Classic Load Balancers_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LBCookieStickinessPolicy

Policies
