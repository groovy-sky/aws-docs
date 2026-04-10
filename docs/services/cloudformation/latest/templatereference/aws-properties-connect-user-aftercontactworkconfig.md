This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::User AfterContactWorkConfig

Configuration settings for after contact work (ACW) timeout.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AfterContactWorkTimeLimit" : Integer
}

```

### YAML

```yaml

  AfterContactWorkTimeLimit: Integer

```

## Properties

`AfterContactWorkTimeLimit`

The ACW timeout duration in seconds. Minimum: 1 second. Maximum: 2,000,000 seconds (24 days). Enter 0 for indefinite ACW time.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Connect::User

AfterContactWorkConfigPerChannel

All content copied from https://docs.aws.amazon.com/.
