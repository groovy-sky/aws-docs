# ModifyInstanceEventStartTime

Modifies the start time for a scheduled Amazon EC2 instance event.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the operation, without actually making the
request, and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**InstanceEventId**

The ID of the event whose date and time you are modifying.

Type: String

Required: Yes

**InstanceId**

The ID of the instance with the scheduled event.

Type: String

Required: Yes

**NotBefore**

The new date and time when the event will take place.

Type: Timestamp

Required: Yes

## Response Elements

The following elements are returned by the service.

**event**

Information about the event.

Type: [InstanceStatusEvent](api-instancestatusevent.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

The following example shows how to modify the event start time for the
specified instance. The event ID is specified by the
`InstanceEventId` parameter and the new date and time is
specified by the `NotBefore` parameter.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ModifyInstanceEventStartTime
&InstanceId=i-1234567890abcdef0
&InstanceEventId=instance-event-0abcdef1234567890
&NotBefore=2019-03-25T10:00:00.000
&AUTHPARAMS
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/modifyinstanceeventstarttime.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/modifyinstanceeventstarttime.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyinstanceeventstarttime.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyinstanceeventstarttime.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyinstanceeventstarttime.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyinstanceeventstarttime.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyinstanceeventstarttime.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyinstanceeventstarttime.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/modifyinstanceeventstarttime.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyinstanceeventstarttime.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyInstanceCreditSpecification

ModifyInstanceEventWindow
