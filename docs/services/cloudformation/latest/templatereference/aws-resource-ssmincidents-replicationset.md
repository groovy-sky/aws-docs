This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMIncidents::ReplicationSet

The `AWS::SSMIncidents::ReplicationSet` resource specifies a set of AWS Regions
that Incident Manager data is replicated to and the AWS Key Management Service (AWS KMS
key used to encrypt the data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SSMIncidents::ReplicationSet",
  "Properties" : {
      "DeletionProtected" : Boolean,
      "Regions" : [ ReplicationRegion, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SSMIncidents::ReplicationSet
Properties:
  DeletionProtected: Boolean
  Regions:
    - ReplicationRegion
  Tags:
    - Tag

```

## Properties

`DeletionProtected`

Determines if the replication set deletion protection is enabled or not. If deletion
protection is enabled, you can't delete the last Region in the replication set.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Regions`

Specifies the Regions of the replication set.

_Required_: Yes

_Type_: Array of [ReplicationRegion](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ssmincidents-replicationset-replicationregion.html)

_Minimum_: `1`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of tags to add to the replication set.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ssmincidents-replicationset-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Create a replication set

The following example creates a replication set.

###### Note

We recommend creating a replication set and response plan using a single
template. For a demonstration, see the examples for
[AWS::SSMIncidents::ResponsePlan](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#aws-resource-ssmincidents-responseplan--examples).

#### JSON

```json

{
   "AWSTemplateFormatVersion": "2010-09-09",
   "Description": "Sample AWS CloudFormation template to create a replication set (JSON).",
   "Resources": {
      "MyReplicationSet": {
         "Type": "AWS::SSMIncidents::ReplicationSet",
         "Properties": {
            "DeletionProtected": true,
            "Regions": [
               {
                  "RegionName": {
                     "Ref": "AWS::Region"
                  }
               }
            ],
            "Tags": [
               {
                  "Key": "MyReplicationSetTagKey",
                  "Value": "MyReplicationSetTagValue"
               }
            ]
         }
      }
   }
}
```

#### YAML

```yaml

---
AWSTemplateFormatVersion: 2010-09-09
Description: "Sample AWS CloudFormation template to create a replication set (YAML)."
Resources:
  MyReplicationSet:
    Type: AWS::SSMIncidents::ReplicationSet
    Properties:
      DeletionProtected: true
      Regions:
        - RegionName:
            Ref: "AWS::Region"
      Tags:
        - Key: MyReplicationSetTagKey
          Value: MyReplicationSetTagValue
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Systems Manager Incident Manager

RegionConfiguration
