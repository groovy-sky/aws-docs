# ServiceConnectTlsConfiguration

The key that encrypts and decrypts your resources for Service Connect TLS.

## Contents

**issuerCertificateAuthority**

The signer certificate authority.

Type: [ServiceConnectTlsCertificateAuthority](api-serviceconnecttlscertificateauthority.md) object

Required: Yes

**kmsKey**

The AWS Key Management Service key.

Type: String

Required: No

**roleArn**

The Amazon Resource Name (ARN) of the IAM role that's associated with the Service
Connect TLS.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/serviceconnecttlsconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/serviceconnecttlsconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/serviceconnecttlsconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ServiceConnectTlsCertificateAuthority

ServiceCurrentRevisionSummary
