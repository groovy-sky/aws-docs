This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Amplify::Domain CertificateSettings

The type of SSL/TLS certificate to use for your custom domain. If a certificate type
isn't specified, Amplify uses the default `AMPLIFY_MANAGED`
certificate.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateType" : String,
  "CustomCertificateArn" : String
}

```

### YAML

```yaml

  CertificateType: String
  CustomCertificateArn: String

```

## Properties

`CertificateType`

The certificate type.

Specify `AMPLIFY_MANAGED` to use the default certificate that Amplify
provisions for you.

Specify `CUSTOM` to use your own certificate that you have already added to
AWS Certificate Manager in your AWS account. Make sure you request (or
import) the certificate in the US East (N. Virginia) Region (us-east-1). For more
information about using ACM, see [Importing certificates into\
AWS Certificate Manager](../../../acm/latest/userguide/import-certificate.md) in the _ACM User_
_guide_.

_Required_: No

_Type_: String

_Allowed values_: `AMPLIFY_MANAGED | CUSTOM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomCertificateArn`

The Amazon resource name (ARN) for the custom certificate that you have already added
to AWS Certificate Manager in your AWS account.

This field is required only when the certificate type is `CUSTOM`.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws:acm:[a-z0-9-]+:\d{12}:certificate\/.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Certificate

SubDomainSetting
