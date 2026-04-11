# DeleteTrafficPolicy

Deletes a traffic policy.

When you delete a traffic policy, Route 53 sets a flag on the policy to indicate that
it has been deleted. However, Route 53 never fully deletes the traffic policy. Note the
following:

- Deleted traffic policies aren't listed if you run [ListTrafficPolicies](api-listtrafficpolicies.md).

- There's no way to get a list of deleted policies.

- If you retain the ID of the policy, you can get information about the policy,
including the traffic policy document, by running [GetTrafficPolicy](api-gettrafficpolicy.md).

## Request Syntax

```nohighlight

DELETE /2013-04-01/trafficpolicy/Id/Version HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[Id](#API_DeleteTrafficPolicy_RequestSyntax)**

The ID of the traffic policy that you want to delete.

Length Constraints: Minimum length of 1. Maximum length of 36.

Required: Yes

**[Version](#API_DeleteTrafficPolicy_RequestSyntax)**

The version number of the traffic policy that you want to delete.

Valid Range: Minimum value of 1. Maximum value of 1000.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ConcurrentModification**

Another user submitted a request to create, update, or delete the object at the same
time that you did. Retry the request.

**message**

HTTP Status Code: 400

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

**NoSuchTrafficPolicy**

No traffic policy exists with the specified ID.

**message**

HTTP Status Code: 404

**TrafficPolicyInUse**

One or more traffic policy instances were created by using the specified traffic
policy.

**message**

HTTP Status Code: 400

## Examples

### Example Request

This example illustrates one usage of DeleteTrafficPolicy.

```

DELETE /2013-04-01/trafficpolicy/12345678-abcd-9876-fedc-1a2b3c4de5f6/2
```

### Example Response

This example illustrates one usage of DeleteTrafficPolicy.

```

HTTP/1.1 200 OK
<?xml version="1.0" encoding="UTF-8"?>
<DeleteTrafficPolicyResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
</DeleteTrafficPolicyResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/deletetrafficpolicy.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/deletetrafficpolicy.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/deletetrafficpolicy.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/deletetrafficpolicy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/deletetrafficpolicy.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/deletetrafficpolicy.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/deletetrafficpolicy.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/deletetrafficpolicy.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/deletetrafficpolicy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/deletetrafficpolicy.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteReusableDelegationSet

DeleteTrafficPolicyInstance
