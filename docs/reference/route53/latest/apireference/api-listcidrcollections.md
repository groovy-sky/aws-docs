# ListCidrCollections

Returns a paginated list of CIDR collections in the AWS account
(metadata only).

## Request Syntax

```nohighlight

GET /2013-04-01/cidrcollection?maxresults=MaxResults&nexttoken=NextToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[maxresults](#API_ListCidrCollections_RequestSyntax)**

The maximum number of CIDR collections to return in the response.

**[nexttoken](#API_ListCidrCollections_RequestSyntax)**

An opaque pagination token to indicate where the service is to begin enumerating
results.

If no value is provided, the listing of results starts from the beginning.

Length Constraints: Maximum length of 1024.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<ListCidrCollectionsResponse>
   <CidrCollections>
      <CollectionSummary>
         <Arn>string</Arn>
         <Id>string</Id>
         <Name>string</Name>
         <Version>long</Version>
      </CollectionSummary>
   </CidrCollections>
   <NextToken>string</NextToken>
</ListCidrCollectionsResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[ListCidrCollectionsResponse](#API_ListCidrCollections_ResponseSyntax)**

Root level tag for the ListCidrCollectionsResponse parameters.

Required: Yes

**[CidrCollections](#API_ListCidrCollections_ResponseSyntax)**

A complex type with information about the CIDR collection.

Type: Array of [CollectionSummary](https://docs.aws.amazon.com/Route53/latest/APIReference/API_CollectionSummary.html) objects

**[NextToken](#API_ListCidrCollections_ResponseSyntax)**

An opaque pagination token to indicate where the service is to begin enumerating
results.

If no value is provided, the listing of results starts from the beginning.

Type: String

Length Constraints: Maximum length of 1024.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

## Examples

### Example request

This example illustrates one usage of ListCidrCollections.

```

GET /2013-04-01/cidrcollection?maxresults=1
```

### Example response

This example illustrates one usage of ListCidrCollections.

```

HTTP/1.1 200
<?xml version="1.0"?>
<ListCidrCollectionsResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">>
   <NextToken>eyJjb2xsZWN0aW9uSWQiOiIwNGJmZjEyZS04NTdjLTFiNmEtNTc2OS0wMTQwMzg4NmE1NzkiLCJjb2xsZWN0aW9uTmFtZSI6ImZlbndhLTEifQ==</NextToken>
     <CidrCollections>
       <member>
         <Arn>arn:aws:route53:::cidrcollection/c8c02a84-aaaa-bbbb-e0d2-d833a2f80106</<Arn>
         <Id>c8c02a84-aaaa-bbbb-e0d2-d833a2f80106</Id>
         <Name>isp-city-cidrs</Name>
         <Version>1</ersion>
       </member>
    </CidrCollections>
</ListCidrCollectionsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/route53-2013-04-01/ListCidrCollections)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/route53-2013-04-01/ListCidrCollections)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53-2013-04-01/ListCidrCollections)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/route53-2013-04-01/ListCidrCollections)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53-2013-04-01/ListCidrCollections)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/route53-2013-04-01/ListCidrCollections)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/route53-2013-04-01/ListCidrCollections)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/route53-2013-04-01/ListCidrCollections)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/route53-2013-04-01/ListCidrCollections)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53-2013-04-01/ListCidrCollections)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ListCidrBlocks

ListCidrLocations
