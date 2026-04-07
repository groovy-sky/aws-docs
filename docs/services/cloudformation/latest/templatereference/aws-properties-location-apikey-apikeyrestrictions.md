This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Location::APIKey ApiKeyRestrictions

API Restrictions on the allowed actions, resources, and referers for an API key
resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowActions" : [ String, ... ],
  "AllowAndroidApps" : [ AndroidApp, ... ],
  "AllowAppleApps" : [ AppleApp, ... ],
  "AllowReferers" : [ String, ... ],
  "AllowResources" : [ String, ... ]
}

```

### YAML

```yaml

  AllowActions:
    - String
  AllowAndroidApps:
    - AndroidApp
  AllowAppleApps:
    - AppleApp
  AllowReferers:
    - String
  AllowResources:
    - String

```

## Properties

`AllowActions`

A list of allowed actions that an API key resource grants permissions to
perform. You must have at least one action for each type of resource. For example,
if you have a place resource, you must include at least one place action.

The following are valid values for the actions.

- **Map actions**

- `geo:GetMap*` \- Allows all actions needed for map rendering.

- **Enhanced Maps actions**

- `geo-maps:GetTile` \- Allows getting map tiles for rendering.

- `geo-maps:GetStaticMap` \- Allows getting static map images.

- **Place actions**

- `geo:SearchPlaceIndexForText` \- Allows finding geo coordinates of a known place.

- `geo:SearchPlaceIndexForPosition` \- Allows getting nearest address to geo coordinates.

- `geo:SearchPlaceIndexForSuggestions` \- Allows suggestions based on an incomplete or misspelled query.

- `geo:GetPlace` \- Allows getting details of a place.

- **Enhanced Places actions**

- `geo-places:Autocomplete` \- Allows auto-completion of search text.

- `geo-places:Geocode` \- Allows finding geo coordinates of a known place.

- `geo-places:GetPlace` \- Allows getting details of a place.

- `geo-places:ReverseGeocode` \- Allows getting nearest address to geo coordinates.

- `geo-places:SearchNearby` \- Allows category based places search around geo coordinates.

- `geo-places:SearchText` \- Allows place or address search based on free-form text.

- `geo-places:Suggest` \- Allows suggestions based on an incomplete or misspelled query.

- **Route actions**

- `geo:CalculateRoute` \- Allows point to point routing.

- `geo:CalculateRouteMatrix` \- Allows matrix routing.

- **Enhanced Routes actions**

- `geo-routes:CalculateIsolines` \- Allows isoline calculation.

- `geo-routes:CalculateRoutes` \- Allows point to point routing.

- `geo-routes:CalculateRouteMatrix` \- Allows matrix routing.

- `geo-routes:OptimizeWaypoints` \- Allows computing the best sequence of waypoints.

- `geo-routes:SnapToRoads` \- Allows snapping GPS points to a likely route.

###### Note

You must use these strings exactly. For example, to provide access to map
rendering, the only valid action is `geo:GetMap*` as an input to
the list. `["geo:GetMap*"]` is valid but
`["geo:GetTile"]` is not. Similarly, you cannot use
`["geo:SearchPlaceIndexFor*"]` \- you must list each of the Place
actions separately.

_Required_: Yes

_Type_: Array of String

_Minimum_: `5 | 1`

_Maximum_: `200 | 24`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowAndroidApps`

Property description not available.

_Required_: No

_Type_: Array of [AndroidApp](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-location-apikey-androidapp.html)

_Minimum_: `0`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowAppleApps`

Property description not available.

_Required_: No

_Type_: Array of [AppleApp](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-location-apikey-appleapp.html)

_Minimum_: `0`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowReferers`

An optional list of allowed HTTP referers for which requests must originate from.
Requests using this API key from other domains will not be allowed.

Requirements:

- Contain only alphanumeric characters (A–Z, a–z, 0–9) or any symbols in this
list ``$\-._+!*`(),;/?:@=&``

- May contain a percent (%) if followed by 2 hexadecimal digits (A-F, a-f, 0-9);
this is used for URL encoding purposes.

- May contain wildcard characters question mark (?) and asterisk (\*).

Question mark (?) will replace any single character (including hexadecimal
digits).

Asterisk (\*) will replace any multiple characters (including multiple
hexadecimal digits).

- No spaces allowed. For example, `https://example.com`.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `253 | 5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowResources`

A list of allowed resource ARNs that a API key bearer can perform actions on.

- The ARN must be the correct ARN for a map, place, or route ARN. You may
include wildcards in the resource-id to match multiple resources of the same
type.

- The resources must be in the same `partition`, `region`,
and `account-id` as the key that is being created.

- Other than wildcards, you must include the full ARN, including the
`arn`, `partition`, `service`,
`region`, `account-id` and `resource-id`
delimited by colons (:).

- No spaces allowed, even with wildcards. For example,
`arn:aws:geo:region:account-id:map/ExampleMap*`.

