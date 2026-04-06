This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CertificateManager::Certificate

The `AWS::CertificateManager::Certificate` resource requests an AWS Certificate Manager (ACM) certificate that you can use to enable
secure connections. For example, you can deploy an ACM certificate to an Elastic Load
Balancer to enable HTTPS support. For more information, see [RequestCertificate](../../../../reference/acm/latest/apireference/api-requestcertificate.md)
in the AWS Certificate Manager API Reference.

###### Important

When you use the `AWS::CertificateManager::Certificate` resource in a
CloudFormation stack, domain validation is handled automatically if all three of the
following are true: The certificate domain is hosted in Amazon Route 53, the domain
resides in your AWS account, and you are using DNS
validation.

However, if the certificate uses email validation, or if the domain is not hosted
in Route 53, then the stack will remain in the `CREATE_IN_PROGRESS`
state. Further stack operations are delayed until you validate the certificate
request, either by acting upon the instructions in the validation email, or by
adding a CNAME record to your DNS configuration. For more information, see [Option 1:\
DNS Validation](https://docs.aws.amazon.com/acm/latest/userguide/dns-validation.html) and [Option 2: Email\
Validation](https://docs.aws.amazon.com/acm/latest/userguide/email-validation.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CertificateManager::Certificate",
  "Properties" : {
      "CertificateAuthorityArn" : String,
      "CertificateExport" : String,
      "CertificateTransparencyLoggingPreference" : String,
      "DomainName" : String,
      "DomainValidationOptions" : [ DomainValidationOption, ... ],
      "KeyAlgorithm" : String,
      "SubjectAlternativeNames" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "ValidationMethod" : String
    }
}

```

### YAML

```yaml

Type: AWS::CertificateManager::Certificate
Properties:
  CertificateAuthorityArn: String
  CertificateExport: String
  CertificateTransparencyLoggingPreference: String
  DomainName: String
  DomainValidationOptions:
    - DomainValidationOption
  KeyAlgorithm: String
  SubjectAlternativeNames:
    - String
  Tags:
    - Tag
  ValidationMethod: String

