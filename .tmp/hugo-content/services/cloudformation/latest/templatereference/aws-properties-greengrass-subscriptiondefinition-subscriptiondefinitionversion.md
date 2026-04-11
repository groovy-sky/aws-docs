This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::SubscriptionDefinition SubscriptionDefinitionVersion

A subscription definition version contains a list of [subscriptions](../userguide/aws-properties-greengrass-subscriptiondefinition-subscription.md).

###### Note

After you create a subscription definition version that contains the subscriptions
you want to deploy, you must add it to your group version. For more information, see
[`AWS::Greengrass::Group`](../userguide/aws-resource-greengrass-group.md).

In an CloudFormation template, `SubscriptionDefinitionVersion` is the
property type of the `InitialVersion` property in the [`AWS::Greengrass::SubscriptionDefinition`](../userguide/aws-resource-greengrass-subscriptiondefinition.md) resource.

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

_Type_: Array of [Subscription](aws-properties-greengrass-subscriptiondefinition-subscription.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [SubscriptionDefinitionVersion](../../../../reference/greengrass/v1/apireference/definitions-subscriptiondefinitionversion.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Subscription

AWS::Greengrass::SubscriptionDefinitionVersion

All content copied from https://docs.aws.amazon.com/.
