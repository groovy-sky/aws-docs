---
title: "AcmCertificateMetadataFilter"
---

# AcmCertificateMetadataFilter

Filters certificates by ACM metadata.

## Contents

###### Note

In the following list, the required parameters are described first.

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**Exported**

Filter by whether the certificate has been exported.

Type: Boolean

Required: No

**ExportOption**

Filter by certificate export option.

Type: String

Valid Values: `ENABLED | DISABLED`

Required: No

**InUse**

Filter by whether the certificate is in use.

Type: Boolean

Required: No

**ManagedBy**

Filter by the entity that manages the certificate.

Type: String

Valid Values: `CLOUDFRONT`

Required: No

**RenewalStatus**

Filter by certificate renewal status.

Type: String

Valid Values: `PENDING_AUTO_RENEWAL | PENDING_VALIDATION | SUCCESS | FAILED`

Required: No

**Status**

Filter by certificate status.

Type: String

Valid Values: `PENDING_VALIDATION | ISSUED | INACTIVE | EXPIRED | VALIDATION_TIMED_OUT | REVOKED | FAILED`

Required: No

**Type**

Filter by certificate type.

Type: String

Valid Values: `IMPORTED | AMAZON_ISSUED | PRIVATE`

Required: No

**ValidationMethod**

Filter by validation method.

Type: String

Valid Values: `EMAIL | DNS | HTTP`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/acm-2015-12-08/AcmCertificateMetadataFilter)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/acm-2015-12-08/AcmCertificateMetadataFilter)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/acm-2015-12-08/AcmCertificateMetadataFilter)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AcmCertificateMetadata

CertificateDetail

All content copied from https://docs.aws.amazon.com/.
