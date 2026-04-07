# ModifyNetworkInterfaceAttribute

Modifies the specified network interface attribute. You can specify only one attribute
at a time. You can use this action to attach and detach security groups from an existing
EC2 instance.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AssociatedSubnetId.N**

A list of subnet IDs to associate with the network interface.

Type: Array of strings

Required: No

**AssociatePublicIpAddress**

Indicates whether to assign a public IPv4 address to a network interface. This option
can be enabled for any network interface but will only apply to the primary network
interface (eth0).

Type: Boolean

Required: No

**Attachment**

Information about the interface attachment. If modifying the `delete on
                termination` attribute, you must specify the ID of the interface
attachment.

Type: [NetworkInterfaceAttachmentChanges](api-networkinterfaceattachmentchanges.md) object

Required: No

**ConnectionTrackingSpecification**

A connection tracking specification.

Type: [ConnectionTrackingSpecificationRequest](api-connectiontrackingspecificationrequest.md) object

Required: No

**Description**

A description for the network interface.

Type: [AttributeValue](api-attributevalue.md) object

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**EnablePrimaryIpv6**

If you’re modifying a network interface in a dual-stack or IPv6-only subnet, you have
the option to assign a primary IPv6 IP address. A primary IPv6 address is an IPv6 GUA
address associated with an ENI that you have enabled to use a primary IPv6 address. Use
this option if the instance that this ENI will be attached to relies on its IPv6 address
not changing. AWS will automatically assign an IPv6 address associated
with the ENI attached to your instance to be the primary IPv6 address. Once you enable
an IPv6 GUA address to be a primary IPv6, you cannot disable it. When you enable an IPv6
GUA address to be a primary IPv6, the first IPv6 GUA will be made the primary IPv6
address until the instance is terminated or the network interface is detached. If you
have multiple IPv6 addresses associated with an ENI attached to your instance and you
enable a primary IPv6 address, the first IPv6 GUA address associated with the ENI
becomes the primary IPv6 address.

Type: Boolean

Required: No

**EnaSrdSpecification**

Updates the ENA Express configuration for the network interface that’s attached to the
instance.

Type: [EnaSrdSpecification](api-enasrdspecification.md) object

Required: No

**NetworkInterfaceId**

The ID of the network interface.

Type: String

Required: Yes

**SecurityGroupId.N**

Changes the security groups for the network interface. The new set of groups you
specify replaces the current set. You must specify at least one group, even if it's just
the default security group in the VPC. You must specify the ID of the security group,
not the name.

Type: Array of strings

Required: No

**SourceDestCheck**

Enable or disable source/destination checks, which ensure that the instance is either
the source or the destination of any traffic that it receives. If the value is
`true`, source/destination checks are enabled; otherwise, they are
disabled. The default value is `true`. You must disable source/destination
checks if the instance runs services such as network address translation, routing, or
firewalls.

Type: [AttributeBooleanValue](api-attributebooleanvalue.md) object

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

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example sets source/destination checking to `false` for the
specified network interface.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ModifyNetworkInterfaceAttribute
&NetworkInterfaceId=eni-ffda3197
&SourceDestCheck.Value=false
&AUTHPARAMS
```

#### Sample Response

```

<ModifyNetworkInterfaceAttributeResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>657a4623-5620-4232-b03b-427e852d71cf</requestId>
    <return>true</return>
</ModifyNetworkInterfaceAttributeResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ModifyNetworkInterfaceAttribute)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ModifyNetworkInterfaceAttribute)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifynetworkinterfaceattribute.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifynetworkinterfaceattribute.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifynetworkinterfaceattribute.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifynetworkinterfaceattribute.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifynetworkinterfaceattribute.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifynetworkinterfaceattribute.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ModifyNetworkInterfaceAttribute)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifynetworkinterfaceattribute.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyManagedPrefixList

ModifyPrivateDnsNameOptions
