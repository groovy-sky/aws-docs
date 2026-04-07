This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::Retriever

Adds a retriever to your Amazon Q Business application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::QBusiness::Retriever",
  "Properties" : {
      "ApplicationId" : String,
      "Configuration" : RetrieverConfiguration,
      "DisplayName" : String,
      "RoleArn" : String,
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::QBusiness::Retriever
Properties:
  ApplicationId: String
  Configuration:
    RetrieverConfiguration
  DisplayName: String
  RoleArn: String
  Tags:
    - Tag
  Type: String

```

## Properties

`ApplicationId`

The identifier of the Amazon Q Business application using the retriever.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9-]{35}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Configuration`

Provides information on how the retriever used for your Amazon Q Business application is
configured.

_Required_: Yes

_Type_: [RetrieverConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-retriever-retrieverconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The name of your retriever.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_-]*$`

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of an IAM role used by Amazon Q Business to access the basic
authentication credentials stored in a Secrets Manager secret.

_Required_: No

_Type_: String

_Pattern_: `^arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}$`

_Minimum_: `0`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of key-value pairs that identify or categorize the retriever. You can also use
tags to help control access to the retriever. Tag keys and values can consist of Unicode
letters, digits, white space, and any of the following symbols: \_ . : / = + - @.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-retriever-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of your retriever.

_Required_: Yes

_Type_: String

_Allowed values_: `NATIVE_INDEX | KENDRA_INDEX`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the application ID and retriever ID. For example:

`{"Ref": "ApplicationId|RetrieverId"}`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The Unix timestamp when the retriever was created.

`RetrieverArn`

The Amazon Resource Name (ARN) of the IAM role associated with the retriever.

`RetrieverId`

The identifier of the retriever used by your Amazon Q Business application.

`Status`

The status of your retriever.

`UpdatedAt`

The Unix timestamp when the retriever was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

KendraIndexConfiguration
