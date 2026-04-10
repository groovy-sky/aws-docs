This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Location::APIKey

The API key resource in your AWS account, which lets you grant actions for Amazon Location resources to the API key bearer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Location::APIKey",
  "Properties" : {
      "Description" : String,
      "ExpireTime" : String,
      "ForceDelete" : Boolean,
      "ForceUpdate" : Boolean,
      "KeyName" : String,
      "NoExpiry" : Boolean,
      "Restrictions" : ApiKeyRestrictions,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Location::APIKey
Properties:
  Description: String
  ExpireTime: String
  ForceDelete: Boolean
  ForceUpdate: Boolean
  KeyName: String
  NoExpiry: Boolean
  Restrictions:
    ApiKeyRestrictions
  Tags:
    - Tag

```

## Properties

`Description`

Updates the description for the API key resource.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExpireTime`

The optional timestamp for when the API key resource will expire in [ISO 8601 format](https://www.iso.org/iso-8601-date-and-time-format.html).

_Required_: No

_Type_: String

_Pattern_: `^([0-2]\d{3})-(0[0-9]|1[0-2])-([0-2]\d|3[01])T([01]\d|2[0-4]):([0-5]\d):([0-6]\d)((\.\d{3})?)Z$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ForceDelete`

ForceDelete bypasses an API key's expiry conditions and deletes the key. Set the parameter `true` to delete the key or to `false` to not preemptively delete the API key.

Valid values: `true`, or `false`.

###### Note

This action is irreversible. Only use ForceDelete if you are certain the key is no longer in use.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ForceUpdate`

The boolean flag to be included for updating `ExpireTime` or Restrictions details.
Must be set to `true` to update an API key resource that has been used in the past 7 days.
`False` if force update is not preferred.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyName`

A custom name for the API key resource.

Requirements:

- Contain only alphanumeric characters (A–Z, a–z, 0–9), hyphens (-), periods
(.), and underscores (\_).

- Must be a unique API key name.

- No spaces allowed. For example, `ExampleAPIKey`.

_Required_: Yes

_Type_: String

_Pattern_: `^[-._\w]+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NoExpiry`

Whether the API key should expire. Set to `true` to set the API key to have no expiration time.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Restrictions`

The API key restrictions for the API key resource.

_Required_: Yes

_Type_: [ApiKeyRestrictions](aws-properties-location-apikey-apikeyrestrictions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Applies one or more tags to the map resource. A tag is a key-value pair that helps manage, identify, search, and filter your resources by labelling them.

_Required_: No

_Type_: Array of [Tag](aws-properties-location-apikey-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the `APIKey`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) for the resource. Used when you need to specify a resource across all AWS.

`CreateTime`

The timestamp for when the API key resource was created in ISO 8601 format: YYYY-MM-DDThh:mm:ss.sssZ.

`KeyArn`

The Amazon Resource Name (ARN) for the API key resource. Used when you need to specify a resource across all AWS.

`UpdateTime`

The timestamp for when the API key resource was last updated in ISO 8601 format: `YYYY-MM-DDThh:mm:ss.sssZ`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Location Service

AndroidApp

All content copied from https://docs.aws.amazon.com/.
