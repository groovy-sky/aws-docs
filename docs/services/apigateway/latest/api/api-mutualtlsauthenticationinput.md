# MutualTlsAuthenticationInput

The mutual TLS authentication configuration for a custom domain name. If specified, API Gateway
performs two-way authentication between the client and the server. Clients must present a
trusted certificate to access your API.

## Contents

**truststoreUri**

An Amazon S3 URL that specifies the truststore for mutual TLS authentication, for example
`s3://bucket-name/key-name`. The truststore can contain certificates from public or private
certificate authorities. To update the truststore, upload a new version to S3, and then update
your custom domain name to use the new version. To update the truststore, you must have
permissions to access the S3 object.

Type: String

Required: No

**truststoreVersion**

The version of the S3 object that contains your truststore. To specify a version, you must have versioning enabled for the S3 bucket

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/mutualtlsauthenticationinput.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/mutualtlsauthenticationinput.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/mutualtlsauthenticationinput.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MutualTlsAuthentication

PatchOperation

All content copied from https://docs.aws.amazon.com/.
