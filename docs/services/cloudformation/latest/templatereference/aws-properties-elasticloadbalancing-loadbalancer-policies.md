This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancing::LoadBalancer Policies

Specifies policies for your Classic Load Balancer.

To associate policies with a listener, use the [PolicyNames](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-policynames)
property for the listener.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Attributes" : [ {Key: Value, ...}, ... ],
  "InstancePorts" : [ String, ... ],
  "LoadBalancerPorts" : [ String, ... ],
  "PolicyName" : String,
  "PolicyType" : String
}

```

### YAML

```yaml

  Attributes:
    -
    Key: Value
  InstancePorts:
    - String
  LoadBalancerPorts:
    - String
  PolicyName: String
  PolicyType: String

```

## Properties

`Attributes`

The policy attributes.

_Required_: Yes

_Type_: Array of Object

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstancePorts`

The instance ports for the policy. Required only for some policy types.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoadBalancerPorts`

The load balancer ports for the policy. Required only for some policy types.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyName`

The name of the policy.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyType`

The name of the policy type.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

#### JSON

```json

"Policies": [{
    "PolicyName": "My-SSLNegotiation-Policy",
    "PolicyType": "SSLNegotiationPolicyType",
    "Attributes": [{
        "Name": "Reference-Security-Policy",
        "Value": "ELBSecurityPolicy-TLS-1-2-2017-01"
    }]
}]
```

#### YAML

```yaml

Policies:
    - PolicyName: My-SSLNegotiation-Policy
      PolicyType: SSLNegotiationPolicyType
      Attributes:
      - Name: Reference-Security-Policy
        Value: ELBSecurityPolicy-TLS-1-2-2017-01
```

## See also

- [CreateLoadBalancerPolicy](https://docs.aws.amazon.com/elasticloadbalancing/2012-06-01/APIReference/API_CreateLoadBalancerPolicy.html)
in the _Elastic Load Balancing API Reference (version 2012-06-01)_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Listeners

Tag
