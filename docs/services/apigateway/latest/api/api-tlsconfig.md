# TlsConfig

Specifies the TLS configuration for an integration.

## Contents

**insecureSkipVerification**

Specifies whether or not API Gateway skips verification that the certificate for an integration endpoint is
issued by a supported certificate authority. This isn’t recommended, but it enables you to
use certificates that are signed by private certificate authorities, or certificates
that are self-signed. If enabled, API Gateway still performs basic certificate
validation, which includes checking the certificate's expiration date, hostname, and
presence of a root certificate authority. Supported only for `HTTP` and
`HTTP_PROXY` integrations.

###### Important

Enabling `insecureSkipVerification` isn't recommended, especially for integrations with public
HTTPS endpoints. If you enable `insecureSkipVerification`, you increase the risk of man-in-the-middle attacks.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apigateway-2015-07-09/TlsConfig)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apigateway-2015-07-09/TlsConfig)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apigateway-2015-07-09/TlsConfig)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ThrottleSettings

UsagePlan
