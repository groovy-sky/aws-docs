# RevokeSecurityGroupIngress

Removes the specified inbound (ingress) rules from a security group.

You can specify rules using either rule IDs or security group rule properties. If you use
rule properties, the values that you specify (for example, ports) must match the existing rule's
values exactly. Each rule has a protocol, from and to ports, and source (CIDR range,
security group, or prefix list). For the TCP and UDP protocols, you must also specify the
destination port or range of ports. For the ICMP protocol, you must also specify the ICMP type
and code. If the security group rule has a description, you do not need to specify the description
to revoke the rule.

For a default VPC, if the values you specify do not match the existing rule's values,
no error is returned, and the output describes the security group rules that were not
revoked.

For a non-default VPC, if the values you specify do not match the existing rule's
values, an `InvalidPermission.NotFound` client error is returned, and no
rules are revoked.

AWS recommends that you describe the security group to verify that the rules were removed.

Rule changes are propagated to instances within the security group as quickly as possible.
However, a small delay might occur.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CidrIp**

The CIDR IP address range. You can't specify this parameter when specifying a source security group.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**FromPort**

If the protocol is TCP or UDP, this is the start of the port range.
If the protocol is ICMP, this is the ICMP type or -1 (all ICMP types).

Type: Integer

Required: No

**GroupId**

The ID of the security group.

Type: String

Required: No

**GroupName**

\[Default VPC\] The name of the security group. You must specify either the
security group ID or the security group name in the request. For security groups in a
nondefault VPC, you must specify the security group ID.

Type: String

Required: No

**IpPermissions.N**

The sets of IP permissions. You can't specify a source security group and a CIDR IP address range in the same set of permissions.

Type: Array of [IpPermission](api-ippermission.md) objects

Required: No

**IpProtocol**

The IP protocol name ( `tcp`, `udp`, `icmp`) or number
(see [Protocol Numbers](http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)).
Use `-1` to specify all.

Type: String

Required: No

**SecurityGroupRuleId.N**

The IDs of the security group rules.

Type: Array of strings

Required: No

**SourceSecurityGroupName**

\[Default VPC\] The name of the source security group. You can't specify this parameter
in combination with the following parameters: the CIDR IP address range, the start of the port range,
the IP protocol, and the end of the port range. The source security group must be in the same VPC.
To revoke a specific rule for an IP protocol and port range, use a set of IP permissions instead.

Type: String

Required: No

**SourceSecurityGroupOwnerId**

Not supported.

Type: String

Required: No

**ToPort**

If the protocol is TCP or UDP, this is the end of the port range.
If the protocol is ICMP, this is the ICMP code or -1 (all ICMP codes).

Type: Integer

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**return**

Returns `true` if the request succeeds; otherwise, returns an error.

Type: Boolean

**revokedSecurityGroupRuleSet**

Details about the revoked security group rules.

Type: Array of [RevokedSecurityGroupRule](api-revokedsecuritygrouprule.md) objects

**unknownIpPermissionSet**

The inbound rules that were unknown to the service. In some cases,
`unknownIpPermissionSet` might be in a different format from the request
parameter.

Type: Array of [IpPermission](api-ippermission.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example revokes TCP port 80 access from the `205.192.0.0/16`
IPv4 address range for the security group named `websrv`. If the security group
is for a VPC, specify the ID of the security group instead of the name.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=RevokeSecurityGroupIngress
&GroupName=websrv
&IpPermissions.1.IpProtocol=tcp
&IpPermissions.1.FromPort=80
&IpPermissions.1.ToPort=80
&IpPermissions.1.IpRanges.1.CidrIp=205.192.0.0/16
&AUTHPARAMS
```

#### Sample Response

```

<RevokeSecurityGroupIngressResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
  <return>true</return>
</RevokeSecurityGroupIngressResponse>
```

### Example 2

This example revokes TCP port 22 (SSH) access from IPv6 range
`2001:db8:1234:1a00::/64`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=RevokeSecurityGroupIngress
&GroupName=websrv
&IpPermissions.1.IpProtocol=tcp
&IpPermissions.1.FromPort=80
&IpPermissions.1.ToPort=80
&IpPermissions.1.Ipv6Ranges.1.CidrIpv6=2001:db8:1234:1a00::/64
&AUTHPARAMS
```

### Example 3

This example revokes TCP port 22 access from the `203.0.113.4/32`
address range for the security group `sg-112233`. The security group rule
includes the description 'Access from workstation 1a2b'. Specifying the description to
revoke the rule is optional, but if you do specify the description, it must match the
existing rule's value exactly.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=RevokeSecurityGroupIngress
&GroupId=sg-112233
&IpPermissions.1.IpProtocol=tcp
&IpPermissions.1.FromPort=22
&IpPermissions.1.ToPort=22
&IpPermissions.1.IpRanges.1.CidrIp=203.0.113.4/32
&IpPermissions.1.IpRanges.1.Description=Access from workstation 1a2b
&AUTHPARAMS
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/revokesecuritygroupingress.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/revokesecuritygroupingress.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/revokesecuritygroupingress.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/revokesecuritygroupingress.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/revokesecuritygroupingress.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/revokesecuritygroupingress.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/revokesecuritygroupingress.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/revokesecuritygroupingress.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/revokesecuritygroupingress.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/revokesecuritygroupingress.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RevokeSecurityGroupEgress

RunInstances
