This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::LoadBalancer

The `AWS::Lightsail::LoadBalancer` resource specifies a load balancer that
can be used with Lightsail instances.

###### Note

You cannot attach a TLS certificate to a load balancer using the
`AWS::Lightsail::LoadBalancer` resource type. Instead, use the
`AWS::Lightsail::LoadBalancerTlsCertificate` resource type to create a certificate
and attach it to a load balancer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lightsail::LoadBalancer",
  "Properties" : {
      "AttachedInstances" : [ String, ... ],
      "HealthCheckPath" : String,
      "InstancePort" : Integer,
      "IpAddressType" : String,
      "LoadBalancerName" : String,
      "SessionStickinessEnabled" : Boolean,
      "SessionStickinessLBCookieDurationSeconds" : String,
      "Tags" : [ Tag, ... ],
      "TlsPolicyName" : String
    }
}

```

### YAML

```yaml

Type: AWS::Lightsail::LoadBalancer
Properties:
  AttachedInstances:
    - String
  HealthCheckPath: String
  InstancePort: Integer
  IpAddressType: String
  LoadBalancerName: String
  SessionStickinessEnabled: Boolean
  SessionStickinessLBCookieDurationSeconds: String
  Tags:
    - Tag
  TlsPolicyName: String

```

## Properties

`AttachedInstances`

The Lightsail instances to attach to the load balancer.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthCheckPath`

The path on the attached instance where the health check will be performed. If no path
is specified, the load balancer tries to make a request to the default (root) page
( `/index.html`).

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstancePort`

The port that the load balancer uses to direct traffic to your Lightsail
instances. For HTTP traffic, specify port `80`. For HTTPS traffic, specify port
`443`.

_Required_: Yes

_Type_: Integer

_Update requires_: Updates are not supported.

`IpAddressType`

The IP address type of the load balancer.

The possible values are `ipv4` for IPv4 only, and `dualstack` for
both IPv4 and IPv6.

_Required_: No

_Type_: String

_Allowed values_: `dualstack | ipv4 | ipv6`

_Update requires_: Updates are not supported.

`LoadBalancerName`

The name of the load balancer.

_Required_: Yes

_Type_: String

_Pattern_: `\w[\w\-]*\w`

_Update requires_: Updates are not supported.

`SessionStickinessEnabled`

A Boolean value indicating whether session stickiness is enabled.

Enable session stickiness (also known as _session affinity_) to bind
a user's session to a specific instance. This ensures that all requests from the user
during the session are sent to the same instance.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SessionStickinessLBCookieDurationSeconds`

The time period, in seconds, after which the load balancer session stickiness cookie
should be considered stale. If you do not specify this parameter, the default value is 0,
which indicates that the sticky session should last for the duration of the browser
session.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md)
in the _AWS CloudFormation User Guide_.

###### Note

The `Value` of `Tags` is optional for Lightsail resources.

_Required_: No

_Type_: Array of [Tag](aws-properties-lightsail-loadbalancer-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TlsPolicyName`

The name of the TLS security policy for the load balancer.

_Required_: No

_Type_: String

_Pattern_: `\w[\w\-]*\w`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a unique identifier for this resource.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`LoadBalancerArn`

The Amazon Resource Name (ARN) of the load balancer.

## Remarks

_Load balancer TLS certificate_

You cannot attach a TLS certificate to a load balancer using the
`AWS::Lightsail::LoadBalancer` resource type. Instead, use the
`AWS::Lightsail::LoadBalancerTlsCertificate` resource type to create
and attach certificates to a load balancer.

_Configuring HTTPS redirection_

HTTPS redirection can only be set using the `HttpsRedirectionEnabled`
parameter on the `AWS::Lightsail::LoadBalancerTlsCertificate` resource
that is attached to the load balancer.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
