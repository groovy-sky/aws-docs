This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::SubscriptionDefinitionVersion

The `AWS::Greengrass::SubscriptionDefinitionVersion` resource represents a
subscription definition version for AWS IoT Greengrass. A subscription definition version
contains a list of subscriptions.

###### Note

To create a subscription definition version, you must specify the ID of the
subscription definition that you want to associate with the version. For information
about creating a subscription definition, see [`AWS::Greengrass::SubscriptionDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-subscriptiondefinition.html).

After you create a subscription definition version that contains the subscriptions
you want to deploy, you must add it to your group version. For more information, see
[`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Greengrass::SubscriptionDefinitionVersion",
  "Properties" : {
      "SubscriptionDefinitionId" : String,
      "Subscriptions" : [ Subscription, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Greengrass::SubscriptionDefinitionVersion
Properties:
  SubscriptionDefinitionId: String
  Subscriptions:
    - Subscription

```

## Properties

`SubscriptionDefinitionId`

The ID of the subscription definition associated with this version. This value is a
GUID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Subscriptions`

The subscriptions in this version.

_Required_: Yes

_Type_: Array of [Subscription](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-greengrass-subscriptiondefinitionversion-subscription.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the subscription definition
version, such as `arn:aws:greengrass:us-east-1:123456789012:/greengrass/definition/subscriptions/1234a5b6-78cd-901e-2fgh-3i45j6k178l9/versions/9876ac30-4bdb-4f9d-95af-b5fdb66be1a2`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

### Subscription Definition Version Snippet

The following snippet defines subscription definition and subscription definition
version subscriptions. The subscription definition version references the
subscription definition and contains a subscription. In this example, the
subscription source is an existing device in the group. The target is a function in
the group that was created in another stack and is referenced using the
`ImportValue` function.

For an example of a complete template, see the [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-subscription-greengrass-group.html) subscription.

#### JSON

```json

"TestSubscriptionDefinition": {
    "Type": "AWS::Greengrass::SubscriptionDefinition",
    "Properties": {
        "Name": "DemoTestSubscriptionDefinition"
    }
},
"TestSubscriptionDefinitionVersion": {
    "Type": "AWS::Greengrass::SubscriptionDefinitionVersion",
    "Properties": {
        "SubscriptionDefinitionId": {
            "Ref": "TestSubscriptionDefinition"
        },
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
```

#### YAML

```yaml

TestSubscriptionDefinition:
  Type: 'AWS::Greengrass::SubscriptionDefinition'
  Properties:
    Name: DemoTestSubscriptionDefinition
TestSubscriptionDefinitionVersion:
  Type: 'AWS::Greengrass::SubscriptionDefinitionVersion'
  Properties:
    SubscriptionDefinitionId: !Ref TestSubscriptionDefinition
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

- [CreateSubscriptionDefinitionVersion](https://docs.aws.amazon.com/greengrass/v1/apireference/createsubscriptiondefinitionversion-post.html) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](https://docs.aws.amazon.com/greengrass/v1/developerguide)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SubscriptionDefinitionVersion

Subscription
