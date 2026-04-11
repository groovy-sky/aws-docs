# ListTrafficPolicies

Gets information about the latest version for every traffic policy that is associated
with the current AWS account. Policies are listed in the order that they
were created in.

For information about how of deleting a traffic policy affects the response from
`ListTrafficPolicies`, see [DeleteTrafficPolicy](api-deletetrafficpolicy.md).

## Request Syntax

```nohighlight

GET /2013-04-01/trafficpolicies?maxitems=MaxItems&trafficpolicyid=TrafficPolicyIdMarker HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[maxitems](#API_ListTrafficPolicies_RequestSyntax)**

(Optional) The maximum number of traffic policies that you want Amazon Route 53 to
return in response to this request. If you have more than `MaxItems` traffic
policies, the value of `IsTruncated` in the response is `true`,
and the value of `TrafficPolicyIdMarker` is the ID of the first traffic
policy that Route 53 will return if you submit another request.

**[trafficpolicyid](#API_ListTrafficPolicies_RequestSyntax)**

(Conditional) For your first request to `ListTrafficPolicies`, don't
include the `TrafficPolicyIdMarker` parameter.

If you have more traffic policies than the value of `MaxItems`,
`ListTrafficPolicies` returns only the first `MaxItems`
traffic policies. To get the next group of policies, submit another request to
`ListTrafficPolicies`. For the value of
`TrafficPolicyIdMarker`, specify the value of
`TrafficPolicyIdMarker` that was returned in the previous
response.

Length Constraints: Minimum length of 1. Maximum length of 36.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<ListTrafficPoliciesResponse>
   <IsTruncated>boolean</IsTruncated>
   <MaxItems>string</MaxItems>
   <TrafficPolicyIdMarker>string</TrafficPolicyIdMarker>
   <TrafficPolicySummaries>
      <TrafficPolicySummary>
         <Id>string</Id>
         <LatestVersion>integer</LatestVersion>
         <Name>string</Name>
         <TrafficPolicyCount>integer</TrafficPolicyCount>
         <Type>string</Type>
      </TrafficPolicySummary>
   </TrafficPolicySummaries>
</ListTrafficPoliciesResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[ListTrafficPoliciesResponse](#API_ListTrafficPolicies_ResponseSyntax)**

Root level tag for the ListTrafficPoliciesResponse parameters.

Required: Yes

**[IsTruncated](#API_ListTrafficPolicies_ResponseSyntax)**

A flag that indicates whether there are more traffic policies to be listed. If the
response was truncated, you can get the next group of traffic policies by submitting
another `ListTrafficPolicies` request and specifying the value of
`TrafficPolicyIdMarker` in the `TrafficPolicyIdMarker` request
parameter.

Type: Boolean

**[MaxItems](#API_ListTrafficPolicies_ResponseSyntax)**

The value that you specified for the `MaxItems` parameter in the
`ListTrafficPolicies` request that produced the current response.

Type: String

**[TrafficPolicyIdMarker](#API_ListTrafficPolicies_ResponseSyntax)**

If the value of `IsTruncated` is `true`,
`TrafficPolicyIdMarker` is the ID of the first traffic policy in the next
group of `MaxItems` traffic policies.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 36.

**[TrafficPolicySummaries](#API_ListTrafficPolicies_ResponseSyntax)**

A list that contains one `TrafficPolicySummary` element for each traffic
policy that was created by the current AWS account.

Type: Array of [TrafficPolicySummary](api-trafficpolicysummary.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

## Examples

### Example Request

This example illustrates one usage of ListTrafficPolicies.

```

GET /2013-04-01/trafficpolicies?maxitems=1
```

### Example Response

This example illustrates one usage of ListTrafficPolicies.

```

HTTP/1.1 200 OK
<?xml version="1.0" encoding="UTF-8"?>
<ListTrafficPoliciesResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <TrafficPolicySummaries>
      <TrafficPolicySummary>
         <Id>12345678-abcd-9876-fedc-1a2b3c4de5f6</Id>
         <Name>MyTrafficPolicy</Name>
         <Type>A</Type>
         <LatestVersion>77</LatestVersion>
         <TrafficPolicyCount>44</TrafficPolicyCount>
      </TrafficPolicySummary>
   </TrafficPolicySummaries>
   <IsTrucated>true</IsTruncated>
   <TrafficPolicyIdMarker>12345678-abcd-9876-fedc-1a2b3c4de5f7</TrafficPolicyIdMarker>
   <MaxItems>1</MaxItems>
</ListTrafficPoliciesResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/listtrafficpolicies.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/listtrafficpolicies.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/listtrafficpolicies.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/listtrafficpolicies.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/listtrafficpolicies.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/listtrafficpolicies.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/listtrafficpolicies.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/listtrafficpolicies.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/listtrafficpolicies.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/listtrafficpolicies.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListTagsForResources

ListTrafficPolicyInstances
