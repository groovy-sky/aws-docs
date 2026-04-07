This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ConnectionFunction

A connection function.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFront::ConnectionFunction",
  "Properties" : {
      "AutoPublish" : Boolean,
      "ConnectionFunctionCode" : String,
      "ConnectionFunctionConfig" : ConnectionFunctionConfig,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CloudFront::ConnectionFunction
Properties:
  AutoPublish: Boolean
  ConnectionFunctionCode: String
  ConnectionFunctionConfig:
    ConnectionFunctionConfig
  Name: String
  Tags:
    - Tag

```

## Properties

`AutoPublish`

A flag that determines whether to automatically publish the function to the
`LIVE` stage when it’s created. To automatically publish to the
`LIVE` stage, set this property to `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectionFunctionCode`

The code for the connection function.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectionFunctionConfig`

Contains configuration information about a CloudFront function.

_Required_: Yes

_Type_: [ConnectionFunctionConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-connectionfunction-connectionfunctionconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The connection function name.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9-_]{1,64}`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A complex type that contains zero or more `Tag` elements.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-connectionfunction-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`ConnectionFunctionArn`

The connection function Amazon Resource Name (ARN).

`CreatedTime`

The connection function created time.

`ETag`

A complex type that contains `Tag` key and `Tag` value.

`Id`

The connection function ID.

`LastModifiedTime`

The connection function last modified time.

`Stage`

The connection function stage.

`Status`

The connection function status.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CloudFrontOriginAccessIdentityConfig

ConnectionFunctionConfig
