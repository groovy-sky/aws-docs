# GetAccountLimit

Gets the specified limit for the current account, for example, the maximum number of
health checks that you can create using the account.

For the default limit, see [Limits](../../../../services/route53/latest/developerguide/dnslimitations.md) in the
_Amazon Route 53 Developer Guide_. To request a higher limit,
[open a case](https://console.aws.amazon.com/support/home).

###### Note

You can also view account limits in AWS Trusted Advisor. Sign in to
the AWS Management Console and open the Trusted Advisor console at [https://console.aws.amazon.com/trustedadvisor/](https://console.aws.amazon.com/trustedadvisor). Then choose **Service limits** in the navigation pane.

## Request Syntax

```nohighlight

GET /2013-04-01/accountlimit/Type HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[Type](#API_GetAccountLimit_RequestSyntax)**

The limit that you want to get. Valid values include the following:

- **MAX\_HEALTH\_CHECKS\_BY\_OWNER**: The maximum
number of health checks that you can create using the current account.

- **MAX\_HOSTED\_ZONES\_BY\_OWNER**: The maximum number
of hosted zones that you can create using the current account.

- **MAX\_REUSABLE\_DELEGATION\_SETS\_BY\_OWNER**: The
maximum number of reusable delegation sets that you can create using the current
account.

- **MAX\_TRAFFIC\_POLICIES\_BY\_OWNER**: The maximum
number of traffic policies that you can create using the current account.

- **MAX\_TRAFFIC\_POLICY\_INSTANCES\_BY\_OWNER**: The
maximum number of traffic policy instances that you can create using the current
account. (Traffic policy instances are referred to as traffic flow policy
records in the Amazon Route 53 console.)

Valid Values: `MAX_HEALTH_CHECKS_BY_OWNER | MAX_HOSTED_ZONES_BY_OWNER | MAX_TRAFFIC_POLICY_INSTANCES_BY_OWNER | MAX_REUSABLE_DELEGATION_SETS_BY_OWNER | MAX_TRAFFIC_POLICIES_BY_OWNER`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<GetAccountLimitResponse>
   <Count>long</Count>
   <Limit>
      <Type>string</Type>
      <Value>long</Value>
   </Limit>
</GetAccountLimitResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[GetAccountLimitResponse](#API_GetAccountLimit_ResponseSyntax)**

Root level tag for the GetAccountLimitResponse parameters.

Required: Yes

**[Count](#API_GetAccountLimit_ResponseSyntax)**

The current number of entities that you have created of the specified type. For
example, if you specified `MAX_HEALTH_CHECKS_BY_OWNER` for the value of
`Type` in the request, the value of `Count` is the current
number of health checks that you have created using the current account.

Type: Long

Valid Range: Minimum value of 0.

**[Limit](#API_GetAccountLimit_ResponseSyntax)**

The current setting for the specified limit. For example, if you specified
`MAX_HEALTH_CHECKS_BY_OWNER` for the value of `Type` in the
request, the value of `Limit` is the maximum number of health checks that you
can create using the current account.

Type: [AccountLimit](https://docs.aws.amazon.com/Route53/latest/APIReference/API_AccountLimit.html) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/route53-2013-04-01/GetAccountLimit)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/route53-2013-04-01/GetAccountLimit)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53-2013-04-01/GetAccountLimit)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/route53-2013-04-01/GetAccountLimit)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53-2013-04-01/GetAccountLimit)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/route53-2013-04-01/GetAccountLimit)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/route53-2013-04-01/GetAccountLimit)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/route53-2013-04-01/GetAccountLimit)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/route53-2013-04-01/GetAccountLimit)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53-2013-04-01/GetAccountLimit)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EnableHostedZoneDNSSEC

GetChange
