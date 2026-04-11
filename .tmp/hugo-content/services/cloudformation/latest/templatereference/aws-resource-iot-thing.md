This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::Thing

Use the `AWS::IoT::Thing` resource to declare an AWS IoT
thing.

For information about working with things, see [How AWS IoT\
Works](../../../iot/latest/developerguide/aws-iot-how-it-works.md) and [Device Registry for AWS IoT](../../../iot/latest/developerguide/thing-registry.md) in the _AWS IoT Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::Thing",
  "Properties" : {
      "AttributePayload" : AttributePayload,
      "ThingName" : String
    }
}

```

### YAML

```yaml

Type: AWS::IoT::Thing
Properties:
  AttributePayload:
    AttributePayload
  ThingName: String

```

## Properties

`AttributePayload`

A string that contains up to three key value pairs. Maximum length of 800. Duplicates
not allowed.

_Required_: No

_Type_: [AttributePayload](aws-properties-iot-thing-attributepayload.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThingName`

The name of the thing to update.

You can't change a thing's name. To change a thing's name, you must create a
new thing, give it the new name, and then delete the old thing.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9:_-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the thing name. For example:

`{ "Ref": "MyThing" }`

For a stack named MyStack, a value similar to the following is returned:

`MyStack-MyThing-AB1CDEFGHIJK`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the AWS IoT thing, such as
`arn:aws:iot:us-east-2:123456789012:thing/MyThing`.

`Id`

The Id of this thing.

## Examples

The following example declares a thing and the values of its attributes.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "MyThing": {
      "Type": "AWS::IoT::Thing",
      "Properties": {
        "ThingName": {
          "Ref": "NameParameter"
        },
        "AttributePayload": {
          "Attributes": {
            "myAttributeA": {
              "Ref": "MyAttributeValueA"
            },
            "myAttributeB": {
              "Ref": "MyAttributeValueB"
            },
            "myAttributeC": {
              "Ref": "MyAttributeValueC"
            }
          }
        }
      }
    }
  },
  "Parameters": {
    "NameParameter": {
      "Type": "String"
    },
    "MyAttributeValueA": {
      "Type": "String",
      "Default": "myStringA123"
    },
    "MyAttributeValueB": {
      "Type": "String",
      "Default": "myStringB123"
    },
    "MyAttributeValueC": {
      "Type": "String",
      "Default": "myStringC123"
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Resources:
  MyThing:
    Type: AWS::IoT::Thing
    Properties:
      ThingName:
        Ref: NameParameter
      AttributePayload:
        Attributes:
          myAttributeA:
            Ref: MyAttributeValueA
          myAttributeB:
            Ref: MyAttributeValueB
          myAttributeC:
            Ref: MyAttributeValueC

Parameters:
  NameParameter:
    Type: String
  MyAttributeValueA:
    Type: String
    Default: myStringA123
  MyAttributeValueB:
    Type: String
    Default: myStringB123
  MyAttributeValueC:
    Type: String
    Default: myStringC123
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AttributePayload

All content copied from https://docs.aws.amazon.com/.
