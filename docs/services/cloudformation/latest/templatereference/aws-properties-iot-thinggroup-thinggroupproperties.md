This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::ThingGroup ThingGroupProperties

Thing group properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttributePayload" : AttributePayload,
  "ThingGroupDescription" : String
}

```

### YAML

```yaml

  AttributePayload:
    AttributePayload
  ThingGroupDescription: String

```

## Properties

`AttributePayload`

The thing group attributes in JSON format.

_Required_: No

_Type_: [AttributePayload](aws-properties-iot-thinggroup-attributepayload.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThingGroupDescription`

The thing group description.

_Required_: No

_Type_: String

_Pattern_: `[\p{Graph}\x20]*`

_Maximum_: `2028`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::IoT::ThingPrincipalAttachment

All content copied from https://docs.aws.amazon.com/.
