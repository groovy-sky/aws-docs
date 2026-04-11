This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Macie::AllowList S3WordsList

Specifies the location and name of an Amazon Simple Storage Service (Amazon S3)
object that lists specific, predefined text to ignore when inspecting data sources for sensitive
data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String,
  "ObjectKey" : String
}

```

### YAML

```yaml

  BucketName: String
  ObjectKey: String

```

## Properties

`BucketName`

The full name of the S3 bucket that contains the object. This value correlates to the
`Name` field of a bucket's properties in Amazon S3.

This value is case sensitive. In addition, don't use wildcard characters or specify
partial values for the name.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObjectKey`

The full name of the S3 object. This value correlates to the `Key` field of
an object's properties in Amazon S3. If the name includes a path, include the
complete path. For example, `AllowLists/Macie/MyList.txt`.

This value is case sensitive. In addition, don't use wildcard characters or specify
partial values for the name.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Criteria

Tag

All content copied from https://docs.aws.amazon.com/.
