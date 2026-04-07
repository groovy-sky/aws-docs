This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchServerless::LifecyclePolicy

Creates a lifecyle policy to be applied to OpenSearch Serverless indexes. Lifecycle policies define
the number of days or hours to retain the data on an OpenSearch Serverless index. For more information,
see [Creating data lifecycle policies](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-lifecycle.html#serverless-lifecycle-create).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::OpenSearchServerless::LifecyclePolicy",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "Policy" : String,
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::OpenSearchServerless::LifecyclePolicy
Properties:
  Description: String
  Name: String
  Policy: String
  Type: String

```

## Properties

`Description`

The description of the lifecycle policy.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the lifecycle policy.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z][a-z0-9-]+$`

_Minimum_: `3`

_Maximum_: `32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Policy`

The JSON policy document without any whitespaces.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0009\u000A\u000D\u0020-\u007E\u00A1-\u00FF]+`

_Minimum_: `1`

_Maximum_: `20480`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of lifecycle policy.

_Required_: Yes

_Type_: String

_Allowed values_: `retention`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the type and name of the lifecycle policy. For more
information about using the `Ref` function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

## Examples

### Create a lifecycle policy that sets a minimum rollover period for all indexes in a collection

The following example specifies an OpenSearch Serverless lifecycle policy that
sets a minimum rollover period for all indexes within
`my-collection`.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Description": "OpenSearch Serverless lifecycle policy template",
  "Resources": {
    "TestLifecyclePolicy": {
      "Type": "AWS::OpenSearchServerless::LifecyclePolicy",
      "Properties": {
        "Name": "test-lifecycle-policy",
        "Type": "retention",
        "Description": "Lifecycle policy for all indexes in my-collection",
        "Policy": {\"Rules\": [{\"Resource\": [\"index/my-collection/*\"],\"ResourceType\": \"index\",\"MinIndexRetention\": \"2d\"}}
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09T00:00:00.000Z
Description: OpenSearch Serverless lifecycle policy template
Resources:
  TestLifecyclePolicy:
    Type: 'AWS::OpenSearchServerless::LifecyclePolicy'
    Properties:
      Name: test-lifecycle-policy
      Type: retention
      Description: Lifecycle policy for all indexes in my-collection
      Policy: {"Rules": [{"Resource": ["index/my-collection/*"],"ResourceType": "index","MinIndexRetention": "2d"}]}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PropertyMapping

AWS::OpenSearchServerless::SecurityConfig
