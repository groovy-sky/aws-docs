This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::MessageTemplateVersion

Creates a new Amazon Q in Connect message template version from the current content
and configuration of a message template. Versions are immutable and monotonically
increasing. Once a version is created, you can reference a specific version of the
message template by passing in
`<messageTemplateArn>:<versionNumber>` as the message
template identifier. An error is displayed if the supplied
`messageTemplateContentSha256` is different from the
`messageTemplateContentSha256` of the message template with
`$LATEST` qualifier. If multiple
`CreateMessageTemplateVersion` requests are made while the message
template remains the same, only the first invocation creates a new version and the
succeeding requests will return the same response as the first invocation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Wisdom::MessageTemplateVersion",
  "Properties" : {
      "MessageTemplateArn" : String,
      "MessageTemplateContentSha256" : String
    }
}

```

### YAML

```yaml

Type: AWS::Wisdom::MessageTemplateVersion
Properties:
  MessageTemplateArn: String
  MessageTemplateContentSha256: String

```

## Properties

`MessageTemplateArn`

The Amazon Resource Name (ARN) of the message template.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[a-z-]*?:wisdom:[a-z0-9-]*?:[0-9]{12}:[a-z-]*?/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}(?:/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12})?$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MessageTemplateContentSha256`

The content SHA256 of the message template.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

## Return values

### Ref

### Fn::GetAtt

`MessageTemplateVersionArn`

The Amazon Resource Name (ARN) of the Message Template Version.

`MessageTemplateVersionNumber`

The version number for this Message Template Version.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::Wisdom::QuickResponse

All content copied from https://docs.aws.amazon.com/.
