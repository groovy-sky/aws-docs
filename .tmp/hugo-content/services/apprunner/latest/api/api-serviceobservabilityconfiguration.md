# ServiceObservabilityConfiguration

Describes the observability configuration of an AWS App Runner service. These are additional observability features, like tracing, that you choose to
enable. They're configured in a separate resource that you associate with your service.

## Contents

**ObservabilityEnabled**

When `true`, an observability configuration resource is associated with the service, and an `ObservabilityConfigurationArn` is
specified.

Type: Boolean

Required: Yes

**ObservabilityConfigurationArn**

The Amazon Resource Name (ARN) of the observability configuration that is associated with the service. Specified only when
`ObservabilityEnabled` is `true`.

Specify an ARN with a name and a revision number to associate that revision. For example:
`arn:aws:apprunner:us-east-1:123456789012:observabilityconfiguration/xray-tracing/3`

Specify just the name to associate the latest revision. For example:
`arn:aws:apprunner:us-east-1:123456789012:observabilityconfiguration/xray-tracing`

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/serviceobservabilityconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/serviceobservabilityconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/serviceobservabilityconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Service

ServiceSummary

All content copied from https://docs.aws.amazon.com/.
