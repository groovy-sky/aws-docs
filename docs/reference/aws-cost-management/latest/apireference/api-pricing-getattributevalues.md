# GetAttributeValues

Returns a list of attribute values. Attributes are similar to the details
in a Price List API offer file. For a list of available attributes, see
[Offer File Definitions](../../../../services/awsaccountbilling/latest/aboutv2/reading-an-offer-pps-defs.md)
in the [AWS Billing and Cost Management User Guide](../../../../services/awsaccountbilling/latest/aboutv2/billing-what-is.md).

## Request Syntax

```nohighlight

{
   "AttributeName": "string",
   "MaxResults": number,
   "NextToken": "string",
   "ServiceCode": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[AttributeName](#API_pricing_GetAttributeValues_RequestSyntax)**

The name of the attribute that you want to retrieve the values for, such as `volumeType`.

Type: String

Required: Yes

**[MaxResults](#API_pricing_GetAttributeValues_RequestSyntax)**

The maximum number of results to return in response.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

**[NextToken](#API_pricing_GetAttributeValues_RequestSyntax)**

The pagination token that indicates the next set of results that you want to retrieve.

Type: String

Required: No

**[ServiceCode](#API_pricing_GetAttributeValues_RequestSyntax)**

The service code for the service whose attributes you want to retrieve. For example, if you want
the retrieve an EC2 attribute, use `AmazonEC2`.

Type: String

Required: Yes

## Response Syntax

```nohighlight

{
   "AttributeValues": [
      {
         "Value": "string"
      }
   ],
   "NextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[AttributeValues](#API_pricing_GetAttributeValues_ResponseSyntax)**

The list of values for an attribute. For example, `Throughput Optimized HDD` and
`Provisioned IOPS` are two available values for the `AmazonEC2` `volumeType`.

Type: Array of [AttributeValue](api-pricing-attributevalue.md) objects

**[NextToken](#API_pricing_GetAttributeValues_ResponseSyntax)**

The pagination token that indicates the next set of results to retrieve.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ExpiredNextTokenException**

The pagination token expired. Try again without a pagination token.

HTTP Status Code: 400

**InternalErrorException**

An error on the server occurred during the processing of your request. Try again later.

HTTP Status Code: 500

**InvalidNextTokenException**

The pagination token is invalid. Try again without a pagination token.

HTTP Status Code: 400

**InvalidParameterException**

One or more parameters had an invalid value.

HTTP Status Code: 400

**NotFoundException**

The requested resource can't be found.

HTTP Status Code: 400

**ThrottlingException**

You've made too many requests exceeding service quotas.

HTTP Status Code: 400

## Examples

### The following is a sample request and response of the GetAttributeValues operation.

This example illustrates one usage of GetAttributeValues.

#### Sample Request

```

POST / HTTP/1.1
Host: api.pricing.<region>.<domain>
x-amz-Date: <Date>
Authorization: AWS4-HMAC-SHA256 Credential=<Credential>, SignedHeaders=contenttype;date;host;user-agent;x-amz-date;x-amz-target;x-amzn-requestid,Signature=<Signature>
User-Agent: <UserAgentString>
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Connection: Keep-Alive
X-Amz-Target: AWSPriceListService.GetAttributeValues
{
    "ServiceCode": "AmazonEC2",
    "AttributeName": "volumeType",
    "NextToken": null,
    "MaxResults": 2
}

```

#### Sample Response

```

HTTP/1.1 200 OK
x-amzn-RequestId: <RequestId>
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Date: <Date>
{
    "AttributeValues": [
        {
            "Value": "Throughput Optimized HDD"
        },
        {
            "Value": "Provisioned IOPS"
        }
    ],
    "NextToken": "GpgauTGIY7LGezucl5LV0w==:7GzYJ0nw0DBTJ2J66EoTIIynE6O1uXwQtTRqioJzQadBnDVgHPzI1en4BUQnPCLpzeBk9RQQAWaFieA4+DapFAGLgk+Z/9/cTw9GldnPOHN98+FdmJP7wKU3QQpQ8MQr5KOeBkIsAqvAQYdL0DkL7tHwPtE5iCEByAmg9gcC/yBU1vAOsf7R3VaNN4M5jMDv3woSWqASSIlBVB6tgW78YL22KhssoItM/jWW+aP6Jqtq4mldxp/ct6DWAl+xLFwHU/CbketimPPXyqHF3/UXDw=="
}

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/pricing-2017-10-15/getattributevalues.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/pricing-2017-10-15/getattributevalues.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/pricing-2017-10-15/getattributevalues.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/pricing-2017-10-15/getattributevalues.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/pricing-2017-10-15/getattributevalues.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/pricing-2017-10-15/getattributevalues.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/pricing-2017-10-15/getattributevalues.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/pricing-2017-10-15/getattributevalues.md)

- [AWS SDK for Python](../../../../services/goto/boto3/pricing-2017-10-15/getattributevalues.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/pricing-2017-10-15/getattributevalues.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeServices

GetPriceListFileUrl
