This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LakeFormation::Tag

The `AWS::LakeFormation::Tag` resource represents an LF-tag, which consists of a key and one or more possible values for the key.
During a stack operation, AWS CloudFormation calls the AWS Lake Formation `CreateLFTag` API to create a tag, and `UpdateLFTag` API to update a tag resource, and a `DeleteLFTag` to delete it.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::LakeFormation::Tag",
  "Properties" : {
      "CatalogId" : String,
      "TagKey" : String,
      "TagValues" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::LakeFormation::Tag
Properties:
  CatalogId: String
  TagKey: String
  TagValues:
    - String

```

## Properties

`CatalogId`

Catalog id string, not less than 1 or more than 255 bytes long, matching the [single-line string pattern](../../../lake-formation/latest/dg/aws-lake-formation-api-aws-lake-formation-api-common.md).

The identifier for the Data Catalog. By default, the account ID.
The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your AWS Lake Formation environment.

_Required_: No

_Type_: String

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TagKey`

UTF-8 string, not less than 1 or more than 255 bytes long, matching the [single-line string pattern](../../../lake-formation/latest/dg/aws-lake-formation-api-aws-lake-formation-api-common.md).

The key-name for the LF-tag.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

_Required_: Yes

_Type_: String

_Pattern_: `^([{a-zA-Z}{\s}{0-9}_.:\/=+\-@%]*)$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TagValues`

An array of UTF-8 strings, not less than 1 or more than 50 strings.

A list of possible values of the corresponding `TagKey` of an LF-tag key-value pair.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Tag’s `TagKey` property value.

For example: `tagKeyName`

## Remarks

_Note the following:_

Only data lake administrators can create LF-tags.

An `LF-tag` can be assigned to Data Catalog resources (databases, tables, and columns) via [AWS::LakeFormation::TagAssociation](../userguide/aws-resource-lakeformation-tagassociation.md) to implement tag-based access control.

## Examples

### Creating a tag resource in a template

The following example demonstrates how to define a tag resource in a template.

#### JSON

```json

{
    "SampleTag": {
        "Type": "AWS::LakeFormation::Tag",
        "Properties": {
            "TagKey": "sample_tag_key",
            "TagValues": ["sample_tag_value1", "sample_tag_value2"]
        }
    }
}

```

#### YAML

```yaml

SampleTag:
    Type: AWS::LakeFormation::Tag
    Properties:
      TagKey: "sample_tag_key"
      TagValues:
        - "sample_tag_value1"
        - "sample_tag_value2"
```

## See also

[Assign an `LF-tag` to a Data Catalog resource - AWS::LakeFormation::TagAssociation](../userguide/aws-properties-lakeformation-tagassociation-resource.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::LakeFormation::Resource

AWS::LakeFormation::TagAssociation

All content copied from https://docs.aws.amazon.com/.
