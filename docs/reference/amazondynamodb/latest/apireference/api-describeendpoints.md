# DescribeEndpoints

Returns the regional endpoint information. For more information on policy permissions,
please see [Internetwork traffic privacy](../../../../services/dynamodb/latest/developerguide/inter-network-traffic-privacy.md#inter-network-traffic-DescribeEndpoints).

## Response Syntax

```nohighlight

{
   "Endpoints": [
      {
         "Address": "string",
         "CachePeriodInMinutes": number
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Endpoints](#API_DescribeEndpoints_ResponseSyntax)**

List of endpoints.

Type: Array of [Endpoint](api-endpoint.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dynamodb-2012-08-10/describeendpoints.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dynamodb-2012-08-10/describeendpoints.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/describeendpoints.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dynamodb-2012-08-10/describeendpoints.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/describeendpoints.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dynamodb-2012-08-10/describeendpoints.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dynamodb-2012-08-10/describeendpoints.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dynamodb-2012-08-10/describeendpoints.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dynamodb-2012-08-10/describeendpoints.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/describeendpoints.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeContributorInsights

DescribeExport

All content copied from https://docs.aws.amazon.com/.
