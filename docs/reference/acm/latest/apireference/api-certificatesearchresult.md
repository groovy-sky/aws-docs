---
title: "CertificateSearchResult"
---

# CertificateSearchResult

Contains information about a certificate returned by the [SearchCertificates](api-searchcertificates.md) action. This structure includes the certificate ARN, X.509 attributes, and ACM metadata.

## Contents

###### Note

In the following list, the required parameters are described first.

**CertificateArn**

The Amazon Resource Name (ARN) of the certificate.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:[\w+=/,.@-]+:acm:[\w+=/,.@-]*:[0-9]+:[\w+=,.@-]+(/[\w+=,.@-]+)*`

Required: No

**CertificateMetadata**

ACM-specific metadata about the certificate.

Type: [CertificateMetadata](api-certificatemetadata.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

**X509Attributes**

X.509 certificate attributes such as subject, issuer, and validity period.

Type: [X509Attributes](api-x509attributes.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/acm-2015-12-08/certificatesearchresult.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/acm-2015-12-08/certificatesearchresult.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/acm-2015-12-08/certificatesearchresult.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CertificateOptions

CertificateSummary

All content copied from https://docs.aws.amazon.com/.
