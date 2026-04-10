This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::Association

The `AWS::SSM::Association` resource creates a State Manager association
for your managed instances. A State Manager association defines the state that you want
to maintain on your instances. For example, an association can specify that anti-virus
software must be installed and running on your instances, or that certain ports must be
closed. For static targets, the association specifies a schedule for when the
configuration is reapplied. For dynamic targets, such as an AWS Resource Groups or an
AWS Auto Scaling Group, State Manager applies the configuration when new
instances are added to the group. The association also specifies actions to take when
applying the configuration. For example, an association for anti-virus software might
run once a day. If the software is not installed, then State Manager installs it. If the
software is installed, but the service is not running, then the association might
instruct State Manager to start the service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SSM::Association",
  "Properties" : {
      "ApplyOnlyAtCronInterval" : Boolean,
      "AssociationName" : String,
      "AutomationTargetParameterName" : String,
      "CalendarNames" : [ String, ... ],
      "ComplianceSeverity" : String,
      "DocumentVersion" : String,
      "InstanceId" : String,
      "MaxConcurrency" : String,
      "MaxErrors" : String,
      "Name" : String,
      "OutputLocation" : InstanceAssociationOutputLocation,
      "Parameters" : [ String, ... ],
      "ScheduleExpression" : String,
      "ScheduleOffset" : Integer,
      "SyncCompliance" : String,
      "Targets" : [ Target, ... ],
      "WaitForSuccessTimeoutSeconds" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::SSM::Association
Properties:
  ApplyOnlyAtCronInterval: Boolean
  AssociationName: String
  AutomationTargetParameterName: String
  CalendarNames:
    - String
  ComplianceSeverity: String
  DocumentVersion: String
  InstanceId: String
  MaxConcurrency: String
  MaxErrors: String
  Name: String
  OutputLocation:
    InstanceAssociationOutputLocation
  Parameters:
    - String
  ScheduleExpression: String
  ScheduleOffset: Integer
  SyncCompliance: String
  Targets:
    - Target
  WaitForSuccessTimeoutSeconds: Integer

```

## Properties

`ApplyOnlyAtCronInterval`

By default, when you create a new association, the system runs it immediately after it
is created and then according to the schedule you specified. Specify this option if you
don't want an association to run immediately after you create it. This parameter is not
supported for rate expressions.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssociationName`

Specify a descriptive name for the association.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-.]{3,128}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutomationTargetParameterName`

Choose the parameter that will define how your automation will branch out. This target
is required for associations that use an Automation runbook and target resources by
using rate controls. Automation is a tool in AWS Systems Manager.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CalendarNames`

The names or Amazon Resource Names (ARNs) of the Change Calendar type documents your
associations are gated under. The associations only run when that Change Calendar is open. For
more information, see [AWS Systems Manager Change\
Calendar](../../../systems-manager/latest/userguide/systems-manager-change-calendar.md) in the _AWS Systems Manager User Guide_.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComplianceSeverity`

The severity level that is assigned to the association.

_Required_: No

_Type_: String

_Allowed values_: `CRITICAL | HIGH | MEDIUM | LOW | UNSPECIFIED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentVersion`

The version of the SSM document to associate with the target.

###### Important

Note the following important information.

- State Manager doesn't support running associations that use a new version
of a document if that document is shared from another account. State Manager
always runs the `default` version of a document if shared from
another account, even though the Systems Manager console shows that a
new version was processed. If you want to run an association using a new
version of a document shared form another account, you must set the document
version to `default`.

- `DocumentVersion` is not valid for documents owned by AWS, such as `AWS-RunPatchBaseline` or
`AWS-UpdateSSMAgent`. If you specify
`DocumentVersion` for an AWS document, the
system returns the following error: "Error occurred during operation
'CreateAssociation'." (RequestToken: <token>, HandlerErrorCode:
GeneralServiceException).

_Required_: No

_Type_: String

_Pattern_: `([$]LATEST|[$]DEFAULT|^[1-9][0-9]*$)`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceId`

