# ListHostedZonesByVPC

Lists all the private hosted zones that a specified VPC is associated with, regardless
of which AWS account or AWS service owns the hosted zones.
The `HostedZoneOwner` structure in the response contains one of the following
values:

- An `OwningAccount` element, which contains the account number of
either the current AWS account or another AWS account. Some services, such as AWS Cloud Map, create
hosted zones using the current account.

- An `OwningService` element, which identifies the AWS
service that created and owns the hosted zone. For example, if a hosted zone was
created by Amazon Elastic File System (Amazon EFS), the value of
`Owner` is `efs.amazonaws.com`.

`ListHostedZonesByVPC` returns the hosted zones associated with the specified VPC and does not reflect the hosted zone
associations to VPCs via Route 53 Profiles. To get the associations to a Profile, call the [ListProfileResourceAssociations](api-route53profiles-listprofileresourceassociations.md) API.

###### Note

When listing private hosted zones, the hosted zone and the Amazon VPC must
belong to the same partition where the hosted zones were created. A partition is a
group of AWS Regions. Each AWS account is scoped to
one partition.

The following are the supported partitions:

- `aws` \- AWS Regions

- `aws-cn` \- China Regions

- `aws-us-gov` \- AWS GovCloud (US) Region

For more information, see [Access Management](../../../../general/general/latest/gr/aws-arns-and-namespaces.md)
in the _AWS General Reference_.

## Request Syntax

```nohighlight

GET /2013-04-01/hostedzonesbyvpc?maxitems=MaxItems&nexttoken=NextToken&vpcid=VPCId&vpcregion=VPCRegion HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[maxitems](#API_ListHostedZonesByVPC_RequestSyntax)**

(Optional) The maximum number of hosted zones that you want Amazon Route 53 to return.
If the specified VPC is associated with more than `MaxItems` hosted zones,
the response includes a `NextToken` element. `NextToken` contains
an encrypted token that identifies the first hosted zone that Route 53 will return if
you submit another request.

**[nexttoken](#API_ListHostedZonesByVPC_RequestSyntax)**

If the previous response included a `NextToken` element, the specified VPC
is associated with more hosted zones. To get more hosted zones, submit another
`ListHostedZonesByVPC` request.

For the value of `NextToken`, specify the value of `NextToken`
from the previous response.

If the previous response didn't include a `NextToken` element, there are no
more hosted zones to get.

Length Constraints: Maximum length of 1024.

**[vpcid](#API_ListHostedZonesByVPC_RequestSyntax)**

The ID of the Amazon VPC that you want to list hosted zones for.

Length Constraints: Maximum length of 1024.

Required: Yes

**[vpcregion](#API_ListHostedZonesByVPC_RequestSyntax)**

For the Amazon VPC that you specified for `VPCId`, the AWS
Region that you created the VPC in.

Length Constraints: Minimum length of 1. Maximum length of 64.

Valid Values: `us-east-1 | us-east-2 | us-west-1 | us-west-2 | eu-west-1 | eu-west-2 | eu-west-3 | eu-central-1 | eu-central-2 | ap-east-1 | me-south-1 | us-gov-west-1 | us-gov-east-1 | us-iso-east-1 | us-iso-west-1 | us-isob-east-1 | me-central-1 | ap-southeast-1 | ap-southeast-2 | ap-southeast-3 | ap-south-1 | ap-south-2 | ap-northeast-1 | ap-northeast-2 | ap-northeast-3 | eu-north-1 | sa-east-1 | ca-central-1 | cn-north-1 | cn-northwest-1 | af-south-1 | eu-south-1 | eu-south-2 | ap-southeast-4 | il-central-1 | ca-west-1 | ap-southeast-5 | mx-central-1 | us-isof-south-1 | us-isof-east-1 | ap-southeast-7 | ap-east-2 | eu-isoe-west-1 | ap-southeast-6 | us-isob-west-1`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<ListHostedZonesByVPCResponse>
   <HostedZoneSummaries>
      <HostedZoneSummary>
         <HostedZoneId>string</HostedZoneId>
         <Name>string</Name>
         <Owner>
            <OwningAccount>string</OwningAccount>
            <OwningService>string</OwningService>
         </Owner>
      </HostedZoneSummary>
   </HostedZoneSummaries>
   <MaxItems>string</MaxItems>
   <NextToken>string</NextToken>
</ListHostedZonesByVPCResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[ListHostedZonesByVPCResponse](#API_ListHostedZonesByVPC_ResponseSyntax)**

Root level tag for the ListHostedZonesByVPCResponse parameters.

Required: Yes

**[HostedZoneSummaries](#API_ListHostedZonesByVPC_ResponseSyntax)**

A list that contains one `HostedZoneSummary` element for each hosted zone
that the specified Amazon VPC is associated with. Each `HostedZoneSummary`
element contains the hosted zone name and ID, and information about who owns the hosted
zone.

Type: Array of [HostedZoneSummary](api-hostedzonesummary.md) objects

**[MaxItems](#API_ListHostedZonesByVPC_ResponseSyntax)**

The value that you specified for `MaxItems` in the most recent
`ListHostedZonesByVPC` request.

Type: String

**[NextToken](#API_ListHostedZonesByVPC_ResponseSyntax)**

The value that you will use for `NextToken` in the next
`ListHostedZonesByVPC` request.

Type: String

Length Constraints: Maximum length of 1024.

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

## Examples

### Example Request

This example illustrates one usage of ListHostedZonesByVPC.

```

GET /2013-04-01/hostedzonesbyvpc?maxitems=10&vpcid=vpc-a1b2c3d4e5&vpcregion=us-west-1 HTTP/1.1
```

### Example Response

This example illustrates one usage of ListHostedZonesByVPC.

```

HTTP/1.1 200 OK
<?xml version="1.0" encoding="UTF-8"?>
<ListHostedZonesByVPCResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <HostedZoneSummaries>
      <HostedZoneSummary>
         <HostedZoneId>Z111111QQQQQQQ</HostedZoneId>
         <Name>example.com</Name>
         <Owner>
           <OwningAccount>123456789012</OwningAccount>
         </Owner>
      </HostedZoneSummary>
      <HostedZoneSummary>
         <HostedZoneId>Z222222RRRRRRR</HostedZoneId>
         <Name>example.net</Name>
         <Owner>
            <OwningAccount>111122223333</OwningAccount>
         </Owner>
      </HostedZoneSummary>
      <HostedZoneSummary>
         <HostedZoneId>Z333333SSSSSSS</HostedZoneId>
         <Name>example.org</Name>
         <Owner>
            <OwningService>efs.amazonaws.com</OwningService>
         </Owner>
      </HostedZoneSummary>
   </HostedZoneSummaries>
   <MaxItems>10</MaxItems>
</ListHostedZonesByVPCResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/listhostedzonesbyvpc.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/listhostedzonesbyvpc.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/listhostedzonesbyvpc.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/listhostedzonesbyvpc.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/listhostedzonesbyvpc.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/listhostedzonesbyvpc.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/listhostedzonesbyvpc.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/listhostedzonesbyvpc.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/listhostedzonesbyvpc.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/listhostedzonesbyvpc.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListHostedZonesByName

ListQueryLoggingConfigs
