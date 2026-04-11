# CreateObservabilityConfiguration

###### Important

AWS App Runner will no longer be open to new
customers starting March 31, 2026. If you would like to use App Runner, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS App Runner availability change](../dg/apprunner-availability-change.md).

Create an AWS App Runner observability configuration resource. App Runner requires this resource when you create or update App Runner services and you want to enable
non-default observability features. You can share an observability configuration across multiple services.

Create multiple revisions of a configuration by calling this action multiple times using the same `ObservabilityConfigurationName`. The
call returns incremental `ObservabilityConfigurationRevision` values. When you create a service and configure an observability configuration
resource, the service uses the latest active revision of the observability configuration by default. You can optionally configure the service to use a
specific revision.

The observability configuration resource is designed to configure multiple features (currently one feature, tracing). This action takes optional
parameters that describe the configuration of these features (currently one parameter, `TraceConfiguration`). If you don't specify a feature
parameter, App Runner doesn't enable the feature.

## Request Syntax

```nohighlight

{
   "ObservabilityConfigurationName": "string",
   "Tags": [
      {
         "Key": "string",
         "Value": "string"
      }
   ],
   "TraceConfiguration": {
      "Vendor": "string"
   }
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ObservabilityConfigurationName](#API_CreateObservabilityConfiguration_RequestSyntax)**

A name for the observability configuration. When you use it for the first time in an AWS Region, App Runner creates revision number
`1` of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration.

###### Note

The name `DefaultConfiguration` is reserved. You can't use it to create a new observability configuration, and you can't create a
revision of it.

When you want to use your own observability configuration for your App Runner service, _create a configuration with a different name_,
and then provide it when you create or update your service.

Type: String

Length Constraints: Minimum length of 4. Maximum length of 32.

Pattern: `[A-Za-z0-9][A-Za-z0-9\-_]{3,31}`

Required: Yes

**[Tags](#API_CreateObservabilityConfiguration_RequestSyntax)**

A list of metadata items that you can associate with your observability configuration resource. A tag is a key-value pair.

Type: Array of [Tag](api-tag.md) objects

Required: No

**[TraceConfiguration](#API_CreateObservabilityConfiguration_RequestSyntax)**

The configuration of the tracing feature within this observability configuration. If you don't specify it, App Runner doesn't enable tracing.

Type: [TraceConfiguration](api-traceconfiguration.md) object

Required: No

## Response Syntax

```nohighlight

{
   "ObservabilityConfiguration": {
      "CreatedAt": number,
      "DeletedAt": number,
      "Latest": boolean,
      "ObservabilityConfigurationArn": "string",
      "ObservabilityConfigurationName": "string",
      "ObservabilityConfigurationRevision": number,
      "Status": "string",
      "TraceConfiguration": {
         "Vendor": "string"
      }
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[ObservabilityConfiguration](#API_CreateObservabilityConfiguration_ResponseSyntax)**

A description of the App Runner observability configuration that's created by this request.

Type: [ObservabilityConfiguration](api-observabilityconfiguration.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServiceErrorException**

An unexpected service exception occurred.

HTTP Status Code: 500

**InvalidRequestException**

One or more input parameters aren't valid. Refer to the API action's document page, correct the input parameters, and try the action again.

HTTP Status Code: 400

**ServiceQuotaExceededException**

App Runner can't create this resource. You've reached your account quota for this resource type.

For App Runner per-resource quotas, see [AWS App Runner endpoints and quotas](../../../../general/latest/gr/apprunner.md) in the
_AWS General Reference_.

HTTP Status Code: 400

## Examples

### Create an observability configuration to enable X-Ray tracing

This example illustrates how to create an observability configuration that enables tracing using AWS X-Ray.

The call returns an `ObservabilityConfiguration` object. In the example, this is the first call to create a configuration named
`xray-tracing`. The revision is set to `1`, and it's the latest revision.

#### Sample Request

```json

$ aws apprunner create-observability-configuration --cli-input-json "`cat`"
{
  "ObservabilityConfigurationName": "xray-tracing",
  "TraceConfiguration": {
    "Vendor": "AWSXRAY"
  }
}
```

#### Sample Response

```json

{
  "ObservabilityConfiguration": {
    "ObservabilityConfigurationArn": "arn:aws:apprunner:us-east-1:123456789012:observabilityconfiguration/xray-tracing/1/2f50e7656d7819fead0f59672e68042e",
    "ObservabilityConfigurationName": "xray-tracing",
    "ObservabilityConfigurationRevision": 1,
    "CreatedAt": "2020-11-03T00:29:17Z",
    "Latest": true,
    "Status": "ACTIVE",
    "TraceConfiguration": {
        "Vendor": "AWSXRAY"
    }
  }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apprunner-2020-05-15/createobservabilityconfiguration.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apprunner-2020-05-15/createobservabilityconfiguration.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/createobservabilityconfiguration.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apprunner-2020-05-15/createobservabilityconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/createobservabilityconfiguration.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apprunner-2020-05-15/createobservabilityconfiguration.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apprunner-2020-05-15/createobservabilityconfiguration.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apprunner-2020-05-15/createobservabilityconfiguration.md)

- [AWS SDK for Python](../../../goto/boto3/apprunner-2020-05-15/createobservabilityconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/createobservabilityconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateConnection

CreateService

All content copied from https://docs.aws.amazon.com/.
