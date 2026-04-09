# CreateAutoScalingConfiguration

###### Important

AWS App Runner will no longer be open to new
customers starting March 31, 2026. If you would like to use App Runner, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS App Runner availability change](../dg/apprunner-availability-change.md).

Create an AWS App Runner automatic scaling configuration resource. App Runner requires this resource when you create or update App Runner services and you require
non-default auto scaling settings. You can share an auto scaling configuration across multiple services.

Create multiple revisions of a configuration by calling this action multiple times using the same `AutoScalingConfigurationName`. The call
returns incremental `AutoScalingConfigurationRevision` values. When you create a service and configure an auto scaling configuration resource,
the service uses the latest active revision of the auto scaling configuration by default. You can optionally configure the service to use a specific
revision.

Configure a higher `MinSize` to increase the spread of your App Runner service over more Availability Zones in the AWS Region. The
tradeoff is a higher minimal cost.

Configure a lower `MaxSize` to control your cost. The tradeoff is lower responsiveness during peak demand.

## Request Syntax

```nohighlight

{
   "AutoScalingConfigurationName": "string",
   "MaxConcurrency": number,
   "MaxSize": number,
   "MinSize": number,
   "Tags": [
      {
         "Key": "string",
         "Value": "string"
      }
   ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[AutoScalingConfigurationName](#API_CreateAutoScalingConfiguration_RequestSyntax)**

A name for the auto scaling configuration. When you use it for the first time in an AWS Region, App Runner creates revision number
`1` of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration.

###### Note

Prior to the release of [Auto scale\
configuration enhancements](../relnotes/release-2023-09-22-auto-scale-config.md), the name `DefaultConfiguration` was reserved.

This restriction is no longer in place. You can now manage `DefaultConfiguration` the same way you manage your custom auto scaling
configurations. This means you can do the following with the `DefaultConfiguration` that App Runner provides:

- Create new revisions of the `DefaultConfiguration`.

- Delete the revisions of the `DefaultConfiguration`.

- Delete the auto scaling configuration for which the App Runner `DefaultConfiguration` was created.

- If you delete the auto scaling configuration you can create another custom auto scaling configuration with the same
`DefaultConfiguration` name. The original `DefaultConfiguration` resource provided by App Runner remains in your account unless
you make changes to it.

Type: String

Length Constraints: Minimum length of 4. Maximum length of 32.

Pattern: `[A-Za-z0-9][A-Za-z0-9\-_]{3,31}`

Required: Yes

**[MaxConcurrency](#API_CreateAutoScalingConfiguration_RequestSyntax)**

The maximum number of concurrent requests that you want an instance to process. If the number of concurrent requests exceeds this limit, App Runner scales
up your service.

Default: `100`

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 200.

Required: No

**[MaxSize](#API_CreateAutoScalingConfiguration_RequestSyntax)**

The maximum number of instances that your service scales up to. At most `MaxSize` instances actively serve traffic for your service.

Default: `25`

Type: Integer

Valid Range: Minimum value of 1.

Required: No

**[MinSize](#API_CreateAutoScalingConfiguration_RequestSyntax)**

The minimum number of instances that App Runner provisions for your service. The service always has at least `MinSize` provisioned instances.
Some of them actively serve traffic. The rest of them (provisioned and inactive instances) are a cost-effective compute capacity reserve and are ready to
be quickly activated. You pay for memory usage of all the provisioned instances. You pay for CPU usage of only the active subset.

App Runner temporarily doubles the number of provisioned instances during deployments, to maintain the same capacity for both old and new code.

Default: `1`

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 25.

Required: No

**[Tags](#API_CreateAutoScalingConfiguration_RequestSyntax)**

A list of metadata items that you can associate with your auto scaling configuration resource. A tag is a key-value pair.

Type: Array of [Tag](api-tag.md) objects

Required: No

## Response Syntax

```nohighlight

{
   "AutoScalingConfiguration": {
      "AutoScalingConfigurationArn": "string",
      "AutoScalingConfigurationName": "string",
      "AutoScalingConfigurationRevision": number,
      "CreatedAt": number,
      "DeletedAt": number,
      "HasAssociatedService": boolean,
      "IsDefault": boolean,
      "Latest": boolean,
      "MaxConcurrency": number,
      "MaxSize": number,
      "MinSize": number,
      "Status": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[AutoScalingConfiguration](#API_CreateAutoScalingConfiguration_ResponseSyntax)**

A description of the App Runner auto scaling configuration that's created by this request.

Type: [AutoScalingConfiguration](api-autoscalingconfiguration.md) object

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

### Create a high availability auto scaling configuration

This example illustrates how to create an auto scaling configuration optimized for high availability by setting `MinSize` to
`5`. With this configuration, App Runner attempts to spread your service instances over the most Availability Zones possible, up to five,
depending on the AWS Region.

The call returns an `AutoScalingConfiguration` object with the other settings set to their defaults. In the example, this is the first
call to create a configuration named `high-availability`. The revision is set to `1`, and it's the latest revision.

#### Sample Request

```json

$ aws apprunner create-auto-scaling-configuration --cli-input-json "`cat`"
{
  "AutoScalingConfigurationName": "high-availability",
  "MinSize": 5
}
```

#### Sample Response

```json

{
  "AutoScalingConfiguration": {
    "AutoScalingConfigurationArn": "arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability/1/2f50e7656d7819fead0f59672e68042e",
    "AutoScalingConfigurationName": "high-availability",
    "AutoScalingConfigurationRevision": 1,
    "CreatedAt": "2020-11-03T00:29:17Z",
    "Latest": true,
    "Status": "ACTIVE",
    "MaxConcurrency": 100,
    "MaxSize": 25,
    "MinSize": 5
  }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apprunner-2020-05-15/createautoscalingconfiguration.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apprunner-2020-05-15/createautoscalingconfiguration.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/createautoscalingconfiguration.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apprunner-2020-05-15/createautoscalingconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/createautoscalingconfiguration.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apprunner-2020-05-15/createautoscalingconfiguration.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apprunner-2020-05-15/createautoscalingconfiguration.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apprunner-2020-05-15/createautoscalingconfiguration.md)

- [AWS SDK for Python](../../../goto/boto3/apprunner-2020-05-15/createautoscalingconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/createautoscalingconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssociateCustomDomain

CreateConnection

All content copied from https://docs.aws.amazon.com/.
