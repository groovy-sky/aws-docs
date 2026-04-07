# CreateSecurityGroup

Creates a security group.

A security group acts as a virtual firewall for your instance to control inbound and outbound traffic.
For more information, see
[Amazon EC2 security groups](../../../../services/ec2/latest/userguide/using-network-security.md) in
the _Amazon EC2 User Guide_ and
[Security groups for your VPC](../../../../services/amazonvpc/latest/userguide/vpc-securitygroups.md) in the
_Amazon VPC User Guide_.

When you create a security group, you specify a friendly name of your choice.
You can't have two security groups for the same VPC with the same name.

You have a default security group for use in your VPC. If you don't specify a security group
when you launch an instance, the instance is launched into the appropriate default security group.
A default security group includes a default rule that grants instances unrestricted network access
to each other.

You can add or remove rules from your security groups using
[AuthorizeSecurityGroupIngress](api-authorizesecuritygroupingress.md),
[AuthorizeSecurityGroupEgress](api-authorizesecuritygroupegress.md),
[RevokeSecurityGroupIngress](api-revokesecuritygroupingress.md), and
[RevokeSecurityGroupEgress](api-revokesecuritygroupegress.md).

For more information about VPC security group limits, see [Amazon VPC Limits](../../../../services/vpc/latest/userguide/amazon-vpc-limits.md).

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**GroupDescription**

A description for the security group.

Constraints: Up to 255 characters in length

Valid characters: a-z, A-Z, 0-9, spaces, and .\_-:/()#,@\[\]+=&;{}!$\*

Type: String

Required: Yes

**GroupName**

The name of the security group. Names are case-insensitive and must be unique within the VPC.

Constraints: Up to 255 characters in length. Can't start with `sg-`.

Valid characters: a-z, A-Z, 0-9, spaces, and .\_-:/()#,@\[\]+=&;{}!$\*

Type: String

Required: Yes

**TagSpecification.N**

The tags to assign to the security group.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

**VpcId**

The ID of the VPC. Required for a nondefault VPC.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**groupId**

The ID of the security group.

Type: String

**requestId**

The ID of the request.

Type: String

**securityGroupArn**

The security group ARN.

Type: String

**tagSet**

The tags assigned to the security group.

Type: Array of [Tag](api-tag.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example creates a security group named `WebServerSG` for the specified VPC.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CreateSecurityGroup
&GroupName=WebServerSG
&GroupDescription=Web Servers
&VpcId=vpc-3325caf2
&AUTHPARAMS
```

#### Sample Response

```

<CreateSecurityGroupResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <return>true</return>
   <groupId>sg-0a42d66a</groupId>
</CreateSecurityGroupResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateSecurityGroup)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateSecurityGroup)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createsecuritygroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createsecuritygroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createsecuritygroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createsecuritygroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createsecuritygroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createsecuritygroup.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateSecurityGroup)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createsecuritygroup.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateSecondarySubnet

CreateSnapshot
