# DescribeIdentityIdFormat

Describes the ID format settings for resources for the specified IAM user, IAM role, or root
user. For example, you can view the resource types that are enabled for longer IDs. This request only
returns information about resource types whose ID formats can be modified; it does not return
information about other resource types. For more information, see [Resource IDs](../../../../services/ec2/latest/userguide/resource-ids.md) in the _Amazon Elastic Compute Cloud User Guide_.

The following resource types support longer IDs: `bundle` \|
`conversion-task` \| `customer-gateway` \| `dhcp-options` \|
`elastic-ip-allocation` \| `elastic-ip-association` \|
`export-task` \| `flow-log` \| `image` \|
`import-task` \| `instance` \| `internet-gateway` \|
`network-acl` \| `network-acl-association` \|
`network-interface` \| `network-interface-attachment` \|
`prefix-list` \| `reservation` \| `route-table` \|
`route-table-association` \| `security-group` \|
`snapshot` \| `subnet` \|
`subnet-cidr-block-association` \| `volume` \| `vpc`
\| `vpc-cidr-block-association` \| `vpc-endpoint` \|
`vpc-peering-connection` \| `vpn-connection` \| `vpn-gateway`.

These settings apply to the principal specified in the request. They do not apply to the
principal that makes the request.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**PrincipalArn**

The ARN of the principal, which can be an IAM role, IAM user, or the root user.

Type: String

Required: Yes

**Resource**

The type of resource: `bundle` \|
`conversion-task` \| `customer-gateway` \| `dhcp-options` \|
`elastic-ip-allocation` \| `elastic-ip-association` \|
`export-task` \| `flow-log` \| `image` \|
`import-task` \| `instance` \| `internet-gateway` \|
`network-acl` \| `network-acl-association` \|
`network-interface` \| `network-interface-attachment` \|
`prefix-list` \| `reservation` \| `route-table` \|
`route-table-association` \| `security-group` \|
`snapshot` \| `subnet` \|
`subnet-cidr-block-association` \| `volume` \| `vpc`
\| `vpc-cidr-block-association` \| `vpc-endpoint` \|
`vpc-peering-connection` \| `vpn-connection` \| `vpn-gateway`

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**statusSet**

Information about the ID format for the resources.

Type: Array of [IdFormat](api-idformat.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example describes the ID format for the IAM role 'EC2Role'.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeIdentityIdFormat
&PrincipalArn=arn:aws:iam::123456789012:role/EC2Role
&AUTHPARAMS
```

#### Sample Response

```

<DescribeIdentityIdFormatResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>7a62c49f-347e-4fc4-9331-6e8eEXAMPLE</requestId>
    <statusSet>
        <item>
            <useLongIds>true</useLongIds>
            <deadline>2016-12-15T12:00:00Z</deadline>
            <resource>reservation</resource>
        </item>
        <item>
            <useLongIds>true</useLongIds>
            <deadline>2016-12-15T12:00:00Z</deadline>
            <resource>instance</resource>
        </item>
        <item>
            <useLongIds>true</useLongIds>
            <deadline>2016-12-15T12:00:00Z</deadline>
            <resource>volume</resource>
        </item>
        <item>
            <useLongIds>true</useLongIds>
            <deadline>2016-12-15T12:00:00Z</deadline>
            <resource>snapshot</resource>
        </item>
        <item>
            <useLongIds>true</useLongIds>
            <resource>network-interface-attachment</resource>
        </item>
        <item>
            <useLongIds>true</useLongIds>
            <resource>network-interface</resource>
        </item>
        <item>
            <useLongIds>true</useLongIds>
            <resource>elastic-ip-allocation</resource>
        </item>
        <item>
            <useLongIds>true</useLongIds>
            <resource>elastic-ip-association</resource>
        </item>
        <item>
            <useLongIds>true</useLongIds>
            <resource>vpc</resource>
        </item>
        <item>
            <useLongIds>true</useLongIds>
            <resource>subnet</resource>
        </item>
        <item>
            <useLongIds>true</useLongIds>
            <resource>route-table</resource>
        </item>
        <item>
            <useLongIds>true</useLongIds>
            <resource>route-table-association</resource>
        </item>
        <item>
            <useLongIds>true</useLongIds>
            <resource>network-acl</resource>
        </item>
        <item>
            <useLongIds>true</useLongIds>
            <resource>network-acl-association</resource>
        </item>
        <item>
            <useLongIds>true</useLongIds>
            <resource>dhcp-options</resource>
        </item>
        <item>
            <useLongIds>true</useLongIds>
            <resource>internet-gateway</resource>
        </item>
        <item>
            <useLongIds>false</useLongIds>
            <resource>vpc-cidr-block-association</resource>
        </item>
        <item>
            <useLongIds>false</useLongIds>
            <resource>vpc-ipv6-cidr-block-association</resource>
        </item>
        <item>
            <useLongIds>true</useLongIds>
            <resource>subnet-ipv6-cidr-block-association</resource>
        </item>
        <item>
            <useLongIds>false</useLongIds>
            <resource>vpc-peering-connection</resource>
        </item>
        <item>
            <useLongIds>true</useLongIds>
            <resource>security-group</resource>
        </item>
        <item>
            <useLongIds>true</useLongIds>
            <resource>flow-log</resource>
        </item>
        <item>
            <useLongIds>true</useLongIds>
            <resource>customer-gateway</resource>
        </item>
        <item>
            <useLongIds>true</useLongIds>
            <resource>vpc-endpoint</resource>
        </item>
        <item>
            <useLongIds>true</useLongIds>
            <resource>vpn-connection</resource>
        </item>
        <item>
            <useLongIds>true</useLongIds>
            <resource>vpn-gateway</resource>
        </item>
    </statusSet>
</DescribeIdentityIdFormatResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describeidentityidformat.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describeidentityidformat.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describeidentityidformat.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describeidentityidformat.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describeidentityidformat.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describeidentityidformat.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describeidentityidformat.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describeidentityidformat.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describeidentityidformat.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describeidentityidformat.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeIamInstanceProfileAssociations

DescribeIdFormat
