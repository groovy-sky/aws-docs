# DetachNetworkInterface

Detaches a network interface from an instance.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/CommonParameters.html).

**AttachmentId**

The ID of the attachment.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**Force**

Specifies whether to force a detachment.

###### Note

- Use the `Force` parameter only as a last resort to detach a
network interface from a failed instance.

- If you use the `Force` parameter to detach a network interface,
you might not be able to attach a different network interface to the same
index on the instance without first stopping and starting the
instance.

- If you force the detachment of a network interface, the [instance\
metadata](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html) might not get updated. This means that the attributes
associated with the detached network interface might still be visible. The
instance metadata will get updated when you stop and start the
instance.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**return**

Is `true` if the request succeeds, and an error otherwise.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/errors-overview.html#CommonErrors).

## Examples

### Example

This example detaches the specified elastic network interface (ENI).

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DetachNetworkInterface
&AttachmentId=eni-attach-d94b09b0
&AUTHPARAMS
```

#### Sample Response

```

<DetachNetworkInterfaceResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>ce540707-0635-46bc-97da-33a8a362a0e8</requestId>
    <return>true</return>
</DetachNetworkInterfaceResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DetachNetworkInterface)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DetachNetworkInterface)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/DetachNetworkInterface)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/DetachNetworkInterface)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/DetachNetworkInterface)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/DetachNetworkInterface)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/DetachNetworkInterface)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/DetachNetworkInterface)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DetachNetworkInterface)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/DetachNetworkInterface)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DetachInternetGateway

DetachVerifiedAccessTrustProvider
