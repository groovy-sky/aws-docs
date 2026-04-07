This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SpotFleet

Specifies a Spot Fleet request.

The Spot Fleet request specifies the total target capacity and the On-Demand target
capacity. Amazon EC2 calculates the difference between the total capacity and On-Demand
capacity, and launches the difference as Spot capacity.

You can submit a single request that includes multiple launch specifications that vary
by instance type, AMI, Availability Zone, or subnet.

By default, the Spot Fleet requests Spot Instances in the Spot Instance pool where the
price per unit is the lowest. Each launch specification can include its own instance
weighting that reflects the value of the instance type to your application
workload.

Alternatively, you can specify that the Spot Fleet distribute the target capacity
across the Spot pools included in its launch specifications. By ensuring that the Spot
Instances in your Spot Fleet are in different Spot pools, you can improve the
availability of your fleet.

You can specify tags for the Spot Fleet request and instances launched by the fleet.
You cannot tag other resource types in a Spot Fleet request because only the
`spot-fleet-request` and `instance` resource types are
supported.

For more information, see [Spot Fleet](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet.html)
in the _Amazon EC2 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::SpotFleet",
  "Properties" : {
      "SpotFleetRequestConfigData" : SpotFleetRequestConfigData,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::SpotFleet
Properties:
  SpotFleetRequestConfigData:
    SpotFleetRequestConfigData
  Tags:
    - Tag

```

## Properties

`SpotFleetRequestConfigData`

Describes the configuration of a Spot Fleet request.

_Required_: Yes

_Type_: [SpotFleetRequestConfigData](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags for a Spot Fleet resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-spotfleet-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the Spot Fleet.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The ID of the Spot Fleet.

## Examples

- [Create a Spot Fleet using a launch template](#aws-resource-ec2-spotfleet--examples--Create_a_Spot_Fleet_using_a_launch_template)

- [Create a Spot Fleet using a launch specification](#aws-resource-ec2-spotfleet--examples--Create_a_Spot_Fleet_using_a_launch_specification)

### Create a Spot Fleet using a launch template

The following example specifies a Spot Fleet with a target capacity of 2 instances
and a launch template that provides the instance configuration.

#### JSON

```json

{
    "Resources": {
        "mySpotFleet": {
            "Type": "AWS::EC2::SpotFleet",
            "Properties": {
                "SpotFleetRequestConfigData": {
                    "IamFleetRole": "arn:aws:iam::123456789012:role/aws-ec2-spot-fleet-tagging-role",
                    "TargetCapacity": 2,
                    "LaunchTemplateConfigs": [
                        {
                            "LaunchTemplateSpecification": {
                                "LaunchTemplateId": {
                                    "Ref": "myLaunchTemplate"
                                },
                                "Version": {
                                    "Fn::GetAtt": [
                                        "myLaunchTemplate",
                                        "DefaultVersionNumber"
                                    ]
                                }
                            }
                        }
                    ]
                }
            }
        },
        "myLaunchTemplate": {
            "Type": "AWS::EC2::LaunchTemplate",
            "Properties": {
                "LaunchTemplateData": {
                    "ImageId": "ami-0a70b9d193ae8a799",
                    "InstanceType": "t2.micro",
                    "SecurityGroupIds": [
                        "sg-12a4c434"
                    ],
                    "TagSpecifications": [
                        {
                            "ResourceType": "instance",
                            "Tags": [
                                {
                                    "Key": "department",
                                    "Value": "dev"
                                }
                            ]
                        }
                    ]
                }
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
   mySpotFleet:
     Type: 'AWS::EC2::SpotFleet'
     Properties:
       SpotFleetRequestConfigData:
         IamFleetRole: arn:aws:iam::123456789012:role/aws-ec2-spot-fleet-tagging-role
         TargetCapacity: 2
         LaunchTemplateConfigs:
           - LaunchTemplateSpecification:
               LaunchTemplateId: !Ref myLaunchTemplate
               Version: !GetAtt myLaunchTemplate.DefaultVersionNumber
   myLaunchTemplate:
     Type: 'AWS::EC2::LaunchTemplate'
     Properties:
       LaunchTemplateData:
          ImageId: ami-0a70b9d193ae8a799
          InstanceType: t2.micro
          SecurityGroupIds:
            - sg-12a4c434
          TagSpecifications:
            - ResourceType: instance
              Tags:
              - Key: department
                Value: dev
```

### Create a Spot Fleet using a launch specification

The following example specifies a Spot Fleet with two launch specifications. The
weighted capacities are the same, so Amazon EC2 launches the same number of instances
for each specification. For more information, see [How Spot Fleet Works](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet.html) in the
_Amazon EC2 User Guide_.

#### JSON

```json

"SpotFleet": {
  "Type": "AWS::EC2::SpotFleet",
  "Properties": {
    "SpotFleetRequestConfigData": {
     "IamFleetRole": { "Fn::GetAtt": [ "IAMFleetRole", "Arn"] },
     "SpotPrice": "1000",
     "TargetCapacity": { "Ref": "TargetCapacity" },
     "LaunchSpecifications": [
     {
         "EbsOptimized": "false",
         "InstanceType": { "Ref": "InstanceType" },
         "ImageId": { "Fn::FindInMap": [ "AWSRegionArch2AMI", { "Ref": "AWS::Region" },
                      { "Fn::FindInMap": [ "AWSInstanceType2Arch", { "Ref": "InstanceType" }, "Arch" ] }
                     ]},
          "SubnetId": { "Ref": "Subnet1" },
          "WeightedCapacity": "8"
     },
     {
         "EbsOptimized": "true",
         "InstanceType": { "Ref": "InstanceType" },
         "ImageId": { "Fn::FindInMap": [ "AWSRegionArch2AMI", { "Ref": "AWS::Region" },
                      { "Fn::FindInMap": [ "AWSInstanceType2Arch", { "Ref": "InstanceType" }, "Arch" ] }
                     ]},
          "Monitoring": { "Enabled": "true" },
          "SecurityGroups": [ { "GroupId": { "Fn::GetAtt": [ "SG0", "GroupId" ] } } ],
          "SubnetId": { "Ref": "Subnet0" },
          "IamInstanceProfile": { "Arn": { "Fn::GetAtt": [ "RootInstanceProfile", "Arn" ] } },
          "WeightedCapacity": "8"
         }
         ]
     }
   }
}
```

#### YAML

```yaml

SpotFleet:
  Type: AWS::EC2::SpotFleet
  Properties:
    SpotFleetRequestConfigData:
      IamFleetRole: !GetAtt [IAMFleetRole, Arn]
      SpotPrice: '1000'
      TargetCapacity:
        Ref: TargetCapacity
      LaunchSpecifications:
      - EbsOptimized: 'false'
        InstanceType:
          Ref: InstanceType
      ImageId:
        Fn::FindInMap:
        - AWSRegionArch2AMI
        - Ref: AWS::Region
        - Fn::FindInMap:
          - AWSInstanceType2Arch
          - Ref: InstanceType
          - Arch
      SubnetId:
        Ref: Subnet1
      WeightedCapacity: '8'
    - EbsOptimized: 'true'
      InstanceType:
        Ref: InstanceType
      ImageId:
        Fn::FindInMap:
        - AWSRegionArch2AMI
        - Ref: AWS::Region
        - Fn::FindInMap:
        - AWSInstanceType2Arch
        - Ref: InstanceType
        - Arch
      Monitoring:
        Enabled: 'true'
        SecurityGroups:
         - GroupId:
             Fn::GetAtt:
             - SG0
             - GroupId
       SubnetId:
          Ref: Subnet0
       IamInstanceProfile:
         Arn:
         Fn::GetAtt:
         - RootInstanceProfile
         - Arn
       WeightedCapacity: '8'
```

## See also

- [AWS::ApplicationAutoScaling::ScalableTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html)

- [AWS::ApplicationAutoScaling::ScalingPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::SnapshotBlockPublicAccess

AcceleratorCountRequest
