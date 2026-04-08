# HttpRedirect

Contains information for HTTP-based domain validation of certificates requested through Amazon CloudFront and issued by ACM.
This field exists only when the certificate type is `AMAZON_ISSUED` and the validation method is `HTTP`.

## Contents

###### Note

In the following list, the required parameters are described first.

**RedirectFrom**

The URL including the domain to be validated. The certificate authority sends
`GET` requests here during validation.

Type: String

Required: No

**RedirectTo**

The URL hosting the validation token. `RedirectFrom` must return this
content or redirect here.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/acm-2015-12-08/httpredirect.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/acm-2015-12-08/httpredirect.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/acm-2015-12-08/httpredirect.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Filters

KeyUsage