For more information about ARN format, see [Amazon Resource Names\
(ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1600 | 8`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Allowing all Enhanced actions](#aws-properties-location-apikey-apikeyrestrictions--examples--Allowing_all_Enhanced_actions)

- [Allowing the GetStaticMap Enhanced Map action](#aws-properties-location-apikey-apikeyrestrictions--examples--Allowing_the_GetStaticMap_Enhanced_Map_action)

- [Allowing the SearchPlaceIndexForText Place action](#aws-properties-location-apikey-apikeyrestrictions--examples--Allowing_the_SearchPlaceIndexForText_Place_action)

### Allowing all Enhanced actions

The following example creates an API key that allows all Enhanced Maps, Enhanced Places and Enhanced Routes actions.

#### JSON

```json

{
  "Resources": {
    "ExampleKey": {
      "Properties": {
        "KeyName": "ExampleKey",
        "NoExpiry": true,
        "Restrictions": {
          "AllowActions": [
            "geo-maps:*",
            "geo-places:*",
            "geo-routes:*"
          ],
          "AllowResources": [
            {
              "Fn::Sub": [
                "arn:aws:geo-maps:${Region}::provider/default",
                { "Region": {"Ref": "AWS::Region"} }
              ]
            },
            {
              "Fn::Sub": [
                "arn:aws:geo-places:${Region}::provider/default",
                { "Region": {"Ref": "AWS::Region"} }
              ]
            },
            {
              "Fn::Sub": [
                "arn:aws:geo-routes:${Region}::provider/default",
                { "Region": {"Ref": "AWS::Region"} }
              ]
            },
          ]
        }
      },
      "Type": "AWS::Location::APIKey"
    }
  }
}
```

#### YAML

```yaml

Resources:
  ExampleKey:
    Type: "AWS::Location::APIKey"
    Properties:
      KeyName: "ExampleKey"
      NoExpiry: true
      Restrictions:
        AllowActions:
          - "geo-maps:*"
          - "geo-places:*"
          - "geo-routes:*"
        AllowResources:
          - Fn::Sub:
              - "arn:aws:geo-maps:${Region}::provider/default"
              - Region:
                  Ref: "AWS::Region"
          - Fn::Sub:
              - "arn:aws:geo-places:${Region}::provider/default"
              - Region:
                  Ref: "AWS::Region"
          - Fn::Sub:
              - "arn:aws:geo-routes:${Region}::provider/default"
              - Region:
                  Ref: "AWS::Region"
```

### Allowing the `GetStaticMap` Enhanced Map action

The following example creates an API key that allows the Enhanced Maps action `GetStaticMap`.

#### JSON

```json

{
  "Resources": {
    "ExampleKey": {
      "Properties": {
        "KeyName": "ExampleKey",
        "NoExpiry": true,
        "Restrictions": {
          "AllowActions": [
            "geo-maps:GetStaticMap",
          ],
          "AllowResources": [
            {
              "Fn::Sub": [
                "arn:aws:geo-maps:${Region}::provider/default",
                { "Region": {"Ref": "AWS::Region"} }
              ]
            }
          ]
        }
      },
      "Type": "AWS::Location::APIKey"
    }
  }
}
```

#### YAML

```yaml

Resources:
  ExampleKey:
    Type: "AWS::Location::APIKey"
    Properties:
      KeyName: "ExampleKey"
      NoExpiry: true
      Restrictions:
        AllowActions:
          - "geo-maps:GetStaticMap"
        AllowResources:
          - Fn::Sub:
              - "arn:aws:geo-maps:${Region}::provider/default"
              - Region:
                  Ref: "AWS::Region"
```

### Allowing the `SearchPlaceIndexForText` Place action

The following example creates an API key that allows the Place action
`SearchPlaceIndexForText` on the `place-name`
Place index.

#### JSON

```json

{
  "Resources": {
    "ExampleKey": {
      "Properties": {
        "KeyName": "ExampleKey",
        "NoExpiry": true,
        "Restrictions": {
          "AllowActions": [
            "geo:SearchPlaceIndexForText"
          ],
          "AllowResources": [
            {
              "Fn::Sub": [
                "arn:aws:geo:${Region}:${Account}:place-index/place-name",
                {
                  "Region": {"Ref": "AWS::Region"},
                  "Account": {"Ref": "AWS::AccountId"}
                }
              ]
            }
          ]
        }
      },
      "Type": "AWS::Location::APIKey"
    }
  }
}
```

#### YAML

```yaml

Resources:
  ExampleKey:
    Type: "AWS::Location::APIKey"
    Properties:
      KeyName: "ExampleKey"
      NoExpiry: true
      Restrictions:
        AllowActions:
          - "geo:SearchPlaceIndexForText"
        AllowResources:
          - Fn::Sub:
              - "arn:aws:geo:${Region}:${Account}:place-index/place-name"
              - Region:
                  Ref: "AWS::Region"
                Account:
                  Ref: "AWS::AccountId"
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AndroidApp

AppleApp
