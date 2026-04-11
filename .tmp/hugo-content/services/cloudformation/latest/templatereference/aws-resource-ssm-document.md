This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::Document

The `AWS::SSM::Document` resource creates a Systems Manager (SSM)
document in AWS Systems Manager. This document defines the actions that Systems Manager performs on your AWS resources.

###### Note

This resource does not support CloudFormation drift detection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SSM::Document",
  "Properties" : {
      "Attachments" : [ AttachmentsSource, ... ],
      "Content" : Json,
      "DocumentFormat" : String,
      "DocumentType" : String,
      "Name" : String,
      "Requires" : [ DocumentRequires, ... ],
      "Tags" : [ Tag, ... ],
      "TargetType" : String,
      "UpdateMethod" : String,
      "VersionName" : String
    }
}

```

### YAML

```yaml

Type: AWS::SSM::Document
Properties:
  Attachments:
    - AttachmentsSource
  Content: Json
  DocumentFormat: String
  DocumentType: String
  Name: String
  Requires:
    - DocumentRequires
  Tags:
    - Tag
  TargetType: String
  UpdateMethod: String
  VersionName: String

```

## Properties

`Attachments`

A list of key-value pairs that describe attachments to a version of a document.

_Required_: No

_Type_: Array of [AttachmentsSource](aws-properties-ssm-document-attachmentssource.md)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Content`

The content for the new SSM document in JSON or YAML. For more information about the
schemas for SSM document content, see [SSM document\
schema features and examples](../../../systems-manager/latest/userguide/document-schemas-features.md) in the _AWS Systems Manager User_
_Guide_.

###### Note

This parameter also supports `String` data types.

_Required_: Yes

_Type_: Json

_Minimum_: `1`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`DocumentFormat`

Specify the document format for the request. `JSON` is the default
format.

_Required_: No

_Type_: String

_Allowed values_: `YAML | JSON | TEXT`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`DocumentType`

The type of document to create.

_Required_: No

_Type_: String

_Allowed values_: `ApplicationConfiguration | ApplicationConfigurationSchema | Automation | Automation.ChangeTemplate | AutoApprovalPolicy | ChangeCalendar | CloudFormation | Command | DeploymentStrategy | ManualApprovalPolicy | Package | Policy | ProblemAnalysis | ProblemAnalysisTemplate | Session`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

A name for the SSM document.

###### Important

You can't use the following strings as document name prefixes. These are reserved by AWS
for use as document name prefixes:

- `aws`

- `amazon`

- `amzn`

- `AWSEC2`

- `AWSConfigRemediation`

- `AWSSupport`

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-.]{3,128}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Requires`

A list of SSM documents required by a document. This parameter is used exclusively by
AWS AppConfig. When a user creates an AWS AppConfig configuration in an SSM document, the user must also
specify a required document for validation purposes. In this case, an
`ApplicationConfiguration` document requires an
`ApplicationConfigurationSchema` document for validation purposes. For more
information, see [What is AWS AppConfig?](../../../appconfig/latest/userguide/what-is-appconfig.md) in the
_AWS AppConfig User Guide_.

_Required_: No

_Type_: Array of [DocumentRequires](aws-properties-ssm-document-documentrequires.md)

_Minimum_: `1`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Tags`

AWS CloudFormation resource tags to apply to the document. Use tags to help you
identify and categorize resources.

_Required_: No

_Type_: Array of [Tag](aws-properties-ssm-document-tag.md)

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetType`

Specify a target type to define the kinds of resources the document can run on. For
example, to run a document on EC2 instances, specify the following value:
`/AWS::EC2::Instance`. If you specify a value of '/' the document can run
on all types of resources. If you don't specify a value, the document can't run on any
resources. For a list of valid resource types, see [AWS resource and property types reference](../userguide/aws-template-resource-type-ref.md) in the
_AWS CloudFormation User Guide_.

_Required_: No

_Type_: String

_Pattern_: `^\/[\w\.\-\:\/]*$`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`UpdateMethod`

If the document resource you specify in your template already exists, this parameter
determines whether a new version of the existing document is created, or the existing
document is replaced. `Replace` is the default method. If you specify
`NewVersion` for the `UpdateMethod` parameter, and the
`Name` of the document does not match an existing resource, a new
document is created. When you specify `NewVersion`, the default version of
the document is changed to the newly created version.

_Required_: No

_Type_: String

_Allowed values_: `Replace | NewVersion`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VersionName`

