This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::Certificate

The `AWS::DMS::Certificate` resource creates an Secure Sockets Layer (SSL) certificate that
encrypts connections between AWS DMS endpoints and the replication instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DMS::Certificate",
  "Properties" : {
      "CertificateIdentifier" : String,
      "CertificatePem" : String,
      "CertificateWallet" : String
    }
}

```

### YAML

```yaml

Type: AWS::DMS::Certificate
Properties:
  CertificateIdentifier: String
  CertificatePem: String
  CertificateWallet: String

```

## Properties

`CertificateIdentifier`

A customer-assigned name for the certificate. Identifiers must begin with a letter and
must contain only ASCII letters, digits, and hyphens. They can't end with a hyphen or
contain two consecutive hyphens.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CertificatePem`

The contents of a `.pem` file, which contains an X.509 certificate.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CertificateWallet`

The location of an imported Oracle Wallet certificate for use with SSL. An example
is: `filebase64("${path.root}/rds-ca-2019-root.sso")`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the certificate.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

## Examples

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Certificate test",
    "Resources": {
        "BasicCertificate": {
            "Type": "AWS::DMS::Certificate",
            "Properties": {
                "CertificatePem": "-----BEGIN CERTIFICATE-----\n MIID/DCCAuSgAwIBAgIBUDANBgkqhkiG9w0BAQsFADCBijELMAkGA1UEBhMCVVMx...mqfEEuC7uUoPofXdBp2ObQ==\n -----END CERTIFICATE-----\n"
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: "Certificate test"
Resources:
  BasicCertificate:
    Type: AWS::DMS::Certificate
    Properties:
      CertificatePem: |-
          -----BEGIN CERTIFICATE-----
           MIID/DCCAuSgAwIBAgABCDEFgkqhkiG9w0BAQsFADCBijEXAMPLE1UEBhMCVVMx...mqfEEuC7uUoPofXdBp2ObQ==
           -----END CERTIFICATE-----
```

## See also

- [ImportCertificate](../../../../reference/dms/latest/apireference/api-importcertificate.md)
in the _AWS Database Migration Service API Reference_

- [Managing AWS resources as a single unit with AWS CloudFormation stacks](../userguide/stacks.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Database Migration Service

AWS::DMS::DataMigration

All content copied from https://docs.aws.amazon.com/.
