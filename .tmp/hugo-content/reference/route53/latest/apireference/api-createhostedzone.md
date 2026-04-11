# CreateHostedZone

Creates a new public or private hosted zone. You create records in a public hosted
zone to define how you want to route traffic on the internet for a domain, such as
example.com, and its subdomains (apex.example.com, acme.example.com). You create records
in a private hosted zone to define how you want to route traffic for a domain and its
subdomains within one or more Amazon Virtual Private Clouds (Amazon VPCs).

###### Important

You can't convert a public hosted zone to a private hosted zone or vice versa.
Instead, you must create a new hosted zone with the same name and create new
resource record sets.

For more information about charges for hosted zones, see [Amazon Route 53 Pricing](http://aws.amazon.com/route53/pricing).

Note the following:

- You can't create a hosted zone for a top-level domain (TLD) such as
.com.

- For public hosted zones, Route 53 automatically creates a default SOA record
and four NS records for the zone. For more information about SOA and NS records,
see [NS and SOA Records\
that Route 53 Creates for a Hosted Zone](../../../../services/route53/latest/developerguide/soa-nsrecords.md) in the
_Amazon Route 53 Developer Guide_.

If you want to use the same name servers for multiple public hosted zones, you
can optionally associate a reusable delegation set with the hosted zone. See the
`DelegationSetId` element.

- If your domain is registered with a registrar other than Route 53,
you must update the name servers with your registrar to make Route 53 the DNS
service for the domain. For more information, see [Migrating DNS Service\
for an Existing Domain to Amazon Route 53](../../../../services/route53/latest/developerguide/migratingdns.md) in the
_Amazon Route 53 Developer Guide_.

When you submit a `CreateHostedZone` request, the initial status of the
hosted zone is `PENDING`. For public hosted zones, this means that the NS and
SOA records are not yet available on all Route 53 DNS servers. When the NS and
SOA records are available, the status of the zone changes to `INSYNC`.

The `CreateHostedZone` request requires the caller to have an
`ec2:DescribeVpcs` permission.

###### Note

When creating private hosted zones, the Amazon VPC must belong to the same
partition where the hosted zone is created. A partition is a group of AWS Regions. Each AWS account is scoped to one
partition.

The following are the supported partitions:

- `aws` \- AWS Regions

- `aws-cn` \- China Regions

- `aws-us-gov` \- AWS GovCloud (US) Region

For more information, see [Access Management](../../../../general/general/latest/gr/aws-arns-and-namespaces.md)
in the _AWS General Reference_.

## Request Syntax

```nohighlight

POST /2013-04-01/hostedzone HTTP/1.1
<?xml version="1.0" encoding="UTF-8"?>
<CreateHostedZoneRequest xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <CallerReference>string</CallerReference>
   <DelegationSetId>string</DelegationSetId>
   <HostedZoneConfig>
      <Comment>string</Comment>
      <PrivateZone>boolean</PrivateZone>
   </HostedZoneConfig>
   <Name>string</Name>
   <VPC>
      <VPCId>string</VPCId>
      <VPCRegion>string</VPCRegion>
   </VPC>
</CreateHostedZoneRequest>
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in XML format.

**[CreateHostedZoneRequest](#API_CreateHostedZone_RequestSyntax)**

Root level tag for the CreateHostedZoneRequest parameters.

Required: Yes

**[CallerReference](#API_CreateHostedZone_RequestSyntax)**

A unique string that identifies the request and that allows failed
`CreateHostedZone` requests to be retried without the risk of executing
the operation twice. You must use a unique `CallerReference` string every
time you submit a `CreateHostedZone` request. `CallerReference`
can be any unique string, for example, a date/time stamp.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**[DelegationSetId](#API_CreateHostedZone_RequestSyntax)**

If you want to associate a reusable delegation set with this hosted zone, the ID that
Amazon Route 53 assigned to the reusable delegation set when you created it.
For more information about reusable delegation sets, see [CreateReusableDelegationSet](api-createreusabledelegationset.md).

If you are using a reusable delegation set to create a public hosted zone for a subdomain,
make sure that the parent hosted zone doesn't use one or more of the same name servers.
If you have overlapping nameservers, the operation will cause a
`ConflictingDomainsExist` error.

Type: String

Length Constraints: Maximum length of 32.

Required: No

**[HostedZoneConfig](#API_CreateHostedZone_RequestSyntax)**

(Optional) A complex type that contains the following optional values:

- For public and private hosted zones, an optional comment

- For private hosted zones, an optional `PrivateZone` element

If you don't specify a comment or the `PrivateZone` element, omit
`HostedZoneConfig` and the other elements.

Type: [HostedZoneConfig](api-hostedzoneconfig.md) object

Required: No

**[Name](#API_CreateHostedZone_RequestSyntax)**

The name of the domain. Specify a fully qualified domain name, for example,
_www.example.com_. The trailing dot is optional; Amazon Route 53 assumes that the domain name is fully qualified. This means that
Route 53 treats _www.example.com_ (without a trailing
dot) and _www.example.com._ (with a trailing dot) as
identical.

If you're creating a public hosted zone, this is the name you have registered with
your DNS registrar. If your domain name is registered with a registrar other than
Route 53, change the name servers for your domain to the set of
`NameServers` that `CreateHostedZone` returns in
`DelegationSet`.

Type: String

Length Constraints: Maximum length of 1024.

Required: Yes

**[VPC](#API_CreateHostedZone_RequestSyntax)**

(Private hosted zones only) A complex type that contains information about the Amazon
VPC that you're associating with this hosted zone.

You can specify only one Amazon VPC when you create a private hosted zone. If you are
associating a VPC with a hosted zone with this request, the paramaters
`VPCId` and `VPCRegion` are also required.

To associate additional Amazon VPCs with the hosted zone, use [AssociateVPCWithHostedZone](api-associatevpcwithhostedzone.md) after you create a hosted zone.

Type: [VPC](api-vpc.md) object

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 201
Location: Location
<?xml version="1.0" encoding="UTF-8"?>
<CreateHostedZoneResponse>
   <ChangeInfo>
      <Comment>string</Comment>
      <Id>string</Id>
      <Status>string</Status>
      <SubmittedAt>timestamp</SubmittedAt>
   </ChangeInfo>
   <DelegationSet>
      <CallerReference>string</CallerReference>
      <Id>string</Id>
      <NameServers>
         <NameServer>string</NameServer>
      </NameServers>
   </DelegationSet>
   <HostedZone>
      <CallerReference>string</CallerReference>
      <Config>
         <Comment>string</Comment>
         <PrivateZone>boolean</PrivateZone>
      </Config>
      <Features>
         <AcceleratedRecoveryStatus>string</AcceleratedRecoveryStatus>
         <FailureReasons>
            <AcceleratedRecovery>string</AcceleratedRecovery>
         </FailureReasons>
      </Features>
      <Id>string</Id>
      <LinkedService>
         <Description>string</Description>
         <ServicePrincipal>string</ServicePrincipal>
      </LinkedService>
      <Name>string</Name>
      <ResourceRecordSetCount>long</ResourceRecordSetCount>
   </HostedZone>
   <VPC>
      <VPCId>string</VPCId>
      <VPCRegion>string</VPCRegion>
   </VPC>
</CreateHostedZoneResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 201 response.

The response returns the following HTTP headers.

**[Location](#API_CreateHostedZone_ResponseSyntax)**

The unique URL representing the new hosted zone.

Length Constraints: Maximum length of 1024.

The following data is returned in XML format by the service.

**[CreateHostedZoneResponse](#API_CreateHostedZone_ResponseSyntax)**

Root level tag for the CreateHostedZoneResponse parameters.

Required: Yes

**[ChangeInfo](#API_CreateHostedZone_ResponseSyntax)**

A complex type that contains information about the `CreateHostedZone`
request.

Type: [ChangeInfo](api-changeinfo.md) object

**[DelegationSet](#API_CreateHostedZone_ResponseSyntax)**

A complex type that describes the name servers for this hosted zone.

Type: [DelegationSet](api-delegationset.md) object

**[HostedZone](#API_CreateHostedZone_ResponseSyntax)**

A complex type that contains general information about the hosted zone.

Type: [HostedZone](api-hostedzone.md) object

**[VPC](#API_CreateHostedZone_ResponseSyntax)**

A complex type that contains information about an Amazon VPC that you associated with
this hosted zone.

Type: [VPC](api-vpc.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ConflictingDomainExists**

The cause of this error depends on the operation that you're performing:

- **Create a public hosted zone:** Two hosted zones
that have the same name or that have a parent/child relationship (example.com
and test.example.com) can't have any common name servers. You tried to create a
hosted zone that has the same name as an existing hosted zone or that's the
parent or child of an existing hosted zone, and you specified a delegation set
that shares one or more name servers with the existing hosted zone. For more
information, see [CreateReusableDelegationSet](api-createreusabledelegationset.md).

- **Create a private hosted zone:** A hosted zone
with the specified name already exists and is already associated with the Amazon
VPC that you specified.

- **Associate VPCs with a private hosted zone:**
The VPC that you specified is already associated with another hosted zone that
has the same name.

HTTP Status Code: 400

**DelegationSetNotAvailable**

You can create a hosted zone that has the same name as an existing hosted zone
(example.com is common), but there is a limit to the number of hosted zones that have
the same name. If you get this error, Amazon Route 53 has reached that limit. If you own
the domain name and Route 53 generates this error, contact Customer Support.

**message**

HTTP Status Code: 400

**DelegationSetNotReusable**

A reusable delegation set with the specified ID does not exist.

**message**

HTTP Status Code: 400

**HostedZoneAlreadyExists**

The hosted zone you're trying to create already exists. Amazon Route 53 returns this
error when a hosted zone has already been created with the specified
`CallerReference`.

**message**

HTTP Status Code: 409

**InvalidDomainName**

The specified domain name is not valid.

**message**

HTTP Status Code: 400

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

**InvalidVPCId**

The VPC ID that you specified either isn't a valid ID or the current account is not
authorized to access this VPC.

**message**

HTTP Status Code: 400

**NoSuchDelegationSet**

A reusable delegation set with the specified ID does not exist.

**message**

HTTP Status Code: 400

**TooManyHostedZones**

This operation can't be completed either because the current account has reached the
limit on the number of hosted zones or because you've reached the limit on the number of
hosted zones that can be associated with a reusable delegation set.

For information about default limits, see [Limits](../../../../services/route53/latest/developerguide/dnslimitations.md) in the
_Amazon Route 53 Developer Guide_.

To get the current limit on hosted zones that can be created by an account, see [GetAccountLimit](api-getaccountlimit.md).

To get the current limit on hosted zones that can be associated with a reusable
delegation set, see [GetReusableDelegationSetLimit](api-getreusabledelegationsetlimit.md).

To request a higher limit, [create a\
case](http://aws.amazon.com/route53-request) with the AWS Support Center.

**message**

HTTP Status Code: 400

## Examples

### Example Request (Public Hosted Zone)

This example illustrates one usage of CreateHostedZone.

```

POST /2013-04-01/hostedzone HTTP/1.1
<?xml version="1.0" encoding="UTF-8"?>
<CreateHostedZoneRequest xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <Name>example.com</Name>
   <CallerReference>myUniqueIdentifier</CallerReference>
   <HostedZoneConfig>
      <Comment>This is my first hosted zone.</Comment>
   </HostedZoneConfig>
   <DelegationSetId>NZ8X2CISAMPLE</DelegationSetId>
</CreateHostedZoneRequest>
```

### Example Response (Public Hosted Zone)

This example illustrates one usage of CreateHostedZone.

```

HTTP/1.1 201 Created
<?xml version="1.0" encoding="UTF-8"?>
<CreateHostedZoneResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <HostedZone>
      <Id>/hostedzone/Z1PA6795UKMFR9</Id>
      <Name>example.com.</Name>
      <CallerReference>myUniqueIdentifier</CallerReference>
      <Config>
         <Comment>This is my first hosted zone.</Comment>
         <PrivateZone>false</PrivateZone>
      </Config>
      <ResourceRecordSetCount>2</ResourceRecordSetCount>
   </HostedZone>
   <ChangeInfo>
      <Id>/change/C1PA6795UKMFR9</Id>
      <Status>PENDING</Status>
      <SubmittedAt>2017-03-15T01:36:41.958Z</SubmittedAt>
   </ChangeInfo>
   <DelegationSet>
      <Id>NZ8X2CISAMPLE</Id>
      <CallerReference>2017-03-01T11:44:14.448Z</Id>
      <NameServers>
         <NameServer>ns-2048.awsdns-64.com</NameServer>
         <NameServer>ns-2049.awsdns-65.net</NameServer>
         <NameServer>ns-2050.awsdns-66.org</NameServer>
         <NameServer>ns-2051.awsdns-67.co.uk</NameServer>
      </NameServers>
   </DelegationSet>
</CreateHostedZoneResponse>
```

### Example Request (Private Hosted Zone)

This example illustrates one usage of CreateHostedZone.

```

POST /2013-04-01/hostedzone HTTP/1.1
<?xml version="1.0" encoding="UTF-8"?>
<CreateHostedZoneRequest xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <Name>example.com</Name>
   <VPC>
      <VPCId>vpc-1a2b3c4d</VPCId>
      <VPCRegion>us-east-2</VPCRegion>
   </VPC>
   <CallerReference>myUniqueIdentifier</CallerReference>
   <HostedZoneConfig>
      <Comment>This is my first hosted zone.</Comment>
   </HostedZoneConfig>
</CreateHostedZoneRequest>
```

### Example Response (Private Hosted Zone)

This example illustrates one usage of CreateHostedZone.

```

HTTP/1.1 201 Created
<?xml version="1.0" encoding="UTF-8"?>
<CreateHostedZoneResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <HostedZone>
      <Id>/hostedzone/Z1D633PJN98FT9</Id>
      <Name>example.com.</Name>
      <VPC>
         <VPCId>vpc-1a2b3c4d</VPCId>
         <VPCRegion>us-east-2</VPCRegion>
      </VPC>
      <CallerReference>myUniqueIdentifier</CallerReference>
      <Config>
         <Comment>This is my first hosted zone.</Comment>
         <PrivateZone>true</PrivateZone>
      </Config>
      <ResourceRecordSetCount>2</ResourceRecordSetCount>
   </HostedZone>
   <ChangeInfo>
      <Id>/change/C1PA6795UKMFR9</Id>
      <Status>PENDING</Status>
      <SubmittedAt>2017-03-15T01:36:41.958Z</SubmittedAt>
   </ChangeInfo>
</CreateHostedZoneResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/createhostedzone.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/createhostedzone.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/createhostedzone.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/createhostedzone.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/createhostedzone.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/createhostedzone.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/createhostedzone.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/createhostedzone.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/createhostedzone.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/createhostedzone.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateHealthCheck

CreateKeySigningKey
