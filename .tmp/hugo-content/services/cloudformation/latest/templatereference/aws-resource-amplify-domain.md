This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Amplify::Domain

Specifies the AWS::Amplify::Domain resource that enables you to connect a custom domain to your app.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Amplify::Domain",
  "Properties" : {
      "AppId" : String,
      "AutoSubDomainCreationPatterns" : [ String, ... ],
      "AutoSubDomainIAMRole" : String,
      "CertificateSettings" : CertificateSettings,
      "DomainName" : String,
      "EnableAutoSubDomain" : Boolean,
      "SubDomainSettings" : [ SubDomainSetting, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Amplify::Domain
Properties:
  AppId: String
  AutoSubDomainCreationPatterns:
    - String
  AutoSubDomainIAMRole: String
  CertificateSettings:
    CertificateSettings
  DomainName: String
  EnableAutoSubDomain: Boolean
  SubDomainSettings:
    - SubDomainSetting

```

## Properties

`AppId`

The unique ID for an Amplify app.

_Required_: Yes

_Type_: String

_Pattern_: `d[a-z0-9]+`

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AutoSubDomainCreationPatterns`

Sets the branch patterns for automatic subdomain creation.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutoSubDomainIAMRole`

The required AWS Identity and Access Management (IAMlong) service role for the Amazon Resource Name
(ARN) for automatically creating subdomains.

_Required_: No

_Type_: String

_Pattern_: `^$|^arn:.+:iam::\d{12}:role.+`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CertificateSettings`

The type of SSL/TLS certificate to use for your custom domain. If you don't specify a
certificate type, Amplify uses the default certificate that it provisions and manages
for you.

_Required_: No

_Type_: [CertificateSettings](aws-properties-amplify-domain-certificatesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainName`

The domain name for the domain association.

_Required_: Yes

_Type_: String

_Pattern_: `^(((?!-)[A-Za-z0-9-]{0,62}[A-Za-z0-9])\.)+((?!-)[A-Za-z0-9-]{1,62}[A-Za-z0-9])(\.)?$`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnableAutoSubDomain`

Enables the automated creation of subdomains for branches.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubDomainSettings`

The setting for the subdomain.

_Required_: Yes

_Type_: Array of [SubDomainSetting](aws-properties-amplify-domain-subdomainsetting.md)

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

ARN for the Domain Association.

`AutoSubDomainCreationPatterns`

Branch patterns for the automatically created subdomain.

`AutoSubDomainIAMRole`

The IAM service role for the subdomain.

`CertificateRecord`

DNS Record for certificate verification.

`DomainName`

Name of the domain.

`DomainStatus`

Status for the Domain Association.

`EnableAutoSubDomain`

Specifies whether the automated creation of subdomains for branches is enabled.

`StatusReason`

Reason for the current status of the domain.

`UpdateStatus`

The status of the domain update operation that is currently in progress. The following
list describes the valid update states.

REQUESTING\_CERTIFICATE

The certificate is in the process of being updated.

PENDING\_VERIFICATION

Indicates that an Amplify managed certificate is in the
process of being verified. This occurs during the creation of a custom
domain or when a custom domain is updated to use a managed
certificate.

IMPORTING\_CUSTOM\_CERTIFICATE

Indicates that an Amplify custom certificate is in the
process of being imported. This occurs during the creation of a custom
domain or when a custom domain is updated to use a custom
certificate.

PENDING\_DEPLOYMENT

Indicates that the subdomain or certificate changes are being
propagated.

AWAITING\_APP\_CNAME

Amplify is waiting for CNAME records corresponding to
subdomains to be propagated. If your custom domain is on Route 53,
Amplify handles this for you automatically. For more
information about custom domains, see [Setting up custom\
domains](../../../amplify/latest/userguide/custom-domains.md) in the _Amplify Hosting User_
_Guide_.

UPDATE\_COMPLETE

The certificate has been associated with a domain.

UPDATE\_FAILED

The certificate has failed to be provisioned or associated, and there is
no existing active certificate to roll back to.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Certificate

All content copied from https://docs.aws.amazon.com/.
