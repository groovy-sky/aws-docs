This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchServerless::Collection

Specifies an OpenSearch Serverless collection. For more information, see [Creating and\
managing Amazon OpenSearch Serverless collections](../../../opensearch-service/latest/developerguide/serverless-manage.md) in the _Amazon_
_OpenSearch Service Developer Guide_.

###### Important

To create a collection successfully, you must associate a KMS key for encryption. You have two options:
1\. Specify the KMS key directly in the CreateCollection request – Use the encryption-config parameter when creating the collection.
2\. Define the KMS key in an encryption security policy – Create a matching [encryption policy](../../../opensearch-service/latest/developerguide/serverless-encryption.md) before or alongside your collection.
Note: If you specify a KMS key in both locations, the key provided in the CreateCollection request takes precedence over the security policy configuration.
When using CloudFormation templates, you can include the encryption policy resource in the same template as the collection resource by using the [DependsOn](../userguide/aws-attribute-dependson.md) attribute to ensure proper resource creation order. Alternatively, ensure the encryption policy or KMS key configuration exists before creating the collection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::OpenSearchServerless::Collection",
  "Properties" : {
      "CollectionGroupName" : String,
      "Description" : String,
      "EncryptionConfig" : EncryptionConfig,
      "Name" : String,
      "StandbyReplicas" : String,
      "Tags" : [ Tag, ... ],
      "Type" : String,
      "VectorOptions" : VectorOptions
    }
}

```

### YAML

```yaml

Type: AWS::OpenSearchServerless::Collection
Properties:
  CollectionGroupName: String
  Description: String
  EncryptionConfig:
    EncryptionConfig
  Name: String
  StandbyReplicas: String
  Tags:
    - Tag
  Type: String
  VectorOptions:
    VectorOptions

```

## Properties

`CollectionGroupName`

The name of the collection group to associate with the collection.

_Required_: No

_Type_: String

_Pattern_: `[a-z][a-z0-9-]+`

_Minimum_: `3`

_Maximum_: `32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A description of the collection.

_Required_: No

_Type_: String

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionConfig`

Encryption settings for the collection.

_Required_: No

_Type_: [EncryptionConfig](aws-properties-opensearchserverless-collection-encryptionconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the collection.

Collection names must meet the following criteria:

- Starts with a lowercase letter

- Unique to your account and AWS Region

- Contains between 3 and 28 characters

- Contains only lowercase letters a-z, the numbers 0-9, and the hyphen (-)

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z][a-z0-9-]{2,31}$`

_Minimum_: `3`

_Maximum_: `32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StandbyReplicas`

Indicates whether to use standby replicas for the collection. You can't update this
property after the collection is already created. If you attempt to modify this property,
the collection continues to use the original value.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An arbitrary set of tags (key–value pairs) to associate with the collection.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-opensearchserverless-collection-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The type of collection. Possible values are `SEARCH`,
`TIMESERIES`, and `VECTORSEARCH`. For more information, see [Choosing a collection type](../../../opensearch-service/latest/developerguide/serverless-overview.md#serverless-usecase).

_Required_: No

_Type_: String

_Allowed values_: `SEARCH | TIMESERIES | VECTORSEARCH`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VectorOptions`

Configuration options for vector search capabilities in the collection.

_Required_: No

_Type_: [VectorOptions](aws-properties-opensearchserverless-collection-vectoroptions.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the collection ID. For more information about using the
`Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

`GetAtt` returns a value for a specified attribute of this type. For more
information, see [Fn::GetAtt](../userguide/intrinsic-function-reference-getatt.md). The following are the available attributes and sample return
values.

`Arn`

The Amazon Resource Name (ARN) of the collection. For example,
`arn:aws:aoss:us-east-1:123456789012:collection/07tjusf2h91cunochc`.

`CollectionEndpoint`

Collection-specific endpoint used to submit index, search, and data upload requests to
an OpenSearch Serverless collection. For example,
`https://07tjusf2h91cunochc.us-east-1.aoss.amazonaws.com`.

`DashboardEndpoint`

The collection-specific endpoint used to access OpenSearch Dashboards. For example,
`https://07tjusf2h91cunochc.us-east-1.aoss.amazonaws.com/_dashboards`.

`Id`

A unique identifier for the collection. For example,
`07tjusf2h91cunochc`.

`KmsKeyArn`

The ARN of the AWS KMS key used to encrypt the collection.

## Examples

### Create a collection

The following example specifies an OpenSearch Serverless collection named
`test-collection`. The collection type is `SEARCH`. The
template also creates a matching encryption policy, which is required in order for
the collection to be created successfully.

For a complete sample policy that creates network, encryption, and access
policies, as well as a matching collection, see [Using AWS CloudFormation to create Amazon OpenSearch Serverless\
collections](../../../opensearch-service/latest/developerguide/serverless-cfn.md) in the _Amazon OpenSearch Service Developer_
_Guide._

###### Note

This example uses public network access, which isn't recommended for production
workloads. We recommend using VPC access to protect your collections. For more
information, see [AWS::OpenSearchServerless::VpcEndpoint](../userguide/aws-resource-opensearchserverless-vpcendpoint.md) and [Access\
Amazon OpenSearch Serverless using an interface endpoint](../../../opensearch-service/latest/developerguide/serverless-vpc.md).

#### JSON

```json

{
   "Description":"OpenSearch Serverless collection template",
   "Resources":{
      "TestCollection":{
         "Type":"AWS::OpenSearchServerless::Collection",
         "Properties":{
            "Name":"test-collection",
            "Type":"SEARCH",
            "Description":"Search collection"
         },
         "DependsOn":"EncryptionPolicy"
      },
      "EncryptionPolicy":{
         "Type":"AWS::OpenSearchServerless::SecurityPolicy",
         "Properties":{
            "Name":"test-encryption-policy",
            "Type":"encryption",
            "Description":"Encryption policy for test collection",
            "Policy":"{\"Rules\":[{\"ResourceType\":\"collection\",\"Resource\":[\"collection/test-collection\"]}],\"AWSOwnedKey\":true}"
         }
      }
   }
```

#### YAML

```yaml

Description: OpenSearch Serverless collection template
Resources:
  TestCollection:
    Type: 'AWS::OpenSearchServerless::Collection'
    Properties:
      Name: test-collection
      Type: SEARCH
      Description: Search collection
    DependsOn: EncryptionPolicy
  EncryptionPolicy:
    Type: 'AWS::OpenSearchServerless::SecurityPolicy'
    Properties:
      Name: test-encryption-policy
      Type: encryption
      Description: Encryption policy for test collection
      Policy: >-
        {"Rules":[{"ResourceType":"collection","Resource":["collection/test-collection"]}],"AWSOwnedKey":true}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::OpenSearchServerless::AccessPolicy

EncryptionConfig

All content copied from https://docs.aws.amazon.com/.
