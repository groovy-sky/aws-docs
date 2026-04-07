# CreateDefaultSubnet

Creates a default subnet with a size `/20` IPv4 CIDR block in the
specified Availability Zone in your default VPC. You can have only one default subnet
per Availability Zone. For more information, see [Create a default\
subnet](../../../../services/vpc/latest/userguide/work-with-default-vpc.md#create-default-subnet) in the _Amazon VPC User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AvailabilityZone**

The Availability Zone in which to create the default subnet.

Either `AvailabilityZone` or `AvailabilityZoneId` must be specified,
but not both.

Type: String

Required: No

**AvailabilityZoneId**

The ID of the Availability Zone.

Either `AvailabilityZone` or `AvailabilityZoneId` must be specified,
but not both.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Ipv6Native**

Indicates whether to create an IPv6 only subnet. If you already have a default subnet
for this Availability Zone, you must delete it before you can create an IPv6 only subnet.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**subnet**

Information about the subnet.

Type: [Subnet](api-subnet.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example creates a default subnet in Availability Zone
`us-east-2a`.

#### Sample Request

```

https://ec2.us-east-2.amazonaws.com/?Action=CreateDefaultSubnet
&AvailabilityZone=us-east-2a
&AUTHPARAMS
```

#### Sample Response

```

<CreateDefaultSubnetResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>12e2fb2e-e566-488a-926d-4655example</requestId>
    <subnet>
        <assignIpv6AddressOnCreation>false</assignIpv6AddressOnCreation>
        <availabilityZone>us-east-2a</availabilityZone>
        <availableIpAddressCount>4091</availableIpAddressCount>
        <cidrBlock>172.31.32.0/20</cidrBlock>
        <defaultForAz>true</defaultForAz>
        <ipv6CidrBlockAssociationSet/>
        <mapPublicIpOnLaunch>true</mapPublicIpOnLaunch>
        <state>available</state>
        <subnetId>subnet-111f7123</subnetId>
        <tagSet/>
        <vpcId>vpc-8eaaeabc</vpcId>
    </subnet>
</CreateDefaultSubnetResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateDefaultSubnet)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateDefaultSubnet)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createdefaultsubnet.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createdefaultsubnet.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createdefaultsubnet.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createdefaultsubnet.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createdefaultsubnet.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createdefaultsubnet.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateDefaultSubnet)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createdefaultsubnet.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateCustomerGateway

CreateDefaultVpc
