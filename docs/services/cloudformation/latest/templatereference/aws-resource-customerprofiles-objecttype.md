This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::ObjectType

Specifies an Amazon Connect Customer Profiles Object Type Mapping.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CustomerProfiles::ObjectType",
  "Properties" : {
      "AllowProfileCreation" : Boolean,
      "Description" : String,
      "DomainName" : String,
      "EncryptionKey" : String,
      "ExpirationDays" : Integer,
      "Fields" : [ FieldMap, ... ],
      "Keys" : [ KeyMap, ... ],
      "MaxProfileObjectCount" : Integer,
      "ObjectTypeName" : String,
      "SourceLastUpdatedTimestampFormat" : String,
      "SourcePriority" : Integer,
      "Tags" : [ Tag, ... ],
      "TemplateId" : String
    }
}

```

### YAML

```yaml

Type: AWS::CustomerProfiles::ObjectType
Properties:
  AllowProfileCreation: Boolean
  Description: String
  DomainName: String
  EncryptionKey: String
  ExpirationDays: Integer
  Fields:
    - FieldMap
  Keys:
    - KeyMap
  MaxProfileObjectCount: Integer
  ObjectTypeName: String
  SourceLastUpdatedTimestampFormat: String
  SourcePriority: Integer
  Tags:
    - Tag
  TemplateId: String

```

## Properties

`AllowProfileCreation`

Indicates whether a profile should be created when data is received if one doesn’t
exist for an object of this type. The default is `FALSE`. If the
AllowProfileCreation flag is set to `FALSE`, then the service tries to fetch
a standard profile and associate this object with the profile. If it is set to
`TRUE`, and if no match is found, then the service creates a new standard
profile.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the profile object type mapping.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainName`

The unique name of the domain.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EncryptionKey`

The customer-provided key to encrypt the profile object that will be created in this
profile object type mapping. If not specified the system will use the encryption key of
the domain.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExpirationDays`

The number of days until the data of this type expires.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `1098`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Fields`

A list of field definitions for the object type mapping.

_Required_: No

_Type_: Array of [FieldMap](aws-properties-customerprofiles-objecttype-fieldmap.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Keys`

A list of keys that can be used to map data to the profile or search for the
profile.

_Required_: No

_Type_: Array of [KeyMap](aws-properties-customerprofiles-objecttype-keymap.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxProfileObjectCount`

The amount of profile object max count assigned to the object type.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObjectTypeName`

The name of the profile object type.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z_][a-zA-Z_0-9-]*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceLastUpdatedTimestampFormat`

The format of your sourceLastUpdatedTimestamp that was previously set up.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourcePriority`

Property description not available.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for this resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-customerprofiles-objecttype-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateId`

A unique identifier for the template mapping. This can be used instead of specifying
the Keys and Fields properties directly.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the DomainName and the ObjectTypeName of the object
type.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The timestamp of when the object type was created.

`LastUpdatedAt`

The timestamp of when the object type was most recently edited.

`MaxAvailableProfileObjectCount`

The amount of provisioned profile object max count available.

## Examples

The following example creates a object type if Domain existed.

#### YAML

```yaml

Type: "AWS::CustomerProfiles::ObjectType"
Properties:
    DomainName: "ExampleDomain"
    ObjectTypeName: "ExampleObjectType"
    AllowProfileCreation: false
    Description: "Description Example"
    ExpirationDays: 1
    Fields:
      - Name: "email"
        ObjectTypeField:
          Source: "_source.email"
          Target: "_profile.BusinessEmail"
          ContentType: "EMAIL_ADDRESS"
    Keys:
      - Name: "_email"
        ObjectTypeKeyList:
          - FieldNames:
            - "email"
            StandardIdentifiers:
              - "PROFILE"
              - "UNIQUE"
```

#### JSON

```json

"Type": "AWS::CustomerProfiles::ObjectType",
"Properties": {
    "DomainName": "ExampleDomain",
    "ObjectTypeName": "ExampleObjectType",
    "AllowProfileCreation": false,
    "Description": "Description Example",
    "ExpirationDays": 1,
    "Fields": [{
        "Name": "email",
        "ObjectTypeField": {
            "Source": "_source.email",
            "Target": "_profile.BusinessEmail",
            "ContentType": "EMAIL_ADDRESS"
        }
    }],
    "Keys": [{
        "Name": "_email",
        "ObjectTypeKeyList": [{
            "FieldNames": [ "email" ],
            "StandardIdentifiers": [ "PROFILE", "UNIQUE" ]
        }]
    }]
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ZendeskSourceProperties

FieldMap

All content copied from https://docs.aws.amazon.com/.
