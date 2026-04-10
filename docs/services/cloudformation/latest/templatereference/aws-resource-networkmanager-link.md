This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkManager::Link

Specifies a link for a site.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::NetworkManager::Link",
  "Properties" : {
      "Bandwidth" : Bandwidth,
      "Description" : String,
      "GlobalNetworkId" : String,
      "Provider" : String,
      "SiteId" : String,
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::NetworkManager::Link
Properties:
  Bandwidth:
    Bandwidth
  Description: String
  GlobalNetworkId: String
  Provider: String
  SiteId: String
  Tags:
    - Tag
  Type: String

```

## Properties

`Bandwidth`

The bandwidth for the link.

_Required_: Yes

_Type_: [Bandwidth](aws-properties-networkmanager-link-bandwidth.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the link.

Constraints: Maximum length of 256 characters.

_Required_: No

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlobalNetworkId`

The ID of the global network.

_Required_: Yes

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Provider`

The provider of the link.

Constraints: Maximum length of 128 characters. Cannot include the following characters: \| \ ^

_Required_: No

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SiteId`

The ID of the site.

_Required_: Yes

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags for the link.

_Required_: No

_Type_: Array of [Tag](aws-properties-networkmanager-link-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the link.

Constraints: Maximum length of 128 characters. Cannot include the following characters: \| \ ^

_Required_: No

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the IDs of the global network and link. For example: `global-network-01231231231231231|link-11112222aaaabbbb1`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The date and time that the link was created.

`LinkArn`

The ARN of the link. For example,
`arn:aws:networkmanager::123456789012:link/global-network-01231231231231231/link-11112222aaaabbbb1`.

`LinkId`

The ID of the link. For example, `link-11112222aaaabbbb1`.

`State`

The state of the link.

## Examples

### Link

The following example creates a link in a global network.

#### JSON

```json

{
    "Type": "AWS::NetworkManager::Link",
    "Properties": {
        "Description": "Broadband link",
        "GlobalNetworkId": {
            "Ref": "GlobalNetwork"
        },
        "SiteId": {
            "Fn::GetAtt": [
                "Site",
                "SiteId"
            ]
        },
        "Bandwidth": {
            "DownloadSpeed": 20,
            "UploadSpeed": 20
        },
        "Provider": "AnyCompany",
        "Type": "Broadband",
        "Tags": [
            {
                "Key": "Name",
                "Value": "broadband-link-1"
            }
        ]
    }
}
```

#### YAML

```yaml

Type: AWS::NetworkManager::Link
Properties:
  Description: "Broadband link"
  GlobalNetworkId: !Ref GlobalNetwork
  SiteId: !GetAtt Site.SiteId
  Bandwidth:
    DownloadSpeed: 20
    UploadSpeed: 20
  Provider: "AnyCompany"
  Type: "Broadband"
  Tags:
    - Key: Name
      Value: broadband-link-1
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Bandwidth

All content copied from https://docs.aws.amazon.com/.
