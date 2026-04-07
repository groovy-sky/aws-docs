# ListTagsForResources

Lists tags for up to 10 health checks or hosted zones.

For information about using tags for cost allocation, see [Using Cost Allocation\
Tags](../../../../services/awsaccountbilling/latest/aboutv2/cost-alloc-tags.md) in the _AWS Billing and Cost Management User Guide_.

## Request Syntax

```nohighlight

POST /2013-04-01/tags/ResourceType HTTP/1.1
<?xml version="1.0" encoding="UTF-8"?>
<ListTagsForResourcesRequest xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <ResourceIds>
      <ResourceId>string</ResourceId>
   </ResourceIds>
</ListTagsForResourcesRequest>
```

## URI Request Parameters

The request uses the following URI parameters.

**[ResourceType](#API_ListTagsForResources_RequestSyntax)**

The type of the resources.

- The resource type for health checks is `healthcheck`.

- The resource type for hosted zones is `hostedzone`.

Valid Values: `healthcheck | hostedzone`

Required: Yes

## Request Body

The request accepts the following data in XML format.

**[ListTagsForResourcesRequest](#API_ListTagsForResources_RequestSyntax)**

Root level tag for the ListTagsForResourcesRequest parameters.

Required: Yes

**[ResourceIds](#API_ListTagsForResources_RequestSyntax)**

A complex type that contains the ResourceId element for each resource for which you
want to get a list of tags.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 10 items.

Length Constraints: Maximum length of 64.

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<ListTagsForResourcesResponse>
   <ResourceTagSets>
      <ResourceTagSet>
         <ResourceId>string</ResourceId>
         <ResourceType>string</ResourceType>
         <Tags>
            <Tag>
               <Key>string</Key>
               <Value>string</Value>
            </Tag>
         </Tags>
      </ResourceTagSet>
   </ResourceTagSets>
</ListTagsForResourcesResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[ListTagsForResourcesResponse](#API_ListTagsForResources_ResponseSyntax)**

Root level tag for the ListTagsForResourcesResponse parameters.

Required: Yes

**[ResourceTagSets](#API_ListTagsForResources_ResponseSyntax)**

A list of `ResourceTagSet` s containing tags associated with the specified
resources.

Type: Array of [ResourceTagSet](https://docs.aws.amazon.com/Route53/latest/APIReference/API_ResourceTagSet.html) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

**NoSuchHealthCheck**

No health check exists with the specified ID.

**message**

HTTP Status Code: 404

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

**ThrottlingException**

The limit on the number of requests per second was exceeded.

HTTP Status Code: 400

## Examples

### Example Request

This example illustrates one usage of ListTagsForResources.

```

GET /2013-04-01/tags/healthcheck HTTP/1.1
<?xml version="1.0" encoding="UTF-8"?>
<ListTagsForResourcesRequest xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <ResourceIds>
      <ResourceId>abcdef11-2222-3333-4444-555555fedcba</ResourceId>
      <ResourceId>aaaaaaaa-1234-5678-9012-bbbbbbcccccc</ResourceId>
   </ResourceIds>
</ListTagsForResourcesRequest>
```

### Example Response

This example illustrates one usage of ListTagsForResources.

```

HTTP/1.1 200 OK
<?xml version="1.0" encoding="UTF-8"?>
<ListTagsForResourcesResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <ResourceTagSets>
      <ResourceTagSet>
         <ResourceType>healthcheck</ResourceType>
         <ResourceId>abcdef11-2222-3333-4444-555555fedcba</ResourceId>
         <Tags>
            <Tag>
               <Key>Owner</Key>
               <Value>dbadmin</Value>
            </Tag>
         </Tags>
      </ResourceTagSet>
      <ResourceTagSet>
         <ResourceType>healthcheck</ResourceType>
         <ResourceId>aaaaaaaa-1234-5678-9012-bbbbbbcccccc</ResourceId>
         <Tags>
            <Tag>
               <Key>Cost Center</Key>
               <Value>80432</Value>
            </Tag>
         </Tags>
      </ResourceTagSet>
   </ResourceTagSets>
</ListTagsForResourcesResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/route53-2013-04-01/ListTagsForResources)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/route53-2013-04-01/ListTagsForResources)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53-2013-04-01/ListTagsForResources)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/route53-2013-04-01/ListTagsForResources)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53-2013-04-01/ListTagsForResources)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/route53-2013-04-01/ListTagsForResources)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/route53-2013-04-01/ListTagsForResources)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/route53-2013-04-01/ListTagsForResources)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/route53-2013-04-01/ListTagsForResources)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53-2013-04-01/ListTagsForResources)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ListTagsForResource

ListTrafficPolicies
