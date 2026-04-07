This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::DomainName DomainNameConfiguration

The `DomainNameConfiguration` property type specifies the configuration
for an API's domain name.

`DomainNameConfiguration` is a property of the [AWS::ApiGatewayV2::DomainName](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-domainname.html) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateArn" : String,
  "CertificateName" : String,
  "EndpointType" : String,
  "IpAddressType" : String,
  "OwnershipVerificationCertificateArn" : String,
  "SecurityPolicy" : String
}

```

### YAML

```yaml

  CertificateArn: String
  CertificateName: String
  EndpointType: String
  IpAddressType: String
  OwnershipVerificationCertificateArn: String
  SecurityPolicy: String

```

## Properties

`CertificateArn`

An AWS-managed certificate that will be used by the edge-optimized endpoint for this domain name. AWS Certificate Manager is the only supported source.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CertificateName`

The user-friendly name of the certificate that will be used by the edge-optimized endpoint for this domain name.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointType`

The endpoint type.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpAddressType`

The IP address types that can invoke the domain name. Use `ipv4` to allow only IPv4 addresses to
invoke your domain name, or use `dualstack` to allow both IPv4 and IPv6 addresses to invoke your domain
name.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OwnershipVerificationCertificateArn`

The Amazon resource name (ARN) for the public certificate issued by AWS Certificate Manager. This ARN is used to validate custom domain ownership. It's required only if you configure mutual TLS and use either an ACM-imported or a private CA certificate ARN as the regionalCertificateArn.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityPolicy`

The Transport Layer Security (TLS) version of the security policy for this domain name. The valid values are `TLS_1_0` and `TLS_1_2`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [DomainNameConfiguration](https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/domainnames-domainname.html#domainnames-domainname-model-domainnameconfiguration) in the _Amazon_
_API Gateway Version 2 API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ApiGatewayV2::DomainName

MutualTlsAuthentication
