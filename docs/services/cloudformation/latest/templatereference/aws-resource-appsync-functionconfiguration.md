This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::FunctionConfiguration

The `AWS::AppSync::FunctionConfiguration` resource defines the functions in
GraphQL APIs to perform certain operations. You can use pipeline resolvers to attach
functions. For more information, see [Pipeline Resolvers](https://docs.aws.amazon.com/appsync/latest/devguide/pipeline-resolvers.html) in
the _AWS AppSync Developer Guide_.

###### Note

When you submit an update, AWS CloudFormation updates resources based on
differences between what you submit and the stack's current template. To cause this
resource to be updated you must change a property value for this resource in the
CloudFormation template. Changing the Amazon S3 file content
without changing a property value will not result in an update operation.

See [Update Behaviors of Stack Resources](../userguide/using-cfn-updating-stacks-update-behaviors.md) in the _AWS CloudFormation User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppSync::FunctionConfiguration",
  "Properties" : {
      "ApiId" : String,
      "Code" : String,
      "CodeS3Location" : String,
      "DataSourceName" : String,
      "Description" : String,
      "FunctionVersion" : String,
      "MaxBatchSize" : Integer,
      "Name" : String,
      "RequestMappingTemplate" : String,
      "RequestMappingTemplateS3Location" : String,
      "ResponseMappingTemplate" : String,
      "ResponseMappingTemplateS3Location" : String,
      "Runtime" : AppSyncRuntime,
      "SyncConfig" : SyncConfig
    }
}

```

### YAML

```yaml

Type: AWS::AppSync::FunctionConfiguration
Properties:
  ApiId: String
  Code: String
  CodeS3Location: String
  DataSourceName: String
  Description: String
  FunctionVersion: String
  MaxBatchSize: Integer
  Name: String
  RequestMappingTemplate: String
  RequestMappingTemplateS3Location: String
  ResponseMappingTemplate: String
  ResponseMappingTemplateS3Location: String
  Runtime:
    AppSyncRuntime
  SyncConfig:
    SyncConfig

```

## Properties

`ApiId`

The AWS AppSync GraphQL API that you want to attach using this
function.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Code`

The `resolver` code that contains the request and response functions. When
code is used, the `runtime` is required. The runtime value must be
`APPSYNC_JS`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CodeS3Location`

The Amazon S3 endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSourceName`

The name of data source this function will attach.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The `Function` description.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FunctionVersion`

The version of the request mapping template. Currently, only the 2018-05-29 version of the template is
supported.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxBatchSize`

The maximum number of resolver request inputs that will be sent to a single AWS Lambda function in a `BatchInvoke` operation.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the function.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequestMappingTemplate`

The `Function` request mapping template. Functions support only the 2018-05-29 version of the
request mapping template.

_Required_: No

_Type_: String

_Pattern_: `^.*$`

_Minimum_: `1`

_Maximum_: `65536`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequestMappingTemplateS3Location`

Describes a Sync configuration for a resolver.

Contains information on which Conflict Detection, as well as Resolution strategy,
should be performed when the resolver is invoked.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResponseMappingTemplate`

The `Function` response mapping template.

_Required_: No

_Type_: String

_Pattern_: `^.*$`

_Minimum_: `1`

_Maximum_: `65536`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResponseMappingTemplateS3Location`

The location of a response mapping template in an Amazon S3 bucket. Use this
if you want to provision with a template file in Amazon S3 rather than
embedding it in your CloudFormation template.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Runtime`

Describes a runtime used by an AWS AppSync resolver or AWS AppSync function. Specifies the name and version of the runtime to use.
Note that if a runtime is specified, code must also be specified.

_Required_: No

_Type_: [AppSyncRuntime](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appsync-functionconfiguration-appsyncruntime.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SyncConfig`

Describes a Sync configuration for a resolver.

Specifies which Conflict Detection strategy and Resolution strategy to use when the resolver is
invoked.

_Required_: No

_Type_: [SyncConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appsync-functionconfiguration-syncconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of an `AWS::AppSync::FunctionConfiguration`
resource to the intrinsic `Ref` function, the function returns the ARN of the
FunctionConfiguration, such as
`arn:aws:appsync:us-east-1:123456789012:apis/graphqlapiid/functions/functionid`.

For more information about using the `Ref` function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref).

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. The
following are the available attributes and sample return values.

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt).

`DataSourceName`

The name of data source this function will attach.

`FunctionArn`

ARN of the function, such as
`arn:aws:appsync:us-east-1:123456789012:apis/graphqlapiid/functions/functionId`.

`FunctionId`

The unique ID of this function.

`Name`

The name of the function.

## Examples

### Function Creation Example

The following example creates a function and associates it with an existing
GraphQL API and a data source by passing the GraphQL API ID and data source name
as a parameter.

#### YAML

```yaml

Parameters:
  graphQlApiId:
    Type: String
  dataSourceName:
    Type: String
  name:
    Type: String
  description:
    Type: String
  functionVersion:
    Type: String
  requestMappingTemplateS3LocationInput:
    Type: String
  responseMappingTemplateS3LocationInput:
    Type: String
Resources:
  FunctionConfiguration:
    Type: AWS::AppSync::FunctionConfiguration
    Properties:
      ApiId:
	Ref: graphQlApiId
      Name:
        Ref: name
      Description:
        Ref: description
      DataSourceName:
        Ref: dataSourceName
      FunctionVersion:
        Ref: functionVersion
      RequestMappingTemplateS3Location:
        Ref: requestMappingTemplateS3LocationInput
      ResponseMappingTemplateS3Location:
        Ref: responseMappingTemplateS3LocationInput

```

#### JSON

```json

{
  "Parameters": {
    "graphQlApiId": {
      "Type": "String"
    },
    "name": {
      "Type": "String"
    },
    "description": {
      "Type": "String"
    },
    "dataSourceName": {
      "Type": "String"
    },
    "functionVersion": {
      "Type": "String"
    },
    "requestMappingTemplateS3LocationInput": {
      "Type": "String"
    },
    "responseMappingTemplateS3LocationInput": {
      "Type": "String"
    }
  },
  "Resources": {
    "FunctionConfiguration": {
      "Type": "AWS::AppSync::FunctionConfiguration",
      "Properties": {
        "ApiId": {
          "Ref": "graphQlApiId"
        },
        "Name": {
          "Ref": "name"
        },
        "Description": {
          "Ref": "description"
        },
        "FunctionVersion": {
          "Ref": "functionVersion"
        },
        "DataSourceName": {
          "Ref": "dataSourceName"
        },
        "RequestMappingTemplateS3Location": {
          "Ref": "requestMappingTemplateS3LocationInput"
        },
        "ResponseMappingTemplateS3Location": {
          "Ref": "responseMappingTemplateS3LocationInput"
        }
      }
    }
  }
}

```

## See also

- [CreateFunction](https://docs.aws.amazon.com/appsync/latest/APIReference/API_CreateFunction.html) operation in the _AWS AppSync API Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::AppSync::DomainNameApiAssociation

AppSyncRuntime
