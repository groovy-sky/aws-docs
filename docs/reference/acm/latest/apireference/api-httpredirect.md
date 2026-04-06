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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/acm-2015-12-08/HttpRedirect)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/acm-2015-12-08/HttpRedirect)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/acm-2015-12-08/HttpRedirect)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Filters

KeyUsage