###### Note

`InstanceId` has been deprecated. To specify an instance ID for an
association, use the `Targets` parameter. If you use the parameter
`InstanceId`, you cannot use the parameters
`AssociationName`, `DocumentVersion`,
`MaxErrors`, `MaxConcurrency`,
`OutputLocation`, or `ScheduleExpression`. To use these
parameters, you must use the `Targets` parameter.

Note that in some examples later in this page, `InstanceIds` is used as
the tag-key name in a `Targets` filter. `InstanceId` is not
used as a parameter.

The ID of the instance that the SSM document is associated with. You must specify the
`InstanceId` or `Targets` property.

_Required_: Conditional

_Type_: String

_Pattern_: `(^i-(\w{8}|\w{17})$)|(^mi-\w{17}$)`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxConcurrency`

The maximum number of targets allowed to run the association at the same time. You can
specify a number, for example 10, or a percentage of the target set, for example 10%. The default
value is 100%, which means all targets run the association at the same time.

If a new managed node starts and attempts to run an association while Systems Manager is running
`MaxConcurrency` associations, the association is allowed to run. During the next
association interval, the new managed node will process its association within the limit
specified for `MaxConcurrency`.

_Required_: No

_Type_: String

_Pattern_: `^([1-9][0-9]{0,6}|[1-9][0-9]%|[1-9]%|100%)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxErrors`

The number of errors that are allowed before the system stops sending requests to run the
association on additional targets. You can specify either an absolute number of errors, for
example 10, or a percentage of the target set, for example 10%. If you specify 3, for example,
the system stops sending requests when the fourth error is received. If you specify 0, then the
system stops sending requests after the first error is returned. If you run an association on 50
managed nodes and set `MaxError` to 10%, then the system stops sending the request
when the sixth error is received.

Executions that are already running an association when `MaxErrors` is reached
are allowed to complete, but some of these executions may fail as well. If you need to ensure
that there won't be more than max-errors failed executions, set `MaxConcurrency` to 1
so that executions proceed one at a time.

_Required_: No

_Type_: String

_Pattern_: `^([1-9][0-9]{0,6}|[0]|[1-9][0-9]%|[0-9]%|100%)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the SSM document that contains the configuration information for the
instance. You can specify `Command` or `Automation` documents. The
documents can be AWS-predefined documents, documents you created, or a
document that is shared with you from another account. For SSM documents that are shared
with you from other AWS accounts, you must specify the complete SSM
document ARN, in the following format:

`arn:partition:ssm:region:account-id:document/document-name`

For example:
`arn:aws:ssm:us-east-2:12345678912:document/My-Shared-Document`

For AWS-predefined documents and SSM documents you created in your
account, you only need to specify the document name. For example, `AWS-ApplyPatchBaseline` or `My-Document`.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-.:/]{3,200}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputLocation`

An Amazon Simple Storage Service (Amazon S3) bucket where you want to store the output
details of the request.

_Required_: No

_Type_: [InstanceAssociationOutputLocation](aws-properties-ssm-association-instanceassociationoutputlocation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

The parameters for the runtime configuration of the document.

_Required_: No

_Type_: Array of String

_Pattern_: `.{1,255}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleExpression`

A cron expression that specifies a schedule when the association runs. The schedule runs in
Coordinated Universal Time (UTC).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleOffset`

