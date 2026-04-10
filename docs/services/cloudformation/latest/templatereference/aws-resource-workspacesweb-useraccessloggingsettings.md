This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesWeb::UserAccessLoggingSettings

This resource specifies user access logging settings that can be associated with a web portal.

In order to receive logs from WorkSpaces Secure Browser, you must have an Amazon Kinesis Data Stream that
starts with "amazon-workspaces-web-\*". Your Amazon Kinesis data stream must either have server-side
encryption turned off, or must use AWS managed keys for server-side encryption.

For more information about setting server-side encryption in Amazon Kinesis, see [How Do I Get\
Started with Server-Side Encryption?](../../../streams/latest/dev/getting-started-with-sse.md).

For more information about setting up user access logging, see [Set up user access logging](../../../workspaces-web/latest/adminguide/user-logging.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WorkSpacesWeb::UserAccessLoggingSettings",
  "Properties" : {
      "KinesisStreamArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::WorkSpacesWeb::UserAccessLoggingSettings
Properties:
  KinesisStreamArn: String
  Tags:
    - Tag

```

## Properties

`KinesisStreamArn`

The ARN of the Kinesis stream.

_Required_: Yes

_Type_: String

_Pattern_: `arn:[\w+=/,.@-]+:kinesis:[a-zA-Z0-9\-]*:[a-zA-Z0-9]{1,12}:stream/.+`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to add to the user access logging settings resource. A tag is a key-value pair.

_Required_: No

_Type_: Array of [Tag](aws-properties-workspacesweb-useraccessloggingsettings-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function,
`Ref` returns the resource's Amazon Resource Name (ARN).

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

`AssociatedPortalArns`

A list of web portal ARNs that this user access logging settings is associated with.

`UserAccessLoggingSettingsArn`

The ARN of the user access logging settings.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
