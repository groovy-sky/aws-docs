# UpdateSecurityGroupRuleDescriptionsIngress

Updates the description of an ingress (inbound) security group rule. You can replace an
existing description, or add a description to a rule that did not have one previously.
You can remove a description for a security group rule by omitting the description
parameter in the request.

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

\[Default VPC\] The name of the security group. You must specify either the
security group ID or the security group name. For security groups in a
nondefault VPC, you must specify the security group ID.

Type: String

Required: No

**IpPermissions.N**

The IP permissions for the security group rule. You must specify either IP permissions
or a description.

Type: Array of [IpPermission](api-ippermission.md) objects

Required: No

**SecurityGroupRuleDescription.N**

The description for the ingress security group rules. You must specify either
a description or IP permissions.

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

### Example 1

This example updates the description for the security group rule that allows
inbound access over port 22 from the `203.0.113.0/16` IPv4 address range. The
description ' `SSH access from ABC office`' replaces any existing description
for the rule.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=UpdateSecurityGroupRuleDescriptionsIngress
&GroupId=sg-112233
&IpPermissions.1.IpProtocol=tcp
&IpPermissions.1.FromPort=22
&IpPermissions.1.ToPort=22
&IpPermissions.1.IpRanges.1.CidrIp=203.0.113.0/16
&IpPermissions.1.IpRanges.1.Description=SSH access from ABC office
&AUTHPARAMS
```

#### Sample Response

```

<UpdateSecurityGroupRuleDescriptionsIngressResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>b4a57536-2e4a-4cbe-82f0-399example</requestId>
    <return>true</return>
</UpdateSecurityGroupRuleDescriptionsIngressResponse>
```

### Example 2

This example removes the description for the specified security group rule.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=UpdateSecurityGroupRuleDescriptionsIngress
&GroupId=sg-112233
&IpPermissions.1.IpProtocol=tcp
&IpPermissions.1.FromPort=22
&IpPermissions.1.ToPort=22
&IpPermissions.1.IpRanges.1.CidrIp=203.0.113.4/32
&AUTHPARAMS
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/updatesecuritygroupruledescriptionsingress.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/updatesecuritygroupruledescriptionsingress.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/updatesecuritygroupruledescriptionsingress.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/updatesecuritygroupruledescriptionsingress.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/updatesecuritygroupruledescriptionsingress.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/updatesecuritygroupruledescriptionsingress.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/updatesecuritygroupruledescriptionsingress.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/updatesecuritygroupruledescriptionsingress.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/updatesecuritygroupruledescriptionsingress.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/updatesecuritygroupruledescriptionsingress.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

UpdateSecurityGroupRuleDescriptionsEgress

WithdrawByoipCidr
