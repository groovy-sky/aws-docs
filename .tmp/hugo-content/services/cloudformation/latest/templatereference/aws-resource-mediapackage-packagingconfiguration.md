This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::PackagingConfiguration

Creates a packaging configuration in a packaging group.

The packaging configuration represents a single delivery point for an asset. It determines the format and setting for the egressing content. Specify only one package format per configuration, such as `HlsPackage`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaPackage::PackagingConfiguration",
  "Properties" : {
      "CmafPackage" : CmafPackage,
      "DashPackage" : DashPackage,
      "HlsPackage" : HlsPackage,
      "Id" : String,
      "MssPackage" : MssPackage,
      "PackagingGroupId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MediaPackage::PackagingConfiguration
Properties:
  CmafPackage:
    CmafPackage
  DashPackage:
    DashPackage
  HlsPackage:
    HlsPackage
  Id: String
  MssPackage:
    MssPackage
  PackagingGroupId: String
  Tags:
    - Tag

```

## Properties

`CmafPackage`

Parameters for CMAF packaging.

_Required_: No

_Type_: [CmafPackage](aws-properties-mediapackage-packagingconfiguration-cmafpackage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DashPackage`

Parameters for DASH-ISO packaging.

_Required_: No

_Type_: [DashPackage](aws-properties-mediapackage-packagingconfiguration-dashpackage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HlsPackage`

Parameters for Apple HLS packaging.

_Required_: No

_Type_: [HlsPackage](aws-properties-mediapackage-packagingconfiguration-hlspackage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

Unique identifier that you assign to the packaging configuration.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MssPackage`

Parameters for Microsoft Smooth Streaming packaging.

_Required_: No

_Type_: [MssPackage](aws-properties-mediapackage-packagingconfiguration-msspackage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PackagingGroupId`

The ID of the packaging group associated with this packaging configuration.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to assign to the packaging configuration.

_Required_: No

_Type_: Array of [Tag](aws-properties-mediapackage-packagingconfiguration-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the packaging configuration.

For example: `{ "Ref": "myPackagingConfiguration" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) for the packaging configuration. You can get this from the response to any request to the packaging configuration.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

CmafEncryption

All content copied from https://docs.aws.amazon.com/.
