# DeleteDeliveryDestination

Deletes a _delivery destination_. A delivery is a connection between a
logical _delivery source_ and a logical _delivery_
_destination_.

You can't delete a delivery destination if any current deliveries are associated with it.
To find whether any deliveries are associated with this delivery destination, use the [DescribeDeliveries](api-describedeliveries.md) operation and check the `deliveryDestinationArn`
field in the results.

## Request Syntax

```nohighlight

{
   "name": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[name](#API_DeleteDeliveryDestination_RequestSyntax)**

The name of the delivery destination that you want to delete. You can find a list of
delivery destination names by using the [DescribeDeliveryDestinations](api-describedeliverydestinations.md) operation.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 60.

Pattern: `[\w-]*`

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ConflictException**

This operation attempted to create a resource that already exists.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

**ServiceQuotaExceededException**

This request exceeds a service quota.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

**ThrottlingException**

The request was throttled because of quota limits.

HTTP Status Code: 400

**ValidationException**

One of the parameters for the request is not valid.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/deletedeliverydestination.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/deletedeliverydestination.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/deletedeliverydestination.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/deletedeliverydestination.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/deletedeliverydestination.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/deletedeliverydestination.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/deletedeliverydestination.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/deletedeliverydestination.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/deletedeliverydestination.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/deletedeliverydestination.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteDelivery

DeleteDeliveryDestinationPolicy
