This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Location::GeofenceCollection

The `AWS::Location::GeofenceCollection` resource specifies the ability to
detect and act when a tracked device enters or exits a defined geographical boundary
known as a geofence.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Location::GeofenceCollection",
  "Properties" : {
      "CollectionName" : String,
      "Description" : String,
      "KmsKeyId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Location::GeofenceCollection
Properties:
  CollectionName: String
  Description: String
  KmsKeyId: String
  Tags:
    - Tag

```

## Properties

`CollectionName`

A custom name for the geofence collection.

Requirements:

- Contain only alphanumeric characters (A–Z, a–z, 0–9), hyphens (-), periods
(.), and underscores (\_).

- Must be a unique geofence collection name.

- No spaces allowed. For example, `ExampleGeofenceCollection`.

_Required_: Yes

_Type_: String

_Pattern_: `^[-._\w]+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

An optional description for the geofence collection.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

A key identifier for an [AWS KMS customer\
managed key](https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html). Enter a key ID, key ARN, alias name, or alias ARN.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Applies one or more tags to the geofence collection. A tag is a key-value pair helps
manage, identify, search, and filter your resources by labelling them.

Format: `"key" : "value"`

Restrictions:

- Maximum 50 tags per resource

- Each resource tag must be unique with a maximum of one value.

- Maximum key length: 128 Unicode characters in UTF-8

- Maximum value length: 256 Unicode characters in UTF-8

- Can use alphanumeric characters (A–Z, a–z, 0–9), and the following characters:
\+ \- = . \_ : / @.

- Cannot use "aws:" as a prefix for a key.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-location-geofencecollection-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `GeofenceCollection` name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) for the geofence collection resource. Used when you
need to specify a resource across all AWS.

- Format example:
`arn:aws:geo:region:account-id:geofence-collection/ExampleGeofenceCollection`

`CollectionArn`

Synonym for `Arn`. The Amazon Resource Name (ARN) for the geofence collection resource. Used when you
need to specify a resource across all AWS.

- Format example:
`arn:aws:geo:region:account-id:geofence-collection/ExampleGeofenceCollection`

`CreateTime`

The timestamp for when the geofence collection resource was created in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format:
`YYYY-MM-DDThh:mm:ss.sssZ`.

`UpdateTime`

The timestamp for when the geofence collection resource was last updated in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format:
`YYYY-MM-DDThh:mm:ss.sssZ`.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
