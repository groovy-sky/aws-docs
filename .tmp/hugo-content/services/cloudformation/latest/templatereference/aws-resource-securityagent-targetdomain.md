This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityAgent::TargetDomain

The `AWS::SecurityAgent::TargetDomain` resource specifies a target domain for security testing. A target domain represents a web domain that you own and want to include in penetration testing scope.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SecurityAgent::TargetDomain",
  "Properties" : {
      "Tags" : [ Tag, ... ],
      "TargetDomainName" : String,
      "VerificationMethod" : String
    }
}

```

### YAML

```yaml

Type: AWS::SecurityAgent::TargetDomain
Properties:
  Tags:
    - Tag
  TargetDomainName: String
  VerificationMethod: String

```

## Properties

`Tags`

The tags to associate with the target domain.

_Required_: No

_Type_: Array of [Tag](aws-properties-securityagent-targetdomain-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetDomainName`

The domain name to register as a target domain.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VerificationMethod`

The method to use for verifying domain ownership. Valid values are DNS\_TXT and HTTP\_ROUTE.

_Required_: Yes

_Type_: String

_Allowed values_: `DNS_TXT | HTTP_ROUTE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the target domain ID. For example:

`{ "Ref": "MyTargetDomain" }`

For the target domain `MyTargetDomain`, `Ref` returns the unique identifier of the target domain.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The date and time when the target domain was registered, in ISO 8601 format. For example: `2024-01-01T00:00:00Z`.

`TargetDomainId`

The unique identifier of the target domain. For example: `td-0123456789abcdef0`.

`VerificationStatus`

The current verification status of the target domain. Valid values are `PENDING`, `VERIFIED`, `FAILED`, and `UNREACHABLE`.

`VerifiedAt`

The date and time when the target domain was last successfully verified, in ISO 8601 format. For example: `2024-01-01T00:00:00Z`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcConfig

DnsVerification

All content copied from https://docs.aws.amazon.com/.
