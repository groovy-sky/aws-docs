This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkManager::Site

Creates a new site in a global network.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::NetworkManager::Site",
  "Properties" : {
      "Description" : String,
      "GlobalNetworkId" : String,
      "Location" : Location,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::NetworkManager::Site
Properties:
  Description: String
  GlobalNetworkId: String
  Location:
    Location
  Tags:
    - Tag

```

## Properties

`Description`

A description of your site.

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

`Location`

The site location. This information is used for visualization in the Network Manager console. If you specify the address, the latitude and longitude are automatically calculated.

- `Address`: The physical address of the site.

- `Latitude`: The latitude of the site.

- `Longitude`: The longitude of the site.

_Required_: No

_Type_: [Location](aws-properties-networkmanager-site-location.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags for the site.

_Required_: No

_Type_: Array of [Tag](aws-properties-networkmanager-site-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the IDs of the global network and the site. For example: `global-network-01231231231231231|site-444555aaabbb11223`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The time that the site was created.

`SiteArn`

The ARN of the site. For example,
`arn:aws:networkmanager::123456789012:site/global-network-01231231231231231/site-444555aaabbb11223`.

`SiteId`

The ID of the site. For example, `site-444555aaabbb11223`.

`State`

The current state of the site.

## Examples

### Site

The following example creates a site in a global network.

#### JSON

```json

{
    "Type": "AWS::NetworkManager::Site",
    "Properties": {
        "Description": "Chicago office",
        "GlobalNetworkId": {
            "Ref": "GlobalNetwork"
        },
        "Location": {
            "Address": "227 W Monroe St, Chicago, IL 60606",
            "Latitude": "41.880520",
            "Longitude": "-87.634720"
        },
        "Tags": [
            {
                "Key": "Network",
                "Value": "north-america"
            }
        ]
    }
}
```

#### YAML

```yaml

Type: AWS::NetworkManager::Site
Properties:
  Description: "Chicago office"
  GlobalNetworkId: !Ref GlobalNetwork
  Location:
    Address: "227 W Monroe St, Chicago, IL 60606"
    Latitude: "41.880520"
    Longitude: "-87.634720"
  Tags:
    - Key: Network
      Value: north-america
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::NetworkManager::LinkAssociation

Location

All content copied from https://docs.aws.amazon.com/.
