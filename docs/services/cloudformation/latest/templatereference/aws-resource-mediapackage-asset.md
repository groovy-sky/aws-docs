This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::Asset

Creates an asset to ingest VOD content.

After it's created, the asset starts ingesting content and generates playback URLs for the packaging configurations associated with it. When ingest is complete, downstream
devices use the appropriate URL to request VOD content from AWS Elemental MediaPackage.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaPackage::Asset",
  "Properties" : {
      "EgressEndpoints" : [ EgressEndpoint, ... ],
      "Id" : String,
      "PackagingGroupId" : String,
      "ResourceId" : String,
      "SourceArn" : String,
      "SourceRoleArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MediaPackage::Asset
Properties:
  EgressEndpoints:
    - EgressEndpoint
  Id: String
  PackagingGroupId: String
  ResourceId: String
  SourceArn: String
  SourceRoleArn: String
  Tags:
    - Tag

```

## Properties

`EgressEndpoints`

List of playback endpoints that are available for this asset.

_Required_: No

_Type_: Array of [EgressEndpoint](aws-properties-mediapackage-asset-egressendpoint.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

Unique identifier that you assign to the asset.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PackagingGroupId`

The ID of the packaging group associated with this asset.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceId`

Unique identifier for this asset, as it's configured in the key provider service.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceArn`

The ARN for the source content in Amazon S3.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceRoleArn`

The ARN for the IAM role that provides AWS Elemental MediaPackage access to the Amazon S3 bucket where the source content is stored. Valid format: arn:aws:iam::{accountID}:role/{name}

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to assign to the asset.

_Required_: No

_Type_: Array of [Tag](aws-properties-mediapackage-asset-tag.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the asset.

For example: `{ "Ref": "myAsset" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) for the asset. You can get this from the response to any request to the asset.

`CreatedAt`

The date and time that the asset was initially submitted for ingest.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Elemental MediaPackage

EgressEndpoint

All content copied from https://docs.aws.amazon.com/.
