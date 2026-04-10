This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution ViewerCertificate

A complex type that determines the distribution's SSL/TLS configuration for
communicating with viewers.

If the distribution doesn't use `Aliases` (also known as alternate domain
names or CNAMEs)—that is, if the distribution uses the CloudFront domain name such as
`d111111abcdef8.cloudfront.net`—set
`CloudFrontDefaultCertificate` to `true` and leave all other
fields empty.

If the distribution uses `Aliases` (alternate domain names or CNAMEs), use
the fields in this type to specify the following settings:

- Which viewers the distribution accepts HTTPS connections from: only viewers
that support [server name indication (SNI)](https://en.wikipedia.org/wiki/Server_Name_Indication) (recommended), or all viewers
including those that don't support SNI.

- To accept HTTPS connections from only viewers that support SNI, set
`SSLSupportMethod` to `sni-only`. This is
recommended. Most browsers and clients support SNI.
(In CloudFormation, the field name is `SslSupportMethod`.
Note the different capitalization.)

- To accept HTTPS connections from all viewers, including those that
don't support SNI, set `SSLSupportMethod` to
`vip`. This is not recommended, and results in additional
monthly charges from CloudFront. (In CloudFormation,
the field name is `SslSupportMethod`. Note the different
capitalization.)

- The minimum SSL/TLS protocol version that the distribution can use to
communicate with viewers. To specify a minimum version, choose a value for
`MinimumProtocolVersion`. For more information, see [Security Policy](../../../amazoncloudfront/latest/developerguide/distribution-web-values-specify.md#DownloadDistValues-security-policy) in the
_Amazon CloudFront Developer Guide_.

- The location of the SSL/TLS certificate, [AWS Certificate Manager\
(ACM)](../../../acm/latest/userguide/acm-overview.md) (recommended) or [AWS Identity and Access Management (IAM)](../../../iam/latest/userguide/id-credentials-server-certs.md). You specify the location by
setting a value in one of the following fields (not both):

- `ACMCertificateArn` (In CloudFormation,
this field name is `AcmCertificateArn`. Note the
different capitalization.)

- `IAMCertificateId` (In CloudFormation,
this field name is `IamCertificateId`. Note the different
capitalization.)

All distributions support HTTPS connections from viewers. To require viewers to use
HTTPS only, or to redirect them from HTTP to HTTPS, use
`ViewerProtocolPolicy` in the `CacheBehavior` or
`DefaultCacheBehavior`. To specify how CloudFront should use SSL/TLS to
communicate with your custom origin, use `CustomOriginConfig`.

For more information, see [Using HTTPS with CloudFront](../../../amazoncloudfront/latest/developerguide/using-https.md) and [Using Alternate Domain Names and HTTPS](../../../amazoncloudfront/latest/developerguide/using-https-alternate-domain-names.md) in the
_Amazon CloudFront Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AcmCertificateArn" : String,
  "CloudFrontDefaultCertificate" : Boolean,
  "IamCertificateId" : String,
  "MinimumProtocolVersion" : String,
  "SslSupportMethod" : String
}

```

### YAML

```yaml

  AcmCertificateArn: String
  CloudFrontDefaultCertificate: Boolean
  IamCertificateId: String
  MinimumProtocolVersion: String
  SslSupportMethod: String

```

## Properties

`AcmCertificateArn`

###### Note

In CloudFormation, this field name is `AcmCertificateArn`. Note the
different capitalization.

If the distribution uses `Aliases` (alternate domain names or CNAMEs) and
the SSL/TLS certificate is stored in [AWS Certificate Manager (ACM)](../../../acm/latest/userguide/acm-overview.md), provide the Amazon Resource Name
(ARN) of the ACM certificate. CloudFront only supports ACM certificates in the US East
(N. Virginia) Region ( `us-east-1`).

If you specify an ACM certificate ARN, you must also specify values for
`MinimumProtocolVersion` and `SSLSupportMethod`. (In CloudFormation, the field name is `SslSupportMethod`. Note
the different capitalization.)

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CloudFrontDefaultCertificate`

If the distribution uses the CloudFront domain name such as
`d111111abcdef8.cloudfront.net`, set this field to
`true`.

If the distribution uses `Aliases` (alternate domain names or CNAMEs), omit this field and specify values for the following fields:

- `AcmCertificateArn` or `IamCertificateId` (specify a
value for one, not both)

- `MinimumProtocolVersion`

- `SslSupportMethod`

_Required_: Conditional

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamCertificateId`

###### Note

This field only supports standard distributions. You can't specify this field for multi-tenant distributions. For more information, see [Unsupported features for SaaS Manager for Amazon CloudFront](../../../amazoncloudfront/latest/developerguide/distribution-config-options.md#unsupported-saas) in the _Amazon CloudFront Developer Guide_.

###### Note

In CloudFormation, this field name is `IamCertificateId`. Note the
different capitalization.

If the distribution uses `Aliases` (alternate domain names or CNAMEs) and
the SSL/TLS certificate is stored in [AWS Identity and Access Management (IAM)](../../../iam/latest/userguide/id-credentials-server-certs.md), provide the ID of the IAM certificate.

If you specify an IAM certificate ID, you must also specify values for
`MinimumProtocolVersion` and `SSLSupportMethod`. (In CloudFormation, the field name is `SslSupportMethod`. Note
the different capitalization.)

_Required_: Conditional

_Type_: String

_Minimum_: `0`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinimumProtocolVersion`

If the distribution uses `Aliases` (alternate domain names or CNAMEs),
specify the security policy that you want CloudFront to use for HTTPS connections with
viewers. The security policy determines two settings:

- The minimum SSL/TLS protocol that CloudFront can use to communicate with
viewers.

- The ciphers that CloudFront can use to encrypt the content that it returns to
viewers.

For more information, see [Security Policy](../../../amazoncloudfront/latest/developerguide/distribution-web-values-specify.md#DownloadDistValues-security-policy) and [Supported Protocols and Ciphers Between Viewers and\
CloudFront](../../../amazoncloudfront/latest/developerguide/secure-connections-supported-viewer-protocols-ciphers.md#secure-connections-supported-ciphers) in the _Amazon CloudFront Developer Guide_.

###### Note

On the CloudFront console, this setting is called **Security**
**Policy**.

When you're using SNI only (you set `SSLSupportMethod` to
`sni-only`), you must specify `TLSv1` or higher. (In CloudFormation, the field name is `SslSupportMethod`. Note
the different capitalization.)

If the distribution uses the CloudFront domain name such as
`d111111abcdef8.cloudfront.net` (you set
`CloudFrontDefaultCertificate` to `true`), CloudFront automatically
sets the security policy to `TLSv1` regardless of the value that you set
here.

_Required_: Conditional

_Type_: String

_Allowed values_: `SSLv3 | TLSv1 | TLSv1_2016 | TLSv1.1_2016 | TLSv1.2_2018 | TLSv1.2_2019 | TLSv1.2_2021 | TLSv1.3_2025 | TLSv1.2_2025`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SslSupportMethod`

###### Note

In CloudFormation, this field name is `SslSupportMethod`. Note the
different capitalization.

If the distribution uses `Aliases` (alternate domain names or CNAMEs),
specify which viewers the distribution accepts HTTPS connections from.

- `sni-only` – The distribution accepts HTTPS connections from only
viewers that support [server name\
indication (SNI)](https://en.wikipedia.org/wiki/Server_Name_Indication). This is recommended. Most browsers and clients
support SNI.

- `vip` – The distribution accepts HTTPS connections from all viewers
including those that don't support SNI. This is not recommended, and results in
additional monthly charges from CloudFront.

- `static-ip` \- Do not specify this value unless your distribution
has been enabled for this feature by the CloudFront team. If you have a use case
that requires static IP addresses for a distribution, contact CloudFront through
the [Support Center](https://console.aws.amazon.com/support/home).

If the distribution uses the CloudFront domain name such as
`d111111abcdef8.cloudfront.net`, don't set a value for this field.

_Required_: Conditional

_Type_: String

_Allowed values_: `sni-only | vip | static-ip`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ViewerCertificate](../../../../reference/cloudfront/latest/apireference/api-viewercertificate.md) in the _Amazon CloudFront API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TrustStoreConfig

ViewerMtlsConfig

All content copied from https://docs.aws.amazon.com/.
