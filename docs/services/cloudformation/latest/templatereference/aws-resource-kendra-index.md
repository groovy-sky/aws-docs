This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::Index

Creates an Amazon Kendra index

Once the index is active you can add documents to your index using the [BatchPutDocument](../../../kendra/latest/dg/batchputdocument.md) operation or using one of the supported data sources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Kendra::Index",
  "Properties" : {
      "CapacityUnits" : CapacityUnitsConfiguration,
      "Description" : String,
      "DocumentMetadataConfigurations" : [ DocumentMetadataConfiguration, ... ],
      "Edition" : String,
      "Name" : String,
      "RoleArn" : String,
      "ServerSideEncryptionConfiguration" : ServerSideEncryptionConfiguration,
      "Tags" : [ Tag, ... ],
      "UserContextPolicy" : String,
      "UserTokenConfigurations" : [ UserTokenConfiguration, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Kendra::Index
Properties:
  CapacityUnits:
    CapacityUnitsConfiguration
  Description: String
  DocumentMetadataConfigurations:
    - DocumentMetadataConfiguration
  Edition: String
  Name: String
  RoleArn: String
  ServerSideEncryptionConfiguration:
    ServerSideEncryptionConfiguration
  Tags:
    - Tag
  UserContextPolicy: String
  UserTokenConfigurations:
    - UserTokenConfiguration

```

## Properties

`CapacityUnits`

Specifies additional capacity units configured for your Enterprise Edition index. You
can add and remove capacity units to fit your usage requirements.

_Required_: No

_Type_: [CapacityUnitsConfiguration](aws-properties-kendra-index-capacityunitsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description for the index.

_Required_: No

_Type_: String

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentMetadataConfigurations`

Specifies the properties of an index field. You can add either a custom or a built-in
field. You can add and remove built-in fields at any time. When a built-in field is
removed it's configuration reverts to the default for the field. Custom fields can't be
removed from an index after they are added.

_Required_: No

_Type_: Array of [DocumentMetadataConfiguration](aws-properties-kendra-index-documentmetadataconfiguration.md)

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Edition`

Indicates whether the index is a Enterprise Edition index, a Developer Edition index,
or a GenAI Enterprise Edition index.

_Required_: Yes

_Type_: String

_Allowed values_: `DEVELOPER_EDITION | ENTERPRISE_EDITION | GEN_AI_ENTERPRISE_EDITION`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the index.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

An IAM role that gives Amazon Kendra permissions to access your Amazon CloudWatch logs
and metrics. This is also the role used when you use the [BatchPutDocument](../../../kendra/latest/dg/batchputdocument.md) operation to
index documents from an Amazon S3 bucket.

_Required_: Yes

_Type_: String

_Pattern_: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

_Minimum_: `1`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerSideEncryptionConfiguration`

The identifier of the AWS KMS customer managed key (CMK) to use to
encrypt data indexed by Amazon Kendra. Amazon Kendra doesn't support asymmetric
CMKs.

_Required_: No

_Type_: [ServerSideEncryptionConfiguration](aws-properties-kendra-index-serversideencryptionconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-kendra-index-tag.md)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserContextPolicy`

The user context policy.

ATTRIBUTE\_FILTER

- All indexed content is searchable and displayable for all users. If you want
to filter search results on user context, you can use the attribute filters of
`_user_id` and `_group_ids` or you can provide user
and group information in `UserContext`.

USER\_TOKEN

- Enables token-based user access control to filter search results on user
context. All documents with no access control and all documents accessible to
the user will be searchable and displayable.

_Required_: No

_Type_: String

_Allowed values_: `ATTRIBUTE_FILTER | USER_TOKEN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserTokenConfigurations`

Defines the type of user token used for the index.

_Required_: No

_Type_: Array of [UserTokenConfiguration](aws-properties-kendra-index-usertokenconfiguration.md)

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the index ID. For example:

`{"Ref": "index-id"}`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the index. For example:
`arn:aws:kendra:us-west-2:111122223333:index/0123456789abcdef`.

`Id`

The identifier for the index. For example:
`f4aeaa10-8056-4b2c-a343-522ca0f41234`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

CapacityUnitsConfiguration

All content copied from https://docs.aws.amazon.com/.
