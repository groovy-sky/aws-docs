# CreateTrafficMirrorFilter

Creates a Traffic Mirror filter.

A Traffic Mirror filter is a set of rules that defines the traffic to mirror.

By default, no traffic is mirrored. To mirror traffic, use [CreateTrafficMirrorFilterRule](api-createtrafficmirrorfilterrule.md)
to add Traffic Mirror rules to the filter. The rules you add define what traffic gets mirrored.
You can also use [ModifyTrafficMirrorFilterNetworkServices](api-modifytrafficmirrorfilternetworkservices.md) to mirror supported network services.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request. For more information, see [How to ensure idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

Required: No

**Description**

The description of the Traffic Mirror filter.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**TagSpecification.N**

The tags to assign to a Traffic Mirror filter.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**clientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request. For more information, see [How to ensure idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

**requestId**

The ID of the request.

Type: String

**trafficMirrorFilter**

Information about the Traffic Mirror filter.

Type: [TrafficMirrorFilter](api-trafficmirrorfilter.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateTrafficMirrorFilter)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateTrafficMirrorFilter)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createtrafficmirrorfilter.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createtrafficmirrorfilter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createtrafficmirrorfilter.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createtrafficmirrorfilter.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createtrafficmirrorfilter.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createtrafficmirrorfilter.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateTrafficMirrorFilter)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createtrafficmirrorfilter.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateTags

CreateTrafficMirrorFilterRule
