This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RolesAnywhere::TrustAnchor SourceData

A union object representing the data field of the TrustAnchor depending on its type

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AcmPcaArn" : String,
  "X509CertificateData" : String
}

```

### YAML

```yaml

  AcmPcaArn: String
  X509CertificateData: String

```

## Properties

`AcmPcaArn`

The root certificate of the AWS Private Certificate Authority specified by this ARN is used in trust
validation for temporary credential requests. Included for trust anchors of type `AWS_ACM_PCA`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`X509CertificateData`

The PEM-encoded data for the certificate anchor. Included for trust anchors of type `CERTIFICATE_BUNDLE`.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `8000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Source

Tag

All content copied from https://docs.aws.amazon.com/.
