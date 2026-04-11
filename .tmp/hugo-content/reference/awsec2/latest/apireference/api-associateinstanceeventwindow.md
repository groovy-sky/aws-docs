# AssociateInstanceEventWindow

Associates one or more targets with an event window. Only one type of target (instance
IDs, Dedicated Host IDs, or tags) can be specified with an event window.

For more information, see [Define event windows for scheduled\
events](../../../../services/ec2/latest/userguide/event-windows.md) in the _Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AssociationTarget**

One or more targets associated with the specified event window.

Type: [InstanceEventWindowAssociationRequest](api-instanceeventwindowassociationrequest.md) object

Required: Yes

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/associateinstanceeventwindow.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/associateinstanceeventwindow.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/associateinstanceeventwindow.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/associateinstanceeventwindow.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/associateinstanceeventwindow.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/associateinstanceeventwindow.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/associateinstanceeventwindow.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/associateinstanceeventwindow.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/associateinstanceeventwindow.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/associateinstanceeventwindow.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AssociateIamInstanceProfile

AssociateIpamByoasn
