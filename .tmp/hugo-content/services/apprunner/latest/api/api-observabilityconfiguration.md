# ObservabilityConfiguration

Describes an AWS App Runner observability configuration resource. Multiple revisions of a configuration have the same
`ObservabilityConfigurationName` and different `ObservabilityConfigurationRevision` values.

The resource is designed to configure multiple features (currently one feature, tracing). This type contains optional members that describe the
configuration of these features (currently one member, `TraceConfiguration`). If a feature member isn't specified, the feature isn't
enabled.

## Contents

**CreatedAt**

The time when the observability configuration was created. It's in Unix time stamp format.

Type: Timestamp

Required: No

**DeletedAt**

The time when the observability configuration was deleted. It's in Unix time stamp format.

Type: Timestamp

Required: No

**Latest**

It's set to `true` for the configuration with the highest `Revision` among all configurations that share the same
`ObservabilityConfigurationName`. It's set to `false` otherwise.

Type: Boolean

Required: No

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

**Status**

The current state of the observability configuration. If the status of a configuration revision is `INACTIVE`, it was deleted and can't be
used. Inactive configuration revisions are permanently removed some time after they are deleted.

Type: String

Valid Values: `ACTIVE | INACTIVE`

Required: No

**TraceConfiguration**

The configuration of the tracing feature within this observability configuration. If not specified, tracing isn't enabled.

Type: [TraceConfiguration](api-traceconfiguration.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/observabilityconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/observabilityconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/observabilityconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NetworkConfiguration

ObservabilityConfigurationSummary

All content copied from https://docs.aws.amazon.com/.
