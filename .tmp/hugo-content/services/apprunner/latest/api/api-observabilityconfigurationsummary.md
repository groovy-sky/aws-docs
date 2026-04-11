# ObservabilityConfigurationSummary

Provides summary information about an AWS App Runner observability configuration resource.

This type contains limited information about an observability configuration. It includes only identification information, without configuration
details. It's returned by the [ListObservabilityConfigurations](api-listobservabilityconfigurations.md) action. Complete configuration information is returned by the [CreateObservabilityConfiguration](api-createobservabilityconfiguration.md), [DescribeObservabilityConfiguration](api-describeobservabilityconfiguration.md), and [DeleteObservabilityConfiguration](api-deleteobservabilityconfiguration.md)
actions using the [ObservabilityConfiguration](api-observabilityconfiguration.md) type.

## Contents

**ObservabilityConfigurationArn**

The Amazon Resource Name (ARN) of this observability configuration.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

Required: No

**ObservabilityConfigurationName**

The customer-provided observability configuration name. It can be used in multiple revisions of a configuration.

Type: String

Length Constraints: Minimum length of 4. Maximum length of 32.

Pattern: `[A-Za-z0-9][A-Za-z0-9\-_]{3,31}`

Required: No

**ObservabilityConfigurationRevision**

The revision of this observability configuration. It's unique among all the active configurations ( `"Status": "ACTIVE"`) that share the
same `ObservabilityConfigurationName`.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/observabilityconfigurationsummary.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/observabilityconfigurationsummary.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/observabilityconfigurationsummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ObservabilityConfiguration

OperationSummary

All content copied from https://docs.aws.amazon.com/.
