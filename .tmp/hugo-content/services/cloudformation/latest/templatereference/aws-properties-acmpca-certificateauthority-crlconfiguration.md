This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ACMPCA::CertificateAuthority CrlConfiguration

Contains configuration information for a certificate revocation list (CRL). Your
private certificate authority (CA) creates base CRLs. Delta CRLs are not supported. You
can enable CRLs for your new or an existing private CA by setting the **Enabled** parameter to `true`. Your private CA
writes CRLs to an S3 bucket that you specify in the **S3BucketName** parameter. You can hide the name of your bucket by
specifying a value for the **CustomCname** parameter. Your
private CA by default copies the CNAME or the S3 bucket name to the **CRL**
**Distribution Points** extension of each certificate it issues. If you want to configure
this default behavior to be something different, you can set the **CrlDistributionPointExtensionConfiguration**
parameter. Your S3 bucket policy must give write permission to AWS Private CA.

AWS Private CA assets that are stored in Amazon S3 can be protected with
encryption. For more information, see [Encrypting Your\
CRLs](../../../privateca/latest/userguide/pcacreateca.md#crl-encryption).

Your private CA uses the value in the **ExpirationInDays** parameter to calculate the **nextUpdate** field in the CRL. The CRL is refreshed prior to a
certificate's expiration date or when a certificate is revoked. When a certificate is
revoked, it appears in the CRL until the certificate expires, and then in one additional
CRL after expiration, and it always appears in the audit report.

A CRL is typically updated approximately 30 minutes after a certificate is revoked. If
for any reason a CRL update fails, AWS Private CA makes further attempts
every 15 minutes.

CRLs contain the following fields:

- **Version**: The current version number defined in
RFC 5280 is V2. The integer value is 0x1.

- **Signature Algorithm**: The name of the algorithm
used to sign the CRL.

- **Issuer**: The X.500 distinguished name of your
private CA that issued the CRL.

- **Last Update**: The issue date and time of this
CRL.

- **Next Update**: The day and time by which the next
CRL will be issued.

- **Revoked Certificates**: List of revoked
certificates. Each list item contains the following information.

- **Serial Number**: The serial number, in
hexadecimal format, of the revoked certificate.

- **Revocation Date**: Date and time the
certificate was revoked.

- **CRL Entry Extensions**: Optional
extensions for the CRL entry.

- **X509v3 CRL Reason Code**: Reason
the certificate was revoked.

- **CRL Extensions**: Optional extensions for the
CRL.

- **X509v3 Authority Key Identifier**:
Identifies the public key associated with the private key used to sign
the certificate.

- **X509v3 CRL Number:**: Decimal sequence
number for the CRL.

- **Signature Algorithm**: Algorithm used by your
private CA to sign the CRL.

- **Signature Value**: Signature computed over the
CRL.

Certificate revocation lists created by AWS Private CA are DER-encoded.
You can use the following OpenSSL command to list a CRL.

`openssl crl -inform DER -text -in crl_path -noout`

For more information, see [Planning a certificate revocation\
list (CRL)](../../../privateca/latest/userguide/crl-planning.md) in the _AWS Private Certificate Authority User Guide_

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CrlDistributionPointExtensionConfiguration" : CrlDistributionPointExtensionConfiguration,
  "CrlType" : String,
  "CustomCname" : String,
  "CustomPath" : String,
  "Enabled" : Boolean,
  "ExpirationInDays" : Integer,
  "S3BucketName" : String,
  "S3ObjectAcl" : String
}

```

### YAML

```yaml

  CrlDistributionPointExtensionConfiguration:
    CrlDistributionPointExtensionConfiguration
  CrlType: String
  CustomCname: String
  CustomPath: String
  Enabled: Boolean
  ExpirationInDays: Integer
  S3BucketName: String
  S3ObjectAcl: String

```

## Properties

`CrlDistributionPointExtensionConfiguration`

Configures the default behavior of the CRL Distribution Point extension for certificates issued by your CA. If this field is not provided, then the CRL Distribution Point extension will be present and contain the default CRL URL.

_Required_: No

_Type_: [CrlDistributionPointExtensionConfiguration](aws-properties-acmpca-certificateauthority-crldistributionpointextensionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CrlType`

