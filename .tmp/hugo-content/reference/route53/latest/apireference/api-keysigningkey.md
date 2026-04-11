# KeySigningKey

A key-signing key (KSK) is a complex type that represents a public/private key pair.
The private key is used to generate a digital signature for the zone signing key (ZSK).
The public key is stored in the DNS and is used to authenticate the ZSK. A KSK is always
associated with a hosted zone; it cannot exist by itself.

## Contents

**CreatedDate**

The date when the key-signing key (KSK) was created.

Type: Timestamp

Required: No

**DigestAlgorithmMnemonic**

A string used to represent the delegation signer digest algorithm. This value must
follow the guidelines provided by [RFC-8624 Section 3.3](https://tools.ietf.org/html/rfc8624).

Type: String

Required: No

**DigestAlgorithmType**

An integer used to represent the delegation signer digest algorithm. This value must
follow the guidelines provided by [RFC-8624 Section\
3.3](https://tools.ietf.org/html/rfc8624).

Type: Integer

Required: No

**DigestValue**

A cryptographic digest of a DNSKEY resource record (RR). DNSKEY records are used to
publish the public key that resolvers can use to verify DNSSEC signatures that are used
to secure certain kinds of information provided by the DNS system.

Type: String

Required: No

**DNSKEYRecord**

A string that represents a DNSKEY record.

Type: String

Required: No

**DSRecord**

A string that represents a delegation signer (DS) record.

Type: String

Required: No

**Flag**

An integer that specifies how the key is used. For key-signing key (KSK), this value
is always 257.

Type: Integer

Required: No

**KeyTag**

An integer used to identify the DNSSEC record for the domain name. The process used to
calculate the value is described in [RFC-4034 Appendix B](https://tools.ietf.org/rfc/rfc4034.txt).

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 65536.

Required: No

**KmsArn**

The Amazon resource name (ARN) used to identify the customer managed key in AWS Key Management Service (AWS KMS). The `KmsArn` must be unique for each
key-signing key (KSK) in a single hosted zone.

You must configure the customer managed key as follows:

Status

Enabled

Key spec

ECC\_NIST\_P256

Key usage

Sign and verify

Key policy

The key policy must give permission for the following actions:

- DescribeKey

- GetPublicKey

- Sign

The key policy must also include the Amazon Route 53 service in the
principal for your account. Specify the following:

- `"Service": "dnssec-route53.amazonaws.com"`

For more information about working with the customer managed key in AWS KMS, see [AWS Key Management Service\
concepts](../../../../services/kms/latest/developerguide/concepts.md).

Type: String

Required: No

**LastModifiedDate**

The last time that the key-signing key (KSK) was changed.

Type: Timestamp

Required: No

**Name**

A string used to identify a key-signing key (KSK). `Name` can include
numbers, letters, and underscores (\_). `Name` must be unique for each
key-signing key in the same hosted zone.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 128.

Required: No

**PublicKey**

The public key, represented as a Base64 encoding, as required by [RFC-4034 Page 5](https://tools.ietf.org/rfc/rfc4034.txt).

Type: String

Required: No

**SigningAlgorithmMnemonic**

A string used to represent the signing algorithm. This value must follow the
guidelines provided by [RFC-8624 Section 3.1](https://tools.ietf.org/html/rfc8624).

Type: String

Required: No

**SigningAlgorithmType**

An integer used to represent the signing algorithm. This value must follow the
guidelines provided by [RFC-8624 Section 3.1](https://tools.ietf.org/html/rfc8624).

Type: Integer

Required: No

**Status**

A string that represents the current key-signing key (KSK) status.

Status can have one of the following values:

ACTIVE

The KSK is being used for signing.

INACTIVE

The KSK is not being used for signing.

DELETING

The KSK is in the process of being deleted.

ACTION\_NEEDED

There is a problem with the KSK that requires you to take action to
resolve. For example, the customer managed key might have been deleted,
or the permissions for the customer managed key might have been
changed.

INTERNAL\_FAILURE

There was an error during a request. Before you can continue to work with
DNSSEC signing, including actions that involve this KSK, you must correct
the problem. For example, you may need to activate or deactivate the
KSK.

Type: String

Length Constraints: Minimum length of 5. Maximum length of 150.

Required: No

**StatusMessage**

The status message provided for the following key-signing key (KSK) statuses:
`ACTION_NEEDED` or `INTERNAL_FAILURE`. The status message
includes information about what the problem might be and steps that you can take to
correct the issue.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 512.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/keysigningkey.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/keysigningkey.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/keysigningkey.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

HostedZoneSummary

LinkedService
