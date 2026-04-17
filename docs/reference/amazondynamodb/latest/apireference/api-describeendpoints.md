---
title: "DescribeEndpoints"
---

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/dynamodb-2012-08-10/DescribeEndpoints)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/dynamodb-2012-08-10/DescribeEndpoints)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/dynamodb-2012-08-10/DescribeEndpoints)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/dynamodb-2012-08-10/DescribeEndpoints)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/dynamodb-2012-08-10/DescribeEndpoints)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/dynamodb-2012-08-10/DescribeEndpoints)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/dynamodb-2012-08-10/DescribeEndpoints)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/dynamodb-2012-08-10/DescribeEndpoints)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/dynamodb-2012-08-10/DescribeEndpoints)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/dynamodb-2012-08-10/DescribeEndpoints)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeContributorInsights

DescribeExport

All content copied from https://docs.aws.amazon.com/.
