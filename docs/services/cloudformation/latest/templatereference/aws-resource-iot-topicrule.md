This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule

Use the `AWS::IoT::TopicRule` resource to declare an AWS IoT
rule. For information about working with AWS IoT rules, see [Rules for AWS IoT](https://docs.aws.amazon.com/iot/latest/developerguide/iot-rules.html) in the _AWS IoT Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::TopicRule",
  "Properties" : {
      "RuleName" : String,
      "Tags" : [ Tag, ... ],
      "TopicRulePayload" : TopicRulePayload
    }
}

```

### YAML

```yaml

Type: AWS::IoT::TopicRule
Properties:
  RuleName: String
  Tags:
    - Tag
  TopicRulePayload:
    TopicRulePayload

```

## Properties

`RuleName`

The name of the rule.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Metadata which can be used to manage the topic rule.

###### Note

For URI Request parameters use format: ...key1=value1&key2=value2...

For the CLI command-line parameter use format: --tags
"key1=value1&key2=value2..."

For the cli-input-json file use format: "tags":
"key1=value1&key2=value2..."

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-topicrule-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopicRulePayload`

The rule payload.

_Required_: Yes

_Type_: [TopicRulePayload](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-topicrule-topicrulepayload.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the topic rule name. For example:

`{ "Ref": "MyTopicRule" }`

For a stack named My-Stack (the - character is omitted), a value similar to the
following is returned:

`MyStackMyTopicRule12ABC3D456EFG`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the AWS IoT rule, such as
`arn:aws:iot:us-east-2:123456789012:rule/MyIoTRule`.

## Examples

The following example declares an AWS IoT rule.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "MyTopicRule": {
      "Type": "AWS::IoT::TopicRule",
      "Properties": {
        "RuleName": {
          "Ref": "NameParameter"
        },
        "TopicRulePayload": {
          "RuleDisabled": "true",
          "Sql": "SELECT temp FROM 'SomeTopic' WHERE temp > 60",
          "Actions": [
            {
              "S3": {
                "BucketName": {
                  "Ref": "amzn-s3-demo-bucket"
                },
                "RoleArn": {
                  "Fn::GetAtt": [
                    "MyRole",
                    "Arn"
                  ]
                },
                "Key": "MyKey.txt"
              }
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

AWSTemplateFormatVersion: '2010-09-09'
Resources:
  MyTopicRule:
    Type: AWS::IoT::TopicRule
    Properties:
      RuleName:
        Ref: NameParameter
      TopicRulePayload:
        RuleDisabled: 'true'
        Sql: "SELECT temp FROM 'SomeTopic' WHERE temp > 60"
        Actions:
          - S3:
              BucketName:
                Ref: amzn-s3-demo-bucket
              RoleArn:
                Fn::GetAtt:
                  - MyRole
                  - Arn
              Key: MyKey.txt
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ThingTypeProperties

Action
