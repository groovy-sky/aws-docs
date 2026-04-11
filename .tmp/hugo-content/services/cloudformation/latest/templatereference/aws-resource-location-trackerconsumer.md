This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Location::TrackerConsumer

The `AWS::Location::TrackerConsumer` resource specifies an association
between a geofence collection and a tracker resource. The geofence collection is
referred to as the _consumer_ of the tracker. This allows the tracker resource
to communicate location data to the linked geofence collection.

###### Note

Currently not supported — Cross-account configurations, such as creating
associations between a tracker resource in one account and a geofence collection in
another account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Location::TrackerConsumer",
  "Properties" : {
      "ConsumerArn" : String,
      "TrackerName" : String
    }
}

```

### YAML

```yaml

Type: AWS::Location::TrackerConsumer
Properties:
  ConsumerArn: String
  TrackerName: String

```

## Properties

`ConsumerArn`

The Amazon Resource Name (ARN) for the geofence collection to be associated to tracker
resource. Used when you need to specify a resource across all AWS.

- Format example:
`arn:aws:geo:region:account-id:geofence-collection/ExampleGeofenceCollectionConsumer`

_Required_: Yes

_Type_: String

_Pattern_: `^arn(:[a-z0-9]+([.-][a-z0-9]+)*){2}(:([a-z0-9]+([.-][a-z0-9]+)*)?){2}:([^/].*)?$`

_Maximum_: `1600`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TrackerName`

The name for the tracker resource.

Requirements:

- Contain only alphanumeric characters (A-Z, a-z, 0-9) , hyphens (-), periods (.), and underscores (\_).

- Must be a unique tracker resource name.

- No spaces allowed. For example, `ExampleTracker`.

_Required_: Yes

_Type_: String

_Pattern_: `^[-._\w]+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `TrackerConsumer` name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Next

All content copied from https://docs.aws.amazon.com/.
