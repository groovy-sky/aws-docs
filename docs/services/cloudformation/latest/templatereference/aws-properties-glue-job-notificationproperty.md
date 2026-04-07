This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Job NotificationProperty

Specifies configuration properties of a notification.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NotifyDelayAfter" : Integer
}

```

### YAML

```yaml

  NotifyDelayAfter: Integer

```

## Properties

`NotifyDelayAfter`

After a job run starts, the number of minutes to wait before sending a job run delay notification.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

JobCommand

AWS::Glue::MLTransform
