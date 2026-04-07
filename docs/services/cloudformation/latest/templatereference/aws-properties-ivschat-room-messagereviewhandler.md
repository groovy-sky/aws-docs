This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IVSChat::Room MessageReviewHandler

The MessageReviewHandler property type specifies configuration information for optional message review.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FallbackResult" : String,
  "Uri" : String
}

```

### YAML

```yaml

  FallbackResult: String
  Uri: String

```

## Properties

`FallbackResult`

Specifies the fallback behavior (whether the message is allowed or denied) if the handler
does not return a valid response, encounters an error, or times out. (For the timeout period,
see [Service\
Quotas](https://docs.aws.amazon.com/ivs/latest/userguide/service-quotas.html).) If allowed, the message is delivered with returned content to all users
connected to the room. If denied, the message is not delivered to any user.

_Default_: `ALLOW`

_Required_: No

_Type_: String

_Allowed values_: `ALLOW | DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Uri`

Identifier of the message review handler. Currently this must be an ARN of a lambda
function.

_Required_: No

_Type_: String

_Pattern_: `^$|^arn:aws:lambda:[a-z0-9-]+:[0-9]{12}:function:.+`

_Minimum_: `0`

_Maximum_: `170`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Room Message Review Handler Template Examples

The following examples specify an Amazon IVS Chat Room with
message review.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "MessageReviewHandlerRole": {
      "Type": "AWS::IAM::Role",
      "Properties": {
        "AssumeRolePolicyDocument": {
          "Version": "2012-10-17",
          "Statement": {
            "Effect": "Allow",
            "Principal": {
              "Service": "lambda.amazonaws.com"
            },
            "Action": "sts:AssumeRole"
          }
        }
      }
    },
    "MessageReviewHandler": {
      "Type": "AWS::Lambda::Function",
      "Properties": {
        "Handler": "index.handler",
        "Role": {
          "Fn::GetAtt": [
            "MessageReviewHandlerRole",
            "Arn"
          ]
        },
        "Runtime": "nodejs12.x",
        "Code": {
          "S3Bucket":"my-bucket",
          "S3Key": "function.zip"
        }
      }
    },
    "MessageReviewHandlerPermission": {
      "Type": "AWS::Lambda::Permission",
      "Properties": {
        "Principal": "ivschat.amazonaws.com",
        "Action": "lambda:InvokeFunction",
        "FunctionName": {
          "Ref": "MessageReviewHandler"
        },
        "SourceAccount": {
          "Ref": "AWS::AccountId"
        },
        "SourceArn": {
          "Fn::Sub": "arn:${AWS::Partition}:ivschat:${AWS::Region}:${AWS::AccountId}:room/*"
        }
      }
    },
    "Room": {
      "Type": "AWS::IVSChat::Room",
      "Properties": {
        "Name": "MyRoom",
        "MessageReviewHandler": {
          "FallbackResult": "ALLOW",
          "Uri": {
            "Fn::GetAtt": [
              "MessageReviewHandler",
              "Arn"
            ]
          }
        }
      },
      "DependsOn": "MessageReviewHandlerPermission"
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  MessageReviewHandlerRole:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: 2012-10-17
        Statement:
          Effect: Allow
          Principal:
            Service: lambda.amazonaws.com
          Action: sts:AssumeRole
  MessageReviewHandler:
    Type: AWS::Lambda::Function
    Properties:
      Handler: index.handler
      Role: !GetAtt MessageReviewHandlerRole.Arn
      Runtime: nodejs12.x
      Code:
        S3Bucket: my-bucket
        S3Key: function.zip
  MessageReviewHandlerPermission:
    Type: AWS::Lambda::Permission
    Properties:
      Principal: ivschat.amazonaws.com
      Action: lambda:InvokeFunction
      FunctionName: !Ref MessageReviewHandler
      SourceAccount: !Ref AWS::AccountId
      SourceArn: !Sub arn:${AWS::Partition}:ivschat:${AWS::Region}:${AWS::AccountId}:room/*
  Room:
    Type: AWS::IVSChat::Room
    Properties:
      Name: MyRoom
      MessageReviewHandler:
        FallbackResult: ALLOW
        Uri: !GetAtt MessageReviewHandler.Arn
    DependsOn: MessageReviewHandlerPermission
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::IVSChat::Room

Tag