```

## Properties

`CertificateAuthorityArn`

The Amazon Resource Name (ARN) of the private certificate authority (CA) that will be
used to issue the certificate. If you do not provide an ARN and you are trying to
request a private certificate, ACM will attempt to issue a public certificate. For
more information about private CAs, see the [AWS Private Certificate Authority](https://docs.aws.amazon.com/privateca/latest/userguide/PcaWelcome.html) user guide. The ARN must have
the following form:

`arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012`

_Required_: No

_Type_: String

_Pattern_: `arn:[\w+=/,.@-]+:acm-pca:[\w+=/,.@-]*:[0-9]+:[\w+=,.@-]+(/[\w+=,.@-]+)*`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

`CertificateExport`

You can opt out of allowing export of your certificate by specifying the
`DISABLED` option. Allow export of your certificate by specifying the
`ENABLED` option.

If you do not specify an export preference in a new CloudFormation template, it is the
same as explicitly denying export of your certificate.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

`CertificateTransparencyLoggingPreference`

You can opt out of certificate transparency logging by specifying the
`DISABLED` option. Opt in by specifying `ENABLED`. This setting
doces not apply to private certificates.

If you do not specify a certificate transparency logging preference on a new
CloudFormation template, or if you remove the logging preference from an existing
template, this is the same as explicitly enabling the preference.

Changing the certificate transparency logging preference will update the existing
resource by calling `UpdateCertificateOptions` on the certificate. This
action will not create a new resource.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

`DomainName`

The fully qualified domain name (FQDN), such as www.example.com, with which you want
to secure an ACM certificate. Use an asterisk (\*) to create a wildcard certificate that
protects several sites in the same domain. For example, `*.example.com`
protects `www.example.com`, `site.example.com`, and
`images.example.com.`

_Required_: Yes

_Type_: String

_Pattern_: `(\*\.)?(((?!-)[A-Za-z0-9-]{0,62}[A-Za-z0-9])\.)+((?!-)[A-Za-z0-9-]{1,62}[A-Za-z0-9])`

_Minimum_: `1`

_Maximum_: `253`

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

`DomainValidationOptions`

Domain information that domain name registrars use to verify your identity.

###### Important

In order for a AWS::CertificateManager::Certificate to be provisioned and
validated in CloudFormation automatically, the \`DomainName\` property needs to be
identical to one of the \`DomainName\` property supplied in DomainValidationOptions,
if the ValidationMethod is \*\*DNS\*\*. Failing to keep them like-for-like will result
in failure to create the domain validation records in Route53.

_Required_: No

_Type_: Array of [DomainValidationOption](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-certificatemanager-certificate-domainvalidationoption.html)

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

`KeyAlgorithm`

Specifies the algorithm of the public and private key pair that your certificate uses
to encrypt data. RSA is the default key algorithm for ACM certificates. Elliptic Curve
Digital Signature Algorithm (ECDSA) keys are smaller, offering security comparable to
RSA keys but with greater computing efficiency. However, ECDSA is not supported by all
network clients. Some AWS services may require RSA keys, or only support ECDSA keys of
a particular size, while others allow the use of either RSA and ECDSA keys to ensure
that compatibility is not broken. Check the requirements for the AWS service where you
plan to deploy your certificate. For more information about selecting an algorithm, see
[Key\
algorithms](https://docs.aws.amazon.com/acm/latest/userguide/acm-certificate-characteristics.html#algorithms-term).

###### Note

Algorithms supported for an ACM certificate request include:

- `RSA_2048`

- `EC_prime256v1`

- `EC_secp384r1`

Other listed algorithms are for imported certificates only.

###### Note

When you request a private PKI certificate signed by a CA from AWS Private CA, the
specified signing algorithm family (RSA or ECDSA) must match the algorithm family of
the CA's secret key.

Default: RSA\_2048

_Required_: No

_Type_: String

_Allowed values_: `RSA_1024 | RSA_2048 | RSA_3072 | RSA_4096 | EC_prime256v1 | EC_secp384r1 | EC_secp521r1`

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

`SubjectAlternativeNames`

Additional FQDNs to be included in the Subject Alternative Name extension of the ACM
certificate. For example, you can add www.example.net to a certificate for which the
`DomainName` field is www.example.com if users can reach your site by
using either name.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

`Tags`

Key-value pairs that can identify the certificate.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-certificatemanager-certificate-tag.html)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

`ValidationMethod`

The method you want to use to validate that you own or control the domain associated
with a public certificate. You can [validate with DNS](https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-validate-dns.html) or
[validate with email](https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-validate-email.html). We recommend that you use DNS validation.

If not specified, this property defaults to email validation.

_Required_: No

_Type_: String

_Allowed values_: `EMAIL | DNS | HTTP`

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the certificate's Amazon Resource Name (ARN).

For more information about using the `Ref` function, see [`Ref`](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/intrinsic-function-reference-ref.html).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/intrinsic-function-reference-getatt.html).

- CertificateArn

- CertificateStatus

## Examples

### Declaring an Amazon Certificate Manager Certificate Resource

The following example shows how to declare an
`AWS::CertificateManager::Certificate` resource to create an ACM
certificate.

#### JSON

```json

{
   "Resources":{
      "MyCertificate":{
         "Type":"AWS::CertificateManager::Certificate",
         "Properties":{
            "DomainName":"example.com",
            "ValidationMethod":"DNS"
         }
      }
   }
}
```

#### YAML

```yaml

Resources:
  MyCertificate:
    Type: "AWS::CertificateManager::Certificate"
    Properties:
      DomainName: example.com
      ValidationMethod: DNS
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ExpiryEventsConfiguration

DomainValidationOption
