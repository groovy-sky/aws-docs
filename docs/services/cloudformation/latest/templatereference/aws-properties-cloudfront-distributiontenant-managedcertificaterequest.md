This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::DistributionTenant ManagedCertificateRequest

An object that represents the request for the Amazon CloudFront managed ACM certificate.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateTransparencyLoggingPreference" : String,
  "PrimaryDomainName" : String,
  "ValidationTokenHost" : String
}

```

### YAML

```yaml

  CertificateTransparencyLoggingPreference: String
  PrimaryDomainName: String
  ValidationTokenHost: String

```

## Properties

`CertificateTransparencyLoggingPreference`

You can opt out of certificate transparency logging by specifying the `disabled`
option. Opt in by specifying `enabled`. For more information, see [Certificate\
Transparency Logging](../../../acm/latest/userguide/acm-concepts.md#concept-transparency) in the _AWS Certificate Manager User_
_Guide_.

_Required_: No

_Type_: String

_Allowed values_: `enabled | disabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimaryDomainName`

The primary domain name associated with the CloudFront managed ACM certificate.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValidationTokenHost`

Specify how the HTTP validation token will be served when requesting the CloudFront managed ACM certificate.

- For `cloudfront`, CloudFront will automatically serve the validation token. Choose this mode if you can point the domain's DNS to CloudFront immediately.

- For `self-hosted`, you serve the validation token from your existing infrastructure. Choose this mode when you need to maintain current traffic flow while your certificate is being issued. You can place the validation token at the well-known path on your existing web server, wait for ACM to validate and issue the certificate, and then update your DNS to point to CloudFront.

_Required_: No

_Type_: String

_Allowed values_: `cloudfront | self-hosted`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeoRestrictionCustomization

Parameter

All content copied from https://docs.aws.amazon.com/.
