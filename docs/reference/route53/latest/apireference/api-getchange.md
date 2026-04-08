# GetChange

Returns the current status of a change batch request. The status is one of the
following values:

- `PENDING` indicates that the changes in this request have not
propagated to all Amazon Route 53 DNS servers managing the hosted zone. This is the initial status of all
change batch requests.

- `INSYNC` indicates that the changes have propagated to all Route 53
DNS servers managing the hosted zone.

## Request Syntax

```nohighlight

GET /2013-04-01/change/Id HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[Id](#API_GetChange_RequestSyntax)**

The ID of the change batch request. The value that you specify here is the value that
`ChangeResourceRecordSets` returned in the `Id` element when
you submitted the request.

Length Constraints: Minimum length of 1. Maximum length of 6500.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<GetChangeResponse>
   <ChangeInfo>
      <Comment>string</Comment>
      <Id>string</Id>
      <Status>string</Status>
      <SubmittedAt>timestamp</SubmittedAt>
   </ChangeInfo>
</GetChangeResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[GetChangeResponse](#API_GetChange_ResponseSyntax)**

Root level tag for the GetChangeResponse parameters.

Required: Yes

**[ChangeInfo](#API_GetChange_ResponseSyntax)**

A complex type that contains information about the specified change batch.

Type: [ChangeInfo](api-changeinfo.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

**NoSuchChange**

A change with the specified change ID does not exist.

HTTP Status Code: 404

## Examples

### Example Request

This example illustrates one usage of GetChange.

```

GET /2013-04-01/change/C2682N5HXP0BZ4
```

### Example Response

This example illustrates one usage of GetChange.

```

HTTP/1.1 200 OK
<?xml version="1.0" encoding="UTF-8"?>
<GetChangeResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <ChangeInfo>
      <Id>C2682N5HXP0BZ4</Id>
      <Status>INSYNC</Status>
      <SubmittedAt>2017-03-10T01:36:41.958Z</SubmittedAt>
   </ChangeInfo>
</GetChangeResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/getchange.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/getchange.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/getchange.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/getchange.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/getchange.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/getchange.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/getchange.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/getchange.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/getchange.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/getchange.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetAccountLimit

GetCheckerIpRanges
