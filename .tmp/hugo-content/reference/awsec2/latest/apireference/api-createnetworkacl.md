# CreateNetworkAcl

Creates a network ACL in a VPC. Network ACLs provide an optional layer of security (in addition to security groups) for the instances in your VPC.

For more information, see [Network ACLs](../../../../services/vpc/latest/userguide/vpc-network-acls.md) in the
_Amazon VPC User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the
request. For more information, see [Ensuring idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**TagSpecification.N**

The tags to assign to the network ACL.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

**VpcId**

The ID of the VPC.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**clientToken**

Unique, case-sensitive identifier to ensure the idempotency of the request. Only returned if a client token was provided in the request.

Type: String

**networkAcl**

Information about the network ACL.

Type: [NetworkAcl](api-networkacl.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example creates a network ACL in the specified IPv6-enabled VPC. The
response includes default IPv4 and IPv6 entries for egress and ingress traffic,
each with a high rule number. These are the last entries we process to decide
whether traffic is allowed in or out of an associated subnet. If the traffic
doesn't match any rules with a lower rule number, then these default entries
ultimately deny the traffic.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CreateNetworkAcl
&VpcId=vpc-11ad4878
&AUTHPARAMS
```

#### Sample Response

```

<CreateNetworkAclResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <networkAcl>
      <networkAclId>acl-5fb85d36</networkAclId>
      <vpcId>vpc-11ad4878</vpcId>
      <default>false</default>
      <entrySet>
         <item>
            <ruleNumber>32767</ruleNumber>
            <protocol>all</protocol>
            <ruleAction>deny</ruleAction>
            <egress>true</egress>
            <cidrBlock>0.0.0.0/0</cidrBlock>
         </item>
         <item>
            <ruleNumber>32767</ruleNumber>
            <protocol>all</protocol>
            <ruleAction>deny</ruleAction>
            <egress>false</egress>
            <cidrBlock>0.0.0.0/0</cidrBlock>
         </item>
        <item>
           <ruleNumber>32768</ruleNumber>
           <protocol>all</protocol>
           <ruleAction>deny</ruleAction>
           <egress>true</egress>
           <ipv6CidrBlock>::/0</ipv6CidrBlock>
        </item>
        <item>
           <ruleNumber>32768</ruleNumber>
           <protocol>all</protocol>
           <ruleAction>deny</ruleAction>
           <egress>false</egress>
           <ipv6CidrBlock>::/0</ipv6CidrBlock>
        </item>
      </entrySet>
      <associationSet/>
      <tagSet/>
   </networkAcl>
</CreateNetworkAclResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/createnetworkacl.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/createnetworkacl.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createnetworkacl.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createnetworkacl.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createnetworkacl.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createnetworkacl.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createnetworkacl.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createnetworkacl.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/createnetworkacl.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createnetworkacl.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateNatGateway

CreateNetworkAclEntry
