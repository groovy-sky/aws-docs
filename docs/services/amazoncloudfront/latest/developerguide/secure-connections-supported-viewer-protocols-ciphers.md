---
title: "Supported protocols and ciphers between viewers and CloudFront"
---

# Supported protocols and ciphers between viewers and CloudFront

When you [require HTTPS between\
viewers and your CloudFront distribution](downloaddistvaluescachebehavior.md#DownloadDistValuesViewerProtocolPolicy), you must choose a [security policy](downloaddistvaluesgeneral.md#DownloadDistValues-security-policy), which
determines the following settings:

- The minimum SSL/TLS protocol that CloudFront uses to communicate with
viewers.

- The ciphers that CloudFront can use to encrypt the communication with
viewers.

To choose a security policy, specify the applicable value for [Security policy (minimum SSL/TLS version)](downloaddistvaluesgeneral.md#DownloadDistValues-security-policy). The following table lists the
protocols and ciphers that CloudFront can use for each security policy.

A viewer must support at least one of the supported ciphers to establish an HTTPS
connection with CloudFront. CloudFront chooses a cipher in the listed order from among the ciphers
that the viewer supports. See also [OpenSSL, s2n, and RFC cipher names](#secure-connections-openssl-rfc-cipher-names).

Security policySSLv3TLSv1TLSv1\_2016TLSv1.1\_2016TLSv1.2\_2018TLSv1.2\_2019TLSv1.2\_2021TLSv1.2\_2025TLSv1.3\_2025**Supported SSL/TLS protocols**TLSv1.3♦♦♦♦♦♦♦♦♦TLSv1.2♦♦♦♦♦♦♦♦TLSv1.1♦♦♦♦TLSv1♦♦♦SSLv3♦**Supported TLSv1.3 ciphers**TLS\_AES\_128\_GCM\_SHA256♦♦♦♦♦♦♦♦♦TLS\_AES\_256\_GCM\_SHA384♦♦♦♦♦♦♦♦♦TLS\_CHACHA20\_POLY1305\_SHA256♦♦♦♦♦♦♦♦**Supported ECDSA ciphers**ECDHE-ECDSA-AES128-GCM-SHA256♦♦♦♦♦♦♦♦ECDHE-ECDSA-AES128-SHA256♦♦♦♦♦♦ECDHE-ECDSA-AES128-SHA♦♦♦♦ECDHE-ECDSA-AES256-GCM-SHA384♦♦♦♦♦♦♦♦ECDHE-ECDSA-CHACHA20-POLY1305♦♦♦♦♦♦♦ECDHE-ECDSA-AES256-SHA384♦♦♦♦♦♦ECDHE-ECDSA-AES256-SHA♦♦♦♦**Supported RSA ciphers**ECDHE-RSA-AES128-GCM-SHA256♦♦♦♦♦♦♦♦ECDHE-RSA-AES128-SHA256♦♦♦♦♦♦ECDHE-RSA-AES128-SHA♦♦♦♦ECDHE-RSA-AES256-GCM-SHA384♦♦♦♦♦♦♦♦ECDHE-RSA-CHACHA20-POLY1305♦♦♦♦♦♦♦ECDHE-RSA-AES256-SHA384♦♦♦♦♦♦ECDHE-RSA-AES256-SHA♦♦♦♦AES128-GCM-SHA256♦♦♦♦♦AES256-GCM-SHA384♦♦♦♦♦AES128-SHA256♦♦♦♦♦AES256-SHA♦♦♦♦AES128-SHA♦♦♦♦DES-CBC3-SHA♦♦RC4-MD5♦

## OpenSSL, s2n, and RFC cipher names

OpenSSL and [s2n](https://github.com/awslabs/s2n) use different
names for ciphers than the TLS standards use ( [RFC 2246](https://tools.ietf.org/html/rfc2246), [RFC 4346](https://tools.ietf.org/html/rfc4346), [RFC 5246](https://tools.ietf.org/html/rfc5246), and [RFC 8446](https://tools.ietf.org/html/rfc8446)). The following table
maps the OpenSSL and s2n names to the RFC name for each cipher.

CloudFront supports both classical and quantum-safe key exchanges. For classical key exchanges using elliptic curves, CloudFront supports the
following:

- `prime256v1`

- `X25519`

- `secp384r1`

For quantum-safe key exchanges, CloudFront supports the following:

- `X25519MLKEM768`

- `SecP256r1MLKEM768`

###### Note

Quantum-safe key exchanges are only supported with TLS 1.3. TLS 1.2 and earlier versions do not support quantum-safe key exchanges.

For more information, see the following topics:

- [Post-Quantum Cryptography](https://aws.amazon.com/security/post-quantum-cryptography)

- [Cryptography algorithms and AWS services](../../../prescriptive-guidance/latest/encryption-best-practices/aws-cryptography-services.md#algorithms)

- [Hybrid key exchange in TLS 1.3](https://datatracker.ietf.org/doc/draft-ietf-tls-hybrid-design)

For more information about certificate requirements for CloudFront, see [Requirements for using SSL/TLS certificates with CloudFront](cnames-and-https-requirements.md).

OpenSSL and s2n cipher nameRFC cipher name**Supported TLSv1.3 ciphers**TLS\_AES\_128\_GCM\_SHA256TLS\_AES\_128\_GCM\_SHA256TLS\_AES\_256\_GCM\_SHA384TLS\_AES\_256\_GCM\_SHA384TLS\_CHACHA20\_POLY1305\_SHA256TLS\_CHACHA20\_POLY1305\_SHA256**Supported ECDSA ciphers**ECDHE-ECDSA-AES128-GCM-SHA256TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256ECDHE-ECDSA-AES128-SHA256TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_CBC\_SHA256ECDHE-ECDSA-AES128-SHATLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_CBC\_SHAECDHE-ECDSA-AES256-GCM-SHA384TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384ECDHE-ECDSA-CHACHA20-POLY1305TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256ECDHE-ECDSA-AES256-SHA384TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_CBC\_SHA384ECDHE-ECDSA-AES256-SHATLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_CBC\_SHA**Supported RSA ciphers**ECDHE-RSA-AES128-GCM-SHA256TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256ECDHE-RSA-AES128-SHA256TLS\_ECDHE\_RSA\_WITH\_AES\_128\_CBC\_SHA256 ECDHE-RSA-AES128-SHATLS\_ECDHE\_RSA\_WITH\_AES\_128\_CBC\_SHAECDHE-RSA-AES256-GCM-SHA384TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 ECDHE-RSA-CHACHA20-POLY1305TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256ECDHE-RSA-AES256-SHA384TLS\_ECDHE\_RSA\_WITH\_AES\_256\_CBC\_SHA384 ECDHE-RSA-AES256-SHATLS\_ECDHE\_RSA\_WITH\_AES\_256\_CBC\_SHAAES128-GCM-SHA256TLS\_RSA\_WITH\_AES\_128\_GCM\_SHA256AES256-GCM-SHA384TLS\_RSA\_WITH\_AES\_256\_GCM\_SHA384AES128-SHA256TLS\_RSA\_WITH\_AES\_128\_CBC\_SHA256AES256-SHATLS\_RSA\_WITH\_AES\_256\_CBC\_SHAAES128-SHATLS\_RSA\_WITH\_AES\_128\_CBC\_SHADES-CBC3-SHA TLS\_RSA\_WITH\_3DES\_EDE\_CBC\_SHA RC4-MD5TLS\_RSA\_WITH\_RC4\_128\_MD5

## Supported signature schemes between viewers and CloudFront

CloudFront supports the following signature schemes for connections between viewers and
CloudFront.

Security
policySignature schemesSSLv3TLSv1TLSv1\_2016TLSv1.1\_2016TLSv1.2\_2018TLSv1.2\_2019 TLSv1.2\_2021TLSv1.2\_2025TLSv1.3\_2025TLS\_SIGNATURE\_SCHEME\_RSA\_PSS\_PSS\_SHA256♦♦♦♦♦♦♦♦♦TLS\_SIGNATURE\_SCHEME\_RSA\_PSS\_PSS\_SHA384♦♦♦♦♦♦♦♦♦TLS\_SIGNATURE\_SCHEME\_RSA\_PSS\_PSS\_SHA512♦♦♦♦♦♦♦♦♦TLS\_SIGNATURE\_SCHEME\_RSA\_PSS\_RSAE\_SHA256♦♦♦♦♦♦♦♦♦TLS\_SIGNATURE\_SCHEME\_RSA\_PSS\_RSAE\_SHA384♦♦♦♦♦♦♦♦♦TLS\_SIGNATURE\_SCHEME\_RSA\_PSS\_RSAE\_SHA512♦♦♦♦♦♦♦♦♦TLS\_SIGNATURE\_SCHEME\_RSA\_PKCS1\_SHA256♦♦♦♦♦♦♦♦♦TLS\_SIGNATURE\_SCHEME\_RSA\_PKCS1\_SHA384♦♦♦♦♦♦♦♦♦TLS\_SIGNATURE\_SCHEME\_RSA\_PKCS1\_SHA512♦♦♦♦♦♦♦♦♦TLS\_SIGNATURE\_SCHEME\_RSA\_PKCS1\_SHA224♦♦♦♦♦♦♦TLS\_SIGNATURE\_SCHEME\_ECDSA\_SHA256♦♦♦♦♦♦♦♦♦TLS\_SIGNATURE\_SCHEME\_ECDSA\_SHA384♦♦♦♦♦♦♦♦♦TLS\_SIGNATURE\_SCHEME\_ECDSA\_SHA512♦♦♦♦♦♦♦♦♦TLS\_SIGNATURE\_SCHEME\_ECDSA\_SHA224♦♦♦♦♦♦♦TLS\_SIGNATURE\_SCHEME\_ECDSA\_SECP256R1\_SHA256♦♦♦♦♦♦♦♦♦TLS\_SIGNATURE\_SCHEME\_ECDSA\_SECP384R1\_SHA384♦♦♦♦♦♦♦♦♦TLS\_SIGNATURE\_SCHEME\_RSA\_PKCS1\_SHA1♦♦♦♦TLS\_SIGNATURE\_SCHEME\_ECDSA\_SHA1♦♦♦♦

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Require HTTPS to an Amazon S3 origin

Supported protocols and ciphers between CloudFront and the origin

All content copied from https://docs.aws.amazon.com/.
