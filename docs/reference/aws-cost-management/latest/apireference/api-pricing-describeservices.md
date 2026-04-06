# DescribeServices

Returns the metadata for one service or a list of the metadata for all services. Use
this without a service code to get the service codes for all services.
Use it with a service code, such as `AmazonEC2`, to get information specific to
that service, such as the attribute
names available for that service. For example, some of the attribute names available for EC2 are
`volumeType`, `maxIopsVolume`, `operation`,
`locationType`, and `instanceCapacity10xlarge`.

## Request Syntax

```nohighlight

{
   "FormatVersion": "string",
   "MaxResults": number,
   "NextToken": "string",
   "ServiceCode": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/CommonParameters.html).

The request accepts the following data in JSON format.

**[FormatVersion](#API_pricing_DescribeServices_RequestSyntax)**

The format version that you want the response to be in.

Valid values are: `aws_v1`

Type: String

Length Constraints: Minimum length of 1. Maximum length of 32.

Required: No

**[MaxResults](#API_pricing_DescribeServices_RequestSyntax)**

The maximum number of results that you want returned in the response.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

**[NextToken](#API_pricing_DescribeServices_RequestSyntax)**

The pagination token that indicates the next set of results that you want to retrieve.

Type: String

Required: No

**[ServiceCode](#API_pricing_DescribeServices_RequestSyntax)**

The code for the service whose information you want to retrieve, such as `AmazonEC2`.
You can use
the `ServiceCode` to filter the results in a `GetProducts` call.
To retrieve a list of all services, leave this blank.

Type: String

Required: No

## Response Syntax

```nohighlight

{
   "FormatVersion": "string",
   "NextToken": "string",
   "Services": [
      {
         "AttributeNames": [ "string" ],
         "ServiceCode": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[FormatVersion](#API_pricing_DescribeServices_ResponseSyntax)**

The format version of the response. For example, `aws_v1`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 32.

**[NextToken](#API_pricing_DescribeServices_ResponseSyntax)**

The pagination token for the next set of retrievable results.

Type: String

**[Services](#API_pricing_DescribeServices_ResponseSyntax)**

The service metadata for the service or services in the response.

Type: Array of [Service](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_pricing_Service.html) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/CommonErrors.html).

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

### The following is a sample request and response of the GetService operation.

This example illustrates one usage of DescribeServices.

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
X-Amz-Target: AWSPriceListService.DescribeServices
{
    "ServiceCode": "AmazonEC2",
    "FormatVersion": "aws_v1",
    "NextToken": null,
    "MaxResults": 1
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
    "FormatVersion": "aws_v1",
    "NextToken": "abcdefg123",
    "Services": [
        {
            "AttributeNames": [
                "volumeType",
                "maxIopsvolume",
                "instanceCapacity10xlarge",
                "locationType",
                "operation"
            ],
            "ServiceCode": "AmazonEC2"
        }
    ]
}

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/pricing-2017-10-15/DescribeServices)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/pricing-2017-10-15/DescribeServices)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/pricing-2017-10-15/DescribeServices)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/pricing-2017-10-15/DescribeServices)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/pricing-2017-10-15/DescribeServices)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/pricing-2017-10-15/DescribeServices)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/pricing-2017-10-15/DescribeServices)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/pricing-2017-10-15/DescribeServices)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/pricing-2017-10-15/DescribeServices)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/pricing-2017-10-15/DescribeServices)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Price List

GetAttributeValues
