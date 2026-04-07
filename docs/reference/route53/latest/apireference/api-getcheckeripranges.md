# GetCheckerIpRanges

Route 53 does not perform authorization for this API because it retrieves information
that is already available to the public.

###### Important

`GetCheckerIpRanges` still works, but we recommend that you download
ip-ranges.json, which includes IP address ranges for all AWS
services. For more information, see [IP Address Ranges\
of Amazon Route 53 Servers](../../../../services/route53/latest/developerguide/route-53-ip-addresses.md) in the _Amazon Route 53 Developer_
_Guide_.

## Request Syntax

```

GET /2013-04-01/checkeripranges HTTP/1.1

```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<GetCheckerIpRangesResponse>
   <CheckerIpRanges>
      <member>string</member>
   </CheckerIpRanges>
</GetCheckerIpRangesResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[GetCheckerIpRangesResponse](#API_GetCheckerIpRanges_ResponseSyntax)**

Root level tag for the GetCheckerIpRangesResponse parameters.

Required: Yes

**[CheckerIpRanges](#API_GetCheckerIpRanges_ResponseSyntax)**

A complex type that contains sorted list of IP ranges in CIDR format for Amazon Route
53 health checkers.

Type: Array of strings

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/route53-2013-04-01/GetCheckerIpRanges)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/route53-2013-04-01/GetCheckerIpRanges)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53-2013-04-01/GetCheckerIpRanges)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/route53-2013-04-01/GetCheckerIpRanges)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53-2013-04-01/GetCheckerIpRanges)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/route53-2013-04-01/GetCheckerIpRanges)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/route53-2013-04-01/GetCheckerIpRanges)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/route53-2013-04-01/GetCheckerIpRanges)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/route53-2013-04-01/GetCheckerIpRanges)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53-2013-04-01/GetCheckerIpRanges)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetChange

GetDNSSEC
