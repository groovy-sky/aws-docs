# ModifyInstanceEventWindow

Modifies the specified event window.

You can define either a set of time ranges or a cron expression when modifying the event
window, but not both.

To modify the targets associated with the event window, use the [AssociateInstanceEventWindow](api-associateinstanceeventwindow.md) and [DisassociateInstanceEventWindow](api-disassociateinstanceeventwindow.md) API.

If AWS has already scheduled an event, modifying an event window won't change the time
of the scheduled event.

For more information, see [Define event windows for scheduled\
events](../../../../services/ec2/latest/userguide/event-windows.md) in the _Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CronExpression**

The cron expression of the event window, for example, `* 0-4,20-23 * *
         1,5`.

Constraints:

- Only hour and day of the week values are supported.

- For day of the week values, you can specify either integers `0` through
`6`, or alternative single values `SUN` through
`SAT`.

- The minute, month, and year must be specified by `*`.

- The hour value must be one or a multiple range, for example, `0-4` or
`0-4,20-23`.

- Each hour range must be >= 2 hours, for example, `0-2` or
`20-23`.

- The event window must be >= 4 hours. The combined total time ranges in the event
window must be >= 4 hours.

For more information about cron expressions, see [cron](https://en.wikipedia.org/wiki/Cron) on the _Wikipedia_
_website_.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**InstanceEventWindowId**

The ID of the event window.

Type: String

Required: Yes

**Name**

The name of the event window.

Type: String

Required: No

**TimeRange.N**

The time ranges of the event window.

Type: Array of [InstanceEventWindowTimeRangeRequest](api-instanceeventwindowtimerangerequest.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**instanceEventWindow**

Information about the event window.

Type: [InstanceEventWindow](api-instanceeventwindow.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ModifyInstanceEventWindow)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ModifyInstanceEventWindow)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyinstanceeventwindow.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyinstanceeventwindow.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyinstanceeventwindow.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyinstanceeventwindow.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyinstanceeventwindow.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyinstanceeventwindow.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ModifyInstanceEventWindow)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyinstanceeventwindow.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyInstanceEventStartTime

ModifyInstanceMaintenanceOptions
