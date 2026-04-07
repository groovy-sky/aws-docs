This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::Faq

Creates an new set of frequently asked question (FAQ) questions and answers.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Kendra::Faq",
  "Properties" : {
      "Description" : String,
      "FileFormat" : String,
      "IndexId" : String,
      "LanguageCode" : String,
      "Name" : String,
      "RoleArn" : String,
      "S3Path" : S3Path,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Kendra::Faq
Properties:
  Description: String
  FileFormat: String
  IndexId: String
  LanguageCode: String
  Name: String
  RoleArn: String
  S3Path:
    S3Path
  Tags:
    - Tag

```

## Properties

`Description`

A description for the FAQ.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FileFormat`

The format of the input file. You can choose between a basic CSV format, a CSV format
that includes customs attributes in a header, and a JSON format that includes custom
attributes.

The format must match the format of the file stored in the S3 bucket identified in
the S3Path parameter.

Valid values are:

- `CSV`

- `CSV_WITH_HEADER`

- `JSON`

_Required_: No

_Type_: String

_Allowed values_: `CSV | CSV_WITH_HEADER | JSON`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IndexId`

The identifier of the index that contains the FAQ.

_Required_: Yes

_Type_: String

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LanguageCode`

The code for a language. This shows a supported language for the FAQ document
as part of the summary information for FAQs. English is supported by default.
For more information on supported languages, including their codes,
see [Adding \
documents in languages other than English](https://docs.aws.amazon.com/kendra/latest/dg/in-adding-languages.html).

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z-]*`

_Minimum_: `2`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name that you assigned the FAQ when you created or updated the FAQ.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The Amazon Resource Name (ARN) of a role with permission to access the S3 bucket that
contains the FAQ.

_Required_: Yes

_Type_: String

_Pattern_: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

_Minimum_: `1`

_Maximum_: `1284`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3Path`

The Amazon Simple Storage Service (Amazon S3) location of the FAQ input data.

_Required_: Yes

_Type_: [S3Path](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kendra-faq-s3path.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kendra-faq-tag.html)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the FAQ identifier. For example:

`{ "Ref": "<faq-id>|<index-id>" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

`arn:aws:kendra:us-west-2:111122223333:index/335c3741-41df-46a6-b5d3-61f85b787884/faq/f61995a6-cd5c-4e99-9cfc-58816d8bfaa7`

`Id`

The identifier for the FAQ. For example:

`f61995a6-cd5c-4e99-9cfc-58816d8bfaa7`

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

WorkDocsConfiguration

S3Path
