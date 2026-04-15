---
title: "AcmCertificateMetadata"
---

# AcmCertificateMetadata

Contains ACM-specific metadata about a certificate.

## Contents

###### Note

In the following list, the required parameters are described first.

**CreatedAt**

The time at which the certificate was requested.

Type: Timestamp

Required: No

**Exported**

Indicates whether the certificate has been exported.

Type: Boolean

Required: No

**ExportOption**

Indicates whether the certificate can be exported.

Type: String

Valid Values: `ENABLED | DISABLED`

Required: No

**ImportedAt**

The date and time when the certificate was imported. This value exists only when the
certificate type is `IMPORTED`.

Type: Timestamp

Required: No

**InUse**

Indicates whether the certificate is currently in use by an AWS service.

Type: Boolean

Required: No

**IssuedAt**

The time at which the certificate was issued. This value exists only when the certificate
type is `AMAZON_ISSUED`.

Type: Timestamp

Required: No

**ManagedBy**

Identifies the AWS service that manages the certificate issued by ACM.

Type: String

Valid Values: `CLOUDFRONT`

Required: No

**RenewalEligibility**

Specifies whether the certificate is eligible for renewal. At this time, only exported
private certificates can be renewed with the [RenewCertificate](api-renewcertificate.md)
command.

Type: String

Valid Values: `ELIGIBLE | INELIGIBLE`

Required: No

**RenewalStatus**

The renewal status of the certificate.

Type: String

Valid Values: `PENDING_AUTO_RENEWAL | PENDING_VALIDATION | SUCCESS | FAILED`

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

**ValidationMethod**

Specifies the domain validation method.

Type: String

Valid Values: `EMAIL | DNS | HTTP`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/acm-2015-12-08/acmcertificatemetadata.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/acm-2015-12-08/acmcertificatemetadata.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/acm-2015-12-08/acmcertificatemetadata.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Data Types

AcmCertificateMetadataFilter

All content copied from https://docs.aws.amazon.com/.
