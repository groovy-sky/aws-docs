# Integrating with AWS KMS

Your instance should have an application that can make AWS KMS API requests with the Attestation Document retrieved from the
NitroTPM. When you make a request with an Attestation Document, AWS KMS validates the measurements in the provided Attestation Document against
the reference measurements in the KMS key policy. Requests are allowed only if the measurements in the Attestation Document match the
reference measurements in the KMS key policy.

When you call the [Decrypt](../../../../reference/kms/latest/apireference/api-decrypt.md),
[DeriveSharedSecret](../../../../reference/kms/latest/apireference/api-derivesharedsecret.md),
[GenerateDataKey](../../../../reference/kms/latest/apireference/api-generatedatakey.md),
[GenerateDataKeyPair](../../../../reference/kms/latest/apireference/api-generatedatakeypair.md),
or [GenerateRandom](../../../../reference/kms/latest/apireference/api-generaterandom.md) API
operations with an Attestation Document, these APIs encrypt the plaintext in the response under the public key
from the Attestation Document, and return ciphertext instead of plaintext. This ciphertext can be decrypted only
by using the matching private key that was generated in the instance.

For more information, see the [Cryptographic attestation for NitroTPM](../../../kms/latest/developerguide/services-nitro-enclaves.md) in the _AWS Key Management Service Developer Guide_.

###### Note

If you are attesting to a third-party service, you must build your own custom mechanisms for receiving,
parsing, and validating Attestation Documents. For more information, see [Validate a NitroTPM Attestation Document](nitrotpm-attestation-document-validate.md).

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Validate Attestation Document

Isolate data from your own operators
