# CancelSpotInstanceRequests

Cancels one or more Spot Instance requests.

###### Important

Canceling a Spot Instance request does not terminate running Spot Instances
associated with the request.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**SpotInstanceRequestId.N**

The IDs of the Spot Instance requests.

Type: Array of strings

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**spotInstanceRequestSet**

The Spot Instance requests.

Type: Array of [CancelledSpotInstanceRequest](api-cancelledspotinstancerequest.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example cancels the specified Spot Instance request.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CancelSpotInstanceRequests
&SpotInstanceRequestId.1=sir-1a2b3c4d
&AUTHPARAMS
```

#### Sample Response

```

<CancelSpotInstanceRequestsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
  <spotInstanceRequestSet>
    <item>
      <spotInstanceRequestId>sir-1a2b3c4d</spotInstanceRequestId>
      <state>cancelled</state>
    </item>
  </spotInstanceRequestSet>
</CancelSpotInstanceRequestsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/cancelspotinstancerequests.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/cancelspotinstancerequests.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/cancelspotinstancerequests.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/cancelspotinstancerequests.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/cancelspotinstancerequests.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/cancelspotinstancerequests.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/cancelspotinstancerequests.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/cancelspotinstancerequests.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/cancelspotinstancerequests.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/cancelspotinstancerequests.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CancelSpotFleetRequests

ConfirmProductInstance
