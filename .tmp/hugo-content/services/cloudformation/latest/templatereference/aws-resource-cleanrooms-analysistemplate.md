This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::AnalysisTemplate

Creates a new analysis template.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CleanRooms::AnalysisTemplate",
  "Properties" : {
      "AnalysisParameters" : [ AnalysisParameter, ... ],
      "Description" : String,
      "ErrorMessageConfiguration" : ErrorMessageConfiguration,
      "Format" : String,
      "MembershipIdentifier" : String,
      "Name" : String,
      "Schema" : AnalysisSchema,
      "Source" : AnalysisSource,
      "SourceMetadata" : AnalysisSourceMetadata,
      "SyntheticDataParameters" : SyntheticDataParameters,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CleanRooms::AnalysisTemplate
Properties:
  AnalysisParameters:
    - AnalysisParameter
  Description: String
  ErrorMessageConfiguration:
    ErrorMessageConfiguration
  Format: String
  MembershipIdentifier: String
  Name: String
  Schema:
    AnalysisSchema
  Source:
    AnalysisSource
  SourceMetadata:
    AnalysisSourceMetadata
  SyntheticDataParameters:
    SyntheticDataParameters
  Tags:
    - Tag

```

## Properties

`AnalysisParameters`

The parameters of the analysis template.

_Required_: No

_Type_: Array of [AnalysisParameter](aws-properties-cleanrooms-analysistemplate-analysisparameter.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the analysis template.

_Required_: No

_Type_: String

_Pattern_: `^[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDBFF-\uDC00\uDFFF\t\r\n]*$`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ErrorMessageConfiguration`

The configuration that specifies the level of detail in error messages returned by analyses using this template.
When set to `DETAILED`, error messages include more information to help troubleshoot issues with PySpark jobs.
Detailed error messages may expose underlying data, including sensitive information. Recommended for faster troubleshooting in development and testing environments.

_Required_: No

_Type_: [ErrorMessageConfiguration](aws-properties-cleanrooms-analysistemplate-errormessageconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Format`

The format of the analysis template.

_Required_: Yes

_Type_: String

_Allowed values_: `SQL | PYSPARK_1_0`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MembershipIdentifier`

The identifier for a membership resource.

_Required_: Yes

_Type_: String

_Pattern_: `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the analysis template.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_](([a-zA-Z0-9_ ]+-)*([a-zA-Z0-9_ ]+))?$`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Schema`

The entire schema object.

_Required_: No

_Type_: [AnalysisSchema](aws-properties-cleanrooms-analysistemplate-analysisschema.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Source`

The source of the analysis template.

_Required_: Yes

_Type_: [AnalysisSource](aws-properties-cleanrooms-analysistemplate-analysissource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceMetadata`

The source metadata for the analysis template.

_Required_: No

_Type_: [AnalysisSourceMetadata](aws-properties-cleanrooms-analysistemplate-analysissourcemetadata.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SyntheticDataParameters`

The parameters used to generate synthetic data for this analysis template.

_Required_: No

_Type_: [SyntheticDataParameters](aws-properties-cleanrooms-analysistemplate-syntheticdataparameters.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An optional label that you can assign to a resource when you create it. Each tag
consists of a key and an optional value, both of which you define. When you use tagging,
you can also use tag-based access control in IAM policies to control access
to this resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-cleanrooms-analysistemplate-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `AnalysisTemplateIdentifier`, such as `a1b2c3d4-5678-90ab-cdef-EXAMPLE2222`. For example:

`{ "Ref": "myAnalysisTemplate" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AnalysisTemplateIdentifier`

Returns the identifier for the analysis template.

Example: `a1b2c3d4-5678-90ab-cdef-EXAMPLE2222`

`Arn`

Returns the Amazon Resource Name (ARN) of the analysis template.

Example: `arn:aws:cleanrooms:us-east-1:111122223333:membership/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111/analysistemplates/a1b2c3d4-5678-90ab-cdef-EXAMPLE2222`

`CollaborationArn`

Returns the unique ARN for the analysis template’s associated collaboration.

Example: `arn:aws:cleanrooms:us-east-1:111122223333:collaboration/a1b2c3d4-5678-90ab-cdef-EXAMPLE33333`

`CollaborationIdentifier`

Returns the unique ID for the associated collaboration of the analysis template.

Example: `a1b2c3d4-5678-90ab-cdef-EXAMPLE33333`

`MembershipArn`

Returns the Amazon Resource Name (ARN) of the member who created the analysis template.

Example: `arn:aws:cleanrooms:us-east-1:111122223333:membership/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`

## Examples

### Create an analysis template

The following example creates an analysis template.

#### JSON

```json

{
    "ExampleAnalysisTemplate": {
        "Type": "AWS::CleanRooms::AnalysisTemplate",
        "Properties": {
            "MembershipIdentifier": "a1b2c3d4-5678-90ab-cdef-EXAMPLE1111",
            "Name": "exampleAnalysisTemplate",
            "Description": "example description",
            "Source": {
                "Text": "SELECT * FROM cta1 WHERE cta1.column1 > :Param1 AND cta1.column2 like :Param2"
            },
            "Format": "SQL",
            "AnalysisParameters": [
                {
                    "Name": "Param1",
                    "Type": "SMALLINT",
                    "DefaultValue": 1
                },
                {
                    "Name": "Param2",
                    "Type": "CHAR"
                }
            ],
            "Tags": [
                {
                    "Key": "Hello",
                    "Value": "World"
                }
            ]
        }
    }
}
```

#### YAML

```yaml

ExampleAnalysisTemplate:
    Type: 'AWS::CleanRooms::AnalysisTemplate'
    Properties:
      MembershipIdentifier: a1b2c3d4-5678-90ab-cdef-EXAMPLE1111
      Name: exampleAnalysisTemplate
      Description: example description
      Source:
        Text: SELECT * FROM cta1 WHERE cta1.column1 > :Param1 AND cta1.column2 like :Param2
      Format: SQL
      AnalysisParameters:
        - Name: Param1
          Type: SMALLINT
          DefaultValue: 1
        - Name: Param2
          Type: CHAR
      Tags:
        - Key: Hello
          Value: World
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Clean Rooms

AnalysisParameter

All content copied from https://docs.aws.amazon.com/.
