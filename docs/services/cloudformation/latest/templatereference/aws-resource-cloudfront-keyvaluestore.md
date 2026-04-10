This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::KeyValueStore

The key value store. Use this to separate data from function code, allowing you to update
data without having to publish a new version of a function. The key value store holds
keys and their corresponding values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFront::KeyValueStore",
  "Properties" : {
      "Comment" : String,
      "ImportSource" : ImportSource,
      "Name" : String
    }
}

```

### YAML

```yaml

Type: AWS::CloudFront::KeyValueStore
Properties:
  Comment: String
  ImportSource:
    ImportSource
  Name: String

```

## Properties

`Comment`

A comment for the key value store.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImportSource`

The import source for the key value store.

_Required_: No

_Type_: [ImportSource](aws-properties-cloudfront-keyvaluestore-importsource.md)

_Update requires_: Updates are not supported.

`Name`

The name of the key value store.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the key value store.

`Id`

The unique Id for the key value store.

`Status`

The current status of the key value store. For more information, see [Key value store statuses](../../../amazoncloudfront/latest/developerguide/kvs-with-functions-create.md#key-value-store-status) in the _Amazon CloudFront Developer Guide._

## Examples

### Key value store

The following example creates a key value store named
DemoKeyValueStore and a CloudFront function named
DemoFunction.

#### JSON

```json

{
    "Resources": {
        "KeyValueStore": {
            "Type": "AWS::CloudFront::KeyValueStore",
            "Properties": {"Name": "DemoKeyValueStore"}
        },
        "Function": {
            "Type": "AWS::CloudFront::Function",
            "Name": "DemoFunction",
            "FunctionConfig": {
                "Comment": "Function with KeyValueStore",
                "Runtime": "cloudfront-js-2.0",
                "KeyValueStoreAssociations": [
                    {
                        "KeyValueStoreARN": {"Fn::Sub": "${KeyValueStore.Arn}"}
                    }
                ]
            },
            "FunctionCode": {"Fn::Sub": "const kvsId = '${KeyValueStore.Id}';\n// ... remaining function code ..."},
            "AutoPublish": true
        }
    }
}
```

#### YAML

```yaml

Resources:
  KeyValueStore:
    Type: 'AWS::CloudFront::KeyValueStore'
    Properties:
      Name: 'DemoKeyValueStore'
  Function:
    Type: 'AWS::CloudFront::Function'
    Properties:
      Name: 'DemoFunction'
      FunctionConfig:
        Comment: 'Function with KeyValueStore'
        Runtime: 'cloudfront-js-2.0'
        KeyValueStoreAssociations:
          - KeyValueStoreARN: !Sub '${KeyValueStore.Arn}'
      FunctionCode: !Sub |
        const kvsId = '${KeyValueStore.Id}';
        // ... Remaining function code ...
      AutoPublish: true
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KeyGroupConfig

ImportSource

All content copied from https://docs.aws.amazon.com/.
