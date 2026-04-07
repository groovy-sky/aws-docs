This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IVSChat::Room

The `AWS::IVSChat::Room` resource specifies an Amazon IVS room that allows clients to connect and pass messages. For more
information, see [CreateRoom](https://docs.aws.amazon.com/ivs/latest/ChatAPIReference/API_CreateRoom.html) in the
_Amazon Interactive Video Service Chat API Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IVSChat::Room",
  "Properties" : {
      "LoggingConfigurationIdentifiers" : [ String, ... ],
      "MaximumMessageLength" : Integer,
      "MaximumMessageRatePerSecond" : Integer,
      "MessageReviewHandler" : MessageReviewHandler,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IVSChat::Room
Properties:
  LoggingConfigurationIdentifiers:
    - String
  MaximumMessageLength: Integer
  MaximumMessageRatePerSecond: Integer
  MessageReviewHandler:
    MessageReviewHandler
  Name: String
  Tags:
    - Tag

```

## Properties

`LoggingConfigurationIdentifiers`

List of logging-configuration identifiers attached to the room.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `128 | 50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumMessageLength`

Maximum number of characters in a single message. Messages are expected to be UTF-8 encoded and this limit applies specifically to rune/code-point count, not number of bytes.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumMessageRatePerSecond`

Maximum number of messages per second that can be sent to the room (by all clients).

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MessageReviewHandler`

Configuration information for optional review of messages.

_Required_: No

_Type_: [MessageReviewHandler](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ivschat-room-messagereviewhandler.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Room name. The value does not need to be unique.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_]*$`

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivschat-room-tag.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ivschat-room-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the room ARN. For example:

`{ "Ref": "myRoom" }`

For the Amazon IVS room `myRoom `, `Ref`
returns the room ARN.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The room ARN. For example:
`arn:aws:ivschat:us-west-2:123456789012:room/abcdABCDefgh`

`Id`

The room ID. For example: `abcdABCDefgh`

## Examples

### Room Template Examples

The following examples specify an Amazon IVS Chat Room.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "Room": {
      "Type": "AWS::IVSChat::Room",
      "Properties": {
        "Name": "MyRoom",
        "Tags": [
          { "Key": "MyKey", "Value": "MyValue" }
        ]
      }
    }
  },
  "Outputs": {
    "RoomArn": {
      "Value": { "Ref": "Room" }
    },
    "RoomId": {
      "Value": {
        "Fn::GetAtt": [
          "Room",
          "Id"
        ]
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  Room:
    Type: AWS::IVSChat::Room
    Properties:
      Name: MyRoom
      Tags:
        - Key: MyKey
          Value: MyValue
Outputs:
  RoomArn:
    Value: !Ref Room
  RoomId:
    Value: !GetAtt Room.Id
```

## See also

- [Getting\
Started with Amazon Interactive Video Service](https://docs.aws.amazon.com/ivs/latest/userguide/getting-started.html)

- [RoomSummary](https://docs.aws.amazon.com/ivs/latest/ChatAPIReference/API_RoomSummary.html) API data type

- [CreateRoom](https://docs.aws.amazon.com/ivs/latest/ChatAPIReference/API_CreateRoom.html) API endpoint

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

MessageReviewHandler