Number of days to wait after the scheduled day to run an association.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `6`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SyncCompliance`

The mode for generating association compliance. You can specify `AUTO` or
`MANUAL`. In `AUTO` mode, the system uses the status of the
association execution to determine the compliance status. If the association execution
runs successfully, then the association is `COMPLIANT`. If the association
execution doesn't run successfully, the association is
`NON-COMPLIANT`.

In `MANUAL` mode, you must specify the `AssociationId` as a
parameter for the `PutComplianceItems` API action. In this case, compliance
data is not managed by State Manager. It is managed by your direct call to the
`PutComplianceItems` API action.

By default, all associations use `AUTO` mode.

_Required_: No

_Type_: String

_Allowed values_: `AUTO | MANUAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Targets`

The targets for the association. You must specify the `InstanceId` or
`Targets` property. You can target all instances in an AWS account by specifying t he `InstanceIds` key with a value of
`*`.

Supported formats include the following.

- `Key=InstanceIds,Values=<instance-id-1>,<instance-id-2>,<instance-id-3>`

- `Key=tag-key,Values=<my-tag-key-1>,<my-tag-key-2>`

To view a JSON and a YAML example that targets all instances, see "Create an
association for all managed instances in an AWS account" on the
Examples page.

_Required_: Conditional

_Type_: Array of [Target](aws-properties-ssm-association-target.md)

_Minimum_: `0`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WaitForSuccessTimeoutSeconds`

The number of seconds the service should wait for the association status to show
"Success" before proceeding with the stack execution. If the association status doesn't
show "Success" after the specified number of seconds, then stack creation fails.

###### Important

When you specify a value for the `WaitForSuccessTimeoutSeconds`, [drift\
detection](../userguide/using-cfn-stack-drift.md) for your CloudFormation stack’s configuration might yield
inaccurate results. If drift detection is important in your scenario, we recommend
that you don’t include `WaitForSuccessTimeoutSeconds` in your
template.

_Required_: No

_Type_: Integer

_Minimum_: `15`

_Maximum_: `172800`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AssociationId`

The association ID.

## Examples

- [Create an association for a specific instance](#aws-resource-ssm-association--examples--Create_an_association_for_a_specific_instance)

- [Create an association for all managed instances in an AWS account](#aws-resource-ssm-association--examples--Create_an_association_for_all_managed_instances_in_an)

- [Create an association for a specific tag](#aws-resource-ssm-association--examples--Create_an_association_for_a_specific_tag)

- [Create an association that associates an automation document with an instance](#aws-resource-ssm-association--examples--Create_an_association_that_associates_an_automation_document_with_an_instance)

- [Create an association that uses rate controls and sends log output to Amazon S3](#aws-resource-ssm-association--examples--Create_an_association_that_uses_rate_controls_and_sends_log_output_to)

- [Create an association that works with Ansible](#aws-resource-ssm-association--examples--Create_an_association_that_works_with_Ansible)

- [Create an association that runs a bash script](#aws-resource-ssm-association--examples--Create_an_association_that_runs_a_bash_script)

- [Create an association that runs a bash script with Systems Manager Automation](#aws-resource-ssm-association--examples--Create_an_association_that_runs_a_bash_script_with_Automation)

### Create an association for a specific instance

The following example creates an association that uses the AWS-RunShellScript SSM document. The association runs a simple command on a
specific instance.

#### JSON

```json

{
    "Resources": {
        "SpecificInstanceIdAssociation": {
            "Type": "AWS::SSM::Association",
            "Properties": {
                "Name": "AWS-RunShellScript",
                "Targets": [
                    {
                        "Key": "InstanceIds",
                        "Values": [
                            "i-1234567890abcdef0"
                        ]
                    }
                ],
                "Parameters": {
                    "commands": [
                        "ls"
                    ],
                    "workingDirectory": [
                        "/"
                    ]
                }
            }
        }
    }
}
```

#### YAML

```yaml

---
Resources:
  SpecificInstanceIdAssociation:
    Type: AWS::SSM::Association
    Properties:
      Name: AWS-RunShellScript
      Targets:
      - Key: InstanceIds
        Values:
        - i-1234567890abcdef0
      Parameters:
        commands:
        - ls
        workingDirectory:
        - "/"
```

### Create an association for all managed instances in an AWS account

The following example creates an association that uses the AWS-UpdateSSMAgent SSM document. The association updates SSM Agent on all managed
instances (instances configured for Systems Manager) in the user's AWS account according to the specified CRON schedule.

#### JSON

```json

