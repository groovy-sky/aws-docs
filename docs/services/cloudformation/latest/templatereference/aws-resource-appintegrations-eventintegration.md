This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppIntegrations::EventIntegration

Creates an event integration. You provide a name, description, and a reference to an
Amazon EventBridge bus in your account and a partner event source that will push events to
that bus. No objects are created in your account, only metadata that is persisted on the
EventIntegration control plane.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppIntegrations::EventIntegration",
  "Properties" : {
      "Description" : String,
      "EventBridgeBus" : String,
      "EventFilter" : EventFilter,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::AppIntegrations::EventIntegration
Properties:
  Description: String
  EventBridgeBus: String
  EventFilter:
    EventFilter
  Name: String
  Tags:
    - Tag

```

## Properties

`Description`

The event integration description.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventBridgeBus`

The Amazon EventBridge bus for the event integration.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9/\._\-]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EventFilter`

The event integration filter.

_Required_: Yes

_Type_: [EventFilter](aws-properties-appintegrations-eventintegration-eventfilter.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the event integration.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9/\._\-]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-appintegrations-eventintegration-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the EventIntegration name. For example:

`{ "Ref": "myEventIntegrationName" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`EventIntegrationArn`

The Amazon Resource Name (ARN) of the event integration.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

EventFilter

All content copied from https://docs.aws.amazon.com/.
