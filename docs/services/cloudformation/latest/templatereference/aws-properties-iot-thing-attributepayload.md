This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::Thing AttributePayload

The AttributePayload property specifies up to three attributes for an AWS IoT as key-value pairs. AttributePayload is a property of the [AWS::IoT::Thing](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thing.html) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Attributes" : {Key: Value, ...}
}

```

### YAML

```yaml

  Attributes:
    Key: Value

```

## Properties

`Attributes`

A JSON string containing up to three key-value pair in JSON format. For example:

`{\"attributes\":{\"string1\":\"string2\"}}`

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9_.,@/:#-]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

The following example declares an attribute payload with three attributes.

#### JSON

```json

 { "AttributePayload":{ "Attributes":{ "myAttributeA":{
            "Ref":"MyAttributeValueA" }, "myAttributeB":{ "Ref":"MyAttributeValueB" },
            "myAttributeC":{ "Ref":"MyAttributeValueC" } } } }
```

#### YAML

```yaml

 AttributePayload: Attributes: myAttributeA: Ref:
            MyAttributeValueA myAttributeB: Ref: MyAttributeValueB myAttributeC: Ref:
            MyAttributeValueC
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::IoT::Thing

AWS::IoT::ThingGroup
