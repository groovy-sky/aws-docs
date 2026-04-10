This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::Input

The AWS::IoTEvents::Input resource creates an input. To monitor your devices and processes,
they must have a way to get telemetry data into AWS IoT Events. This is done by sending messages
as _inputs_ to AWS IoT Events. For more information, see
[How to Use AWS IoT Events](../../../iotevents/latest/developerguide/how-to-use-iotevents.md) in the _AWS IoT Events Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTEvents::Input",
  "Properties" : {
      "InputDefinition" : InputDefinition,
      "InputDescription" : String,
      "InputName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTEvents::Input
Properties:
  InputDefinition:
    InputDefinition
  InputDescription: String
  InputName: String
  Tags:
    - Tag

```

## Properties

`InputDefinition`

The definition of the input.

_Required_: Yes

_Type_: [InputDefinition](aws-properties-iotevents-input-inputdefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputDescription`

A brief description of the input.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputName`

The name of the input.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z][a-zA-Z0-9_]*$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-iotevents-input-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the input. For example:

`{"Ref": "myInput"}`

For the AWS IoT Events input `myInput`, `Ref` returns the name of the input.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Input

The following example creates an input.

#### JSON

```json

{
  "Description": "Input Template for CloudFormation",
  "Resources": {
    "myInput": {
      "Type": "AWS::IoTEvents::Input",
      "Properties": {
        "InputName": "myInput",
        "InputDescription": "My Input created by CloudFormation",
        "InputDefinition": {
          "Attributes": [
            {
              "JsonPath": "foo"
            },
            {
              "JsonPath": "bar"
            }
          ]
        }
      }
    }
  }
}
```

#### YAML

```yaml

---
Description: "Input Template for CloudFormation"
Resources:
  myInput:
    Type: "AWS::IoTEvents::Input"
    Properties:
      InputName: "myInput"
      InputDescription: "My Input created by CloudFormation"
      InputDefinition:
        Attributes:
          -
            JsonPath: "foo"
          -
            JsonPath: "bar"

```

## See also

- [How to Use AWS IoT Events](../../../iotevents/latest/developerguide/how-to-use-iotevents.md) in the _AWS IoT Events Developer Guide_

- [CreateInput](../../../../reference/iotevents/latest/apireference/api-createinput.md)
in the _AWS IoT Events API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TransitionEvent

Attribute

All content copied from https://docs.aws.amazon.com/.
