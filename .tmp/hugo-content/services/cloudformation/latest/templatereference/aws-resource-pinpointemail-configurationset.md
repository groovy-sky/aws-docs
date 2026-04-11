This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PinpointEmail::ConfigurationSet

Create a configuration set. _Configuration sets_ are groups of
rules that you can apply to the emails you send using Amazon Pinpoint. You apply a configuration
set to an email by including a reference to the configuration set in the headers of the
email. When you apply a configuration set to an email, all of the rules in that
configuration set are applied to the email.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::PinpointEmail::ConfigurationSet",
  "Properties" : {
      "DeliveryOptions" : DeliveryOptions,
      "Name" : String,
      "ReputationOptions" : ReputationOptions,
      "SendingOptions" : SendingOptions,
      "Tags" : [ Tags, ... ],
      "TrackingOptions" : TrackingOptions
    }
}

```

### YAML

```yaml

Type: AWS::PinpointEmail::ConfigurationSet
Properties:
  DeliveryOptions:
    DeliveryOptions
  Name: String
  ReputationOptions:
    ReputationOptions
  SendingOptions:
    SendingOptions
  Tags:
    - Tags
  TrackingOptions:
    TrackingOptions

```

## Properties

`DeliveryOptions`

An object that defines the dedicated IP pool that is used to send emails that you send
using the configuration set.

_Required_: No

_Type_: [DeliveryOptions](aws-properties-pinpointemail-configurationset-deliveryoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the configuration set.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReputationOptions`

An object that defines whether or not Amazon Pinpoint collects reputation metrics for the emails
that you send that use the configuration set.

_Required_: No

_Type_: [ReputationOptions](aws-properties-pinpointemail-configurationset-reputationoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SendingOptions`

An object that defines whether or not Amazon Pinpoint can send email that you send using the
configuration set.

_Required_: No

_Type_: [SendingOptions](aws-properties-pinpointemail-configurationset-sendingoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An object that defines the tags (keys and values) that you want to associate with the
configuration set.

_Required_: No

_Type_: [Array](aws-properties-pinpointemail-configurationset-tags.md) of [Tags](aws-properties-pinpointemail-configurationset-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrackingOptions`

An object that defines the open and click tracking options for emails that you send
using the configuration set.

_Required_: No

_Type_: [TrackingOptions](aws-properties-pinpointemail-configurationset-trackingoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{ "Ref": "myConfigurationSet" }`

For the Amazon Pinpoint configuration set `myConfigurationSet`, Ref returns the
name of the configuration set.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Pinpoint Email

DeliveryOptions

All content copied from https://docs.aws.amazon.com/.
