---
title: "X509Attributes"
---

# X509Attributes

Contains X.509 certificate attributes extracted from the certificate.

## Contents

###### Note

In the following list, the required parameters are described first.

**ExtendedKeyUsages**

Contains a list of Extended Key Usage X.509 v3 extension objects. Each object specifies a
purpose for which the certificate public key can be used and consists of a name and an object
identifier (OID).

Type: Array of strings

Valid Values: `TLS_WEB_SERVER_AUTHENTICATION | TLS_WEB_CLIENT_AUTHENTICATION | CODE_SIGNING | EMAIL_PROTECTION | TIME_STAMPING | OCSP_SIGNING | IPSEC_END_SYSTEM | IPSEC_TUNNEL | IPSEC_USER | ANY | NONE | CUSTOM`

Required: No

**Issuer**

The distinguished name of the certificate issuer.

Type: [DistinguishedName](api-distinguishedname.md) object

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

**NotAfter**

The time after which the certificate is not valid.

Type: Timestamp

Required: No

**NotBefore**

The time before which the certificate is not valid.

Type: Timestamp

Required: No

**SerialNumber**

The serial number assigned by the certificate authority.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 59.

Pattern: `[0-9a-f]{2}(:[0-9a-f]{2}){1,19}`

Required: No

**Subject**

The distinguished name of the certificate subject.

Type: [DistinguishedName](api-distinguishedname.md) object

Required: No

**SubjectAlternativeNames**

One or more domain names (subject alternative names)
included in the certificate. This
list contains the domain names that are bound to the public key that is contained in the
certificate. The subject alternative names include the canonical domain name (CN) of the
certificate and additional domain names that can be used to connect to the website.

Type: Array of [GeneralName](api-generalname.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/acm-2015-12-08/x509attributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/acm-2015-12-08/x509attributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/acm-2015-12-08/x509attributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

X509AttributeFilter

Common Parameters

All content copied from https://docs.aws.amazon.com/.
