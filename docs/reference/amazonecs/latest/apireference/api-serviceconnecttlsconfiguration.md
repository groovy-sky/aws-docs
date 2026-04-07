# ServiceConnectTlsConfiguration

The key that encrypts and decrypts your resources for Service Connect TLS.

## Contents

**issuerCertificateAuthority**

The signer certificate authority.

Type: [ServiceConnectTlsCertificateAuthority](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ServiceConnectTlsCertificateAuthority.html) object

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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ecs-2014-11-13/ServiceConnectTlsConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ecs-2014-11-13/ServiceConnectTlsConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ecs-2014-11-13/ServiceConnectTlsConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ServiceConnectTlsCertificateAuthority

ServiceCurrentRevisionSummary
