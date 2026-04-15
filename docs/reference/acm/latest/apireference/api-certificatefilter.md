---
title: "CertificateFilter"
---

# CertificateFilter

Defines a filter for searching certificates by ARN, X.509 attributes, or ACM metadata.

## Contents

###### Note

In the following list, the required parameters are described first.

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**AcmCertificateMetadataFilter**

Filter by ACM certificate metadata.

Type: [AcmCertificateMetadataFilter](api-acmcertificatemetadatafilter.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

**CertificateArn**

Filter by certificate ARN.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:[\w+=/,.@-]+:acm:[\w+=/,.@-]*:[0-9]+:[\w+=,.@-]+(/[\w+=,.@-]+)*`

Required: No

**X509AttributeFilter**

Filter by X.509 certificate attributes.

Type: [X509AttributeFilter](api-x509attributefilter.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/acm-2015-12-08/certificatefilter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/acm-2015-12-08/certificatefilter.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/acm-2015-12-08/certificatefilter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CertificateDetail

CertificateFilterStatement

All content copied from https://docs.aws.amazon.com/.
