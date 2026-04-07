This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsGuru::ResourceCollection

A collection of AWS resources supported by DevOps Guru. The one type of AWS resource
collection supported is AWS CloudFormation stacks. DevOps Guru can be configured to analyze
only the AWS resources that are defined in the stacks.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DevOpsGuru::ResourceCollection",
  "Properties" : {
      "ResourceCollectionFilter" : ResourceCollectionFilter
    }
}

```

### YAML

```yaml

Type: AWS::DevOpsGuru::ResourceCollection
Properties:
  ResourceCollectionFilter:
    ResourceCollectionFilter

```

## Properties

`ResourceCollectionFilter`

Information about a filter used to specify which AWS resources are analyzed for anomalous behavior by DevOps Guru.

_Required_: Yes

_Type_: [ResourceCollectionFilter](aws-properties-devopsguru-resourcecollection-resourcecollectionfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function,
`Ref` returns Amazon Resource Name (ARN) of the `ResourceCollection`.

For more information about using the `Ref` function, see
[Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. The following are the available
attributes and sample return values. For more information about using `Fn::GetAtt`, see
[Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

`ResourceCollectionType`

The type of AWS resource collections to return. The one valid value is `CLOUD_FORMATION` for AWS CloudFormation stacks.

## Examples

- [Create a resource collection using two CloudFormation stacks](#aws-resource-devopsguru-resourcecollection--examples--Create_a_resource_collection_using_two_CloudFormation_stacks)

- [Create a resource collection with all CloudFormation stacks in your account](#aws-resource-devopsguru-resourcecollection--examples--Create_a_resource_collection_with_all_CloudFormation_stacks_in_your_account)

- [Monitor DevOps Guru resources using tags](#aws-resource-devopsguru-resourcecollection--examples--Monitor_DevOps_Guru_resources_using_tags)

### Create a resource collection using two CloudFormation stacks

#### JSON

```json

{
  "Resources": {
    "MyResourceCollection": {
      "Type": "AWS::DevOpsGuru::ResourceCollection",
      "Properties": {
        "ResourceCollectionFilter": {
          "CloudFormation": {
            "StackNames": [
              "StackA",
              "StackB"
            ]
          }
        }
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  MyResourceCollection:
    Type: AWS::DevOpsGuru::ResourceCollection
    Properties:
      ResourceCollectionFilter:
        CloudFormation:
          StackNames:
          - StackA
          - StackB
```

### Create a resource collection with all CloudFormation stacks in your account

#### JSON

```json

{
  "Resources": {
    "MyResourceCollection": {
      "Type": "AWS::DevOpsGuru::ResourceCollection",
      "Properties": {
        "ResourceCollectionFilter": {
          "CloudFormation": {
            "StackNames": [
              "*"
            ]
          }
        }
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  MyResourceCollection:
    Type: AWS::DevOpsGuru::ResourceCollection
    Properties:
      ResourceCollectionFilter:
        CloudFormation:
          StackNames:
          - "*"
```

### Monitor DevOps Guru resources using tags

#### JSON

```json

{
"DevOpsGuruResourceCollection": {
  "Type": "AWS::DevOpsGuru::ResourceCollection",
  "Condition": "DevOpsGuruEnable",
  "Properties": {
    "ResourceCollectionFilter": {
      "Tags": [{
        "AppBoundaryKey": "devops-guru-workshop",
        "TagValues": [
          "devops-guru-serverless",
          "devops-guru-aurora"
          ]
        }]
      }
    }
  }
}
```

#### YAML

```yaml

DevOpsGuruResourceCollection:
    Type: AWS::DevOpsGuru::ResourceCollection
    Condition: DevOpsGuruEnable
    Properties:
      ResourceCollectionFilter:
        Tags:
          - AppBoundaryKey: devops-guru-workshop
            TagValues:
              - devops-guru-serverless
	      - devops-guru-aurora

```

#### JSON

```json

{
 "DevOpsGuruResourceCollection": {
    "Type": "AWS::DevOpsGuru::ResourceCollection",
    "Condition": "DevOpsGuruEnable",
    "Properties": {
      "ResourceCollectionFilter": {
        "Tags": [
          {
            "AppBoundaryKey": "devops-guru-workshop",
            "TagValues": [
              "*"
            ]
          }
        ]
      }
    }
  }
}
```

#### YAML

```yaml

DevOpsGuruResourceCollection:
    Type: AWS::DevOpsGuru::ResourceCollection
    Condition: DevOpsGuruEnable
    Properties:
      ResourceCollectionFilter:
        Tags:
          - AppBoundaryKey: devops-guru-workshop
            TagValues:
              - "*"

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SnsChannelConfig

CloudFormationCollectionFilter
