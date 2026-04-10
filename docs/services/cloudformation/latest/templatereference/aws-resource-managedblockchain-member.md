This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ManagedBlockchain::Member

Creates a member within a Managed Blockchain network.

Applies only to Hyperledger Fabric.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ManagedBlockchain::Member",
  "Properties" : {
      "InvitationId" : String,
      "MemberConfiguration" : MemberConfiguration,
      "NetworkConfiguration" : NetworkConfiguration,
      "NetworkId" : String
    }
}

```

### YAML

```yaml

Type: AWS::ManagedBlockchain::Member
Properties:
  InvitationId: String
  MemberConfiguration:
    MemberConfiguration
  NetworkConfiguration:
    NetworkConfiguration
  NetworkId: String

```

## Properties

`InvitationId`

The unique identifier of the invitation to join the network sent to the account that creates the member.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MemberConfiguration`

Configuration properties of the member.

_Required_: Yes

_Type_: [MemberConfiguration](aws-properties-managedblockchain-member-memberconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkConfiguration`

Configuration properties of the network to which the member belongs.

_Required_: No

_Type_: [NetworkConfiguration](aws-properties-managedblockchain-member-networkconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkId`

The unique identifier of the network to which the member belongs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the member ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`MemberId`

The unique identifier of the member.

`NetworkId`

The unique identifier of the network to which the member belongs.

## Examples

- [Create a new Managed Blockchain network for Hyperledger Fabric and the first member](#aws-resource-managedblockchain-member--examples--Create_a_new_Managed_Blockchain_network_for_Hyperledger_Fabric_and_the_first_member)

- [Create an additional member in an existing Hyperledger Fabric network](#aws-resource-managedblockchain-member--examples--Create_an_additional_member_in_an_existing_Hyperledger_Fabric_network)

### Create a new Managed Blockchain network for Hyperledger Fabric and the first member

When you create the first member in a Managed Blockchain network, you also specify parameters to create the network.

#### YAML

```yaml

Description: "Basic initial member template"
Parameters:
  MemberName:
    Type: String
  MemberDescription:
    Type: String
  MemberAdminUsername:
    Type: String
  MemberAdminPassword:
    Type: String
  NetworkName:
    Type: String
  NetworkDescription:
    Type: String
  Edition:
    Type: String
  ThresholdPercentage:
    Type: Number
  ThresholdComparator:
    Type: String
  ProposalDurationInHours:
    Type: Number

Resources:
  Member:
    Type: "AWS::ManagedBlockchain::Member"
    Properties:
      NetworkConfiguration:
        Name: !Ref NetworkName
        Description: !Ref NetworkDescription
        Framework: "HYPERLEDGER_FABRIC"
        FrameworkVersion: "1.2"
        NetworkFrameworkConfiguration:
          NetworkFabricConfiguration:
            Edition: !Ref Edition
        VotingPolicy:
          ApprovalThresholdPolicy:
            ThresholdPercentage: !Ref ThresholdPercentage
            ProposalDurationInHours: !Ref ProposalDurationInHours
            ThresholdComparator: !Ref ThresholdComparator
      MemberConfiguration:
        Name: !Ref MemberName
        Description: !Ref MemberDescription
        MemberFrameworkConfiguration:
          MemberFabricConfiguration:
            AdminUsername: !Ref MemberAdminUsername
            AdminPassword: !Ref MemberAdminPassword

```

#### JSON

```json

{
  "Description": "Basic initial member template",
  "Parameters": {
    "MemberName": {
      "Type": "String"
    },
    "MemberDescription": {
      "Type": "String"
    },
    "MemberAdminUsername": {
      "Type": "String"
    },
    "MemberAdminPassword": {
      "Type": "String"
    },
    "NetworkName": {
      "Type": "String"
    },
    "NetworkDescription": {
      "Type": "String"
    },
    "Edition": {
      "Type": "String"
    },
    "ThresholdPercentage": {
      "Type": "Number"
    },
    "ThresholdComparator": {
      "Type": "String"
    },
    "ProposalDurationInHours": {
      "Type": "Number"
    }
  },
  "Resources": {
    "Member": {
      "Type": "AWS::ManagedBlockchain::Member",
      "Properties": {
        "NetworkConfiguration": {
          "Name": "NetworkName",
          "Description": "NetworkDescription",
          "Framework": "HYPERLEDGER_FABRIC",
          "FrameworkVersion": "1.2",
          "NetworkFrameworkConfiguration": {
            "NetworkFabricConfiguration": {
              "Edition": "Edition"
            }
          },
          "VotingPolicy": {
            "ApprovalThresholdPolicy": {
              "ThresholdPercentage": "ThresholdPercentage",
              "ProposalDurationInHours": "ProposalDurationInHours",
              "ThresholdComparator": "ThresholdComparator"
            }
          }
        },
        "MemberConfiguration": {
          "Name": "MemberName",
          "Description": "MemberDescription",
          "MemberFrameworkConfiguration": {
            "MemberFabricConfiguration": {
              "AdminUsername": "MemberAdminUsername",
              "AdminPassword": "MemberAdminPassword"
            }
          }
        }
      }
    }
  }
}
```

### Create an additional member in an existing Hyperledger Fabric network

#### YAML

```yaml

Description: "Basic additional member template"
Parameters:
  MemberName:
    Type: String
  MemberDescription:
    Type: String
  MemberAdminUsername:
    Type: String
  MemberAdminPassword:
    Type: String
  NetworkId:
    Type: String
  InvitationId:
    Type: String
Resources:
  Member:
    Type: 'AWS::ManagedBlockchain::Member'
    Properties:
      MemberConfiguration:
        Name: !Ref MemberName
        Description: !Ref MemberDescription
        MemberFrameworkConfiguration:
          MemberFabricConfiguration:
            AdminUsername: !Ref MemberAdminUsername
            AdminPassword: !Ref MemberAdminPassword
      NetworkId: !Ref NetworkId
      InvitationId: !Ref InvitationId

```

#### JSON

```json

{
           "Description": "Basic additional member template",
  "Parameters": {
      "MemberName": {
          "Type": "String"
      },
      "MemberDescription": {
          "Type": "String"
      },
      "MemberAdminUsername": {
          "Type": "String"
      },
      "MemberAdminPassword": {
          "Type": "String"
      },
      "NetworkId": {
          "Type": "String"
      },
      "InvitationId": {
          "Type": "String"
      }
  },
  "Resources": {
      "Member": {
          "Type": "AWS::ManagedBlockchain::Member",
          "Properties": {
              "MemberConfiguration": {
                  "Name": {
                      "Ref": "MemberName"
                  },
                  "MemberFrameworkConfiguration": {
                      "MemberFabricConfiguration": {
                          "AdminUsername": {
                              "Ref": "MemberAdminUsername"
                          },
                          "AdminPassword": {
                              "Ref": "MemberAdminPassword"
                          }
                      }
                  }
              },
              "NetworkId": {
                  "Ref": "NetworkId"
              },
              "InvitationId": {
                  "Ref": "InvitationId"
              }
          }
      }
  }
}

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

ApprovalThresholdPolicy

All content copied from https://docs.aws.amazon.com/.
