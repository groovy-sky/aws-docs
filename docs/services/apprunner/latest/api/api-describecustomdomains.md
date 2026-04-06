# DescribeCustomDomains

###### Important

AWS App Runner will no longer be open to new
customers starting March 31, 2026. If you would like to use App Runner, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS App Runner availability change](../dg/apprunner-availability-change.md).

Return a description of custom domain names that are associated with an AWS App Runner service.

## Request Syntax

```nohighlight

{
   "MaxResults": number,
   "NextToken": "string",
   "ServiceArn": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](https://docs.aws.amazon.com/apprunner/latest/api/CommonParameters.html).

The request accepts the following data in JSON format.

**[MaxResults](#API_DescribeCustomDomains_RequestSyntax)**

The maximum number of results that each response (result page) can include. It's used for a paginated request.

If you don't specify `MaxResults`, the request retrieves all available results in a single response.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 5.

Required: No

**[NextToken](#API_DescribeCustomDomains_RequestSyntax)**

A token from a previous result page. It's used for a paginated request. The request retrieves the next result page. All other parameter values must be
identical to the ones that are specified in the initial request.

If you don't specify `NextToken`, the request retrieves the first result page.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 51200.

Pattern: `.*`

Required: No

**[ServiceArn](#API_DescribeCustomDomains_RequestSyntax)**

The Amazon Resource Name (ARN) of the App Runner service that you want associated custom domain names to be described for.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

Required: Yes

## Response Syntax

```nohighlight

{
   "CustomDomains": [
      {
         "CertificateValidationRecords": [
            {
               "Name": "string",
               "Status": "string",
               "Type": "string",
               "Value": "string"
            }
         ],
         "DomainName": "string",
         "EnableWWWSubdomain": boolean,
         "Status": "string"
      }
   ],
   "DNSTarget": "string",
   "NextToken": "string",
   "ServiceArn": "string",
   "VpcDNSTargets": [
      {
         "DomainName": "string",
         "VpcId": "string",
         "VpcIngressConnectionArn": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[CustomDomains](#API_DescribeCustomDomains_ResponseSyntax)**

A list of descriptions of custom domain names that are associated with the service. In a paginated request, the request returns up to
`MaxResults` records per call.

Type: Array of [CustomDomain](api-customdomain.md) objects

**[DNSTarget](#API_DescribeCustomDomains_ResponseSyntax)**

The App Runner subdomain of the App Runner service. The associated custom domain names are mapped to this target name.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 51200.

Pattern: `.*`

**[NextToken](#API_DescribeCustomDomains_ResponseSyntax)**

The token that you can pass in a subsequent request to get the next result page. It's returned in a paginated request.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 51200.

Pattern: `.*`

**[ServiceArn](#API_DescribeCustomDomains_ResponseSyntax)**

The Amazon Resource Name (ARN) of the App Runner service whose associated custom domain names you want to describe.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

**[VpcDNSTargets](#API_DescribeCustomDomains_ResponseSyntax)**

DNS Target records for the custom domains of this Amazon VPC.

Type: Array of [VpcDNSTarget](https://docs.aws.amazon.com/apprunner/latest/api/API_VpcDNSTarget.html) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](https://docs.aws.amazon.com/apprunner/latest/api/CommonErrors.html).

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

### Get descriptions of custom domain names associated with a service

This example illustrates how to get descriptions and status of the custom domain names associated with an App Runner service.

#### Sample Request

```json

$ aws apprunner describe-custom-domains --cli-input-json "`cat`"
{
  "ServiceArn": "arn:aws:apprunner:us-east-1:123456789012:service/python-app/8fe1e10304f84fd2b0df550fe98a71fa"
}
```

#### Sample Response

```json

{
  "CustomDomains": [
    {
      "CertificateValidationRecords": [
        {
          "Name": "_70d3f50a94f7c72dc28784cf55db2f6b.example.com",
          "Status": "PENDING_VALIDATION",
          "Type": "CNAME",
          "Value": "_1270c137383c6307b6832db02504c4b0.bsgbmzkfwj.acm-validations.aws."
        },
        {
          "Name": "_287870d3f50a94f7c72dc4cf55db2f6b.www.example.com",
          "Status": "PENDING_VALIDATION",
          "Type": "CNAME",
          "Value": "_832db01270c137383c6307b62504c4b0.mzkbsgbfwj.acm-validations.aws."
        }
      ],
      "DomainName": "example.com",
      "EnableWWWSubdomain": true,
      "Status": "PENDING_CERTIFICATE_DNS_VALIDATION"
    },
    {
      "CertificateValidationRecords": [
        {
          "Name": "_a94f784c70d3f507c72dc28f55db2f6b.deals.example.com",
          "Status": "SUCCESS",
          "Type": "CNAME",
          "Value": "_2db02504c1270c137383c6307b6834b0.bsgbmzkfwj.acm-validations.aws."
        }
      ],
      "DomainName": "deals.example.com",
      "EnableWWWSubdomain": false,
      "Status": "ACTIVE"
    }
  ],
  "DNSTarget": "psbqam834h.us-east-1.awsapprunner.com",
  "ServiceArn": "arn:aws:apprunner:us-east-1:123456789012:service/python-app/8fe1e10304f84fd2b0df550fe98a71fa",
  "VpcDNSTargets": [
      {
         "DomainName": "psbqam834h.us-east-1.awsapprunner.com",
         "VpcId": "vpc-4a5b6c7d",
         "VpcIngressConnectionArn": "arn:aws:apprunner:us-east-1:123456789012:vpcingressconnection/my-ingress-connection-name/3f2eb10e2c494674952026f646844e3d"
      }
   ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/apprunner-2020-05-15/DescribeCustomDomains)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/apprunner-2020-05-15/DescribeCustomDomains)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apprunner-2020-05-15/DescribeCustomDomains)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/apprunner-2020-05-15/DescribeCustomDomains)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apprunner-2020-05-15/DescribeCustomDomains)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/apprunner-2020-05-15/DescribeCustomDomains)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/apprunner-2020-05-15/DescribeCustomDomains)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/apprunner-2020-05-15/DescribeCustomDomains)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/apprunner-2020-05-15/DescribeCustomDomains)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apprunner-2020-05-15/DescribeCustomDomains)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DescribeAutoScalingConfiguration

DescribeObservabilityConfiguration
