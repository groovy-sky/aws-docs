# ListObservabilityConfigurations

###### Important

AWS App Runner will no longer be open to new
customers starting March 31, 2026. If you would like to use App Runner, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS App Runner availability change](../dg/apprunner-availability-change.md).

Returns a list of active AWS App Runner observability configurations in your AWS account. You can query the revisions for a specific
configuration name or the revisions for all active configurations in your account. You can optionally query only the latest revision of each requested
name.

To retrieve a full description of a particular configuration revision, call [DescribeObservabilityConfiguration](api-describeobservabilityconfiguration.md) and provide one
of the ARNs returned by `ListObservabilityConfigurations`.

## Request Syntax

```nohighlight

{
   "LatestOnly": boolean,
   "MaxResults": number,
   "NextToken": "string",
   "ObservabilityConfigurationName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[LatestOnly](#API_ListObservabilityConfigurations_RequestSyntax)**

Set to `true` to list only the latest revision for each requested configuration name.

Set to `false` to list all revisions for each requested configuration name.

Default: `true`

Type: Boolean

Required: No

**[MaxResults](#API_ListObservabilityConfigurations_RequestSyntax)**

The maximum number of results to include in each response (result page). It's used for a paginated request.

If you don't specify `MaxResults`, the request retrieves all available results in a single response.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

**[NextToken](#API_ListObservabilityConfigurations_RequestSyntax)**

A token from a previous result page. It's used for a paginated request. The request retrieves the next result page. All other parameter values must be
identical to the ones that are specified in the initial request.

If you don't specify `NextToken`, the request retrieves the first result page.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Pattern: `.*`

Required: No

**[ObservabilityConfigurationName](#API_ListObservabilityConfigurations_RequestSyntax)**

The name of the App Runner observability configuration that you want to list. If specified, App Runner lists revisions that share this name. If not specified,
App Runner returns revisions of all active configurations.

Type: String

Length Constraints: Minimum length of 4. Maximum length of 32.

Pattern: `[A-Za-z0-9][A-Za-z0-9\-_]{3,31}`

Required: No

## Response Syntax

```nohighlight

{
   "NextToken": "string",
   "ObservabilityConfigurationSummaryList": [
      {
         "ObservabilityConfigurationArn": "string",
         "ObservabilityConfigurationName": "string",
         "ObservabilityConfigurationRevision": number
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NextToken](#API_ListObservabilityConfigurations_ResponseSyntax)**

The token that you can pass in a subsequent request to get the next result page. It's returned in a paginated request.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Pattern: `.*`

**[ObservabilityConfigurationSummaryList](#API_ListObservabilityConfigurations_ResponseSyntax)**

A list of summary information records for observability configurations. In a paginated request, the request returns up to `MaxResults`
records for each call.

Type: Array of [ObservabilityConfigurationSummary](api-observabilityconfigurationsummary.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServiceErrorException**

An unexpected service exception occurred.

HTTP Status Code: 500

**InvalidRequestException**

One or more input parameters aren't valid. Refer to the API action's document page, correct the input parameters, and try the action again.

HTTP Status Code: 400

## Examples

### Paginated listing of App Runner observability configurations

This example illustrates how to list all App Runner observability configurations in your AWS account. Up to five observability
configurations are listed in each response. We set `LatestOnly` to `false` to get all revisions of all active
configurations.

In this example, the response includes two results and there aren't additional ones, so no `NextToken` is returned.

#### Sample Request

```json

$ aws apprunner list-observability-configurations --cli-input-json "`cat`"
{
  "LatestOnly": false,
  "MaxResults": 5
}
```

#### Sample Response

```json

{
  "ObservabilityConfigurationSummaryList": [
    {
      "ObservabilityConfigurationArn": "arn:aws:apprunner:us-east-1:123456789012:observabilityconfiguration/xray-tracing/1/2f50e7656d7819fead0f59672e68042e",
      "ObservabilityConfigurationName": "xray-tracing",
      "ObservabilityConfigurationRevision": 1
    },
    {
      "ObservabilityConfigurationArn": "arn:aws:apprunner:us-east-1:123456789012:observabilityconfiguration/xray-tracing/2/e76562f50d78042e819fead0f59672e6",
      "ObservabilityConfigurationName": "xray-tracing",
      "ObservabilityConfigurationRevision": 2
    }
  ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apprunner-2020-05-15/listobservabilityconfigurations.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apprunner-2020-05-15/listobservabilityconfigurations.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/listobservabilityconfigurations.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apprunner-2020-05-15/listobservabilityconfigurations.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/listobservabilityconfigurations.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apprunner-2020-05-15/listobservabilityconfigurations.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apprunner-2020-05-15/listobservabilityconfigurations.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apprunner-2020-05-15/listobservabilityconfigurations.md)

- [AWS SDK for Python](../../../goto/boto3/apprunner-2020-05-15/listobservabilityconfigurations.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/listobservabilityconfigurations.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListConnections

ListOperations

All content copied from https://docs.aws.amazon.com/.
