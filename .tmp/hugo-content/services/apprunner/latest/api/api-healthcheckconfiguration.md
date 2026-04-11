# HealthCheckConfiguration

Describes the settings for the health check that AWS App Runner performs to monitor the health of a service.

## Contents

**HealthyThreshold**

The number of consecutive checks that must succeed before App Runner decides that the service is healthy.

Default: `1`

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 20.

Required: No

**Interval**

The time interval, in seconds, between health checks.

Default: `5`

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 20.

Required: No

**Path**

The URL that health check requests are sent to.

`Path` is only applicable when you set `Protocol` to `HTTP`.

Default: `"/"`

Type: String

Length Constraints: Minimum length of 1.

Required: No

**Protocol**

The IP protocol that App Runner uses to perform health checks for your service.

If you set `Protocol` to `HTTP`, App Runner sends health check requests to the HTTP path specified by `Path`.

Default: `TCP`

Type: String

Valid Values: `TCP | HTTP`

Required: No

**Timeout**

The time, in seconds, to wait for a health check response before deciding it failed.

Default: `2`

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 20.

Required: No

**UnhealthyThreshold**

The number of consecutive checks that must fail before App Runner decides that the service is unhealthy.

Default: `5`

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 20.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/healthcheckconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/healthcheckconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/healthcheckconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionConfiguration

ImageConfiguration

All content copied from https://docs.aws.amazon.com/.
