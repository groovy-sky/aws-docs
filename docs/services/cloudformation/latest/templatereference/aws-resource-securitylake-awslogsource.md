This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityLake::AwsLogSource

Adds a natively supported AWS service as an AWS source. Enables source types for member accounts in required AWS Regions, based on the parameters you specify. You can choose any source type in any Region for either accounts that are part of a trusted organization or standalone accounts. Once you add an AWS service as a source, Security Lake starts collecting logs and events from it.

###### Important

If you want to create multiple sources using
`AWS::SecurityLake::AwsLogSource`, you must use the
`DependsOn` attribute to create the sources sequentially. With the
`DependsOn` attribute you can specify that the creation
of a specific `AWSLogSource` follows another. When you add a
`DependsOn` attribute to a resource, that resource is
created only after the creation of the resource specified in the
`DependsOn` attribute. For an example, see [Add AWS log sources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-awslogsource.html#aws-resource-securitylake-awslogsource--examples).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SecurityLake::AwsLogSource",
  "Properties" : {
      "Accounts" : [ String, ... ],
      "DataLakeArn" : String,
      "SourceName" : String,
      "SourceVersion" : String
    }
}

```

### YAML

```yaml

Type: AWS::SecurityLake::AwsLogSource
Properties:
  Accounts:
    - String
  DataLakeArn: String
  SourceName: String
  SourceVersion: String

```

## Properties

`Accounts`

Specify the AWS account information where you want to enable Security Lake.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataLakeArn`

The Amazon Resource Name (ARN) used to create the data lake.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceName`

The name for a AWS source. This must be a Regionally unique value. For the list of sources supported by Amazon Security Lake see [Collecting data from AWS services](https://docs.aws.amazon.com/security-lake/latest/userguide/internal-sources.html) in the Amazon Security Lake User Guide.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceVersion`

