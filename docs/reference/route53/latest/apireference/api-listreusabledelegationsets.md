# ListReusableDelegationSets

Retrieves a list of the reusable delegation sets that are associated with the current
AWS account.

## Request Syntax

```nohighlight

GET /2013-04-01/delegationset?marker=Marker&maxitems=MaxItems HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[marker](#API_ListReusableDelegationSets_RequestSyntax)**

If the value of `IsTruncated` in the previous response was
`true`, you have more reusable delegation sets. To get another group,
submit another `ListReusableDelegationSets` request.

For the value of `marker`, specify the value of `NextMarker`
from the previous response, which is the ID of the first reusable delegation set that
Amazon Route 53 will return if you submit another request.

If the value of `IsTruncated` in the previous response was
`false`, there are no more reusable delegation sets to get.

Length Constraints: Maximum length of 64.

**[maxitems](#API_ListReusableDelegationSets_RequestSyntax)**

The number of reusable delegation sets that you want Amazon Route 53 to return in the
response to this request. If you specify a value greater than 100, Route 53 returns only
the first 100 reusable delegation sets.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<ListReusableDelegationSetsResponse>
   <DelegationSets>
      <DelegationSet>
         <CallerReference>string</CallerReference>
         <Id>string</Id>
         <NameServers>
            <NameServer>string</NameServer>
         </NameServers>
      </DelegationSet>
   </DelegationSets>
   <IsTruncated>boolean</IsTruncated>
   <Marker>string</Marker>
   <MaxItems>string</MaxItems>
   <NextMarker>string</NextMarker>
</ListReusableDelegationSetsResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[ListReusableDelegationSetsResponse](#API_ListReusableDelegationSets_ResponseSyntax)**

Root level tag for the ListReusableDelegationSetsResponse parameters.

Required: Yes

**[DelegationSets](#API_ListReusableDelegationSets_ResponseSyntax)**

A complex type that contains one `DelegationSet` element for each reusable
delegation set that was created by the current AWS account.

Type: Array of [DelegationSet](api-delegationset.md) objects

**[IsTruncated](#API_ListReusableDelegationSets_ResponseSyntax)**

A flag that indicates whether there are more reusable delegation sets to be
listed.

Type: Boolean

**[Marker](#API_ListReusableDelegationSets_ResponseSyntax)**

For the second and subsequent calls to `ListReusableDelegationSets`,
`Marker` is the value that you specified for the `marker`
parameter in the request that produced the current response.

Type: String

Length Constraints: Maximum length of 64.

**[MaxItems](#API_ListReusableDelegationSets_ResponseSyntax)**

The value that you specified for the `maxitems` parameter in the call to
`ListReusableDelegationSets` that produced the current response.

Type: String

**[NextMarker](#API_ListReusableDelegationSets_ResponseSyntax)**

If `IsTruncated` is `true`, the value of `NextMarker`
identifies the next reusable delegation set that Amazon Route 53 will return if you
submit another `ListReusableDelegationSets` request and specify the value of
`NextMarker` in the `marker` parameter.

Type: String

Length Constraints: Maximum length of 64.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

## Examples

### Example Request

This example illustrates one usage of ListReusableDelegationSets.

```

GET /2013-04-01/delegationset?maxitems=2
```

### Example Response

This example illustrates one usage of ListReusableDelegationSets.

```

HTTP/1.1 200 OK
<?xml version="1.0" encoding="UTF-8"?>
<ListReusableDelegationSetsResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <DelegationSets>
      <DelegationSet>
         <Id>/delegationset/N1PA6795SAMPLE</Id>
         <CallerReference>2017-03-15T01:36:41.958Z</CallerReference>
         <NameServers>
            <NameServer>ns-2042.awsdns-64.com</NameServer>
            <NameServer>ns-2043.awsdns-65.net</NameServer>
            <NameServer>ns-2044.awsdns-66.org</NameServer>
            <NameServer>ns-2045.awsdns-67.co.uk</NameServer>
         </NameServers>
      </DelegationSet>
      <DelegationSet>
         <Id>/delegationset/N1PA7000SAMPLE</Id>
         <CallerReference>2017-03-16T01:37:42.959Z</CallerReference>
         <NameServers>
            <NameServer>ns-2046.awsdns-68.com</NameServer>
            <NameServer>ns-2047.awsdns-69.net</NameServer>
            <NameServer>ns-2048.awsdns-70.org</NameServer>
            <NameServer>ns-2049.awsdns-71.co.uk</NameServer>
         </NameServers>
      </DelegationSet>
   </DelegationSets>
   <IsTruncated>true</IsTruncated>
   <NextMarker>N1PA6797SAMPLE</NextMarker>
   <MaxItems>2</MaxItems>
</ListReusableDelegationSetsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/route53-2013-04-01/ListReusableDelegationSets)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/route53-2013-04-01/ListReusableDelegationSets)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53-2013-04-01/ListReusableDelegationSets)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/route53-2013-04-01/ListReusableDelegationSets)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53-2013-04-01/ListReusableDelegationSets)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/route53-2013-04-01/ListReusableDelegationSets)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/route53-2013-04-01/ListReusableDelegationSets)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/route53-2013-04-01/ListReusableDelegationSets)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/route53-2013-04-01/ListReusableDelegationSets)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53-2013-04-01/ListReusableDelegationSets)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ListResourceRecordSets

ListTagsForResource