{
    "Resources": {
        "AllInstanceIdsAssociation": {
            "Type": "AWS::SSM::Association",
            "Properties": {
                "AssociationName": "UpdateSSMAgent",
                "Name": "AWS-UpdateSSMAgent",
                "ScheduleExpression": "cron(0 2 ? * SUN *)",
                "Targets": [
                    {
                        "Key": "InstanceIds",
                        "Values": [
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

---
Resources:
  AllInstanceIdsAssociation:
    Type: AWS::SSM::Association
    Properties:
      AssociationName: UpdateSSMAgent
      Name: AWS-UpdateSSMAgent
      ScheduleExpression: cron(0 2 ? * SUN *)
      Targets:
      - Key: InstanceIds
        Values:
        - "*"
```

### Create an association for a specific tag

The following example creates an association that uses the AWS-UpdateSSMAgent SSM document. The association updates SSM Agent on all managed
instances that are assigned a tag key of `Environment` and value of
`Production`. The association runs every seven days according to
the specified rate expression.

#### JSON

```json

{
    "Resources": {
        "TaggedInstancesAssociation": {
            "Type": "AWS::SSM::Association",
            "Properties": {
                "AssociationName": "UpdateSSMAgent",
                "Name": "AWS-UpdateSSMAgent",
                "ScheduleExpression": "rate(7 days)",
                "Targets": [
                    {
                        "Key": "tag:Environment",
                        "Values": [
                            "Production"
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

---
Resources:
  TaggedInstancesAssociation:
    Type: AWS::SSM::Association
    Properties:
      AssociationName: UpdateSSMAgent
      Name: AWS-UpdateSSMAgent
      ScheduleExpression: rate(7 days)
      Targets:
      - Key: tag:Environment
        Values:
        - Production
```

### Create an association that associates an automation document with an instance

The following example creates an association that assigns the AWS-StopEC2Instance automation document to a specific instance.

###### Note

This example specifies the following Amazon Resource Name (ARN):
`arn:${AWS::Partition}:iam::aws:policy/AmazonEC2FullAccess`.
This policy provides more than the required permissions to stop the
instance. We recommend that you use a policy with more restrictive
permissions.

#### YAML

```yaml

---
Resources:
  AutomationExecutionRole:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: 2012-10-17
        Statement:
          - Effect: Allow
            Principal:
              Service: ssm.amazonaws.com
            Action:
              - sts:AssumeRole
      Path: /
      ManagedPolicyArns:
        - !Sub arn:${AWS::Partition}:iam::aws:policy/AmazonEC2FullAccess
  AutomationAssociation:
    Type: AWS::SSM::Association
    Properties:
      Name: AWS-StopEC2Instance
      Parameters:
        AutomationAssumeRole:
          - !GetAtt AutomationExecutionRole.Arn
      Targets:
        - Key: MyInstanceId
          Values:
            - i-1234567890abcdef0
      AutomationTargetParameterName: InstanceId
```

### Create an association that uses rate controls and sends log output to Amazon S3

The following example creates an association that uses rate controls. The
association attempts to update SSM Agent on only 20% of instances at one time.
Systems Manager stops the association from running on any additional
instances if the execution fails on 5% of the total number of instances. Systems Manager also logs the association output to Amazon S3.

#### JSON

```json

{
    "Resources": {
        "RateControlAssociation": {
            "Type": "AWS::SSM::Association",
            "Properties": {
                "Name": "AWS-UpdateSSMAgent",
                "Targets": [
                    {
                        "Key": "InstanceIds",
                        "Values": [
                            "*"
                        ]
                    }
                ],
                "MaxConcurrency": "20%",
                "MaxErrors": "5%"
            },
            "OutputLocation": {
                "S3Location": {
                    "OutputS3BucketName": "amzn-s3-demo-bucket",
                    "OutputS3KeyPrefix": "my-agent-update-output"
                }
            }
        }
    }
}
```

#### YAML

```yaml

---
Resources:
  RateControlAssociation:
    Type: 'AWS::SSM::Association'
    Properties:
      Name: AWS-UpdateSSMAgent
      Targets:
        - Key: InstanceIds
          Values:
            - '*'
      MaxConcurrency: 20%
      MaxErrors: 5%
OutputLocation:
  S3Location:
    OutputS3BucketName: amzn-s3-demo-bucket
    OutputS3KeyPrefix: my-agent-update-output

```

### Create an association that works with Ansible

The following example creates an association that uses Ansible and Systems Manager to deploy Nginx. This template copies the Ansible Playbook from
a Github repo. The target is based on instance ID.

#### YAML

```yaml

---
Description: "Deploy single Amazon Linux 2 EC2 instance"
Parameters:
  # Using SSM Parameter Store to fetch the Latest AMI for Amazon Linux 2. Eliminates the need for AMI Mappings.
  LatestAmiId:
    Type: 'AWS::SSM::Parameter::Value<AWS::EC2::Image::Id>'
    Default: "/aws/service/ami-amazon-linux-latest/amzn2-ami-hvm-x86_64-gp2"
  GitHubOwner:
    Type: 'String'
  GitHubRepo:
    Type: 'String'
  GitHubBranch:
    Type: 'String'
Resources:
  amzn-s3-demo-bucket:
    Type: AWS::S3::Bucket
  SSMInstanceRole:
    Type : AWS::IAM::Role
    Properties:
      Policies:
        - PolicyDocument:
            Version: '2012-10-17'
            Statement:
              - Action:
                  - s3:GetObject
                Resource:
                  - !Sub 'arn:aws:s3:::aws-ssm-${AWS::Region}/*'
                  - !Sub 'arn:aws:s3:::aws-windows-downloads-${AWS::Region}/*'
                  - !Sub 'arn:aws:s3:::amazon-ssm-${AWS::Region}/*'
                  - !Sub 'arn:aws:s3:::amazon-ssm-packages-${AWS::Region}/*'
                  - !Sub 'arn:aws:s3:::${AWS::Region}-birdwatcher-prod/*'
                  - !Sub 'arn:aws:s3:::patch-baseline-snapshot-${AWS::Region}/*'
                Effect: Allow
          PolicyName: ssm-custom-s3-policy
        - PolicyDocument:
            Version: '2012-10-17'
            Statement:
              - Action:
                  - s3:GetObject
                  - s3:PutObject
                  - s3:PutObjectAcl
                  - s3:ListBucket
                Resource:
                  - !Sub 'arn:${AWS::Partition}:s3:::amzn-s3-demo-bucket/*'
                  - !Sub 'arn:${AWS::Partition}:s3:::amzn-s3-demo-bucket'
                Effect: Allow
          PolicyName: s3-instance-bucket-policy
      Path: /
      ManagedPolicyArns:
        - !Sub 'arn:${AWS::Partition}:iam::aws:policy/AmazonSSMManagedInstanceCore'
      AssumeRolePolicyDocument:
        Version: "2012-10-17"
        Statement:
        - Effect: "Allow"
          Principal:
            Service:
            - "ec2.amazonaws.com"
            - "ssm.amazonaws.com"
          Action: "sts:AssumeRole"
  SSMInstanceProfile:
    Type: "AWS::IAM::InstanceProfile"
    Properties:
      Roles:
      - !Ref SSMInstanceRole
  EC2Instance:
    Type: "AWS::EC2::Instance"
    Properties:
      ImageId: !Ref LatestAmiId
      InstanceType: "t3.small"
      IamInstanceProfile: !Ref SSMInstanceProfile
  AnsibleAssociation:
    Type: AWS::SSM::Association
    Properties:
      # Here using the AWS-ApplyAnsiblePlaybooks
      Name: AWS-ApplyAnsiblePlaybooks
      # Targeting Instance by InstanceId passed from the Logical ID of Instance being created
      # in CloudFormation
      Targets:
        - Key: InstanceIds
          Values: [ !Ref EC2Instance ]
      OutputLocation:
        S3Location:
          OutputS3BucketName: !Ref amzn-s3-demo-bucket
          OutputS3KeyPrefix: 'logs/'
      Parameters:
        # Getting an Ansible Playbook from a GitHub Location
        SourceType:
          - 'GitHub'
        # At a minimum must include the following GitHub repo information, if using a private repo
        # would want to include the GitHub Token option
        SourceInfo:
          -  !Sub |
              {"owner":"${GitHubOwner}",
              "repository":"${GitHubRepo}",
              "path":"",
              "getOptions":"branch:${GitHubBranch}"}
        # Installing Ansible and its dependencies
        InstallDependencies:
          - 'True'
        # Playbook file we want to run
        PlaybookFile:
          - 'playbook.yml'
        ExtraVariables:
          - 'SSM=True'
        Check:
          - 'False'
        Verbose:
          - '-v'
Outputs:
  WebServerPublic:
    Value: !GetAtt 'EC2Instance.PublicDnsName'
    Description: Public DNS for WebServer
```

### Create an association that runs a bash script

The following example creates an association that runs a bash script. Target
is based on tags.

#### YAML

```yaml

---
Description: "Deploy single Amazon Linux 2 EC2 instance and install Nginx by a State Manager association"
Parameters:
  LatestAmiId:
    Type: 'AWS::SSM::Parameter::Value*<AWS::EC2::Image::Id>'
    Default: "/aws/service/ami-amazon-linux-latest/amzn2-ami-hvm-x86_64-gp2"
Resources:
  amzn-s3-demo-bucket:
    Type: AWS::S3::Bucket
  # Role that allows SSM Agent to communicate with SSM and allows use of all features of SSM
  SSMInstanceRole:
    Type : AWS::IAM::Role
    Properties:
      Policies:
        - PolicyDocument:
            Version: '2012-10-17'
            Statement:
              - Action:
                  - s3:GetObject
                Resource:
                  - !Sub 'arn:aws:s3:::aws-ssm-${AWS::Region}/*'
                  - !Sub 'arn:aws:s3:::aws-windows-downloads-${AWS::Region}/*'
                  - !Sub 'arn:aws:s3:::amazon-ssm-${AWS::Region}/*'
                  - !Sub 'arn:aws:s3:::amazon-ssm-packages-${AWS::Region}/*'
                  - !Sub 'arn:aws:s3:::${AWS::Region}-birdwatcher-prod/*'
                  - !Sub 'arn:aws:s3:::patch-baseline-snapshot-${AWS::Region}/*'
                Effect: Allow
          PolicyName: ssm-custom-s3-policy
        - PolicyDocument:
            Version: '2012-10-17'
            Statement:
              - Action:
                  - s3:GetObject
                  - s3:PutObject
                  - s3:PutObjectAcl
                  - s3:ListBucket
                Resource:
                  - !Sub 'arn:${AWS::Partition}:s3:::amzn-s3-demo-bucket/*'
                  - !Sub 'arn:${AWS::Partition}:s3:::amzn-s3-demo-bucket'
                Effect: Allow
          PolicyName: s3-instance-bucket-policy
      Path: /
      ManagedPolicyArns:
        - !Sub 'arn:${AWS::Partition}:iam::aws:policy/AmazonSSMManagedInstanceCore'
        - !Sub 'arn:${AWS::Partition}:iam::aws:policy/CloudWatchAgentServerPolicy'
      AssumeRolePolicyDocument:
        Version: "2012-10-17"
        Statement:
        - Effect: "Allow"
          Principal:
            Service:
            - "ec2.amazonaws.com"
            - "ssm.amazonaws.com"
          Action: "sts:AssumeRole"
  SSMInstanceProfile:
    Type: "AWS::IAM::InstanceProfile"
    Properties:
      Roles:
      - !Ref SSMInstanceRole
  EC2Instance:
    Type: "AWS::EC2::Instance"
    Properties:
      ImageId: !Ref LatestAmiId
      InstanceType: "t3.medium"
      IamInstanceProfile: !Ref SSMInstanceProfile
      Tags:
      - Key: 'nginx'
        Value: 'yes'
  NginxAssociation:
    DependsOn: EC2Instance
    # CloudFormation Resource Type that creates State Manager Associations
    Type: AWS::SSM::Association
    Properties:
      # Command Document that this Association will run
      Name: AWS-RunShellScript
      # Targeting Instance by Tags
      Targets:
        - Key: tag:nginx
          Values:
            - 'yes'
      # The passing in the S3 Bucket that is created in the template that logs will be sent to
      OutputLocation:
        S3Location:
          OutputS3BucketName: !Ref amzn-s3-demo-bucket
          OutputS3KeyPrefix: 'logs/'
      # Parameters for the AWS-RunShellScript, in this case commands to install nginx
      Parameters:
        commands:
          - |
              sudo amazon-linux-extras install nginx1 -y
              sudo service nginx start
Outputs:
  WebServerPublic:
    Value: !GetAtt 'EC2Instance.PublicDnsName'
    Description: Public DNS for WebServer
```

### Create an association that runs a bash script with Systems Manager Automation

The following example creates an assocation that runs a bash script using
State Manager and Automation with multiple steps. Target is based on
tags.

#### YAML

```yaml

---
Description: "Deploy single Amazon Linux 2 EC2 instance"
Parameters:
  LatestAmiId:
    Type: 'AWS::SSM::Parameter::Value<AWS::EC2::Image::Id>'
    Default: "/aws/service/ami-amazon-linux-latest/amzn2-ami-hvm-x86_64-gp2"
Resources:
  amzn-s3-demo-bucket:
    Type: AWS::S3::Bucket
  nginxInstallAutomation:
    Type: AWS::SSM::Document
    Properties:
      DocumentType: Automation
      Content:
        schemaVersion: "0.3"
        description: "Updates AMI with Linux distribution packages and installs Nginx software"
        assumeRole: "{{AutomationAssumeRole}}"
        parameters:
          InstanceId:
            description: "ID of the Instance."
            type: "String"
          AutomationAssumeRole:
            default: ""
            description: "(Optional) The ARN of the role that allows Automation to perform the actions on your behalf."
            type: "String"
        mainSteps:
        - name: "updateOSSoftware"
          action: "aws:runCommand"
          maxAttempts: 3
          timeoutSeconds: 3600
          inputs:
            DocumentName: "AWS-RunShellScript"
            InstanceIds:
            - "{{InstanceId}}"
            CloudWatchOutputConfig:
              CloudWatchOutputEnabled: "true"
            Parameters:
              commands:
                - |
                   #!/bin/bash
                   sudo yum update -y
                   needs-restarting -r
                   if [ $? -eq 1 ]
                   then
                           exit 194
                   else
                           exit 0
                   fi
        - name: "InstallNginx"
          action: "aws:runCommand"
          inputs:
            DocumentName: "AWS-RunShellScript"
            InstanceIds:
            - "{{InstanceId}}"
            CloudWatchOutputConfig:
              CloudWatchOutputEnabled: "true"
            Parameters:
              commands:
                - |
                    sudo amazon-linux-extras install nginx1 -y
                    sudo service nginx start
        - name: "TestInstall"
          action: "aws:runCommand"
          maxAttempts: 3
          timeoutSeconds: 3600
          onFailure: "Abort"
          inputs:
           DocumentName: "AWS-RunShellScript"
           InstanceIds:
            - "{{InstanceId}}"
           Parameters:
            commands:
                - |
                   curl localhost
  SSMExecutionRole:
    Type : AWS::IAM::Role
    Properties:
      Policies:
        - PolicyDocument:
            Version: '2012-10-17'
            Statement:
              - Action:
                  - ssm:StartAssociationsOnce
                  - ssm:CreateAssociation
                  - ssm:CreateAssociationBatch
                  - ssm:UpdateAssociation
                Resource: '*'
                Effect: Allow
          PolicyName: ssm-association
      Path: /
      ManagedPolicyArns:
        - !Sub 'arn:${AWS::Partition}:iam::aws:policy/service-role/AmazonSSMAutomationRole'
      AssumeRolePolicyDocument:
        Version: "2012-10-17"
        Statement:
        - Effect: "Allow"
          Principal:
            Service:
            - "ec2.amazonaws.com"
            - "ssm.amazonaws.com"
          Action: "sts:AssumeRole"
  SSMInstanceRole:
    Type : AWS::IAM::Role
    Properties:
      Policies:
        - PolicyDocument:
            Version: '2012-10-17'
            Statement:
              - Action:
                  - s3:GetObject
                Resource:
                  - !Sub 'arn:aws:s3:::aws-ssm-${AWS::Region}/*'
                  - !Sub 'arn:aws:s3:::aws-windows-downloads-${AWS::Region}/*'
                  - !Sub 'arn:aws:s3:::amazon-ssm-${AWS::Region}/*'
                  - !Sub 'arn:aws:s3:::amazon-ssm-packages-${AWS::Region}/*'
                  - !Sub 'arn:aws:s3:::${AWS::Region}-birdwatcher-prod/*'
                  - !Sub 'arn:aws:s3:::patch-baseline-snapshot-${AWS::Region}/*'
                Effect: Allow
          PolicyName: ssm-custom-s3-policy
        - PolicyDocument:
            Version: '2012-10-17'
            Statement:
              - Action:
                  - s3:GetObject
                  - s3:PutObject
                  - s3:PutObjectAcl
                  - s3:ListBucket
                Resource:
                  - !Sub 'arn:${AWS::Partition}:s3:::amzn-s3-demo-bucket/*'
                  - !Sub 'arn:${AWS::Partition}:s3:::amzn-s3-demo-bucket'
                Effect: Allow
          PolicyName: s3-instance-bucket-policy
      Path: /
      ManagedPolicyArns:
        - !Sub 'arn:${AWS::Partition}:iam::aws:policy/AmazonSSMManagedInstanceCore'
        - !Sub 'arn:${AWS::Partition}:iam::aws:policy/CloudWatchAgentServerPolicy'
      AssumeRolePolicyDocument:
        Version: "2012-10-17"
        Statement:
        - Effect: "Allow"
          Principal:
            Service:
            - "ec2.amazonaws.com"
            - "ssm.amazonaws.com"
          Action: "sts:AssumeRole"
  SSMInstanceProfile:
    Type: "AWS::IAM::InstanceProfile"
    Properties:
      Roles:
      - !Ref SSMInstanceRole
  EC2Instance:
    Type: "AWS::EC2::Instance"
    Properties:
      ImageId: !Ref LatestAmiId
      InstanceType: "t3.medium"
      IamInstanceProfile: !Ref SSMInstanceProfile
      Tags:
       - Key: nginx
         Value: Yes
  NginxAssociation:
    DependsOn: EC2Instance
    Type: AWS::SSM::Association
    Properties:
      Name: !Ref nginxInstallAutomation
      OutputLocation:
        S3Location:
          OutputS3BucketName: !Ref amzn-s3-demo-bucket
          OutputS3KeyPrefix: 'logs/'
      AutomationTargetParameterName: InstanceId
      Parameters:
        AutomationAssumeRole:
          - !GetAtt 'SSMExecutionRole.Arn'
      Targets:
        - Key: tag:nginx
          Values:
            - Yes
Outputs:
  WebServerPublic:
    Value: !GetAtt 'EC2Instance.PublicDnsName'
    Description: Public DNS for WebServer
```

## See also

- [Reference: Cron and rate expressions for Systems Manager](../../../systems-manager/latest/userguide/reference-cron-and-rate-expressions.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Systems Manager

InstanceAssociationOutputLocation

All content copied from https://docs.aws.amazon.com/.
