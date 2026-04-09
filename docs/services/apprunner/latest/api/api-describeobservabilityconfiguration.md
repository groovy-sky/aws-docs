# DescribeObservabilityConfiguration

###### Important

AWS App Runner will no longer be open to new
customers starting March 31, 2026. If you would like to use App Runner, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS App Runner availability change](../dg/apprunner-availability-change.md).

Return a full description of an AWS App Runner observability configuration resource.

## Request Syntax

```nohighlight

{
   "ObservabilityConfigurationArn": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ObservabilityConfigurationArn](#API_DescribeObservabilityConfiguration_RequestSyntax)**

The Amazon Resource Name (ARN) of the App Runner observability configuration that you want a description for.

The ARN can be a full observability configuration ARN, or a partial ARN ending with either `.../name
                  ` or
`.../name/revision
                  `. If a revision isn't specified, the latest active revision is
described.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

Required: Yes

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

**[ObservabilityConfiguration](#API_DescribeObservabilityConfiguration_ResponseSyntax)**

A full description of the App Runner observability configuration that you specified in this request.

Type: [ObservabilityConfiguration](api-observabilityconfiguration.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServiceErrorException**

An unexpected service exception occurred.

HTTP Status Code: 500

**InvalidRequestException**

One or more input parameters aren't valid. Refer to the API action's document page, correct the input parameters, and try the action again.

HTTP Status Code: 400

**ResourceNotFoundException**

A resource doesn't exist for the specified Amazon Resource Name (ARN) in your AWS account.

HTTP Status Code: 400

## Examples

### Describe the latest active revision of an observability configuration

This example illustrates how to get a description of the latest active revision of an App Runner observability configuration. To describe the latest
active revision, specify an ARN that ends with the configuration name, without the revision component.

In the example, two revisions exist. Therefore, revision `2` (the latest) is described. The resulting object shows `"Latest":
            true`.

###### Note

The two revisions in our examples are identical, because App Runner doesn't yet support enough observability functionality to demonstrate two
significantly different revisions (for example, multiple tracing vendors). We're including the two examples only to demonstrate the revisioning
behavior during retrieval.

#### Sample Request

```json

$ aws apprunner describe-observability-configuration --cli-input-json "`cat`"
{
  "ObservabilityConfigurationArn": "arn:aws:apprunner:us-east-1:123456789012:observabilityconfiguration/xray-tracing"
}
```

#### Sample Response

```json

{
  "ObservabilityConfiguration": {
    "ObservabilityConfigurationArn": "arn:aws:apprunner:us-east-1:123456789012:observabilityconfiguration/xray-tracing/2/e76562f50d78042e819fead0f59672e6",
    "ObservabilityConfigurationName": "xray-tracing",
    "ObservabilityConfigurationRevision": 2,
    "CreatedAt": "2021-02-25T17:42:59Z",
    "Latest": true,
    "Status": "ACTIVE",
    "TraceConfiguration": {
        "Vendor": "AWSXRAY"
    }
  }
}
```

### Describe a specific revision of an observability configuration

This example illustrates how to get a description of a specific revision of an App Runner observability configuration. To describe a specific revision,
specify an ARN that includes the revision number.

In the example, several revisions exist and revision `1` is queried. The resulting object shows `"Latest": false`.

#### Sample Request

```json

$ aws apprunner describe-observability-configuration --cli-input-json "`cat`"
{
  "ObservabilityConfigurationArn": "arn:aws:apprunner:us-east-1:123456789012:observabilityconfiguration/xray-tracing/1"
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
    "Latest": false,
    "Status": "ACTIVE",
    "TraceConfiguration": {
        "Vendor": "AWSXRAY"
    }
  }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apprunner-2020-05-15/describeobservabilityconfiguration.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apprunner-2020-05-15/describeobservabilityconfiguration.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/describeobservabilityconfiguration.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apprunner-2020-05-15/describeobservabilityconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/describeobservabilityconfiguration.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apprunner-2020-05-15/describeobservabilityconfiguration.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apprunner-2020-05-15/describeobservabilityconfiguration.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apprunner-2020-05-15/describeobservabilityconfiguration.md)

- [AWS SDK for Python](../../../goto/boto3/apprunner-2020-05-15/describeobservabilityconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/describeobservabilityconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeCustomDomains

DescribeService

All content copied from https://docs.aws.amazon.com/.
