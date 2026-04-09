# ListOperations

###### Important

AWS App Runner will no longer be open to new
customers starting March 31, 2026. If you would like to use App Runner, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS App Runner availability change](../dg/apprunner-availability-change.md).

Return a list of operations that occurred on an AWS App Runner service.

The resulting list of [OperationSummary](api-operationsummary.md) objects is sorted in reverse chronological order. The first object on the list represents the
last started operation.

## Request Syntax

```nohighlight

{
   "MaxResults": number,
   "NextToken": "string",
   "ServiceArn": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[MaxResults](#API_ListOperations_RequestSyntax)**

The maximum number of results to include in each response (result page). It's used for a paginated request.

If you don't specify `MaxResults`, the request retrieves all available results in a single response.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 20.

Required: No

**[NextToken](#API_ListOperations_RequestSyntax)**

A token from a previous result page. It's used for a paginated request. The request retrieves the next result page. All other parameter values must be
identical to the ones specified in the initial request.

If you don't specify `NextToken`, the request retrieves the first result page.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 51200.

Pattern: `.*`

Required: No

**[ServiceArn](#API_ListOperations_RequestSyntax)**

The Amazon Resource Name (ARN) of the App Runner service that you want a list of operations for.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

Required: Yes

## Response Syntax

```nohighlight

{
   "NextToken": "string",
   "OperationSummaryList": [
      {
         "EndedAt": number,
         "Id": "string",
         "StartedAt": number,
         "Status": "string",
         "TargetArn": "string",
         "Type": "string",
         "UpdatedAt": number
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NextToken](#API_ListOperations_ResponseSyntax)**

The token that you can pass in a subsequent request to get the next result page. It's returned in a paginated request.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 51200.

Pattern: `.*`

**[OperationSummaryList](#API_ListOperations_ResponseSyntax)**

A list of operation summary information records. In a paginated request, the request returns up to `MaxResults` records for each
call.

Type: Array of [OperationSummary](api-operationsummary.md) objects

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

### List operations that occurred on a service

This example illustrates how to list all operations that occurred on an App Runner service so far. In this example, the service is new and only a single
operation of type `CREATE_SERVICE` has occurred.

#### Sample Request

```json

$ aws apprunner list-operations --cli-input-json "`cat`"
{
  "ServiceArn": "arn:aws:apprunner:us-east-1:123456789012:service/python-app/8fe1e10304f84fd2b0df550fe98a71fa"
}
```

#### Sample Response

```json

{
  "OperationSummaryList": [
    {
      "EndedAt": 1606156217,
      "Id": "17fe9f55-7e91-4097-b243-fcabbb69a4cf",
      "StartedAt": 1606156014,
      "Status": "SUCCEEDED",
      "TargetArn": "arn:aws:apprunner:us-east-1:123456789012:service/python-app/8fe1e10304f84fd2b0df550fe98a71fa",
      "Type": "CREATE_SERVICE",
      "UpdatedAt": 1606156217
    }
  ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apprunner-2020-05-15/listoperations.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apprunner-2020-05-15/listoperations.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/listoperations.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apprunner-2020-05-15/listoperations.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/listoperations.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apprunner-2020-05-15/listoperations.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apprunner-2020-05-15/listoperations.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apprunner-2020-05-15/listoperations.md)

- [AWS SDK for Python](../../../goto/boto3/apprunner-2020-05-15/listoperations.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/listoperations.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListObservabilityConfigurations

ListServices

All content copied from https://docs.aws.amazon.com/.
