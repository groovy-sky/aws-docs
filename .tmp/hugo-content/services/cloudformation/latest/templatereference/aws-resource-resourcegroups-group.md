This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ResourceGroups::Group

Creates a resource group with the specified name and description. You can optionally
include either a resource query or a service configuration. For more information about
constructing a resource query, see [Build queries and groups in\
AWS Resource Groups](../../../arg/latest/userguide/getting-started-query.md) in the _AWS Resource Groups User Guide_. For more information
about service-linked groups and service configurations, see [Service configurations for Resource Groups](../../../../reference/arg/latest/apireference/about-slg.md).

**Minimum permissions**

To run this command, you must have the following permissions:

- `resource-groups:CreateGroup`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ResourceGroups::Group",
  "Properties" : {
      "Configuration" : [ ConfigurationItem, ... ],
      "Description" : String,
      "Name" : String,
      "ResourceQuery" : ResourceQuery,
      "Resources" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ResourceGroups::Group
Properties:
  Configuration:
    - ConfigurationItem
  Description: String
  Name: String
  ResourceQuery:
    ResourceQuery
  Resources:
    - String
  Tags:
    - Tag

```

## Properties

`Configuration`

The service configuration currently associated with the resource group and in effect
for the members of the resource group. A `Configuration` consists of one or
more `ConfigurationItem` entries. For information about service
configurations for resource groups and how to construct them, see [Service\
configurations for resource groups](../../../../reference/arg/latest/apireference/about-slg.md) in the _AWS Resource Groups_
_User Guide_.

###### Note

You can include either a `Configuration` or a
`ResourceQuery`, but not both.

_Required_: Conditional

_Type_: Array of [ConfigurationItem](aws-properties-resourcegroups-group-configurationitem.md)

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the resource group.

_Required_: No

_Type_: String

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of a resource group. The name must be unique within the AWS
Region in which you create the resource. To create multiple resource groups based on the
same CloudFormation stack, you must generate unique names for each.

_Required_: Yes

_Type_: String

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceQuery`

The resource query structure that is used to dynamically determine which AWS resources are members of the associated resource group. For more
information about queries and how to construct them, see [Build queries and groups in\
AWS Resource Groups](../../../arg/latest/userguide/gettingstarted-query.md) in the _AWS Resource Groups User_
_Guide_

###### Note

- You can include either a `ResourceQuery` or a
`Configuration`, but not both.

- You can specify the group's membership either by using a
`ResourceQuery` or by using a list of `Resources`,
but not both.

_Required_: Conditional

_Type_: [ResourceQuery](aws-properties-resourcegroups-group-resourcequery.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Resources`

A list of the Amazon Resource Names (ARNs) of AWS resources that you
want to add to the specified group.

###### Note

- You can specify the group membership either by using a list of
`Resources` or by using a `ResourceQuery`, but not
both.

- You can include a `Resources` property only if you also specify
a `Configuration` property.

_Required_: Conditional

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tag key and value pairs that are attached to the resource group.

_Required_: No

_Type_: Array of [Tag](aws-properties-resourcegroups-group-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

The name of the new resource group.

### Fn::GetAtt

`Arn`

The ARN of the new resource group.

## Examples

The following examples show both the JSON and YAML templates that you can use to
create resource groups with the specified characteristics.

- [Creating a CloudFormation stack-based resource group using defaults](#aws-resource-resourcegroups-group--examples--Creating_a_CloudFormation_stack-based_resource_group_using_defaults)

- [Creating a CloudFormation stack-based resource group with specific resources](#aws-resource-resourcegroups-group--examples--Creating_a_CloudFormation_stack-based_resource_group_with_specific_resources)

- [Creating a CloudFormation stack-based resource group based on another stack](#aws-resource-resourcegroups-group--examples--Creating_a_CloudFormation_stack-based_resource_group_based_on_another_stack)

- [Creating a tag-based resource group](#aws-resource-resourcegroups-group--examples--Creating_a_tag-based_resource_group)

- [Creating a resource group for EC2 capacity reservations](#aws-resource-resourcegroups-group--examples--Creating_a_resource_group_for_EC2_capacity_reservations)

- [Creating a resource group for EC2 capacity reservations with initial members](#aws-resource-resourcegroups-group--examples--Creating_a_resource_group_for_EC2_capacity_reservations_with_initial_members)

- [Creating an EC2 hosts resource group](#aws-resource-resourcegroups-group--examples--Creating_an_EC2_hosts_resource_group)

- [Creating an EC2 hosts resource group with some initial members](#aws-resource-resourcegroups-group--examples--Creating_an_EC2_hosts_resource_group_with_some_initial_members)

### Creating a CloudFormation stack-based resource group using defaults

This example creates a CloudFormation stack-based resource group using
defaults. It includes all supported resource types and uses the same
`StackIdentifier` as the CloudFormation stack that defines the
group.

#### JSON

```json

{
    "ResourceGroup": {
        "Type": "AWS::ResourceGroups::Group",
        "Properties": {
            "Name": "MyResourceGroup",
            "Description": "A basic, empty resource group that is created by CFN",
        }
    }
}
```

#### YAML

```yaml

ResourceGroup:
  Type: "AWS::ResourceGroups::Group"
  Properties:
    Name: "MyMinimalResourceGroup"
    Description: "A basic, empty resource group that is created by CFN"
```

### Creating a CloudFormation stack-based resource group with specific resources

This example creates a CloudFormation stack-based resource group that's
similar to the group shown in the previous example. The difference is that it
can include only resources of the specified types: EC2 instances and DynamoDB
tables.

#### JSON

```json

{
    "CloudFormationStackGroupForSelectedResourceTypes": {
        "Type": "AWS::ResourceGroups::Group",
        "Properties": {
            "Name": "MyCloudFormationResourceGroup-Filters",
            "Description": "A basic resource group that can hold only EC2 instances or DynamoDB tables",
            "ResourceQuery": {
                "Query": {
                    "ResourceTypeFilters": [
                        "AWS::EC2::Instance",
                        "AWS::DynamoDB::Table"
                    ]
                }
            }
        }
    }
}
```

#### YAML

```yaml

CloudFormationStackGroupForSelectedResourceTypes:
  Type: "AWS::ResourceGroups::Group"
  Properties:
    Name: "MyCloudFormationResourceGroup-Filters"
    Description: "A basic resource group that can hold only EC2 instances or DynamoDB tables"
    ResourceQuery:
      Query:
        ResourceTypeFilters:
          - "AWS::EC2::Instance"
          - "AWS::DynamoDB::Table"
```

### Creating a CloudFormation stack-based resource group based on another stack

This example creates a resource group that's built from another stack. The
`StackIdentifier` value specifies that resources in the stack
with the ARN `arn:aws:cloudformation:us-east-1:123456789012:stack/stack-name/9b6f8604-4a39-490c-870b-44b0ebdd38b9`
are included in the group. The group itself (not its individual member
resources) is tagged with `Env=Prod`.

#### JSON

```json

{
    "CloudFormationStackGroupForAnotherStack": {
        "Type": "AWS::ResourceGroups::Group",
        "Properties": {
            "Name": "MyCloudFormationResourceGroupForAnotherStack",
            "Description": "A resource group that is based on another CFN stack",
            "ResourceQuery": {
                "Type": "CLOUDFORMATION_STACK_1_0",
                "Query": {
                    "ResourceTypeFilters": [
                        "AWS::AllSupported"
                    ],
                    "StackIdentifier": "arn:aws:cloudformation:us-east-1:123456789012:stack/stack-name/9b6f8604-4a39-490c-870b-44b0ebdd38b9"
                }
            },
            "Tags": [
                {
                    "Key": "Env",
                    "Value": "Prod"
                }
            ]
        }
    }
}
```

#### YAML

```yaml

CloudFormationStackGroupForAnotherStack:
  Type: "AWS::ResourceGroups::Group"
  Properties:
    Name: "MyCloudFormationResourceGroupForAnotherStack"
    Description: "A group that is based on CFN another stack"
    ResourceQuery:
      Type: "CLOUDFORMATION_STACK_1_0"
      Query:
        ResourceTypeFilters:
          - "AWS::AllSupported"
        StackIdentifier: "arn:aws:cloudformation:us-east-1:123456789012:stack/stack-name/9b6f8604-4a39-490c-870b-44b0ebdd38b9"
    Tags:
      -
        Key: "Env"
        Value: "Prod"
```

### Creating a tag-based resource group

This example shows how to create a resource group whose membership is based on
a tag query. Any resources that have tags that match the query, in this case the
key `Usage` and the value `Integration Tests`, are members
of this group.

#### JSON

```json

{
    "TagBasedGroup": {
        "Type": "AWS::ResourceGroups::Group",
        "Properties": {
            "Name": "MyTagBasedResourceGroup",
            "Description": "A group that is based on a tag query",
            "ResourceQuery": {
                "Type": "TAG_FILTERS_1_0",
                "Query": {
                    "ResourceTypeFilters": [
                        "AWS::AllSupported"
                    ],
                    "TagFilters": [
                        {
                            "Key": "Usage",
                            "Values": [
                                "Integration Tests"
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

TagBasedGroup:
  Type: "AWS::ResourceGroups::Group"
  Properties:
    Name: "MyTagBasedResourceGroup"
    Description: "A group that is based on a tag query"
    ResourceQuery:
      Type:
        "TAG_FILTERS_1_0"
      Query:
        ResourceTypeFilters:
          - "AWS::AllSupported"
        TagFilters:
          -
            Key: "Usage"
            Values:
              - "Integration Tests"
```

### Creating a resource group for EC2 capacity reservations

This example creates a capacity reservations resource group by specifying a
configuration and limiting the allowed resource types. This group initially has
no members.

#### JSON

```json

{
   "CapacityReservationsGroupWithoutResources": {
      "Type": "AWS::ResourceGroups::Group",
      "Properties": {
         "Name": "MyCapacityReservationsGroup",
         "Description": "A resource group for EC2 capacity reservations",
         "Configuration": [
            {
               "Type": "AWS::EC2::CapacityReservationPool",
               "Parameters": []
            },
            {
               "Type": "AWS::ResourceGroups::Generic",
               "Parameters": [
                  {
                     "Name": "allowed-resource-types",
                     "Values": [
                        "AWS::EC2::CapacityReservation"
                     ]
                  }
               ]
            }
         ]
      }
   }
}
```

#### YAML

```yaml

CapacityReservationsGroupWithoutResources:
  Type: "AWS::ResourceGroups::Group"
  Properties:
    Name: "MyCapacityReservationsGroup"
    Description: "A resource group for EC2 capacity reservations"
    Configuration:
    - Type: "AWS::EC2::CapacityReservationPool"
      Parameters: []
    - Type: "AWS::ResourceGroups::Generic"
      Parameters:
      - Name: "allowed-resource-types"
        Values:
          - "AWS::EC2::CapacityReservation"
```

### Creating a resource group for EC2 capacity reservations with initial members

This example creates a capacity reservations resource group identical to the
previous example, except this one includes two initial capacity reservations as
specified in the `Resources` property.

#### JSON

```json

{
   "CapacityReservationsGroupWithResources": {
      "Type": "AWS::ResourceGroups::Group",
      "Properties": {
         "Name": "MyCapacityReservationsGroup",
         "Description": "A resource group for EC2 capacity reservations",
         "Configuration": [
            {
               "Type": "AWS::EC2::CapacityReservationPool",
               "Parameters": []
            },
            {
               "Type": "AWS::ResourceGroups::Generic",
               "Parameters": [
                  {
                     "Name": "allowed-resource-types",
                     "Values": [
                        "AWS::EC2::CapacityReservation"
                     ]
                  }
               ]
            }
         ],
         "Resources": [
            "arn:aws:ec2:us-east-1:123456789012:capacity-reservation/cr-0d4953834cd1a96a3",
            "arn:aws:ec2:us-east-1:123456789012:capacity-reservation/cr-0069a2275c16b7333"
         ]
      }
   }
}
```

#### YAML

```yaml

CapacityReservationsGroupWithResources:
  Type: "AWS::ResourceGroups::Group"
  Properties:
    Name: "MyCapacityReservationsGroup"
    Description: "A resource group for EC2 capacity reservations"
    Configuration:
    - Type: "AWS::EC2::CapacityReservationPool"
      Parameters: []
    - Type: "AWS::ResourceGroups::Generic"
      Parameters:
      - Name: "allowed-resource-types"
        Values:
          - "AWS::EC2::CapacityReservation"
    Resources:
      - "arn:aws:ec2:us-east-1:123456789012:capacity-reservation/cr-0d4953834cd1a96a3"
      - "arn:aws:ec2:us-east-1:123456789012:capacity-reservation/cr-0069a2275c16b7333"
```

### Creating an EC2 hosts resource group

This example creates a host resource group. Instances in the group are
automatically configured to accept any host based license configuration, to
allow only the `c5` host family, to not auto-release hosts, and to be
delete protected. The group initially has no members.

#### JSON

```json

{
  "HostResourceGroup": {
    "Type": "AWS::ResourceGroups::Group",
    "Properties": {
      "Name": "MyHostResourceGroup",
      "Description": "A resource group for EC2 dedicated hosts",
      "Configuration": [
        {
          "Type": "AWS::EC2::HostManagement",
          "Parameters": [
            {
              "Name": "any-host-based-license-configuration",
              "Values": [
                "true"
              ]
            },
            {
              "Name": "allowed-host-families",
              "Values": [
                "c5"
              ]
            },
            {
              "Name": "auto-release-host",
              "Values": [
                "false"
              ]
            }
          ]
        },
        {
          "Type": "AWS::ResourceGroups::Generic",
          "Parameters": [
            {
              "Name": "allowed-resource-types",
              "Values": [
                "AWS::EC2::Host"
              ]
            },
            {
              "Name": "deletion-protection",
              "Values": [
                "UNLESS_EMPTY"
              ]
            }
          ]
        }
      ]
    }
  }
}
```

#### YAML

```yaml

HostResourceGroup:
  Type: "AWS::ResourceGroups::Group"
  Properties:
    Name: "MyHostResourceGroup"
    Description: "A resource group for EC2 dedicated hosts"
    Configuration:
      - Type: "AWS::EC2::HostManagement"
        Parameters:
          - Name: "any-host-based-license-configuration"
            Values:
              - "true"
          - Name: "allowed-host-families"
            Values:
              - "c5"
          - Name: "auto-release-host"
            Values:
              - "false"
      - Type: "AWS::ResourceGroups::Generic"
        Parameters:
          - Name: "allowed-resource-types"
            Values:
              - "AWS::EC2::Host"
          - Name: "deletion-protection"
            Values:
              - "UNLESS_EMPTY"
```

### Creating an EC2 hosts resource group with some initial members

This example creates an EC2 host resource group that is configured to work
with a specific license configuration. The group initially contains the two
hosts specified by the `Resources` property

#### JSON

```json

{
  "HostResourceGroupWithResources": {
    "Type": "AWS::ResourceGroups::Group",
    "Properties": {
      "Name": "MyHostResourceGroup",
      "Description": "A resource group for EC2 dedicated hosts",
      "Configuration": [
        {
          "Type": "AWS::EC2::HostManagement",
          "Parameters": [
            {
              "Name": "allowed-resource-types",
              "Values": [
                "AWS::EC2::Host"
              ]
            },
            {
              "Name": "deletion-protection",
              "Values": [
                "UNLESS_EMPTY"
              ]
            }
          ]
        }
      ],
      "Resources": [
        "arn:aws:ec2:us-east-1:123456789012:dedicated-host/h-0375ef7462b26b11f",
        "arn:aws:ec2:us-east-1:123456789012:dedicated-host/h-0501ab6812c719123"
      ]
    }
  }
}
```

#### YAML

```yaml

HostResourceGroupWithResources:
  Type: "AWS::ResourceGroups::Group"
  Properties:
    Name: "MyHostResourceGroup"
    Description: "A resource group for EC2 dedicated hosts"
    Configuration:
      - Type: "AWS::EC2::HostManagement"
        Parameters:
          - Name: "allowed-host-based-license-configurations"
            Values:
              - "arn:aws:license-manager:us-east-1:123456789012:license-configuration:lic-42bc5628e8edcee52ed797d5bebf0879"
        Parameters:
          - Name: "allowed-resource-types"
            Values:
              - "AWS::EC2::Host"
          - Name: "deletion-protection"
            Values:
              - "UNLESS_EMPTY"
    Resources:
      - "arn:aws:ec2:us-east-1:123456789012:dedicated-host/h-0375ef7462b26b11f"
      - "arn:aws:ec2:us-east-1:123456789012:dedicated-host/h-0501ab6812c719123"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Resource Groups

ConfigurationItem

All content copied from https://docs.aws.amazon.com/.
