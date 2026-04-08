# ModifyIdentityIdFormat

Modifies the ID format of a resource for a specified IAM user, IAM role, or the root
user for an account; or all IAM users, IAM roles, and the root user for an account. You can
specify that resources should receive longer IDs (17-character IDs) when they are created.

This request can only be used to modify longer ID settings for resource types that are
within the opt-in period. Resources currently in their opt-in period include:
`bundle` \| `conversion-task` \| `customer-gateway` \| `dhcp-options` \|
`elastic-ip-allocation` \| `elastic-ip-association` \|
`export-task` \| `flow-log` \| `image` \|
`import-task` \| `internet-gateway` \| `network-acl`
\| `network-acl-association` \| `network-interface` \|
`network-interface-attachment` \| `prefix-list` \|
`route-table` \| `route-table-association` \|
`security-group` \| `subnet` \|
`subnet-cidr-block-association` \| `vpc` \|
`vpc-cidr-block-association` \| `vpc-endpoint` \| `vpc-peering-connection` \| `vpn-connection` \| `vpn-gateway`.

For more information, see [Resource IDs](../../../../services/ec2/latest/userguide/resource-ids.md) in the
_Amazon Elastic Compute Cloud User Guide_.

This setting applies to the principal specified in the request; it does not apply to the
principal that makes the request.

Resources created with longer IDs are visible to all IAM roles and users, regardless of these
settings and provided that they have permission to use the relevant `Describe`
command for the resource type.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**PrincipalArn**

The ARN of the principal, which can be an IAM user, IAM role, or the root user. Specify
`all` to modify the ID format for all IAM users, IAM roles, and the root user of
the account.

Type: String

Required: Yes

**Resource**

The type of resource: `bundle` \| `conversion-task` \| `customer-gateway` \| `dhcp-options` \|
`elastic-ip-allocation` \| `elastic-ip-association` \|
`export-task` \| `flow-log` \| `image` \|
`import-task` \| `internet-gateway` \| `network-acl`
\| `network-acl-association` \| `network-interface` \|
`network-interface-attachment` \| `prefix-list` \|
`route-table` \| `route-table-association` \|
`security-group` \| `subnet` \|
`subnet-cidr-block-association` \| `vpc` \|
`vpc-cidr-block-association` \| `vpc-endpoint` \| `vpc-peering-connection` \| `vpn-connection` \| `vpn-gateway`.

Alternatively, use the `all-current` option to include all resource types that are
currently within their opt-in period for longer IDs.

Type: String

Required: Yes

**UseLongIds**

Indicates whether the resource should use longer IDs (17-character IDs)

Type: Boolean

Required: Yes

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

This example sets the UseLongIds parameter to true for instances launched by
the IAM role 'EC2Role'. Instances launched by the IAM role receive longer IDs.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ModifyIdentityFormat
&Resource=instance
&UseLongIds=true
&PrincipalArn=arn:aws:iam::123456789012:role/EC2Role
&AUTHPARAMS
```

#### Sample Response

```

<ModifyIdentityIdFormatResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>92c1af09-cb4c-410e-8a96-EXAMPLE</requestId>
  <return>true</return>
</ModifyIdentityIdFormatResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/modifyidentityidformat.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/modifyidentityidformat.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyidentityidformat.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyidentityidformat.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyidentityidformat.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyidentityidformat.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyidentityidformat.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyidentityidformat.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/modifyidentityidformat.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyidentityidformat.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyHosts

ModifyIdFormat
