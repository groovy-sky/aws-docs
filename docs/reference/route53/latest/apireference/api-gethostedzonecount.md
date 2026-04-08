# GetHostedZoneCount

Retrieves the number of hosted zones that are associated with the current AWS account.

## Request Syntax

```

GET /2013-04-01/hostedzonecount HTTP/1.1

```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<GetHostedZoneCountResponse>
   <HostedZoneCount>long</HostedZoneCount>
</GetHostedZoneCountResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[GetHostedZoneCountResponse](#API_GetHostedZoneCount_ResponseSyntax)**

Root level tag for the GetHostedZoneCountResponse parameters.

Required: Yes

**[HostedZoneCount](#API_GetHostedZoneCount_ResponseSyntax)**

The total number of public and private hosted zones that are associated with the
current AWS account.

Type: Long

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

## Examples

### Example Request

This example illustrates one usage of GetHostedZoneCount.

```

GET /2013-04-01/hostedzonecount
```

### Example Response

This example illustrates one usage of GetHostedZoneCount.

```

HTTP/1.1 200 OK
<?xml version="1.0" encoding="UTF-8"?>
<GetHostedZoneCountResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <HostedZoneCount>42</HostedZoneCount>
</GetHostedZoneCountResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/gethostedzonecount.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/gethostedzonecount.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/gethostedzonecount.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/gethostedzonecount.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/gethostedzonecount.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/gethostedzonecount.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/gethostedzonecount.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/gethostedzonecount.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/gethostedzonecount.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/gethostedzonecount.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetHostedZone

GetHostedZoneLimit
