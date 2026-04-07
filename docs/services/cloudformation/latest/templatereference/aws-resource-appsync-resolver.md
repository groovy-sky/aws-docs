This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::Resolver

The `AWS::AppSync::Resolver` resource defines the logical GraphQL resolver
that you attach to fields in a schema. Request and response templates for resolvers are
written in Apache Velocity Template Language (VTL) format. For more information about
resolvers, see [Resolver\
Mapping Template Reference](https://docs.aws.amazon.com/appsync/latest/devguide/resolver-mapping-template-reference.html).

###### Note

When you submit an update, AWS CloudFormation updates resources based on
differences between what you submit and the stack's current template. To cause this
resource to be updated you must change a property value for this resource in the
CloudFormation template. Changing the Amazon S3 file
content without changing a property value will not result in an update
operation.

See [Update Behaviors of Stack Resources](../userguide/using-cfn-updating-stacks-update-behaviors.md) in the _AWS CloudFormation User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppSync::Resolver",
  "Properties" : {
      "ApiId" : String,
      "CachingConfig" : CachingConfig,
      "Code" : String,
      "CodeS3Location" : String,
      "DataSourceName" : String,
      "FieldName" : String,
      "Kind" : String,
      "MaxBatchSize" : Integer,
      "MetricsConfig" : String,
      "PipelineConfig" : PipelineConfig,
      "RequestMappingTemplate" : String,
      "RequestMappingTemplateS3Location" : String,
      "ResponseMappingTemplate" : String,
      "ResponseMappingTemplateS3Location" : String,
      "Runtime" : AppSyncRuntime,
      "SyncConfig" : SyncConfig,
      "TypeName" : String
    }
}

```

### YAML

```yaml

Type: AWS::AppSync::Resolver
Properties:
  ApiId: String
  CachingConfig:
    CachingConfig
  Code: String
  CodeS3Location: String
  DataSourceName: String
  FieldName: String
  Kind: String
  MaxBatchSize: Integer
  MetricsConfig: String
  PipelineConfig:
    PipelineConfig
  RequestMappingTemplate: String
  RequestMappingTemplateS3Location: String
  ResponseMappingTemplate: String
  ResponseMappingTemplateS3Location: String
  Runtime:
    AppSyncRuntime
  SyncConfig:
    SyncConfig
  TypeName: String

```

## Properties

`ApiId`

The AWS AppSync GraphQL API to which you want to attach this
resolver.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CachingConfig`

The caching configuration for the resolver.

_Required_: No

_Type_: [CachingConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appsync-resolver-cachingconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

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

The resolver data source name.

_Required_: No

_Type_: String

_Pattern_: `[_A-Za-z][_0-9A-Za-z]*`

_Minimum_: `1`

_Maximum_: `65536`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldName`

The GraphQL field on a type that invokes the resolver.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Kind`

The resolver type.

- **UNIT**: A UNIT resolver type. A UNIT resolver is the default
resolver type. You can use a UNIT resolver to run a GraphQL query against a single data source.

- **PIPELINE**: A PIPELINE resolver type. You can use a PIPELINE
resolver to invoke a series of `Function` objects in a serial manner. You can use a pipeline
resolver to run a GraphQL query against multiple data sources.

_Required_: No

_Type_: String

_Allowed values_: `UNIT | PIPELINE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxBatchSize`

The maximum number of resolver request inputs that will be sent to a single AWS Lambda function in a `BatchInvoke` operation.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricsConfig`

Enables or disables enhanced resolver metrics for specified resolvers. Note that
`MetricsConfig` won't be used unless the
`resolverLevelMetricsBehavior` value is set to
`PER_RESOLVER_METRICS`. If the `resolverLevelMetricsBehavior`
is set to `FULL_REQUEST_RESOLVER_METRICS` instead, `MetricsConfig`
will be ignored. However, you can still set its value.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PipelineConfig`

Functions linked with the pipeline resolver.

_Required_: No

_Type_: [PipelineConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appsync-resolver-pipelineconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequestMappingTemplate`

The request mapping template.

Request mapping templates are optional when using a Lambda data source. For all other
data sources, a request mapping template is required.

_Required_: No

_Type_: String

_Pattern_: `^.*$`

_Minimum_: `1`

_Maximum_: `65536`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequestMappingTemplateS3Location`

The location of a request mapping template in an Amazon S3 bucket. Use this
if you want to provision with a template file in Amazon S3 rather than
embedding it in your CloudFormation template.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResponseMappingTemplate`

The response mapping template.

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

_Type_: [AppSyncRuntime](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appsync-resolver-appsyncruntime.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SyncConfig`

The `SyncConfig` for a resolver attached to a versioned data source.

_Required_: No

_Type_: [SyncConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appsync-resolver-syncconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TypeName`

The GraphQL type that invokes this resolver.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of an `AWS::AppSync::Resolver` resource to the
intrinsic `Ref` function, the function returns the ARN of the Resolver, such
as
`arn:aws:appsync:us-east-1:123456789012:apis/graphqlapiid/types/typename/resolvers/resolvername`.

For more information about using the `Ref` function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref).

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. The
following are the available attributes and sample return values.

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt).

`FieldName`

The GraphQL field on a type that invokes the resolver.

`ResolverArn`

ARN of the resolver, such as
`arn:aws:appsync:us-east-1:123456789012:apis/graphqlapiid/types/typename/resolvers/resolvername`.

`TypeName`

The GraphQL type that invokes this resolver.

## Examples

### Resolver Creation Example

The following example creates a resolver and associates it with an existing
GraphQL API and a data source by passing the GraphQL API ID and data source name
as a parameter.

#### YAML

```yaml

Parameters:
  graphQlApiId:
    Type: String
  dataSourceName:
    Type: String
  fieldName:
    Type: String
  typeName:
    Type: String
  requestMappingTemplateS3LocationInput:
    Type: String
  responseMappingTemplateS3LocationInput:
    Type: String
Resources:
  Resolver:
    Type: AWS::AppSync::Resolver
    Properties:
      ApiId:
	Ref: graphQlApiId
      TypeName:
        Ref: typeName
      FieldName:
        Ref: fieldName
      DataSourceName:
        Ref: dataSourceName
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
    "dataSourceName": {
      "Type": "String"
    },
    "fieldName": {
      "Type": "String"
    },
    "typeName": {
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
    "Resolver": {
      "Type": "AWS::AppSync::Resolver",
      "Properties": {
        "ApiId": {
          "Ref": "graphQlApiId"
        },
        "TypeName": {
          "Ref": "typeName"
        },
        "FieldName": {
          "Ref": "fieldName"
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

- [CreateResolver](https://docs.aws.amazon.com/appsync/latest/APIReference/API_CreateResolver.html) operation in the _AWS AppSync API Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::AppSync::GraphQLSchema

AppSyncRuntime
