This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::SubscriptionDefinition

The `AWS::Greengrass::SubscriptionDefinition` resource represents a
subscription definition for AWS IoT Greengrass. Subscription definitions are used to
organize your subscription definition versions.

Subscription definitions can reference multiple subscription definition versions. All
subscription definition versions must be associated with a subscription definition. Each
subscription definition version can contain one or more subscriptions.

###### Note

When you create a subscription definition, you can optionally include an initial
subscription definition version. To associate a subscription definition version later,
create an [`AWS::Greengrass::SubscriptionDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-subscriptiondefinitionversion.html) resource and
specify the ID of this subscription definition.

After you create the subscription definition version that contains the subscriptions
you want to deploy, you must add it to your group version. For more information, see
[`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Greengrass::SubscriptionDefinition",
  "Properties" : {
      "InitialVersion" : SubscriptionDefinitionVersion,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Greengrass::SubscriptionDefinition
Properties:
  InitialVersion:
    SubscriptionDefinitionVersion
  Name: String
  Tags:
    - Tag

```

## Properties

`InitialVersion`

The subscription definition version to include when the subscription definition is
created. A subscription definition version contains a list of [`subscription`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-subscriptiondefinition-subscription.html) property types.

###### Note

To associate a subscription definition version after the subscription definition is
created, create an [`AWS::Greengrass::SubscriptionDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-subscriptiondefinitionversion.html) resource and specify the ID of this subscription definition.

_Required_: No

_Type_: [SubscriptionDefinitionVersion](aws-properties-greengrass-subscriptiondefinition-subscriptiondefinitionversion.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the subscription definition.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Application-specific metadata to attach to the subscription definition. You can use tags
in IAM policies to control access to AWS IoT Greengrass resources. You
can also use tags to categorize your resources. For more information, see [Tagging Your\
AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/v1/developerguide/tagging.html) in the _AWS IoT Greengrass Version 1 Developer Guide_.

This `Json` property type is processed as a map of key-value pairs. It uses
the following format, which is different from most `Tags` implementations in
CloudFormation templates.

```json

"Tags": {
    "KeyName0": "value",
    "KeyName1": "value",
    "KeyName2": "value"
}
```

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the subscription definition, such as
`1234a5b6-78cd-901e-2fgh-3i45j6k178l9`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the `SubscriptionDefinition`, such as
`arn:aws:greengrass:us-east-1:123456789012:/greengrass/definition/subscriptions/1234a5b6-78cd-901e-2fgh-3i45j6k178l9`.

`Id`

The ID of the `SubscriptionDefinition`, such as
`1234a5b6-78cd-901e-2fgh-3i45j6k178l9`.

`LatestVersionArn`

The ARN of the last `SubscriptionDefinitionVersion` that was added to the
`SubscriptionDefinition`, such as `arn:aws:greengrass:us-east-1:123456789012:/greengrass/definition/subscriptions/1234a5b6-78cd-901e-2fgh-3i45j6k178l9/versions/9876ac30-4bdb-4f9d-95af-b5fdb66be1a2`.

`Name`

The name of the `SubscriptionDefinition`, such as
`MySubscriptionDefinition`.

## Examples

### Subscription Definition Snippet

The following snippet defines a subscription definition subscription with an
initial version that contains a subscription. In this example, the subscription
source is an existing device in the group. The target is a function in the group that
was created in another stack and is referenced using the `ImportValue`
function.

For an example of a complete template, see the [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-subscription-greengrass-group.html) subscription.

#### JSON

```json

"TestSubscriptionDefinition": {
    "Type": "AWS::Greengrass::SubscriptionDefinition",
    "Properties": {
        "Name": "DemoTestSubscriptionDefinition",
        "InitialVersion": {
            "Subscriptions": [
                {
                    "Id": "TestSubscription1",
                    "Source": {
                        "Fn::Join": [
                            ":",
                            [
                                "arn:aws:iot",
                                {
                                    "Ref": "AWS::Region"
                                },
                                {
                                    "Ref": "AWS::AccountId"
                                },
                                "thing/TestDevice1"
                            ]
                        ]
                    },
                    "Subject": "some/topic",
                    "Target": {
                        "Fn::ImportValue": "TestCanaryLambdaVersionArn"
                    }
                }
            ]
        }
    }
}
```

#### YAML

```yaml

TestSubscriptionDefinition:
  Type: 'AWS::Greengrass::SubscriptionDefinition'
  Properties:
    Name: DemoTestSubscriptionDefinition
    InitialVersion:
      Subscriptions:
        - Id: TestSubscription1
          Source: !Join
            - ':'
            - - 'arn:aws:iot'
              - !Ref 'AWS::Region'
              - !Ref 'AWS::AccountId'
              - thing/TestDevice1
          Subject: some/topic
          Target: !ImportValue TestCanaryLambdaVersionArn
```

## See also

- [CreateSubscriptionDefinition](https://docs.aws.amazon.com/greengrass/v1/apireference/createsubscriptiondefinition-post.html) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](https://docs.aws.amazon.com/greengrass/v1/developerguide)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SecretsManagerSecretResourceData

Subscription
