# UpdateSecurityGroupRuleDescriptionsEgress

Updates the description of an egress (outbound) security group rule. You
can replace an existing description, or add a description to a rule that did not have one
previously. You can remove a description for a security group rule by omitting the
description parameter in the request.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**GroupId**

The ID of the security group. You must specify either the security group ID or the
security group name in the request. For security groups in a nondefault VPC, you must
specify the security group ID.

Type: String

Required: No

**GroupName**

\[Default VPC\] The name of the security group. You must specify either the security group
ID or the security group name.

Type: String

Required: No

**IpPermissions.N**

The IP permissions for the security group rule. You must specify either the IP permissions
or the description.

Type: Array of [IpPermission](api-ippermission.md) objects

Required: No

**SecurityGroupRuleDescription.N**

The description for the egress security group rules. You must specify either the
description or the IP permissions.

Type: Array of [SecurityGroupRuleDescription](api-securitygroupruledescription.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**return**

Returns `true` if the request succeeds; otherwise, returns an error.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example updates the description for the security group rule that allows
outbound access over port 80 to the `205.192.0.0/16` IPv4 address range. The
description ' `Outbound HTTP access to server 2`' replaces any existing
description for the rule.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=UpdateSecurityGroupRuleDescriptionsEgress
&GroupId=sg-112233
&IpPermissions.1.IpProtocol=tcp
&IpPermissions.1.FromPort=80
&IpPermissions.1.ToPort=80
&IpPermissions.1.IpRanges.1.CidrIp=205.192.0.0/16
&IpPermissions.1.IpRanges.1.Description=Outbound HTTP access to server 2
&AUTHPARAMS
```

#### Sample Response

```

<UpdateSecurityGroupRuleDescriptionsEgressResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>1480cf25-4fbe-4168-aa9c-365example</requestId>
    <return>true</return>
</UpdateSecurityGroupRuleDescriptionsEgressResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/UpdateSecurityGroupRuleDescriptionsEgress)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/UpdateSecurityGroupRuleDescriptionsEgress)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/updatesecuritygroupruledescriptionsegress.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/updatesecuritygroupruledescriptionsegress.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/updatesecuritygroupruledescriptionsegress.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/updatesecuritygroupruledescriptionsegress.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/updatesecuritygroupruledescriptionsegress.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/updatesecuritygroupruledescriptionsegress.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/UpdateSecurityGroupRuleDescriptionsEgress)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/updatesecuritygroupruledescriptionsegress.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

UpdateInterruptibleCapacityReservationAllocation

UpdateSecurityGroupRuleDescriptionsIngress
