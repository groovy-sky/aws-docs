# CertificateDetail

Contains metadata about an ACM certificate. This structure is returned in the
response to a [DescribeCertificate](api-describecertificate.md) request.

## Contents

###### Note

In the following list, the required parameters are described first.

**CertificateArn**

The Amazon Resource Name (ARN) of the certificate. For more information about ARNs,
see [Amazon Resource Names\
(ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the _AWS General Reference_.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:[\w+=/,.@-]+:acm:[\w+=/,.@-]*:[0-9]+:[\w+=,.@-]+(/[\w+=,.@-]+)*`

Required: No

**CertificateAuthorityArn**

The Amazon Resource Name (ARN) of the private certificate authority (CA) that issued
the certificate. This has the following format:

`arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012`

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:[\w+=/,.@-]+:acm:[\w+=/,.@-]*:[0-9]+:[\w+=,.@-]+(/[\w+=,.@-]+)*`

Required: No

**CreatedAt**

The time at which the certificate was requested.

Type: Timestamp

Required: No

**DomainName**

The fully qualified domain name for the certificate, such as www.example.com or
example.com.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 253.

Pattern: `(\*\.)?(((?!-)[A-Za-z0-9-]{0,62}[A-Za-z0-9])\.)+((?!-)[A-Za-z0-9-]{1,62}[A-Za-z0-9])`

Required: No

**DomainValidationOptions**

Contains information about the initial validation of each domain name that occurs as a
result of the [RequestCertificate](api-requestcertificate.md) request. This field exists only when
the certificate type is `AMAZON_ISSUED`.

Type: Array of [DomainValidation](api-domainvalidation.md) objects

Array Members: Minimum number of 1 item. Maximum number of 1000 items.

Required: No

**ExtendedKeyUsages**

Contains a list of Extended Key Usage X.509 v3 extension objects. Each object specifies a
purpose for which the certificate public key can be used and consists of a name and an object
identifier (OID).

Type: Array of [ExtendedKeyUsage](api-extendedkeyusage.md) objects

Required: No

**FailureReason**

The reason the certificate request failed. This value exists only when the certificate
status is `FAILED`. For more information, see [Certificate Request\
Failed](../../../../services/acm/latest/userguide/troubleshooting.md#troubleshooting-failed) in the _AWS Certificate Manager User Guide_.

Type: String

Valid Values: `NO_AVAILABLE_CONTACTS | ADDITIONAL_VERIFICATION_REQUIRED | DOMAIN_NOT_ALLOWED | INVALID_PUBLIC_DOMAIN | DOMAIN_VALIDATION_DENIED | CAA_ERROR | PCA_LIMIT_EXCEEDED | PCA_INVALID_ARN | PCA_INVALID_STATE | PCA_REQUEST_FAILED | PCA_NAME_CONSTRAINTS_VALIDATION | PCA_RESOURCE_NOT_FOUND | PCA_INVALID_ARGS | PCA_INVALID_DURATION | PCA_ACCESS_DENIED | SLR_NOT_FOUND | OTHER`

Required: No

**ImportedAt**

The date and time when the certificate was imported. This value exists only when the
certificate type is `IMPORTED`.

Type: Timestamp

Required: No

**InUseBy**

A list of ARNs for the AWS resources that are using the certificate. A certificate
can be used by multiple AWS resources.

Type: Array of strings

Required: No

**IssuedAt**

The time at which the certificate was issued. This value exists only when the certificate
type is `AMAZON_ISSUED`.

Type: Timestamp

Required: No

**Issuer**

The name of the certificate authority that issued and signed the certificate.

Type: String

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

Type: Array of [KeyUsage](api-keyusage.md) objects

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

**Options**

Value that specifies whether to add the certificate to a transparency log. Certificate
transparency makes it possible to detect SSL certificates that have been mistakenly or
maliciously issued. A browser might respond to certificate that has not been logged by
showing an error message. The logs are cryptographically secure.

Type: [CertificateOptions](api-certificateoptions.md) object

Required: No

**RenewalEligibility**

Specifies whether the certificate is eligible for renewal. At this time, only exported
private certificates can be renewed with the [RenewCertificate](api-renewcertificate.md)
command.

Type: String

Valid Values: `ELIGIBLE | INELIGIBLE`

Required: No

**RenewalSummary**

Contains information about the status of ACM's [managed renewal](https://docs.aws.amazon.com/acm/latest/userguide/acm-renewal.html) for the certificate. This
field exists only when the certificate type is `AMAZON_ISSUED`.

Type: [RenewalSummary](api-renewalsummary.md) object

Required: No

**RevocationReason**

The reason the certificate was revoked. This value exists only when the certificate
status is `REVOKED`.

Type: String

Valid Values: `UNSPECIFIED | KEY_COMPROMISE | CA_COMPROMISE | AFFILIATION_CHANGED | SUPERCEDED | SUPERSEDED | CESSATION_OF_OPERATION | CERTIFICATE_HOLD | REMOVE_FROM_CRL | PRIVILEGE_WITHDRAWN | A_A_COMPROMISE`

Required: No

**RevokedAt**

The time at which the certificate was revoked. This value exists only when the certificate
status is `REVOKED`.

Type: Timestamp

Required: No

**Serial**

The serial number of the certificate.

Type: String

Required: No

**SignatureAlgorithm**

The algorithm that was used to sign the certificate.

Type: String

Required: No

**Status**

The status of the certificate.

A certificate enters status PENDING\_VALIDATION upon being requested, unless it fails for
any of the reasons given in the troubleshooting topic [Certificate request fails](https://docs.aws.amazon.com/acm/latest/userguide/troubleshooting-failed.html). ACM makes
repeated attempts to validate a certificate for 72 hours and then times out. If a certificate
shows status FAILED or VALIDATION\_TIMED\_OUT, delete the request, correct the issue with [DNS validation](../../../../services/acm/latest/userguide/dns-validation.md) or [Email validation](../../../../services/acm/latest/userguide/email-validation.md), and
try again. If validation succeeds, the certificate enters status ISSUED.

Type: String

Valid Values: `PENDING_VALIDATION | ISSUED | INACTIVE | EXPIRED | VALIDATION_TIMED_OUT | REVOKED | FAILED`

Required: No

**Subject**

The name of the entity that is associated with the public key contained in the
certificate.

Type: String

Required: No

**SubjectAlternativeNames**

One or more domain names (subject alternative names)
included in the certificate. This
list contains the domain names that are bound to the public key that is contained in the
certificate. The subject alternative names include the canonical domain name (CN) of the
certificate and additional domain names that can be used to connect to the website.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 100 items.

Length Constraints: Minimum length of 1. Maximum length of 253.

Pattern: `(\*\.)?(((?!-)[A-Za-z0-9-]{0,62}[A-Za-z0-9])\.)+((?!-)[A-Za-z0-9-]{1,62}[A-Za-z0-9])`

Required: No

**Type**

The source of the certificate. For certificates provided by ACM, this value is
`AMAZON_ISSUED`. For certificates that you imported with [ImportCertificate](api-importcertificate.md), this value is `IMPORTED`. ACM does not provide
[managed renewal](https://docs.aws.amazon.com/acm/latest/userguide/acm-renewal.html) for
imported certificates. For more information about the differences between certificates that
you import and those that ACM provides, see [Importing Certificates](../../../../services/acm/latest/userguide/import-certificate.md) in the
_AWS Certificate Manager User Guide_.

Type: String

Valid Values: `IMPORTED | AMAZON_ISSUED | PRIVATE`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/acm-2015-12-08/CertificateDetail)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/acm-2015-12-08/CertificateDetail)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/acm-2015-12-08/CertificateDetail)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Data Types

CertificateOptions
