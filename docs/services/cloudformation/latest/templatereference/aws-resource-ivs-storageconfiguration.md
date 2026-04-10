This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IVS::StorageConfiguration

The `AWS::IVS::StorageConfiguration` resource specifies an Amazon IVS
storage configuration. A storage configuration describes an S3 location where recorded videos will be stored. For more information,
see [Auto-Record to Amazon S3 (Low-Latency Streaming)](../../../ivs/latest/lowlatencyuserguide/record-to-s3.md)
in the _Amazon IVS Low-Latency Streaming User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IVS::StorageConfiguration",
  "Properties" : {
      "Name" : String,
      "S3" : S3StorageConfiguration,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IVS::StorageConfiguration
Properties:
  Name: String
  S3:
    S3StorageConfiguration
  Tags:
    - Tag

```

## Properties

`Name`

Storage cnfiguration name.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_]*$`

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3`

An S3 storage configuration contains information about where recorded video will be stored. See the
[S3StorageConfiguration](../userguide/aws-properties-ivs-storageconfiguration-s3storageconfiguration.md)
property type for more information.

_Required_: Yes

_Type_: [S3StorageConfiguration](aws-properties-ivs-storageconfiguration-s3storageconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-ivs-storageconfiguration-tag.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-ivs-storageconfiguration-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the storage-configuration ARN. For example:

`{ "Ref": "myStorageConfiguration" }`

For the Amazon IVS storage configuration
`"myStorageConfiguration"`, `Ref` returns the
storage-configuration ARN.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The storage-configuration ARN. For example:
`arn:aws:ivs:us-west-2:123456789012:storage-configuration/abcdABCDefgh`

## Examples

### StorageConfiguration Template Examples

The following examples specify an Amazon IVS storage configuration.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "StorageConfiguration": {
            "Type": "AWS::IVS::StorageConfiguration",
            "Properties": {
                "Name": "myStorageConfiguration",
                "S3": { "BucketName": "my-bucket" },
                "Tags": [
                    {
                        "Key": "MyKey",
                        "Value": "MyValue"
                    }
                ]
            }
        }
    }
}

```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  StorageConfiguration:
    Type: AWS::IVS::StorageConfiguration
    Properties:
      Name: myStorageConfiguration
      S3:
        BucketName: my-bucket
      Tags:
        - Key: myKey
          Value: myValue

```

## See also

- [Getting\
Started with IVS Real-Time Streaming](../../../ivs/latest/realtimeuserguide/getting-started.md)

- [Server-Side Composition (Real-Time Streaming)](../../../ivs/latest/realtimeuserguide/server-side-composition.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ThumbnailConfiguration

S3StorageConfiguration

All content copied from https://docs.aws.amazon.com/.
