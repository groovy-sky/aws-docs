# GetReusableDelegationSetLimit

Gets the maximum number of hosted zones that you can associate with the specified
reusable delegation set.

For the default limit, see [Limits](../../../../services/route53/latest/developerguide/dnslimitations.md) in the
_Amazon Route 53 Developer Guide_. To request a higher limit,
[open a case](https://console.aws.amazon.com/support/home).

## Request Syntax

```nohighlight

GET /2013-04-01/reusabledelegationsetlimit/Id/Type HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[Id](#API_GetReusableDelegationSetLimit_RequestSyntax)**

The ID of the delegation set that you want to get the limit for.

Length Constraints: Maximum length of 32.

Required: Yes

**[Type](#API_GetReusableDelegationSetLimit_RequestSyntax)**

Specify `MAX_ZONES_BY_REUSABLE_DELEGATION_SET` to get the maximum number of
hosted zones that you can associate with the specified reusable delegation set.

Valid Values: `MAX_ZONES_BY_REUSABLE_DELEGATION_SET`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<GetReusableDelegationSetLimitResponse>
   <Count>long</Count>
   <Limit>
      <Type>string</Type>
      <Value>long</Value>
   </Limit>
</GetReusableDelegationSetLimitResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[GetReusableDelegationSetLimitResponse](#API_GetReusableDelegationSetLimit_ResponseSyntax)**

Root level tag for the GetReusableDelegationSetLimitResponse parameters.

Required: Yes

**[Count](#API_GetReusableDelegationSetLimit_ResponseSyntax)**

The current number of hosted zones that you can associate with the specified reusable
delegation set.

Type: Long

Valid Range: Minimum value of 0.

**[Limit](#API_GetReusableDelegationSetLimit_ResponseSyntax)**

The current setting for the limit on hosted zones that you can associate with the
specified reusable delegation set.

Type: [ReusableDelegationSetLimit](api-reusabledelegationsetlimit.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

**NoSuchDelegationSet**

A reusable delegation set with the specified ID does not exist.

**message**

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/getreusabledelegationsetlimit.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/getreusabledelegationsetlimit.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/getreusabledelegationsetlimit.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/getreusabledelegationsetlimit.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/getreusabledelegationsetlimit.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/getreusabledelegationsetlimit.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/getreusabledelegationsetlimit.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/getreusabledelegationsetlimit.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/getreusabledelegationsetlimit.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/getreusabledelegationsetlimit.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetReusableDelegationSet

GetTrafficPolicy
