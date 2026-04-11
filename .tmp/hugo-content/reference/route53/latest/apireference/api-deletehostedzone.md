# DeleteHostedZone

Deletes a hosted zone.

If the hosted zone was created by another service, such as AWS Cloud Map, see
[Deleting Public Hosted Zones That Were Created by Another Service](../../../../services/route53/latest/developerguide/deletehostedzone.md#delete-public-hosted-zone-created-by-another-service) in the
_Amazon Route 53 Developer Guide_ for information
about how to delete it. (The process is the same for public and private hosted zones
that were created by another service.)

If you want to keep your domain registration but you want to stop routing internet
traffic to your website or web application, we recommend that you delete resource record
sets in the hosted zone instead of deleting the hosted zone.

###### Important

If you delete a hosted zone, you can't undelete it. You must create a new hosted
zone and update the name servers for your domain registration, which can require up
to 48 hours to take effect. (If you delegated responsibility for a subdomain to a
hosted zone and you delete the child hosted zone, you must update the name servers
in the parent hosted zone.) In addition, if you delete a hosted zone, someone could
hijack the domain and route traffic to their own resources using your domain
name.

If you want to avoid the monthly charge for the hosted zone, you can transfer DNS
service for the domain to a free DNS service. When you transfer DNS service, you have to
update the name servers for the domain registration. If the domain is registered with
Route 53, see [UpdateDomainNameservers](api-domains-updatedomainnameservers.md) for information about how to replace Route 53 name servers with name servers for the new DNS service. If the domain is
registered with another registrar, use the method provided by the registrar to update
name servers for the domain registration. For more information, perform an internet
search on "free DNS service."

You can delete a hosted zone only if it contains only the default SOA and NS records
and has DNSSEC signing disabled. If the hosted zone contains other records or has DNSSEC
enabled, you must delete the records and disable DNSSEC before deletion. Attempting to
delete a hosted zone with additional records or DNSSEC enabled returns a
`HostedZoneNotEmpty` error. For information about deleting records, see
[ChangeResourceRecordSets](api-changeresourcerecordsets.md).

To verify that the hosted zone has been deleted, do one of the following:

- Use the `GetHostedZone` action to request information about the
hosted zone.

- Use the `ListHostedZones` action to get a list of the hosted zones
associated with the current AWS account.

## Request Syntax

```nohighlight

DELETE /2013-04-01/hostedzone/Id HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[Id](#API_DeleteHostedZone_RequestSyntax)**

The ID of the hosted zone you want to delete.

Length Constraints: Maximum length of 32.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<DeleteHostedZoneResponse>
   <ChangeInfo>
      <Comment>string</Comment>
      <Id>string</Id>
      <Status>string</Status>
      <SubmittedAt>timestamp</SubmittedAt>
   </ChangeInfo>
</DeleteHostedZoneResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[DeleteHostedZoneResponse](#API_DeleteHostedZone_ResponseSyntax)**

Root level tag for the DeleteHostedZoneResponse parameters.

Required: Yes

**[ChangeInfo](#API_DeleteHostedZone_ResponseSyntax)**

A complex type that contains the ID, the status, and the date and time of a request to
delete a hosted zone.

Type: [ChangeInfo](api-changeinfo.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**HostedZoneNotEmpty**

The hosted zone contains resource records that are not SOA or NS records.

**message**

HTTP Status Code: 400

**InvalidDomainName**

The specified domain name is not valid.

**message**

HTTP Status Code: 400

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

**NoSuchHostedZone**

No hosted zone exists with the ID that you specified.

**message**

HTTP Status Code: 404

**PriorRequestNotComplete**

If Amazon Route 53 can't process a request before the next request arrives, it will
reject subsequent requests for the same hosted zone and return an `HTTP 400
				error` ( `Bad request`). If Route 53 returns this error repeatedly
for the same request, we recommend that you wait, in intervals of increasing duration,
before you try the request again.

HTTP Status Code: 400

## Examples

### Example Request

This example illustrates one usage of DeleteHostedZone.

```

DELETE /2013-04-01/hostedzone/Z1PA6795UKMFR9
```

### Example Response

This example illustrates one usage of DeleteHostedZone.

```

HTTP/1.1 200 OK
<?xml version="1.0" encoding="UTF-8"?>
<DeleteHostedZoneResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <ChangeInfo>
      <Id>/change/C1PA6795UKMFR9</Id>
      <Status>PENDING</Status>
      <SubmittedAt>2017-03-10T01:36:41.958Z</SubmittedAt>
   </ChangeInfo>
</DeleteHostedZoneResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/deletehostedzone.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/deletehostedzone.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/deletehostedzone.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/deletehostedzone.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/deletehostedzone.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/deletehostedzone.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/deletehostedzone.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/deletehostedzone.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/deletehostedzone.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/deletehostedzone.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteHealthCheck

DeleteKeySigningKey
