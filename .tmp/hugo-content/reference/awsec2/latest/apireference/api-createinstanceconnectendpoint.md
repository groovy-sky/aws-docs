# CreateInstanceConnectEndpoint

Creates an EC2 Instance Connect Endpoint.

An EC2 Instance Connect Endpoint allows you to connect to an instance, without
requiring the instance to have a public IPv4 or public IPv6 address. For more
information, see [Connect to your instances using EC2 Instance Connect Endpoint](../../../../services/ec2/latest/userguide/connect-using-ec2-instance-connect-endpoint.md) in the
_Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**IpAddressType**

The IP address type of the endpoint.

If no value is specified, the default value is determined by the IP address type of
the subnet:

- `dualstack` \- If the subnet has both IPv4 and IPv6 CIDRs

- `ipv4` \- If the subnet has only IPv4 CIDRs

- `ipv6` \- If the subnet has only IPv6 CIDRs

###### Note

`PreserveClientIp` is only supported on IPv4 EC2 Instance Connect
Endpoints. To use `PreserveClientIp`, the value for
`IpAddressType` must be `ipv4`.

Type: String

Valid Values: `ipv4 | dualstack | ipv6`

Required: No

**PreserveClientIp**

Indicates whether the client IP address is preserved as the source. The following are the possible values.

- `true` \- Use the client IP address as the source.

- `false` \- Use the network interface IP address as the source.

###### Note

`PreserveClientIp` is only supported on IPv4 EC2 Instance Connect
Endpoints. To use `PreserveClientIp`, the value for
`IpAddressType` must be `ipv4`.

Default: `false`

Type: Boolean

Required: No

**SecurityGroupId.N**

One or more security groups to associate with the endpoint. If you don't specify a security group,
the default security group for your VPC will be associated with the endpoint.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 16 items.

Required: No

**SubnetId**

The ID of the subnet in which to create the EC2 Instance Connect Endpoint.

Type: String

Required: Yes

**TagSpecification.N**

The tags to apply to the EC2 Instance Connect Endpoint during creation.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**clientToken**

Unique, case-sensitive idempotency token provided by the client in the the request.

Type: String

**instanceConnectEndpoint**

Information about the EC2 Instance Connect Endpoint.

Type: [Ec2InstanceConnectEndpoint](api-ec2instanceconnectendpoint.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example: Create an EC2 Instance Connect Endpoint

This example creates an EC2 Instance Connect Endpoint in the specified
subnet.

#### Sample Request

```https

https://ec2.amazonaws.com/?Action=CreateInstanceConnectEndpoint
&SubnetId=subnet-0123456789example
&AUTHPARAMS
```

#### Sample Response

```https

<CreateInstanceConnectEndpointResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>a9effc88-e86c-4730-b1e8-30ba091babd1</requestId>
  <instanceConnectEndpoint>
    <createdAt>2023-06-06T20:01:31.000Z</createdAt>
    <instanceConnectEndpointArn>arn:aws:ec2:region:account-id:instance-connect-endpoint/eice-0123456789example</instanceConnectEndpointArn>
    <instanceConnectEndpointId>eice-0123456789example</instanceConnectEndpointId>
    <networkInterfaceIdSet />
    <ownerId>account-id</ownerId>
    <preserveClientIp>false</preserveClientIp>
    <securityGroupIdSet>
      <item>sg-0123456789example</item>
    </securityGroupIdSet>
    <state>create-in-progress</state>
    <stateMessage />
    <subnetId>subnet-0123456789example</subnetId>
    <tagSet />
    <vpcId>vpc-0123456789example</vpcId>
  </instanceConnectEndpoint>
</CreateInstanceConnectEndpointResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/createinstanceconnectendpoint.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/createinstanceconnectendpoint.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createinstanceconnectendpoint.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createinstanceconnectendpoint.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createinstanceconnectendpoint.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createinstanceconnectendpoint.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createinstanceconnectendpoint.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createinstanceconnectendpoint.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/createinstanceconnectendpoint.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createinstanceconnectendpoint.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateImageUsageReport

CreateInstanceEventWindow
