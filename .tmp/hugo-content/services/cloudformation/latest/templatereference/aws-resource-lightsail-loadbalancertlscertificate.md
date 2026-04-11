This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::LoadBalancerTlsCertificate

The `AWS::Lightsail::LoadBalancerTlsCertificate` resource specifies a TLS
certificate that can be used with a Lightsail load balancer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lightsail::LoadBalancerTlsCertificate",
  "Properties" : {
      "CertificateAlternativeNames" : [ String, ... ],
      "CertificateDomainName" : String,
      "CertificateName" : String,
      "HttpsRedirectionEnabled" : Boolean,
      "IsAttached" : Boolean,
      "LoadBalancerName" : String
    }
}

```

### YAML

```yaml

Type: AWS::Lightsail::LoadBalancerTlsCertificate
Properties:
  CertificateAlternativeNames:
    - String
  CertificateDomainName: String
  CertificateName: String
  HttpsRedirectionEnabled: Boolean
  IsAttached: Boolean
  LoadBalancerName: String

```

## Properties

`CertificateAlternativeNames`

An array of alternative domain names and subdomain names for your SSL/TLS
certificate.

In addition to the primary domain name, you can have up to nine alternative domain names.
Wildcards (such as `*.example.com`) are not supported.

_Required_: No

_Type_: Array of String

_Update requires_: Updates are not supported.

`CertificateDomainName`

The domain name for the SSL/TLS certificate. For example, `example.com` or `www.example.com`.

_Required_: Yes

_Type_: String

_Update requires_: Updates are not supported.

`CertificateName`

The name of the SSL/TLS certificate.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`HttpsRedirectionEnabled`

A Boolean value indicating whether HTTPS redirection is enabled for the load
balancer that the TLS certificate is attached to.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsAttached`

A Boolean value indicating whether the SSL/TLS certificate is attached to a Lightsail load balancer.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoadBalancerName`

The name of the load balancer that the SSL/TLS certificate is attached to.

_Required_: Yes

_Type_: String

_Pattern_: `\w[\w\-]*\w`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a unique identifier for this resource.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`LoadBalancerTlsCertificateArn`

The Amazon Resource Name (ARN) of the SSL/TLS certificate.

`Status`

The validation status of the SSL/TLS certificate.

Valid Values: `PENDING_VALIDATION` \| `ISSUED` \|
`INACTIVE` \| `EXPIRED` \| `VALIDATION_TIMED_OUT` \|
`REVOKED` \| `FAILED` \| `UNKNOWN`

## Remarks

_Attaching certificates to load balancers_

Use the `IsAttached` parameter to attach a certificate to a load balancer.
The certificate must be in a valid state before it can be attached.

_Replacing certificates attached to load balancers_

After a certificate is attached to a load balancer, it cannot be detached. It can only be replaced.
If the `isAttached` parameter is changed from `true` to
`false` for a certificate, it won’t be detached from the load balancer and
the stack will drift. You can replace a certificate by changing the
`isAttached` parameter of a different certificate to `true` and
changing the current certificate’s `isAttached` parameter to
`false`.

_Maximum attached certificates_

Don't attach more than one certificate to a load balancer. If you attach multiple certificates
to a load balancer, the behavior is unpredictable, and any one of the certificates might be in effect.
This will cause the stack to drift because only one of the
certificates is attached to the load balancer, but the template shows
multiple.

_Configuring HTTPS redirection_

The `HttpsRedirectionEnabled` parameter can only be set on a certificate
that is in a valid state and is also attached to a load balancer.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::Lightsail::StaticIp

All content copied from https://docs.aws.amazon.com/.
