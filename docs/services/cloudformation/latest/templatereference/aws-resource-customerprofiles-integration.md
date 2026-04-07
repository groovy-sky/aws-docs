This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::Integration

Specifies an Amazon Connect Customer Profiles Integration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CustomerProfiles::Integration",
  "Properties" : {
      "DomainName" : String,
      "EventTriggerNames" : [ String, ... ],
      "FlowDefinition" : FlowDefinition,
      "ObjectTypeName" : String,
      "ObjectTypeNames" : [ ObjectTypeMapping, ... ],
      "Scope" : String,
      "Tags" : [ Tag, ... ],
      "Uri" : String
    }
}

```

### YAML

```yaml

Type: AWS::CustomerProfiles::Integration
Properties:
  DomainName: String
  EventTriggerNames:
    - String
  FlowDefinition:
    FlowDefinition
  ObjectTypeName: String
  ObjectTypeNames:
    - ObjectTypeMapping
  Scope: String
  Tags:
    - Tag
  Uri: String

```

## Properties

`DomainName`

The unique name of the domain.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EventTriggerNames`

A list of unique names for active event triggers associated with the integration.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `64 | 1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FlowDefinition`

The configuration that controls how Customer Profiles retrieves data from the
source.

_Required_: No

_Type_: [FlowDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-customerprofiles-integration-flowdefinition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObjectTypeName`

The name of the profile object type mapping to use.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z_][a-zA-Z_0-9-]*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObjectTypeNames`

The object type mapping.

_Required_: No

_Type_: Array of [ObjectTypeMapping](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-customerprofiles-integration-objecttypemapping.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scope`

The scope or boundary of the integration item's applicability.

_Required_: No

_Type_: String

_Allowed values_: `PROFILE | DOMAIN`

_Pattern_: `^[a-zA-Z_][a-zA-Z_0-9-]*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for this resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-customerprofiles-integration-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Uri`

The URI of the S3 bucket or any other type of data source.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the DomainName and the Uri of the integration.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The timestamp of when the integration was created.

`LastUpdatedAt`

The timestamp of when the integration was most recently edited.

## Examples

The following example creates an integration if Domain existed.

#### YAML

```yaml

Type: "AWS::CustomerProfiles::Integration"
Properties:
    DomainName: "ExampleDomain"
    ObjectTypeName: "CTR"
    Uri: "arn:aws:connect:us-east-1:123456789012:instance/11111111-1111-1111-1111-111111111111"
```

#### JSON

```json

"Type": "AWS::CustomerProfiles::Integration",
"Properties": {
    "DomainName": "ExampleDomain",
    "ObjectTypeName": "CTR",
    "Uri": "arn:aws:connect:us-east-1:123456789012:instance/11111111-1111-1111-1111-111111111111" } } }
}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

ConnectorOperator
