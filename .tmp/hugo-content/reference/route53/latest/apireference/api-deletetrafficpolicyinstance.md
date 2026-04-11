# DeleteTrafficPolicyInstance

Deletes a traffic policy instance and all of the resource record sets that Amazon
Route 53 created when you created the instance.

###### Note

In the Route 53 console, traffic policy instances are known as policy
records.

## Request Syntax

```nohighlight

DELETE /2013-04-01/trafficpolicyinstance/Id HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[Id](#API_DeleteTrafficPolicyInstance_RequestSyntax)**

The ID of the traffic policy instance that you want to delete.

###### Important

When you delete a traffic policy instance, Amazon Route 53 also deletes all of the
resource record sets that were created when you created the traffic policy
instance.

Length Constraints: Minimum length of 1. Maximum length of 36.

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

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

**NoSuchTrafficPolicyInstance**

No traffic policy instance exists with the specified ID.

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

This example illustrates one usage of DeleteTrafficPolicyInstance.

```

DELETE /2013-04-01/trafficpolicyinstance/12131415-abac-5432-caba-6f5e4d3c2b1a
```

### Example Response

This example illustrates one usage of DeleteTrafficPolicyInstance.

```

HTTP/1.1 200 OK
<?xml version="1.0" encoding="UTF-8"?>
<DeleteTrafficPolicyInstanceResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
</DeleteTrafficPolicyInstanceResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/deletetrafficpolicyinstance.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/deletetrafficpolicyinstance.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/deletetrafficpolicyinstance.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/deletetrafficpolicyinstance.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/deletetrafficpolicyinstance.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/deletetrafficpolicyinstance.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/deletetrafficpolicyinstance.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/deletetrafficpolicyinstance.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/deletetrafficpolicyinstance.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/deletetrafficpolicyinstance.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteTrafficPolicy

DeleteVPCAssociationAuthorization
