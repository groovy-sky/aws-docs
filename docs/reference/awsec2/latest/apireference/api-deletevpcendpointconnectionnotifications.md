# DeleteVpcEndpointConnectionNotifications

Deletes the specified VPC endpoint connection notifications.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ConnectionNotificationId.N**

The IDs of the notifications.

Type: Array of strings

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

**unsuccessful**

Information about the notifications that could not be deleted
successfully.

Type: Array of [UnsuccessfulItem](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_UnsuccessfulItem.html) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example deletes connection notification
`vpce-nfn-04bcb952bc8af7123`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeleteVpcEndpointConnectionNotifications
&ConnectionNotificationId.1=vpce-nfn-04bcb952bc8af7123
&AUTHPARAMS
```

#### Sample Response

```

<DeleteVpcEndpointConnectionNotificationsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>2bf45d2e-a871-4375-9a93-f4188example</requestId>
    <unsuccessful/>
</DeleteVpcEndpointConnectionNotificationsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DeleteVpcEndpointConnectionNotifications)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DeleteVpcEndpointConnectionNotifications)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/DeleteVpcEndpointConnectionNotifications)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/DeleteVpcEndpointConnectionNotifications)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/DeleteVpcEndpointConnectionNotifications)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/DeleteVpcEndpointConnectionNotifications)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/DeleteVpcEndpointConnectionNotifications)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/DeleteVpcEndpointConnectionNotifications)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DeleteVpcEndpointConnectionNotifications)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/DeleteVpcEndpointConnectionNotifications)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DeleteVpcEncryptionControl

DeleteVpcEndpoints
