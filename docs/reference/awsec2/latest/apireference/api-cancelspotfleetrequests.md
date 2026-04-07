# CancelSpotFleetRequests

Cancels the specified Spot Fleet requests.

After you cancel a Spot Fleet request, the Spot Fleet launches no new instances.

You must also specify whether a canceled Spot Fleet request should terminate its instances. If you
choose to terminate the instances, the Spot Fleet request enters the
`cancelled_terminating` state. Otherwise, the Spot Fleet request enters
the `cancelled_running` state and the instances continue to run until they
are interrupted or you terminate them manually.

###### Important

**Terminating an instance is permanent and irreversible.**

After you terminate an instance, you can no longer connect to it, and it can't be recovered.
All attached Amazon EBS volumes that are configured to be deleted on termination are also permanently
deleted and can't be recovered. All data stored on instance store volumes is permanently lost.
For more information, see [How instance termination works](../../../../services/ec2/latest/userguide/how-ec2-instance-termination-works.md).

Before you terminate an instance, ensure that you have backed up all data that you need to
retain after the termination to persistent storage.

###### Restrictions

- You can delete up to 100 fleets in a single request. If you exceed the specified
number, no fleets are deleted.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**SpotFleetRequestId.N**

The IDs of the Spot Fleet requests.

Constraint: You can specify up to 100 IDs in a single request.

Type: Array of strings

Required: Yes

**TerminateInstances**

Indicates whether to terminate the associated instances when the Spot Fleet request is canceled.
The default is to terminate the instances.

To let the instances continue to run after the Spot Fleet request is canceled, specify
`no-terminate-instances`.

Type: Boolean

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**successfulFleetRequestSet**

Information about the Spot Fleet requests that are successfully canceled.

Type: Array of [CancelSpotFleetRequestsSuccessItem](api-cancelspotfleetrequestssuccessitem.md) objects

**unsuccessfulFleetRequestSet**

Information about the Spot Fleet requests that are not successfully canceled.

Type: Array of [CancelSpotFleetRequestsErrorItem](api-cancelspotfleetrequestserroritem.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example cancels Spot Fleet request
sfr-123f8fc2-cb31-425e-abcd-example2710 and terminates all instances that were
launched by the request.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CancelSpotFleetRequests
&SpotFleetRequestId.1=sfr-123f8fc2-cb31-425e-abcd-example2710
&TerminateInstances=true
&AUTHPARAMS
```

#### Sample Response

```

<CancelSpotFleetRequestsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>e12d2fe5-6503-4b4b-911c-example</requestId>
    <unsuccessfulFleetRequestSet/>
    <successfulFleetRequestSet>
        <item>
            <spotFleetRequestId>sfr-123f8fc2-cb31-425e-abcd-example2710</spotFleetRequestId>
            <currentSpotFleetRequestState>cancelled_terminating</currentSpotFleetRequestState>
            <previousSpotFleetRequestState>active</previousSpotFleetRequestState>
        </item>
    </successfulFleetRequestSet>
</CancelSpotFleetRequestsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CancelSpotFleetRequests)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CancelSpotFleetRequests)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/cancelspotfleetrequests.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/cancelspotfleetrequests.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/cancelspotfleetrequests.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/cancelspotfleetrequests.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/cancelspotfleetrequests.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/cancelspotfleetrequests.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CancelSpotFleetRequests)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/cancelspotfleetrequests.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CancelReservedInstancesListing

CancelSpotInstanceRequests
