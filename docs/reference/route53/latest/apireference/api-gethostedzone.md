# GetHostedZone

Gets information about a specified hosted zone including the four name servers
assigned to the hosted zone.

`` returns the VPCs associated with the specified hosted zone and does not reflect the VPC
associations by Route 53 Profiles. To get the associations to a Profile, call the [ListProfileAssociations](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53profiles_ListProfileAssociations.html) API.

## Request Syntax

```nohighlight

GET /2013-04-01/hostedzone/Id HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[Id](#API_GetHostedZone_RequestSyntax)**

The ID of the hosted zone that you want to get information about.

Length Constraints: Maximum length of 32.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<GetHostedZoneResponse>
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
   <VPCs>
      <VPC>
         <VPCId>string</VPCId>
         <VPCRegion>string</VPCRegion>
      </VPC>
   </VPCs>
</GetHostedZoneResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[GetHostedZoneResponse](#API_GetHostedZone_ResponseSyntax)**

Root level tag for the GetHostedZoneResponse parameters.

Required: Yes

**[DelegationSet](#API_GetHostedZone_ResponseSyntax)**

A complex type that lists the Amazon Route 53 name servers for the specified hosted
zone.

Type: [DelegationSet](https://docs.aws.amazon.com/Route53/latest/APIReference/API_DelegationSet.html) object

**[HostedZone](#API_GetHostedZone_ResponseSyntax)**

A complex type that contains general information about the specified hosted
zone.

Type: [HostedZone](https://docs.aws.amazon.com/Route53/latest/APIReference/API_HostedZone.html) object

**[VPCs](#API_GetHostedZone_ResponseSyntax)**

A complex type that contains information about the VPCs that are associated with the
specified hosted zone.

Type: Array of [VPC](https://docs.aws.amazon.com/Route53/latest/APIReference/API_VPC.html) objects

Array Members: Minimum number of 1 item.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](https://docs.aws.amazon.com/Route53/latest/APIReference/CommonErrors.html).

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

**NoSuchHostedZone**

No hosted zone exists with the ID that you specified.

**message**

HTTP Status Code: 404

## Examples

### Example Request

This example illustrates one usage of GetHostedZone.

```

GET /2013-04-01/hostedzone/Z1PA6795UKMFR9
```

### Example Response (Public Hosted Zone, Default Delegation Set Assigned by Route 53)

This example illustrates one usage of GetHostedZone.

```

HTTP/1.1 200 OK
<?xml version="1.0" encoding="UTF-8"?>
<GetHostedZoneResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <HostedZone>
      <Id>/hostedzone/Z1PA6795UKMFR9</Id>
      <Name>example.com.</Name>
      <CallerReference>2017-03-01T11:22:14Z</CallerReference>
      <Config>
         <Comment>This is my first hosted zone.</Comment>
         <PrivateZone>false</PrivateZone>
      </Config>
      <ResourceRecordSetCount>17</ResourceRecordSetCount>
   </HostedZone>
   <DelegationSet>
      <NameServers>
         <NameServer>ns-2048.awsdns-64.com</NameServer>
         <NameServer>ns-2049.awsdns-65.net</NameServer>
         <NameServer>ns-2050.awsdns-66.org</NameServer>
         <NameServer>ns-2051.awsdns-67.co.uk</NameServer>
      </NameServers>
   </DelegationSet>
</GetHostedZoneResponse>
```

### Example Response (Public Hosted Zone, Reusable Delegation Set)

This example illustrates one usage of GetHostedZone.

```

HTTP/1.1 200 OK
<?xml version="1.0" encoding="UTF-8"?>
<GetHostedZoneResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <HostedZone>
      <Id>/hostedzone/Z1PA6795UKMFR9</Id>
      <Name>example.com.</Name>
      <CallerReference>2017-03-02T10:44:04Z</CallerReference>
      <Config>
         <Comment>This is my first hosted zone.</Comment>
         <PrivateZone>false</PrivateZone>
      </Config>
      <ResourceRecordSetCount>17</ResourceRecordSetCount>
   </HostedZone>
   <DelegationSet>
      <Id>NU241VPSAMPLE</Id>
      <CallerReference>2017-03-01T11:22:14Z</CallerReference>
      <NameServers>
         <NameServer>ns-2048.awsdns-64.com</NameServer>
         <NameServer>ns-2049.awsdns-65.net</NameServer>
         <NameServer>ns-2050.awsdns-66.org</NameServer>
         <NameServer>ns-2051.awsdns-67.co.uk</NameServer>
      </NameServers>
   </DelegationSet>
</GetHostedZoneResponse>
```

### Example Response (Private Hosted Zone)

This example illustrates one usage of GetHostedZone.

```

HTTP/1.1 200 OK
<?xml version="1.0" encoding="UTF-8"?>
<GetHostedZoneResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <HostedZone>
      <Id>/hostedzone/Z1PA6795UKMFR9</Id>
      <Name>example.com.</Name>
      <CallerReference>myUniqueIdentifier</CallerReference>
      <Config>
         <Comment>This is my first hosted zone.</Comment>
         <PrivateZone>true</PrivateZone>
      </Config>
      <ResourceRecordSetCount>17</ResourceRecordSetCount>
   </HostedZone>
   <VPCs>
      <VPC>
         <VPCRegion>us-east-2</VPCRegion>
         <VPCId>vpc-1a2b3c4d</VPCId>
      </VPC>
   </VPCs>
</GetHostedZoneResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/route53-2013-04-01/GetHostedZone)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/route53-2013-04-01/GetHostedZone)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53-2013-04-01/GetHostedZone)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/route53-2013-04-01/GetHostedZone)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53-2013-04-01/GetHostedZone)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/route53-2013-04-01/GetHostedZone)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/route53-2013-04-01/GetHostedZone)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/route53-2013-04-01/GetHostedZone)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/route53-2013-04-01/GetHostedZone)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53-2013-04-01/GetHostedZone)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetHealthCheckStatus

GetHostedZoneCount