The version for a AWS source. For more details about source versions supported by Amazon Security Lake see [OCSF source identification](https://docs.aws.amazon.com/security-lake/latest/userguide/open-cybersecurity-schema-framework.html#ocsf-source-identification) in the Amazon Security Lake User Guide. This must be a Regionally unique value.

_Required_: Yes

_Type_: String

_Pattern_: `^(latest|[0-9]\.[0-9])$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `ref` function, `ref` returns the `AwsLogSource` name. For example, `VPC_FLOW`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Add AWS log sources

After deploying Security Lake, use this example to add AWS log sources.

#### JSON

```json

{
"AWSTemplateFormatVersion": "2010-09-09",
"Description": "Security Lake Already Deployed",
"Resources": {
    "SecurityLakeSourcesRoute53": {
        "Type": "AWS::SecurityLake::AwsLogSource",
        "Properties": {
            "DataLakeArn": {
                "Fn::Sub": "arn:aws:securitylake:${AWS::Partition}:${AWS::AccountId}:data-lake/default"
            },
            "SourceName": "ROUTE53",
            "SourceVersion": "2.0"
        }
    },
    "SecurityLakeSourcesSecurityHub": {
        "Type": "AWS::SecurityLake::AwsLogSource",
        "Properties": {
            "DataLakeArn": {
                "Fn::Sub": "arn:aws:securitylake:${AWS::Partition}:${AWS::AccountId}:data-lake/default"
            },
            "SourceName": "SH_FINDINGS",
            "SourceVersion": "2.0"
        },
        "DependsOn": "SecurityLakeSourcesRoute53"
    },
    "SecurityLakeSourcesVPCFlow": {
        "Type": "AWS::SecurityLake::AwsLogSource",
        "Properties": {
            "DataLakeArn": {
                "Fn::Sub": "arn:aws:securitylake:${AWS::Partition}:${AWS::AccountId}:data-lake/default"
            },
            "SourceName": "VPC_FLOW",
            "SourceVersion": "2.0"
        },
        "DependsOn": "SecurityLakeSourcesSecurityHub"
    },
    "SecurityLakeSourcesCloudTrailMgmt": {
        "Type": "AWS::SecurityLake::AwsLogSource",
        "Properties": {
            "DataLakeArn": {
                "Fn::Sub": "arn:aws:securitylake:${AWS::Partition}:${AWS::AccountId}:data-lake/default"
            },
            "SourceName": "CLOUD_TRAIL_MGMT",
            "SourceVersion": "2.0"
        },
        "DependsOn": "SecurityLakeSourcesVPCFlow"
    },
    "SecurityLakeSourcesLambdaExecution": {
        "Type": "AWS::SecurityLake::AwsLogSource",
        "Properties": {
            "DataLakeArn": {
                "Fn::Sub": "arn:aws:securitylake:${AWS::Partition}:${AWS::AccountId}:data-lake/default"
            },
            "SourceName": "LAMBDA_EXECUTION",
            "SourceVersion": "2.0"
        },
        "DependsOn": "SecurityLakeSourcesCloudTrailMgmt"
    },
    "SecurityLakeSourcesS3": {
        "Type": "AWS::SecurityLake::AwsLogSource",
        "Properties": {
            "DataLakeArn": {
                "Fn::Sub": "arn:aws:securitylake:${AWS::Partition}:${AWS::AccountId}:data-lake/default"
            },
            "SourceName": "S3_DATA",
            "SourceVersion": "2.0"
        },
        "DependsOn": "SecurityLakeSourcesLambdaExecution"
    },
    "SecurityLakeSourcesEKSAudit": {
        "Type": "AWS::SecurityLake::AwsLogSource",
        "Properties": {
            "DataLakeArn": {
                "Fn::Sub": "arn:aws:securitylake:${AWS::Partition}:${AWS::AccountId}:data-lake/default"
            },
            "SourceName": "EKS_AUDIT",
            "SourceVersion": "2.0"
        },
        "DependsOn": "SecurityLakeSourcesS3"
    }
}
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Description: Security Lake Already Deployed
Resources:
  SecurityLakeSourcesRoute53:
    Type: AWS::SecurityLake::AwsLogSource
    Properties:
      DataLakeArn: !Sub arn:aws:securitylake:${AWS::Partition}:${AWS::AccountId}:data-lake/default
      SourceName: ROUTE53
      SourceVersion: "2.0"
  SecurityLakeSourcesSecurityHub:
    Type: AWS::SecurityLake::AwsLogSource
    Properties:
      DataLakeArn: !Sub arn:aws:securitylake:${AWS::Partition}:${AWS::AccountId}:data-lake/default
      SourceName: SH_FINDINGS
      SourceVersion: "2.0"
    DependsOn: SecurityLakeSourcesRoute53
  SecurityLakeSourcesVPCFlow:
    Type: AWS::SecurityLake::AwsLogSource
    Properties:
      DataLakeArn: !Sub arn:aws:securitylake:${AWS::Partition}:${AWS::AccountId}:data-lake/default
      SourceName: VPC_FLOW
      SourceVersion: "2.0"
    DependsOn: SecurityLakeSourcesSecurityHub
  SecurityLakeSourcesCloudTrailMgmt:
    Type: AWS::SecurityLake::AwsLogSource
    Properties:
      DataLakeArn: !Sub arn:aws:securitylake:${AWS::Partition}:${AWS::AccountId}:data-lake/default
      SourceName: CLOUD_TRAIL_MGMT
      SourceVersion: "2.0"
    DependsOn: SecurityLakeSourcesVPCFlow
  SecurityLakeSourcesLambdaExecution:
    Type: AWS::SecurityLake::AwsLogSource
    Properties:
      DataLakeArn: !Sub arn:aws:securitylake:${AWS::Partition}:${AWS::AccountId}:data-lake/default
      SourceName: LAMBDA_EXECUTION
      SourceVersion: "2.0"
    DependsOn: SecurityLakeSourcesCloudTrailMgmt
  SecurityLakeSourcesS3:
    Type: AWS::SecurityLake::AwsLogSource
    Properties:
      DataLakeArn: !Sub arn:aws:securitylake:${AWS::Partition}:${AWS::AccountId}:data-lake/default
      SourceName: S3_DATA
      SourceVersion: "2.0"
    DependsOn: SecurityLakeSourcesLambdaExecution
  SecurityLakeSourcesEKSAudit:
    Type: AWS::SecurityLake::AwsLogSource
    Properties:
      DataLakeArn: !Sub arn:aws:securitylake:${AWS::Partition}:${AWS::AccountId}:data-lake/default
      SourceName: EKS_AUDIT
      SourceVersion: "2.0"
    DependsOn: SecurityLakeSourcesS3

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Security Lake

AWS::SecurityLake::DataLake
