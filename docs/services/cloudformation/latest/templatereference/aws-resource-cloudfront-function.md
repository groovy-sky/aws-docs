This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Function

Creates a CloudFront function.

To create a function, you provide the function code and some configuration information
about the function. The response contains an Amazon Resource Name (ARN) that uniquely
identifies the function, and the function’s stage.

By default, when you create a function, it’s in the `DEVELOPMENT` stage. In this
stage, you can [test the function](../../../amazoncloudfront/latest/developerguide/test-function.md) in the CloudFront console (or with
`TestFunction` in the CloudFront API).

When you’re ready to use your function with a CloudFront distribution, publish the
function to the `LIVE` stage. You can do this in the CloudFront console, with
`PublishFunction` in the CloudFront API, or by updating the
`AWS::CloudFront::Function` resource with the `AutoPublish`
property set to `true`. When the function is published to the
`LIVE` stage, you can attach it to a distribution’s cache behavior, using the
function’s ARN.

To automatically publish the function to the `LIVE` stage when it’s
created, set the `AutoPublish` property to `true`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFront::Function",
  "Properties" : {
      "AutoPublish" : Boolean,
      "FunctionCode" : String,
      "FunctionConfig" : FunctionConfig,
      "FunctionMetadata" : FunctionMetadata,
      "Name" : String
    }
}

```

### YAML

```yaml

Type: AWS::CloudFront::Function
Properties:
  AutoPublish: Boolean
  FunctionCode: String
  FunctionConfig:
    FunctionConfig
  FunctionMetadata:
    FunctionMetadata
  Name: String

```

## Properties

`AutoPublish`

A flag that determines whether to automatically publish the function to the
`LIVE` stage when it’s created. To automatically publish to the
`LIVE` stage, set this property to `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FunctionCode`

The function code. For more information about writing a CloudFront function, see [Writing\
function code for CloudFront Functions](../../../amazoncloudfront/latest/developerguide/writing-function-code.md) in the
_Amazon CloudFront Developer Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FunctionConfig`

Contains configuration information about a CloudFront function.

_Required_: Yes

_Type_: [FunctionConfig](aws-properties-cloudfront-function-functionconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FunctionMetadata`

Contains metadata about a CloudFront function.

_Required_: No

_Type_: [FunctionMetadata](aws-properties-cloudfront-function-functionmetadata.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name to identify the function.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9-_]{1,64}`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`FunctionARN`

The ARN of the function. For example:

`arn:aws:cloudfront::123456789012:function/ExampleFunction`.

To get the function ARN, use the following syntax:

`!GetAtt
			Function_Logical_ID.FunctionMetadata.FunctionARN`

`FunctionMetadata.FunctionARN`

The Amazon Resource Name (ARN) of the function. The ARN uniquely identifies the
function.

## Examples

### Create a CloudFront function

The following examples show how to create a basic CloudFront function.

#### YAML

```yaml

Resources:
  CloudFrontFunction:
    Type: AWS::CloudFront::Function
    Properties:
      Name: MyFunctionName
      FunctionConfig:
        Comment: A basic CloudFront function
        Runtime: cloudfront-js-2.0
      FunctionCode: |
        function handler(event) {
            // NOTE: This example function is for a viewer request event trigger.
            // Choose viewer request for the event trigger when you associate this function with a distribution.
            var response = {
                statusCode: 200,
                statusDescription: 'OK',
                headers: {
                    'cloudfront-functions': { value: 'generated-by-CloudFront-Functions' }
                }
            };
            return response;
        }
      AutoPublish: true
```

#### JSON

```json

{
    "Resources": {
        "CloudFrontFunction": {
            "Type": "AWS::CloudFront::Function",
            "Properties": {
                "Name": "MyFunctionNameJSON",
                "FunctionConfig": {
                    "Comment": "A basic CloudFront function",
                    "Runtime": "cloudfront-js-2.0"
                },
                "FunctionCode": "function handler(event) {\n    // NOTE: This example function is for a viewer request event trigger.\n    // Choose viewer request for the event trigger when you associate this function with a distribution.\n    var response = {\n        statusCode: 200,\n        statusDescription: 'OK',\n        headers: {\n            'cloudfront-functions': { value: 'generated-by-CloudFront-Functions' }\n        }\n    };\n    return response;\n}\n",
                "AutoPublish": true
            }
        }
    }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WebAclCustomization

FunctionConfig

All content copied from https://docs.aws.amazon.com/.
