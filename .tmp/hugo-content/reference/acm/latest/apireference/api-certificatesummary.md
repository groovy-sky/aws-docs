# CertificateSummary

This structure is returned in the response object of [ListCertificates](api-listcertificates.md) action.

## Contents

###### Note

In the following list, the required parameters are described first.

**CertificateArn**

Amazon Resource Name (ARN) of the certificate. This is of the form:

`arn:aws:acm:region:123456789012:certificate/12345678-1234-1234-1234-123456789012`

For more information about ARNs, see [Amazon Resource Names (ARNs)](../../../../general/latest/gr/aws-arns-and-namespaces.md).

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:[\w+=/,.@-]+:acm:[\w+=/,.@-]*:[0-9]+:[\w+=,.@-]+(/[\w+=,.@-]+)*`

Required: No

**CreatedAt**

The time at which the certificate was requested.

Type: Timestamp

Required: No

**DomainName**

Fully qualified domain name (FQDN), such as www.example.com or example.com, for the
certificate.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 253.

Pattern: `(\*\.)?(((?!-)[A-Za-z0-9-]{0,62}[A-Za-z0-9])\.)+((?!-)[A-Za-z0-9-]{1,62}[A-Za-z0-9])`

Required: No

**Exported**

Indicates whether the certificate has been exported. This value exists only when the
certificate type is `PRIVATE`.

Type: Boolean

Required: No

**ExportOption**

Indicates if export is enabled for the certificate.

Type: String

Valid Values: `ENABLED | DISABLED`

Required: No

**ExtendedKeyUsages**

Contains a list of Extended Key Usage X.509 v3 extension objects. Each object specifies a
purpose for which the certificate public key can be used and consists of a name and an object
identifier (OID).

Type: Array of strings

Valid Values: `TLS_WEB_SERVER_AUTHENTICATION | TLS_WEB_CLIENT_AUTHENTICATION | CODE_SIGNING | EMAIL_PROTECTION | TIME_STAMPING | OCSP_SIGNING | IPSEC_END_SYSTEM | IPSEC_TUNNEL | IPSEC_USER | ANY | NONE | CUSTOM`

Required: No

**HasAdditionalSubjectAlternativeNames**

When called by [ListCertificates](api-listcertificates.md), indicates
whether the full list of subject alternative names has been included in the response. If
false, the response includes all of the subject alternative names included in the
certificate. If true, the response only includes the first 100 subject alternative names
included in the certificate. To display the full list of subject alternative names, use
[DescribeCertificate](api-describecertificate.md).

Type: Boolean

Required: No

**ImportedAt**

The date and time when the certificate was imported. This value exists only when the
certificate type is `IMPORTED`.

Type: Timestamp

Required: No

**InUse**

Indicates whether the certificate is currently in use by any AWS resources.

Type: Boolean

Required: No

**IssuedAt**

The time at which the certificate was issued. This value exists only when the certificate
type is `AMAZON_ISSUED`.

Type: Timestamp

Required: No

**KeyAlgorithm**

The algorithm that was used to generate the public-private key pair.

Type: String

Valid Values: `RSA_1024 | RSA_2048 | RSA_3072 | RSA_4096 | EC_prime256v1 | EC_secp384r1 | EC_secp521r1`

Required: No

**KeyUsages**

A list of Key Usage X.509 v3 extension objects. Each object is a string value that
identifies the purpose of the public key contained in the certificate. Possible extension
values include DIGITAL\_SIGNATURE, KEY\_ENCHIPHERMENT, NON\_REPUDIATION, and more.

Type: Array of strings

Valid Values: `DIGITAL_SIGNATURE | NON_REPUDIATION | KEY_ENCIPHERMENT | DATA_ENCIPHERMENT | KEY_AGREEMENT | CERTIFICATE_SIGNING | CRL_SIGNING | ENCIPHER_ONLY | DECIPHER_ONLY | ANY | CUSTOM`

Required: No

**ManagedBy**

Identifies the AWS service that manages the certificate issued by ACM.

Type: String

Valid Values: `CLOUDFRONT`

Required: No

**NotAfter**

The time after which the certificate is not valid.

Type: Timestamp

Required: No

**NotBefore**

The time before which the certificate is not valid.

Type: Timestamp

Required: No

**RenewalEligibility**

Specifies whether the certificate is eligible for renewal. At this time, only exported
private certificates can be renewed with the [RenewCertificate](api-renewcertificate.md)
command.

Type: String

Valid Values: `ELIGIBLE | INELIGIBLE`

Required: No

**RevokedAt**

The time at which the certificate was revoked. This value exists only when the certificate
status is `REVOKED`.

Type: Timestamp

Required: No

**Status**

The status of the certificate.

A certificate enters status PENDING\_VALIDATION upon being requested, unless it fails for
any of the reasons given in the troubleshooting topic [Certificate request fails](../../../../services/acm/latest/userguide/troubleshooting-failed.md). ACM makes
repeated attempts to validate a certificate for 72 hours and then times out. If a certificate
shows status FAILED or VALIDATION\_TIMED\_OUT, delete the request, correct the issue with [DNS validation](../../../../services/acm/latest/userguide/dns-validation.md) or [Email validation](../../../../services/acm/latest/userguide/email-validation.md), and
try again. If validation succeeds, the certificate enters status ISSUED.

Type: String

Valid Values: `PENDING_VALIDATION | ISSUED | INACTIVE | EXPIRED | VALIDATION_TIMED_OUT | REVOKED | FAILED`

Required: No

**SubjectAlternativeNameSummaries**

One or more domain names (subject alternative names)
included in the certificate. This
list contains the domain names that are bound to the public key that is contained in the
certificate. The subject alternative names include the canonical domain name (CN) of the
certificate and additional domain names that can be used to connect to the website.

When called by [ListCertificates](api-listcertificates.md), this
parameter will only return the first 100 subject alternative names included in the
certificate. To display the full list of subject alternative names, use [DescribeCertificate](api-describecertificate.md).

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 100 items.

Length Constraints: Minimum length of 1. Maximum length of 253.

Pattern: `(\*\.)?(((?!-)[A-Za-z0-9-]{0,62}[A-Za-z0-9])\.)+((?!-)[A-Za-z0-9-]{1,62}[A-Za-z0-9])`

Required: No

**Type**

The source of the certificate. For certificates provided by ACM, this value is
`AMAZON_ISSUED`. For certificates that you imported with [ImportCertificate](api-importcertificate.md), this value is `IMPORTED`. ACM does not provide
[managed renewal](../../../../services/acm/latest/userguide/acm-renewal.md) for
imported certificates. For more information about the differences between certificates that
you import and those that ACM provides, see [Importing Certificates](../../../../services/acm/latest/userguide/import-certificate.md) in the
_AWS Certificate Manager User Guide_.

Type: String

Valid Values: `IMPORTED | AMAZON_ISSUED | PRIVATE`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/acm-2015-12-08/certificatesummary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/acm-2015-12-08/certificatesummary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/acm-2015-12-08/certificatesummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CertificateSearchResult

CommonNameFilter

All content copied from https://docs.aws.amazon.com/.
