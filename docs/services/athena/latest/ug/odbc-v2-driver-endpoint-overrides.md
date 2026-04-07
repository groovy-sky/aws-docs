# Endpoint overrides

## Athena endpoint override

The `endpointOverride ClientConfiguration` class uses this value override
the default HTTP endpoint for the Amazon Athena client. For more information, see [AWS\
Client configuration](https://docs.aws.amazon.com/sdk-for-cpp/v1/developer-guide/client-config.html) in the _AWS SDK for C++ Developer Guide_.

**Connection string name****Parameter type****Default value****Connection string example**EndpointOverrideOptional`none``EndpointOverride=athena.us-west-2.amazonaws.com;`

## Athena streaming endpoint override

The `ClientConfiguration.endpointOverride` method uses this value to
override the default HTTP endpoint for the Amazon Athena streaming client. For more
information, [AWS Client\
configuration](https://docs.aws.amazon.com/sdk-for-cpp/v1/developer-guide/client-config.html) in the _AWS SDK for C++ Developer Guide_. The
Athena Streaming service is available through port 444.

**Connection string name****Parameter type****Default value****Connection string example**StreamingEndpointOverrideOptional`none``StreamingEndpointOverride=athena.us-west-1.amazonaws.com:444;`

## AWS STS endpoint override

The `ClientConfiguration.endpointOverride` method uses this value to
override the default HTTP endpoint for the AWS STS client. For more information, see
[AWS Client configuration](https://docs.aws.amazon.com/sdk-for-cpp/v1/developer-guide/client-config.html) in the _AWS SDK for C++ Developer Guide_.

**Connection string name****Parameter type****Default value****Connection string example**StsEndpointOverrideOptional`none``StsEndpointOverride=sts.us-west-1.amazonaws.com;`

## Lake Formation endpoint override

The `ClientConfiguration.endpointOverride` method uses this value to
override the default HTTP endpoint for the Lake Formation client. For more information, see [AWS\
Client configuration](https://docs.aws.amazon.com/sdk-for-cpp/v1/developer-guide/client-config.html) in the _AWS SDK for C++ Developer Guide_.

**Connection string name****Parameter type****Default value****Connection string example**LakeFormationEndpointOverrideOptional`none``LakeFormationEndpointOverride=lakeformation.us-west-1.amazonaws.com;`

## SSO endpoint override

The `ClientConfiguration.endpointOverride` method uses this value to
override the default HTTP endpoint for the SSO client. For more information, see [AWS\
Client configuration](https://docs.aws.amazon.com/sdk-for-cpp/v1/developer-guide/client-config.html) in the _AWS SDK for C++ Developer Guide_.

**Connection string name****Parameter type****Default value****Connection string example**

SSOEndpointOverride

Optional`none``SSOEndpointOverride=portal.sso.us-east-2.amazonaws.com;`

## SSO OIDC endpoint override

The `ClientConfiguration.endpointOverride` method uses this value to
override the default HTTP endpoint for the SSO OIDC client. For more information, see
[AWS Client configuration](https://docs.aws.amazon.com/sdk-for-cpp/v1/developer-guide/client-config.html) in the _AWS SDK for C++ Developer Guide_.

**Connection string name****Parameter type****Default value****Connection string example**

SSOOIDCEndpointOverride

Optional`none``SSOOIDCEndpointOverride=oidc.us-east-2.amazonaws.com`

## SSO Admin endpoint override

The `ClientConfiguration.endpointOverride` method uses this value to
override the default HTTP endpoint for SSO Admin client. For more information, see
[ClientConfiguration](https://docs.aws.amazon.com/sdk-for-cpp/v1/developer-guide/client-config.html).

**Connection string name****Parameter type****Default value****Connection string example**SSOAdminEndpointOverrideOptionalnone`SSOAdminEndpointOverride=sso.us-east-2.amazonaws.com`

## S3 endpoint override

The `ClientConfiguration.endpointOverride` method uses this value to
override the default HTTP endpoint for S3 client. The endpoint that the driver will use
to download query results when it uses the Amazon S3 fetcher. If this parameter is not
specified, the driver uses a default Amazon S3 endpoint. For more information, see [ClientConfiguration](https://docs.aws.amazon.com/sdk-for-cpp/v1/developer-guide/client-config.html).

**Connection string name****Parameter type****Default value****Connection string example**S3EndpointOverrideOptionalnone`S3EndpointOverride=s3.us-east-2.amazonaws.com`

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Common
auth parameters

Advanced
