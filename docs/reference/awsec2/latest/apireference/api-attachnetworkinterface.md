# AttachNetworkInterface

Attaches a network interface to an instance.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DeviceIndex**

The index of the device for the network interface attachment.

Type: Integer

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**EnaQueueCount**

The number of ENA queues to be created with the instance.

Type: Integer

Required: No

**EnaSrdSpecification**

Configures ENA Express for the network interface that this action attaches to the
instance.

Type: [EnaSrdSpecification](api-enasrdspecification.md) object

Required: No

**InstanceId**

The ID of the instance.

Type: String

Required: Yes

**NetworkCardIndex**

The index of the network card. Some instance types support multiple network cards. The
primary network interface must be assigned to network card index 0. The default is
network card index 0.

Type: Integer

Required: No

**NetworkInterfaceId**

The ID of the network interface.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**attachmentId**

The ID of the network interface attachment.

Type: String

**networkCardIndex**

The index of the network card.

Type: Integer

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example attaches the specified network interface to the specified
instance.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=AttachNetworkInterface
&DeviceIndex=1
&InstanceId=i-1234567890abcdef0
&NetworkInterfaceId=eni-ffda3197
&AUTHPARAMS
```

#### Sample Response

```

<AttachNetworkInterfaceResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>ace8cd1e-e685-4e44-90fb-92014d907212</requestId>
    <attachmentId>eni-attach-d94b09b0</attachmentId>
</AttachNetworkInterfaceResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/AttachNetworkInterface)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/AttachNetworkInterface)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/attachnetworkinterface.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/attachnetworkinterface.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/attachnetworkinterface.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/attachnetworkinterface.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/attachnetworkinterface.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/attachnetworkinterface.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/AttachNetworkInterface)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/attachnetworkinterface.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AttachInternetGateway

AttachVerifiedAccessTrustProvider
