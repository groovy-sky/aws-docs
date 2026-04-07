This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::EmailIdentity FeedbackAttributes

Used to enable or disable feedback forwarding for an identity. This setting determines
what happens when an identity is used to send an email that results in a bounce or
complaint event.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EmailForwardingEnabled" : Boolean
}

```

### YAML

```yaml

  EmailForwardingEnabled: Boolean

```

## Properties

`EmailForwardingEnabled`

Sets the feedback forwarding configuration for the identity.

If the value is `true`, you receive email notifications when bounce or
complaint events occur. These notifications are sent to the address that you specified
in the `Return-Path` header of the original email.

You're required to have a method of tracking bounces and complaints. If you haven't
set up another mechanism for receiving bounce or complaint notifications (for example,
by setting up an event destination), you receive an email notification when these events
occur (even if this setting is disabled).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DkimSigningAttributes

MailFromAttributes
