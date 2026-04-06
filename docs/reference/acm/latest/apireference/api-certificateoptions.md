# CertificateOptions

Structure that contains options for your certificate. You can use this structure to
specify whether to opt in to or out of certificate transparency logging and export your
certificate.

Some browsers require that public certificates issued for your domain be recorded in a
log. Certificates that are not logged typically generate a browser error. Transparency
makes it possible for you to detect SSL/TLS certificates that have been mistakenly or
maliciously issued for your domain. For general information, see [Certificate Transparency Logging](../../../../services/acm/latest/userguide/acm-concepts.md#concept-transparency).

You can export public ACM certificates to use with AWS services as well as outside
AWS Cloud. For more information, see [AWS Certificate Manager exportable public certificate](../../../../services/acm/latest/userguide/acm-exportable-certificates.md).

## Contents

###### Note

In the following list, the required parameters are described first.

**CertificateTransparencyLoggingPreference**

You can opt out of certificate transparency logging by specifying the
`DISABLED` option. Opt in by specifying `ENABLED`.

Type: String

Valid Values: `ENABLED | DISABLED`

Required: No

**Export**

You can opt in to allow the export of your certificates by specifying
`ENABLED`. You cannot update the value of `Export` after the
the certificate is created.

Type: String

Valid Values: `ENABLED | DISABLED`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/acm-2015-12-08/CertificateOptions)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/acm-2015-12-08/CertificateOptions)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/acm-2015-12-08/CertificateOptions)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CertificateDetail

CertificateSummary
