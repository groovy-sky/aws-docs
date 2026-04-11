This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::LayerVersionPermission

The `AWS::Lambda::LayerVersionPermission` resource adds permissions to the resource-based policy of
a version of an [Lambda\
layer](../../../lambda/latest/dg/configuration-layers.md). Use this action to grant layer usage permission to other accounts. You can grant permission to a
single account, all AWS accounts, or all accounts in an organization.

###### Important

Since the release of the [UpdateReplacePolicy](../userguide/aws-attribute-updatereplacepolicy.md) both `UpdateReplacePolicy` and `DeletionPolicy` are required to protect your Resources/LayerPermissions from deletion.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lambda::LayerVersionPermission",
  "Properties" : {
      "Action" : String,
      "LayerVersionArn" : String,
      "OrganizationId" : String,
      "Principal" : String
    }
}

```

### YAML

```yaml

Type: AWS::Lambda::LayerVersionPermission
Properties:
  Action: String
  LayerVersionArn: String
  OrganizationId: String
  Principal: String

```

## Properties

`Action`

The API action that grants access to the layer. For example, `lambda:GetLayerVersion`.

_Required_: Yes

_Type_: String

_Pattern_: `lambda:GetLayerVersion`

_Minimum_: `0`

_Maximum_: `22`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LayerVersionArn`

The name or Amazon Resource Name (ARN) of the layer.

_Required_: Yes

_Type_: String

_Pattern_: `(arn:[a-zA-Z0-9-]+:lambda:[a-zA-Z0-9-]+:\d{12}:layer:[a-zA-Z0-9-_]+)|[a-zA-Z0-9-_]+`

_Minimum_: `1`

_Maximum_: `140`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OrganizationId`

With the principal set to `*`, grant permission to all accounts in the specified
organization.

_Required_: No

_Type_: String

_Pattern_: `o-[a-z0-9]{10,32}`

_Minimum_: `0`

_Maximum_: `34`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Principal`

An account ID, or `*` to grant layer usage permission to all
accounts in an organization, or all AWS accounts (if `organizationId` is not specified).
For the last case, make sure that you really do want all AWS accounts to have usage permission to this layer.

_Required_: Yes

_Type_: String

_Pattern_: `\d{12}|\*|arn:(aws[a-zA-Z-]*):iam::\d{12}:root`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the layer version ARN and statement ID, such as
`arn:aws:lambda:us-east-2:123456789012:layer:my-layer:1#engineering-org`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

### Layer Version Permission

Grant layer use permission to accounts in organization `o-t194hfs8cz`.

#### JSON

```json

"MyLayerPermission": {
    "Type": "AWS::Lambda::LayerVersionPermission",
    "Properties": {
        "Action": "lambda:GetLayerVersion",
        "LayerVersionArn": "arn:aws:lambda:us-east-2:123456789012:layer:my-layer:1",
        "OrganizationId": "o-t194hfs8cz",
        "Principal": "*"
    }
}
```

#### YAML

```yaml

MyLayerPermission:
  Type: AWS::Lambda::LayerVersionPermission
  Properties:
    Action: lambda:GetLayerVersion
    LayerVersionArn: arn:aws:lambda:us-east-2:123456789012:layer:my-layer:1
    OrganizationId: o-t194hfs8cz
    Principal: *
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Content

AWS::Lambda::Permission

All content copied from https://docs.aws.amazon.com/.
