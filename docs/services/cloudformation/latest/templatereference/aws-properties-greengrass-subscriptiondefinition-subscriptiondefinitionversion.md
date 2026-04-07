This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::SubscriptionDefinition SubscriptionDefinitionVersion

A subscription definition version contains a list of [subscriptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-subscriptiondefinition-subscription.html).

###### Note

After you create a subscription definition version that contains the subscriptions
you want to deploy, you must add it to your group version. For more information, see
[`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html).

In an CloudFormation template, `SubscriptionDefinitionVersion` is the
property type of the `InitialVersion` property in the [`AWS::Greengrass::SubscriptionDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-subscriptiondefinition.html) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Subscriptions" : [ Subscription, ... ]
}

```

### YAML

```yaml

  Subscriptions:
    - Subscription

```

## Properties

`Subscriptions`

The subscriptions in this version.

_Required_: Yes

_Type_: Array of [Subscription](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-greengrass-subscriptiondefinition-subscription.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [SubscriptionDefinitionVersion](https://docs.aws.amazon.com/greengrass/v1/apireference/definitions-subscriptiondefinitionversion.html) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](https://docs.aws.amazon.com/greengrass/v1/developerguide)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Subscription

AWS::Greengrass::SubscriptionDefinitionVersion
