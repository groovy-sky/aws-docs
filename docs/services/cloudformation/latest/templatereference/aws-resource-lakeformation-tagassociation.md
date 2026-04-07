This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LakeFormation::TagAssociation

The `AWS::LakeFormation::TagAssociation` resource represents an assignment of an LF-tag to a Data Catalog resource (database, table, or column).
During a stack operation, CloudFormation calls AWS Lake Formation `AddLFTagsToResource` API to create a `TagAssociation` resource and calls the `RemoveLFTagsToResource` API to delete it.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::LakeFormation::TagAssociation",
  "Properties" : {
      "LFTags" : [ LFTagPair, ... ],
      "Resource" : Resource
    }
}

```

### YAML

```yaml

Type: AWS::LakeFormation::TagAssociation
Properties:
  LFTags:
    - LFTagPair
  Resource:
    Resource

```

## Properties

`LFTags`

A structure containing an LF-tag key-value pair.

_Required_: Yes

_Type_: Array of [LFTagPair](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lakeformation-tagassociation-lftagpair.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Resource`

UTF-8 string (valid values: `DATABASE | TABLE`).

The resource for which the LF-tag policy applies.

_Required_: Yes

_Type_: [Resource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lakeformation-tagassociation-resource.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a combination of the `ResourceIdentifier` and `TagsIdentifier` seperated with a pipe.

For example: `{"Catalog":null,"Database":{"CatalogId":null,"Name":"ExampleDbName"},"Table":null,"TableWithColumns":null}|[{"CatalogId":null,"TagKey":"tagKey1","TagValues":null}]`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ResourceIdentifier`

Json encoding of the input resource.

###### Examples

- Database: `{"Catalog":null,"Database":{"CatalogId":"123456789012","Name":"ExampleDbName"},"Table":null,"TableWithColumns":null}`

- Table: `{"Catalog":null,"Database":null,"Table":{"CatalogId":"123456789012","DatabaseName":"ExampleDbName","Name":"ExampleTableName","TableWildcard":null},"TableWithColumns":null}`

- Columns: `{"Catalog":null,"Database":null,"Table":null,"TableWithColumns":{"CatalogId":"123456789012","DatabaseName":"ExampleDbName","Name":"ExampleTableName","ColumnNames":["ExampleColName1","ExampleColName2"]}}`

`TagsIdentifier`

Json encoding of the input LFTags list.

For example: `[{"CatalogId":null,"TagKey":"tagKey1","TagValues":null},{"CatalogId":null,"TagKey":"tagKey2","TagValues":null}]`

## Remarks

Note that for a valid `TagAssociation` to be created, a database or table or column(s) must be specified.

## Examples

- [Examples for creating an LF-tag for a database](#aws-resource-lakeformation-tagassociation--examples--Examples_for_creating_an_LF-tag_for_a_database)

- [Examples for creating and LF-tag for a table](#aws-resource-lakeformation-tagassociation--examples--Examples_for_creating_and_LF-tag_for_a_table)

- [Examples for creating an LF-tag for columns](#aws-resource-lakeformation-tagassociation--examples--Examples_for_creating_an_LF-tag_for_columns)

### Examples for creating an LF-tag for a database

The following example demonstrates how to create an LF-tag on a database.

#### JSON

```json

{
    "SampleTagOnDatabase": {
        "Type": "AWS::LakeFormation::TagAssociation",
        "Properties": {
            "Resource": {
                "Database": {
                    "CatalogId": "12345678910",
                    "Name": "sample_db"
                }
            },
            "LFTags": [
                {
                    "CatalogId": "12345678910",
                    "TagKey": "sample_tag_key",
                    "TagValues": ["sample_tag_value1"]
                }
            ]
        }
    }
}
```

#### YAML

```yaml

SampleTagOnDatabase:
    Type: AWS::LakeFormation::TagAssociation
    Properties:
      Resource:
        Database:
          CatalogId: "12345678910"
          Name: "sample_db"
      LFTags:
        - CatalogId: "12345678910"
          TagKey: "sample_tag_key"
          TagValues:
            - "sample_tag_value1"

```

### Examples for creating and LF-tag for a table

The following example demonstrates how to create an LF-tag on a table.

#### JSON

```json

{
    "SampleTagOnTable": {
        "Type": "AWS::LakeFormation::TagAssociation",
        "Properties": {
            "Resource": {
                "Table": {
                    "CatalogId": "12345678910",
                    "DatabaseName": "sample_db",
                    "Name": "sample_tbl"
                }
            },
            "LFTags": [
                {
                    "CatalogId": "12345678910",
                    "TagKey": "sample_tag_key",
                    "TagValues": ["sample_tag_value1"]
                }
            ]
        }
    }
}
```

#### YAML

```yaml

SampleTagOnTable:
    Type: AWS::LakeFormation::TagAssociation
    Properties:
      Resource:
        Table:
          CatalogId: "12345678910"
          DatabaseName: "sample_db"
          Name: "sample_tbl"
      LFTags:
        - CatalogId: "12345678910"
          TagKey: "sample_tag_key"
          TagValues:
            - "sample_tag_value1"

```

### Examples for creating an LF-tag for columns

The following example demonstrates how to create LF-tags on columns.

#### JSON

```json

{
    "SampleTagOnColumn": {
        "Type": "AWS::LakeFormation::TagAssociation",
        "Properties": {
            "Resource": {
                "TableWithColumns": {
                    "CatalogId": "12345678910",
                    "DatabaseName": "sample_db",
                    "Name": "sample_tbl",
                    "ColumnNames": ["sample_col1", "sample_col2"]
                }
            },
            "LFTags": [
                {
                    "CatalogId": "12345678910",
                    "TagKey": "sample_tag_key",
                    "TagValues": ["sample_tag_value1"]
                }
            ]
        }
    }
}
```

#### YAML

```yaml

TestTagAssociation:
    Type: AWS::LakeFormation::TagAssociation
    Properties:
      Resource:
        TableWithColumns:
          CatalogId: "12345678910"
          DatabaseName: "sample_db"
          Name: "sample_tbl"
          ColumnNames:
            - "sample_col1"
            - "sample_col2"
      LFTags:
        - CatalogId: "12345678910"
          TagKey: "sample_tag_key"
          TagValues:
            - "sample_tag_value1"
```

## See also

[Permission requirements for tag assignments.](https://docs.aws.amazon.com/lake-formation/latest/dg/TBAC-assigning-tags.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::LakeFormation::Tag

DatabaseResource