An optional field specifying the version of the artifact you are creating with the document.
For example, `Release12.1`. This value is unique across all versions of a document,
and can't be changed.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-.]{1,128}$`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Systems Manager document name, such as
`MyNewSSMDocument`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

- [Create an Automation runbook that runs commands on an EC2 Linux instance](#aws-resource-ssm-document--examples--Create_an_Automation_runbook_that_runs_commands_on_an_EC2_Linux_instance)

- [Create a document that runs commands on an EC2 Linux instance](#aws-resource-ssm-document--examples--Create_a_document_that_runs_commands_on_an_EC2_Linux_instance)

- [Join a managed instance to a directory in AWS Directory Service](#aws-resource-ssm-document--examples--Join_a_managed_instance_to_a_directory_in)

- [Associate an SSM document with an instance](#aws-resource-ssm-document--examples--Associate_an_SSM_document_with_an_instance)

- [Create a Systems Manager document for Session Manager preferences](#aws-resource-ssm-document--examples--Create_a_document_for_Session_Manager_preferences)

- [Create a Systems Manager document with JSON content](#aws-resource-ssm-document--examples--Create_a_document_with_JSON_content)

- [Create a Systems Manager Distributor package](#aws-resource-ssm-document--examples--Create_a_Distributor_package)

- [Create a Systems Manager Change Calendar document](#aws-resource-ssm-document--examples--Create_a_Change_Calendar_document)

### Create an Automation runbook that runs commands on an EC2 Linux instance

The following runbook runs the commands you specify on your target Amazon EC2 Linux instance. You specify the commands parameter value when
starting the runbook.

#### YAML

```yaml

document:
    Type: AWS::SSM::Document
    Properties:
      Content:
        schemaVersion: '0.3'
        description: 'Run a script on Linux instances.'
        parameters:
          AutomationAssumeRole:
            type: String
            description: (Optional) The ARN of the role that allows Automation to perform the actions on your behalf.
            default: ''
          commands:
            type: String
            description: "(Required) The commands to run or the path to an existing script on the instance."
            default: 'echo Hello World'
          InstanceId:
            type: String
            description: "(Required) The instance ID you want to run commands on."
            default: ''
        mainSteps:
        - name: sayHello
          action: 'aws:runCommand'
          inputs:
            DocumentName: AWS-RunShellScript
            InstanceIds:
              - '{{InstanceId}}'
            Parameters:
              commands:
              - "{{ commands }}"
      DocumentFormat: YAML
      DocumentType: Automation
      Name: 'CFN_runbook_example'
```

#### JSON

```json

"document": {
         "Type": "AWS::SSM::Document",
         "Properties": {
            "Content": {
               "schemaVersion": "0.3",
               "description": "Run a script on Linux instances.",
               "parameters": {
                  "AutomationAssumeRole": {
                     "type": "String",
                     "description": "(Optional) The ARN of the role that allows Automation to perform the actions on your behalf.",
                     "default": ""
                  },
                  "commands": {
                     "type": "String",
                     "description": "(Required) The commands to run or the path to an existing script on the instance.",
                     "default": "echo Hello World"
                  },
                  "InstanceId": {
                     "type": "String",
                     "description": "(Required) The instance ID you want to run commands on.",
                     "default": ""
                  }
               },
               "mainSteps": [
                  {
                     "name": "sayHello",
                     "action": "aws:runCommand",
                     "inputs": {
                        "DocumentName": "AWS-RunShellScript",
                        "Parameters": {
                           "InstanceIds": [
                              "{{InstanceId}}"
                           ],
                           "commands": [
                              "{{ commands }}"
                           ]
                        }
                     }
                  }
               ]
            },
            "DocumentType": "Automation",
            "Name": "CFN_runbook_example"
         }
      }
```

### Create a document that runs commands on an EC2 Linux instance

The following SSM document runs the commands you specify on your target
Amazon EC2 Linux instance. You specify the commands parameter
value when you run the document using Run Command.

#### YAML

```yaml

document:
  Type: AWS::SSM::Document
  Properties:
    Content:
      schemaVersion: '2.2'
      description: 'Run a script on Linux instances.'
      parameters:
        commands:
          type: String
          description: "(Required) The commands to run or the path to an existing script
        on the instance."
          default: 'echo Hello World'
      mainSteps:
      - action: aws:runShellScript
        name: runCommands
        inputs:
          timeoutSeconds: '60'
          runCommand:
          - "{{ commands }}"
    DocumentFormat: YAML
    DocumentType: Command
    Name: 'CFN_2.2_command_example'
```

#### JSON

```json

"document": {
  "Type": "AWS::SSM::Document",
  "Properties": {
    "Content": {
      "schemaVersion": "2.2",
      "description": "Run a script on Linux instances.",
      "parameters": {
        "commands": {
          "type": "String",
          "description": "(Required) The commands to run or the path to an existing script on the instance.",
          "default": "echo Hello World"
        }
      },
      "mainSteps": [
        {
          "action": "aws:runShellScript",
          "name": "runCommands",
          "inputs": {
            "timeoutSeconds": "60",
            "runCommand": [
              "{{ commands }}"
            ]
          }
        }
      ]
    },
    "DocumentType": "Command",
    "Name": "CFN_2.2_command_ex"
  }
}
```

### Join a managed instance to a directory in AWS Directory Service

The following SSM document joins instances to a directory in AWS Directory Service. The three runtime configuration parameters specify which
directory the instance joins. You specify these parameter values when you
associate the document with an instance.

#### YAML

```yaml

document:
  Type: AWS::SSM::Document
  Properties:
    Content:
      schemaVersion: '1.2'
      description: Join instances to an AWS Directory Service domain.
      parameters:
        directoryId:
          type: String
          description: "(Required) The ID of the AWS Directory Service directory."
        directoryName:
          type: String
          description: "(Required) The name of the directory. For example, test.example.com"
        dnsIpAddresses:
          type: StringList
          default: []
          description: "(Optional) The IP addresses of the DNS servers in the directory.
            Required when DHCP is not configured. For more information, see https://docs.aws.amazon.com/directoryservice/latest/admin-guide/simple_ad_dns.html"
          allowedPattern: "((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)"
      runtimeConfig:
        aws:domainJoin:
          properties:
            directoryId: "{{ directoryId}}"
            directoryName: "{{ directoryName }}"
            dnsIpAddresses: "{{ dnsIpAddresses }}"
```

#### JSON

```json

"document" : {
    "Type": "AWS::SSM::Document",
    "Properties": {
        "Content": {
            "schemaVersion": "1.2",
            "description": "Join instances to an AWS Directory Service domain.",
            "parameters": {
                "directoryId": {
                    "type": "String",
                    "description": "(Required) The ID of the AWS Directory Service directory."
                },
                "directoryName": {
                    "type": "String",
                    "description": "(Required) The name of the directory. For example, test.example.com"
                },
                "dnsIpAddresses": {
                    "type": "StringList",
                    "default": [],
                    "description": "(Optional) The IP addresses of the DNS servers in the directory. Required when DHCP is not configured. For more information, see https://docs.aws.amazon.com/directoryservice/latest/admin-guide/simple_ad_dns.html",
                    "allowedPattern": "((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)"
                }
            },
            "runtimeConfig": {
                "aws:domainJoin": {
                    "properties": {
                        "directoryId": "{{ directoryId}}",
                        "directoryName": "{{ directoryName }}",
                        "dnsIpAddresses": "{{ dnsIpAddresses }}"
                    }
                }
            }
        }
    }
}
```

### Associate an SSM document with an instance

The following example shows how to associate an SSM document with an instance.
The `DocumentName` property specifies the SSM document and the
`AssociationParameters` property specifies values for the runtime
configuration parameters.

#### YAML

```yaml

myEC2:
  Type: AWS::EC2::Instance
  Properties:
    ImageId:
      Ref: myImageId
    InstanceType: t2.micro
    SsmAssociations:
    - DocumentName:
        Ref: document
      AssociationParameters:
      - Key: directoryId
        Value:
        - Ref: myDirectory
      - Key: directoryName
        Value:
        - testDirectory.example.com
      - Key: dnsIpAddresses
        Value:
          Fn::GetAtt:
          - myDirectory
          - DnsIpAddresses
    IamInstanceProfile:
      Ref: myInstanceProfile
    NetworkInterfaces:
    - DeviceIndex: '0'
      AssociatePublicIpAddress: 'true'
      SubnetId:
        Ref: mySubnet
    KeyName:
      Ref: myKeyName
```

#### JSON

```json

"myEC2" : {
    "Type": "AWS::EC2::Instance",
    "Properties": {
        "ImageId": {
            "Ref": "myImageId"
        },
        "InstanceType": "t2.micro",
        "SsmAssociations": [
            {
                "DocumentName": {
                    "Ref": "document"
                },
                "AssociationParameters": [
                    {
                        "Key": "directoryId",
                        "Value": [
                            {
                                "Ref": "myDirectory"
                            }
                        ]
                    },
                    {
                        "Key": "directoryName",
                        "Value": [
                            "testDirectory.example.com"
                        ]
                    },
                    {
                        "Key": "dnsIpAddresses",
                        "Value": {
                            "Fn::GetAtt": [
                                "myDirectory",
                                "DnsIpAddresses"
                            ]
                        }
                    }
                ]
            }
        ],
        "IamInstanceProfile": {
            "Ref": "myInstanceProfile"
        },
        "NetworkInterfaces": [
            {
                "DeviceIndex": "0",
                "AssociatePublicIpAddress": "true",
                "SubnetId": {
                    "Ref": "mySubnet"
                }
            }
        ],
        "KeyName": {
            "Ref": "myKeyName"
        }
    }
}
```

### Create a Systems Manager document for Session Manager preferences

The following example creates a Systems Manager `Session` type
document for Session Manager preferences. Before using this example template
replace the placeholder values.

#### JSON

```json

{
   "Resources":{
      "SessionPreferencesDocument":{
         "Type":"AWS::SSM::Document",
         "Properties":{
            "Name":"SSM-SessionManagerRunShell",
            "Content":{
               "schemaVersion":"1.0",
               "description":"Document to hold regional settings for Session Manager",
               "sessionType":"Standard_Stream",
               "inputs":{
                  "s3BucketName":"amzn-s3-demo-bucket",
                  "s3KeyPrefix":"amzn-s3-demo-bucket-prefix",
                  "s3EncryptionEnabled":true,
                  "cloudWatchLogGroupName":"MyLogGroupName",
                  "cloudWatchEncryptionEnabled":true,
                  "cloudWatchStreamingEnabled":false,
                  "kmsKeyId":"331dba25-425c-446c-abf9-daf42EXAMPLE",
                  "runAsEnabled":false,
                  "runAsDefaultUser":"MyDefaultRunAsUser",
                  "idleSessionTimeout":"20",
                  "shellProfile":{
                     "windows":"my-windowscommands",
                     "linux":"my-linux-commands"
                  }
               }
            },
            "DocumentType":"Session"
         }
      }
   },
   "Outputs":{
      "DocumentName":{
         "Description":"Session Manager preferences document",
         "Value":"SSM-SessionManagerRunShell"
      }
   }
}
```

#### YAML

```yaml

Resources:
  SessionPreferencesDocument:
    Type: AWS::SSM::Document
    Properties:
      Name: SSM-SessionManagerRunShell
      Content:
        schemaVersion: '1.0'
        description: Document to hold regional settings for Session Manager
        sessionType: Standard_Stream
        inputs:
          s3BucketName: 'amzn-s3-demo-bucket'
          s3KeyPrefix: 'amzn-s3-demo-bucket-prefix'
          s3EncryptionEnabled: true
          cloudWatchLogGroupName: 'MyLogGroupName'
          cloudWatchEncryptionEnabled: true
          cloudWatchStreamingEnabled: false
          kmsKeyId: '"331dba25-425c-446c-abf9-daf42EXAMPLE"'
          runAsEnabled: false
          runAsDefaultUser: 'MyDefaultRunAsUser'
          idleSessionTimeout: '20'
          shellProfile:
            windows: my-windows-commands
            linux: my-linux-commands
      DocumentFormat: YAML
      DocumentType: Session
Outputs:
  DocumentName:
    Description: "Session Manager preferences document"
    Value: SSM-SessionManagerRunShell
```

### Create a Systems Manager document with JSON content

The following example creates a new Systems Manager command document with
JSON content.

#### JSON

```json

{
   "Type":"AWS::SSM::Document",
   "Properties":{
      "Content":"{\"schemaVersion\": \"2.2\",  \"description\": \"Command Document Example JSON\nTemplate\",  \"parameters\": {    \"Message\": {      \"type\": \"String\", \"description\":\n\"Example\",      \"default\": \"Hello World\"    }  },  \"mainSteps\": [    { \"action\":\n\"aws:runPowerShellScript\",      \"name\": \"example\",      \"inputs\": {        \"runCommand\":\n[ \"Write-Output {{Message}}\" ]      }    }  ]}",
      "DocumentType":"Command",
      "DocumentFormat":"JSON"
   }
}
```

#### YAML

```yaml

---
Type: "AWS::SSM::Document"
Properties:
  Content: "{\"schemaVersion\": \"2.2\",  \"description\": \"Command Document Example JSON Template\",  \"parameters\": {    \"Message\": {      \"type\": \"String\", \"description\": \"Example\",      \"default\": \"Hello World\"    }  },  \"mainSteps\": [    { \"action\": \"aws:runPowerShellScript\",      \"name\": \"example\",      \"inputs\": {        \"runCommand\": [ \"Write-Output {{Message}}\" ]      }    }  ]}"
  DocumentFormat: JSON
  DocumentType: Command
```

### Create a Systems Manager Distributor package

The following example creates a new Systems Manager Distributor
package.

#### JSON

```json

{
        "Resources": {
        "ExamplePackageDocument": {
            "Type": "AWS::SSM::Document",
            "Properties": {
                "Content": "{\"files\": {\"NewPackage_WINDOWS.zip\": {\"checksums\": {\"sha256\": \"36aeb0ec2c706013cf8c68163459678f7f6daa9489cd3f91d52799331EXAMPLE\"}}}, \"publisher\": \"MyPublisherName\", \"schemaVersion\": \"2.0\", \"packages\": {\"_any\": {\"_any\": {\"x86_64\": {\"file\": \"MyNewPackage_WINDOWS.zip\"}}}}, \"version\": \"1.0\"}",
                "DocumentType": "Package",
                "Attachments": [
                    {
                        "Key": "SourceUrl",
                        "Values": [
                            "s3://example-package-path/valid-package"
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

Resources:
  ExamplePackageDocument:
    Type: AWS::SSM::Document
    Properties:
      Content: '{\"files\": {\"MyNewPackage_WINDOWS.zip\": {\"checksums\":
        {\"sha256\": \"36aeb0ec2c706013cf8c68163459678f7f6daa9489cd3f91d52799331EXAMPLE\"}}},
        \"publisher\": \"MyPublisherName\", \"schemaVersion\":
        \"2.0\", \"packages\": {\"_any\": {\"_any\": {\"x86_64\": {\"file\": \"MyNewPackage_WINDOWS.zip\"}}}},
        \"version\": \"1.0\"}'
      DocumentType: Package
      Attachments:
      - Key: SourceUrl
        Values:
        - "s3://example-package-path/valid-package"
```

### Create a Systems Manager Change Calendar document

The following example creates a new Systems Manager Change Calendar
document.

#### JSON

```json

{
   "Resources":{
      "ExampleChangeCalendarDocument":{
         "Type":"AWS::SSM::Document",
         "Properties":{
            "Content":"BEGIN:VCALENDAR\r\nPRODID:-//AWS//Change Calendar 1.0//EN\r\nVERSION:2.0\r\nX-CALENDAR-TYPE:DEFAULT_OPEN\r\nX-WR-CALDESC:test\r\nBEGIN:VTODO\r\nDTSTAMP:20200320T004207Z\r\nUID:3b5af39a-d0b3-4049-a839-d7bb8af01f92\r\nSUMMARY:Add events to this calendar.\r\nEND:VTODO\r\nEND:VCALENDAR\r\n",
            "DocumentType":"ChangeCalendar",
            "DocumentFormat":"TEXT"
         }
      }
   }
}
```

#### YAML

```yaml

Resources:
  ExampleChangeCalendarDocument:
    Type: 'AWS::SSM::Document'
    Properties:
      Content: "BEGIN:VCALENDAR\r\nPRODID:-//AWS//Change Calendar 1.0//EN\r\nVERSION:2.0\r\nX-CALENDAR-TYPE:DEFAULT_OPEN\r\nX-WR-CALDESC:test\r\nBEGIN:VTODO\r\nDTSTAMP:20200320T004207Z\r\nUID:3b5af39a-d0b3-4049-a839-d7bb8af01f92\r\nSUMMARY:Add events to this calendar.\r\nEND:VTODO\r\nEND:VCALENDAR\r\n"
      DocumentType: ChangeCalendar
      DocumentFormat: TEXT
```

## See also

- [AWS Systems Manager Documents](../../../systems-manager/latest/userguide/sysman-ssm-docs.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Target

AttachmentsSource

All content copied from https://docs.aws.amazon.com/.