Specifies the type of CRL. This setting determines the maximum number of certificates that the certificate authority can issue and revoke. For more information, see [AWS Private CA quotas](../../../../general/latest/gr/pca.md#limits_pca).

- `COMPLETE` \- The default setting. AWS Private CA maintains a single CRL file for all unexpired certificates issued by a CA that have been revoked for any reason. Each certificate that AWS Private CA issues is bound to a specific CRL through the CRL distribution point (CDP) defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280).

- `PARTITIONED` \- Compared to complete CRLs, partitioned CRLs dramatically increase the number of certificates your private CA can issue.

###### Important

When using partitioned CRLs, you must validate that the CRL's associated
issuing distribution point (IDP) URI matches the certiﬁcate's CDP URI to ensure
the right CRL has been fetched. AWS Private CA marks the IDP extension as critical,
which your client must be able to process.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomCname`

Name inserted into the certificate **CRL Distribution**
**Points** extension that enables the use of an alias for the CRL
distribution point. Use this value if you don't want the name of your S3 bucket to be
public.

###### Note

The content of a Canonical Name (CNAME) record must conform to [RFC2396](https://www.ietf.org/rfc/rfc2396.txt) restrictions on the
use of special characters in URIs. Additionally, the value of the CNAME must not
include a protocol prefix such as "http://" or "https://".

_Required_: No

_Type_: String

_Pattern_: `[-a-zA-Z0-9;/?:@&=+$,%_.!~*()']*`

_Minimum_: `0`

_Maximum_: `253`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomPath`

Designates a custom file path in S3 for CRL(s). For example, `http://<CustomName>/<CustomPath>/<CrlPartition_GUID>.crl`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Boolean value that specifies whether certificate revocation lists (CRLs) are enabled.
You can use this value to enable certificate revocation for a new CA when you call the
`CreateCertificateAuthority` operation or for an existing CA when you
call the `UpdateCertificateAuthority` operation.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExpirationInDays`

Validity period of the CRL in days.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `5000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BucketName`

Name of the S3 bucket that contains the CRL. If you do not provide a value for the
**CustomCname** argument, the name of your S3 bucket
is placed into the **CRL Distribution Points** extension of
the issued certificate. You can change the name of your bucket by calling the [UpdateCertificateAuthority](../../../../reference/privateca/latest/apireference/api-updatecertificateauthority.md) operation. You must specify a [bucket\
policy](../../../privateca/latest/userguide/crl-planning.md#s3-policies) that allows AWS Private CA to write the CRL to your bucket.

###### Note

The `S3BucketName` parameter must conform to the [S3\
bucket naming rules](../../../s3/latest/userguide/bucketnamingrules.md).

_Required_: No

_Type_: String

_Pattern_: `[-a-zA-Z0-9._/]+`

_Minimum_: `3`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3ObjectAcl`

Determines whether the CRL will be publicly readable or privately held in the CRL
Amazon S3 bucket. If you choose PUBLIC\_READ, the CRL will be accessible over the public
internet. If you choose BUCKET\_OWNER\_FULL\_CONTROL, only the owner of the CRL S3 bucket
can access the CRL, and your PKI clients may need an alternative method of
access.

If no value is specified, the default is PUBLIC\_READ.

_Note:_ This default can cause CA creation to fail in some
circumstances. If you have have enabled the Block Public Access (BPA) feature in your S3
account, then you must specify the value of this parameter as
`BUCKET_OWNER_FULL_CONTROL`, and not doing so results in an error. If you
have disabled BPA in S3, then you can specify either
`BUCKET_OWNER_FULL_CONTROL` or `PUBLIC_READ` as the
value.

For more information, see [Blocking public access to\
the S3 bucket](../../../privateca/latest/userguide/pcacreateca.md#s3-bpa).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccessMethod

CrlDistributionPointExtensionConfiguration

All content copied from https://docs.aws.amazon.com/.
