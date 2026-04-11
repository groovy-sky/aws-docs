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

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/tlsconfig.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/tlsconfig.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/tlsconfig.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ThrottleSettings

UsagePlan

All content copied from https://docs.aws.amazon.com/.
