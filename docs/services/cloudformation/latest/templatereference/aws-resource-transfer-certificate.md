This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::Certificate

Imports the signing and encryption certificates that you need to create local (AS2)
profiles and partner profiles.

You can import both the certificate and its chain in the `Certificate`
parameter.

After importing a certificate, AWS Transfer Family automatically creates a Amazon CloudWatch metric called `DaysUntilExpiry` that tracks the number of
days until the certificate expires. The metric is based on the `InactiveDate`
parameter and is published daily in the `AWS/Transfer` namespace.

###### Important

It can take up to a full day after importing a certificate for Transfer Family
to emit the `DaysUntilExpiry` metric to your account.

###### Note

If you use the `Certificate` parameter to upload both the certificate
and its chain, don't use the `CertificateChain` parameter.

**CloudWatch monitoring**

The `DaysUntilExpiry` metric includes the following specifications:

- **Units:** Count (days)

- **Dimensions:** `CertificateId` (always present), `Description` (if
provided during certificate import)

- **Statistics:** Minimum, Maximum, Average

- **Frequency:** Published daily

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Transfer::Certificate",
  "Properties" : {
      "ActiveDate" : String,
      "Certificate" : String,
      "CertificateChain" : String,
      "Description" : String,
      "InactiveDate" : String,
      "PrivateKey" : String,
      "Tags" : [ Tag, ... ],
      "Usage" : String
    }
}

```

### YAML

```yaml

Type: AWS::Transfer::Certificate
Properties:
  ActiveDate: String
  Certificate: String
  CertificateChain: String
  Description: String
  InactiveDate: String
  PrivateKey: String
  Tags:
    - Tag
  Usage: String

```

## Properties

`ActiveDate`

An optional date that specifies when the certificate becomes active. If you do not specify a value, `ActiveDate` takes the same value as
`NotBeforeDate`, which is specified by the CA.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Certificate`

The file name for the certificate.

_Required_: Yes

_Type_: String

_Pattern_: `^[\t\n\r\u0020-\u00FF]+$`

_Minimum_: `1`

_Maximum_: `16384`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CertificateChain`

The list of certificates that make up the chain for the certificate.

_Required_: No

_Type_: String

_Pattern_: `^[\t\n\r\u0020-\u00FF]+$`

_Minimum_: `1`

_Maximum_: `2097152`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The name or description that's used to identity the certificate.

_Required_: No

_Type_: String

_Pattern_: `^[\u0021-\u007E]+$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InactiveDate`

An optional date that specifies when the certificate becomes inactive. If you do not specify a value, `InactiveDate` takes the same value as
`NotAfterDate`, which is specified by the CA.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivateKey`

The file that contains the private key for the certificate that's being
imported.

_Required_: No

_Type_: String

_Pattern_: `^[\t\n\r\u0020-\u00FF]+$`

_Minimum_: `1`

_Maximum_: `16384`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Key-value pairs that can be used to group and search for certificates.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-transfer-certificate-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Usage`

Specifies how this certificate is used. It can be used in the following ways:

- `SIGNING`: For signing AS2 messages

- `ENCRYPTION`: For encrypting AS2 messages

- `TLS`: For securing AS2 communications sent over HTTPS

_Required_: Yes

_Type_: String

_Allowed values_: `SIGNING | ENCRYPTION | TLS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `certificateId` , such as
`cert-1c698edce1654f869` .

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The unique Amazon Resource Name (ARN) for the certificate.

`CertificateId`

An array of identifiers for the imported certificates. You use this identifier for
working with profiles and partner profiles.

`NotAfterDate`

The final date that the certificate is valid.

`NotBeforeDate`

The earliest date that the certificate is valid.

`Serial`

The serial number for the certificate.

`Status`

The certificate can be either `ACTIVE` , `PENDING_ROTATION` ,
or `INACTIVE` . `PENDING_ROTATION` means that this certificate
will replace the current certificate when it expires.

`Type`

If a private key has been specified for the certificate, its type is
`CERTIFICATE_WITH_PRIVATE_KEY` . If there is no private key, the type is
`CERTIFICATE` .

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
