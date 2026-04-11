# ListVPCAssociationAuthorizations

Gets a list of the VPCs that were created by other accounts and that can be associated
with a specified hosted zone because you've submitted one or more
`CreateVPCAssociationAuthorization` requests.

The response includes a `VPCs` element with a `VPC` child
element for each VPC that can be associated with the hosted zone.

## Request Syntax

```nohighlight

GET /2013-04-01/hostedzone/Id/authorizevpcassociation?maxresults=MaxResults&nexttoken=NextToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[Id](#API_ListVPCAssociationAuthorizations_RequestSyntax)**

The ID of the hosted zone for which you want a list of VPCs that can be associated
with the hosted zone.

Length Constraints: Maximum length of 32.

Required: Yes

**[maxresults](#API_ListVPCAssociationAuthorizations_RequestSyntax)**

_Optional_: An integer that specifies the maximum number of VPCs
that you want Amazon Route 53 to return. If you don't specify a value for
`MaxResults`, Route 53 returns up to 50 VPCs per page.

**[nexttoken](#API_ListVPCAssociationAuthorizations_RequestSyntax)**

_Optional_: If a response includes a `NextToken`
element, there are more VPCs that can be associated with the specified hosted zone. To
get the next page of results, submit another request, and include the value of
`NextToken` from the response in the `nexttoken` parameter in
another `ListVPCAssociationAuthorizations` request.

Length Constraints: Maximum length of 1024.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<ListVPCAssociationAuthorizationsResponse>
   <HostedZoneId>string</HostedZoneId>
   <NextToken>string</NextToken>
   <VPCs>
      <VPC>
         <VPCId>string</VPCId>
         <VPCRegion>string</VPCRegion>
      </VPC>
   </VPCs>
</ListVPCAssociationAuthorizationsResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[ListVPCAssociationAuthorizationsResponse](#API_ListVPCAssociationAuthorizations_ResponseSyntax)**

Root level tag for the ListVPCAssociationAuthorizationsResponse parameters.

Required: Yes

**[HostedZoneId](#API_ListVPCAssociationAuthorizations_ResponseSyntax)**

The ID of the hosted zone that you can associate the listed VPCs with.

Type: String

Length Constraints: Maximum length of 32.

**[NextToken](#API_ListVPCAssociationAuthorizations_ResponseSyntax)**

When the response includes a `NextToken` element, there are more VPCs that
can be associated with the specified hosted zone. To get the next page of VPCs, submit
another `ListVPCAssociationAuthorizations` request, and include the value of
the `NextToken` element from the response in the `nexttoken`
request parameter.

Type: String

Length Constraints: Maximum length of 1024.

**[VPCs](#API_ListVPCAssociationAuthorizations_ResponseSyntax)**

The list of VPCs that are authorized to be associated with the specified hosted
zone.

Type: Array of [VPC](api-vpc.md) objects

Array Members: Minimum number of 1 item.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

**InvalidPaginationToken**

The value that you specified to get the second or subsequent page of results is
invalid.

HTTP Status Code: 400

**NoSuchHostedZone**

No hosted zone exists with the ID that you specified.

**message**

HTTP Status Code: 404

## Examples

### Example Request

This example illustrates one usage of ListVPCAssociationAuthorizations.

```

GET /2013-04-01/hostedzone/Z1PA6795UKMFR9/authorizevpcassociation&maxresults=1 HTTP/1.1
```

### Example Response

This example illustrates one usage of ListVPCAssociationAuthorizations.

```

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<ListVPCAssociationAuthorizationsResponse>
   <HostedZoneId>Z1PA6795UKMFR9</HostedZoneId>
   <NextToken>Z222222VVVVVVV</NextToken>
   <VPCs>
      <VPC>
         <VPCId>vpc-a1b2c3d4e5</VPCId>
         <VPCRegion>us-east-2</VPCRegion>
      </VPC>
   </VPCs>
</ListVPCAssociationAuthorizationsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/listvpcassociationauthorizations.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/listvpcassociationauthorizations.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/listvpcassociationauthorizations.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/listvpcassociationauthorizations.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/listvpcassociationauthorizations.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/listvpcassociationauthorizations.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/listvpcassociationauthorizations.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/listvpcassociationauthorizations.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/listvpcassociationauthorizations.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/listvpcassociationauthorizations.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListTrafficPolicyVersions

TestDNSAnswer
