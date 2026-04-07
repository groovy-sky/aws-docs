This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RolesAnywhere::CRL

The `AWS::RolesAnywhere::CRL` resource Property description not available. for RolesAnywhere.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RolesAnywhere::CRL",
  "Properties" : {
      "CrlData" : String,
      "Enabled" : Boolean,
      "Name" : String,
      "Tags" : [ Tag, ... ],
      "TrustAnchorArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::RolesAnywhere::CRL
Properties:
  CrlData: String
  Enabled: Boolean
  Name: String
  Tags:
    - Tag
  TrustAnchorArn: String

```

## Properties

`CrlData`

The x509 v3 specified certificate revocation list (CRL).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Specifies whether the certificate revocation list (CRL) is enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the certificate revocation list (CRL).

_Required_: Yes

_Type_: String

_Pattern_: `[ a-zA-Z0-9-_]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of tags to attach to the certificate revocation list (CRL).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-rolesanywhere-crl-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrustAnchorArn`

The ARN of the TrustAnchor the certificate revocation list (CRL) will provide revocation for.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(-[^:]+)?:rolesanywhere(:.*){2}(:trust-anchor.*)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns `CrlId`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CrlId`

The unique primary identifier of the Crl

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IAM Roles Anywhere

Tag
