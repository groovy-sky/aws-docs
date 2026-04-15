---
title: "X509AttributeFilter"
---

# X509AttributeFilter

Filters certificates by X.509 attributes.

## Contents

###### Note

In the following list, the required parameters are described first.

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**ExtendedKeyUsage**

Filter by extended key usage.

Type: String

Valid Values: `TLS_WEB_SERVER_AUTHENTICATION | TLS_WEB_CLIENT_AUTHENTICATION | CODE_SIGNING | EMAIL_PROTECTION | TIME_STAMPING | OCSP_SIGNING | IPSEC_END_SYSTEM | IPSEC_TUNNEL | IPSEC_USER | ANY | NONE | CUSTOM`

Required: No

**KeyAlgorithm**

Filter by key algorithm.

Type: String

Valid Values: `RSA_1024 | RSA_2048 | RSA_3072 | RSA_4096 | EC_prime256v1 | EC_secp384r1 | EC_secp521r1`

Required: No

**KeyUsage**

Filter by key usage.

Type: String

Valid Values: `DIGITAL_SIGNATURE | NON_REPUDIATION | KEY_ENCIPHERMENT | DATA_ENCIPHERMENT | KEY_AGREEMENT | CERTIFICATE_SIGNING | CRL_SIGNING | ENCIPHER_ONLY | DECIPHER_ONLY | ANY | CUSTOM`

Required: No

**NotAfter**

Filter by certificate expiration date. The start date is inclusive.

Type: [TimestampRange](api-timestamprange.md) object

Required: No

**NotBefore**

Filter by certificate validity start date. The start date is inclusive.

Type: [TimestampRange](api-timestamprange.md) object

Required: No

**SerialNumber**

Filter by serial number.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 59.

Pattern: `[0-9a-f]{2}(:[0-9a-f]{2}){1,19}`

Required: No

**Subject**

Filter by certificate subject.

Type: [SubjectFilter](api-subjectfilter.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

**SubjectAlternativeName**

Filter by subject alternative names.

Type: [SubjectAlternativeNameFilter](api-subjectalternativenamefilter.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/acm-2015-12-08/x509attributefilter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/acm-2015-12-08/x509attributefilter.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/acm-2015-12-08/x509attributefilter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TimestampRange

X509Attributes

All content copied from https://docs.aws.amazon.com/.
