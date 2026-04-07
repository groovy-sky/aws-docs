# ModifyVpcEndpointConnectionNotification

Modifies a connection notification for VPC endpoint or VPC endpoint service. You
can change the SNS topic for the notification, or the events for which to be notified.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ConnectionEvents.N**

The events for the endpoint. Valid values are `Accept`,
`Connect`, `Delete`, and `Reject`.

Type: Array of strings

Required: No

**ConnectionNotificationArn**

The ARN for the SNS topic for the notification.

Type: String

Required: No

**ConnectionNotificationId**

The ID of the notification.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**return**

Returns `true` if the request succeeds; otherwise, it returns an error.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

The following example modifies notification
`vpce-nfn-abccb952bc8af7123` by modifying the endpoint events and
the SNS topic ARN.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ModifyVpcEndpointConnectionNotification
&ConnectionNotificationId=vpce-nfn-abccb952bc8af7123
&ConnectionNotificationArn=arn:aws:sns:us-east-1:123456789012:mytopic
&ConnectionEvents.1=Accept
&ConnectionEvents.2=Reject
&AUTHPARAMS
```

#### Sample Response

```

<ModifyVpcEndpointConnectionNotificationResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>08d80840-f750-42db-a6f8-2cd32example</requestId>
    <return>true</return>
</ModifyVpcEndpointConnectionNotificationResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ModifyVpcEndpointConnectionNotification)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ModifyVpcEndpointConnectionNotification)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyvpcendpointconnectionnotification.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyvpcendpointconnectionnotification.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyvpcendpointconnectionnotification.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyvpcendpointconnectionnotification.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyvpcendpointconnectionnotification.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyvpcendpointconnectionnotification.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ModifyVpcEndpointConnectionNotification)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyvpcendpointconnectionnotification.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyVpcEndpoint

ModifyVpcEndpointServiceConfiguration
