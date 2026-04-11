# KeyPairInfo

Describes a key pair.

## Contents

**createTime**

If you used Amazon EC2 to create the key pair, this is the date and time when the key
was created, in [ISO\
8601 date-time format](https://www.iso.org/iso-8601-date-and-time-format.html), in the UTC time zone.

If you imported an existing key pair to Amazon EC2, this is the date and time the key
was imported, in [ISO\
8601 date-time format](https://www.iso.org/iso-8601-date-and-time-format.html), in the UTC time zone.

Type: Timestamp

Required: No

**keyFingerprint**

If you used [CreateKeyPair](api-createkeypair.md) to create the key pair:

- For RSA key pairs, the key fingerprint is the SHA-1 digest of the DER encoded private key.

- For ED25519 key pairs, the key fingerprint is the base64-encoded SHA-256 digest, which
is the default for OpenSSH, starting with [OpenSSH 6.8](http://www.openssh.com/txt/release-6.8).

If you used [ImportKeyPair](api-importkeypair.md) to provide AWS the public key:

- For RSA key pairs, the key fingerprint is the MD5 public key fingerprint as specified in section 4 of RFC4716.

- For ED25519 key pairs, the key fingerprint is the base64-encoded SHA-256
digest, which is the default for OpenSSH, starting with [OpenSSH 6.8](http://www.openssh.com/txt/release-6.8).

Type: String

Required: No

**keyName**

The name of the key pair.

Type: String

Required: No

**keyPairId**

The ID of the key pair.

Type: String

Required: No

**keyType**

The type of key pair.

Type: String

Valid Values: `rsa | ed25519`

Required: No

**publicKey**

The public key material.

Type: String

Required: No

**TagSet.N**

Any tags applied to the key pair.

Type: Array of [Tag](api-tag.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/keypairinfo.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/keypairinfo.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/keypairinfo.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Ipv6Range

LastError
