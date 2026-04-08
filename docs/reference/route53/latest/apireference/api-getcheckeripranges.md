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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/getcheckeripranges.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/getcheckeripranges.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/getcheckeripranges.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/getcheckeripranges.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/getcheckeripranges.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/getcheckeripranges.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/getcheckeripranges.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/getcheckeripranges.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/getcheckeripranges.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/getcheckeripranges.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetChange

GetDNSSEC
