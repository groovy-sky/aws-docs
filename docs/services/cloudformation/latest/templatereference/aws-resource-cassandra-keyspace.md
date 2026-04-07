This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cassandra::Keyspace

You can use the `AWS::Cassandra::Keyspace` resource to create a new keyspace
in Amazon Keyspaces (for Apache Cassandra). For more information, see [Create a\
keyspace](https://docs.aws.amazon.com/keyspaces/latest/devguide/getting-started.keyspaces.html) in the _Amazon Keyspaces_
_Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Cassandra::Keyspace",
  "Properties" : {
      "ClientSideTimestampsEnabled" : Boolean,
      "KeyspaceName" : String,
      "ReplicationSpecification" : ReplicationSpecification,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Cassandra::Keyspace
Properties:
  ClientSideTimestampsEnabled: Boolean
  KeyspaceName: String
  ReplicationSpecification:
    ReplicationSpecification
  Tags:
    - Tag

```

## Properties

`ClientSideTimestampsEnabled`

Indicates whether client-side timestamps are enabled (true) or disabled (false) for all tables in the keyspace.
To add a Region to a single-Region keyspace with at least one table, the value must be set to true.
After you've enabled client-side timestamps for a table, you can’t disable it again.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyspaceName`

The name of the keyspace to be created. The keyspace name is case sensitive. If you don't specify a name, AWS
CloudFormation generates a unique ID and uses that ID for the keyspace name. For more
information, see [Name\
type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).

_Length constraints:_ Minimum length of 1. Maximum length of 48.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_]{1,47}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReplicationSpecification`

Specifies the `ReplicationStrategy`
of a keyspace. The options are:

- `SINGLE_REGION` for a single Region keyspace (optional) or

- `MULTI_REGION` for a multi-Region keyspace

If no `ReplicationStrategy` is provided, the default is `SINGLE_REGION`. If you choose `MULTI_REGION`,
you must also provide a `RegionList` with the AWS Regions that the keyspace is replicated in.

_Required_: No

_Type_: [ReplicationSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cassandra-keyspace-replicationspecification.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cassandra-keyspace-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the keyspace. For example:

`{ "Ref": "MyNewKeyspace" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

This section includes code examples that demonstrate how to create a keyspace using the different options.

- [Create a new keyspace with tags](#aws-resource-cassandra-keyspace--examples--Create_a_new_keyspace_with_tags)

- [Create a new multi-Region keyspace](#aws-resource-cassandra-keyspace--examples--Create_a_new_multi-Region_keyspace)

- [Add a new AWS Region to a single-Region keyspace](#aws-resource-cassandra-keyspace--examples--Add_a_new_Region_to_a_single-Region_keyspace)

- [Add a new AWS Region to a multi-Region keyspace](#aws-resource-cassandra-keyspace--examples--Add_a_new_Region_to_a_multi-Region_keyspace)

### Create a new keyspace with tags

The following example creates a new keyspace named
`MyNewKeyspace` with the following tags: `{'key1':'val1', 'key2':'val2'}`.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "MyNewKeyspace": {
      "Type": "AWS::Cassandra::Keyspace",
      "Properties": {
        "KeyspaceName": "MyNewKeyspace",
        "Tags": [{"Key":"tag1","Value":"val1"}, {"Key":"tag2","Value":"val2"}]
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  MyNewKeyspace:
    Type: AWS::Cassandra::Keyspace
    Properties:
      KeyspaceName: MyNewKeyspace
      Tags:
      - Key: tag1
      Value: val1
      - Key: tag2
      Value: val2
```

### Create a new multi-Region keyspace

The following example creates a new multi-Region keyspace named
`MultiRegionKeyspace` that is replicated across three AWS Regions.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "MultiRegionKeyspace": {
      "Type": "AWS::Cassandra::Keyspace",
      "Properties": {
        "KeyspaceName": "MultiRegionKeyspace",
        "ReplicationSpecification": {
          "ReplicationStrategy": "MULTI_REGION",
          "RegionList": ["us-east-1", "us-west-2", "eu-west-1"]
        }
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  MultiRegionKeyspace:
    Type: AWS::Cassandra::Keyspace
    Properties:
      KeyspaceName: MultiRegionKeyspace
      ReplicationSpecification:
        ReplicationStrategy: MULTI_REGION
        RegionList:
          - us-east-1
          - us-west-2
          - eu-west-1
```

### Add a new AWS Region to a single-Region keyspace

The following example shows how to add the Region `eu-west-1` to a single-Region keyspace named
`MyNewKeyspace` that is currently available in `us-east-1`.

#### JSON

```json

{
   "AWSTemplateFormatVersion":"2010-09-09",
   "Resources":{
      "MultiRegionKeyspace":{
         "Type":"AWS::Cassandra::Keyspace",
         "Properties":{
            "KeyspaceName":"MyNewKeyspace",
            "ReplicationSpecification":{
               "ReplicationStrategy":"MULTI_REGION",
               "RegionList":[
                  "us-east-1",
                  "eu-west-1"
               ]
            },
            "ClientSideTimestampsEnabled":true
         }
      }
   }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  MultiRegionKeyspace:
    Type: AWS::Cassandra::Keyspace
    Properties:
      KeyspaceName: MyNewKeyspace
      ReplicationSpecification:
        ReplicationStrategy: MULTI_REGION
        RegionList:
          - us-east-1
          - eu-west-1
     ClientSideTimestampsEnabled: true
```

### Add a new AWS Region to a multi-Region keyspace

The following example shows how to add the new Region `us-west-2` to the existing multi-Region keyspace named
`MultiRegionKeyspace` that is already replicated across `us-east-1` and `eu-west-1`.

#### JSON

```json

{
   "AWSTemplateFormatVersion":"2010-09-09",
   "Resources":{
      "MultiRegionKeyspace":{
         "Type":"AWS::Cassandra::Keyspace",
         "Properties":{
            "KeyspaceName":"MultiRegionKeyspace",
            "ReplicationSpecification":{
               "ReplicationStrategy":"MULTI_REGION",
               "RegionList":[
                  "us-east-1",
                  "eu-west-1",
                  "us-west-2"
               ]
            }
         }
      }
   }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  MultiRegionKeyspace:
    Type: AWS::Cassandra::Keyspace
    Properties:
      KeyspaceName: MultiRegionKeyspace
      ReplicationSpecification:
        ReplicationStrategy: MULTI_REGION
        RegionList:
          - us-east-1
          - us-west-1
          - us-west-2
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Keyspaces (for Apache Cassandra)

ReplicationSpecification
