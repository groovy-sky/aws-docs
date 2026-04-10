This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PinpointEmail::ConfigurationSet TrackingOptions

An object that defines the tracking options for a configuration set. When you use
Amazon Pinpoint to send an email, it contains an invisible image that's used to track when
recipients open your email. If your email contains links, those links are changed
slightly in order to track when recipients click them.

These images and links include references to a domain operated by AWS. You can
optionally configure Amazon Pinpoint to use a domain that you operate for these images and
links.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomRedirectDomain" : String
}

```

### YAML

```yaml

  CustomRedirectDomain: String

```

## Properties

`CustomRedirectDomain`

The domain that you want to use for tracking open and click events.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tags

AWS::PinpointEmail::ConfigurationSetEventDestination

All content copied from https://docs.aws.amazon.com/.
