This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchServerless::AccessPolicy

Creates a data access policy for OpenSearch Serverless. Access policies limit access to collections
and the resources within them, and allow a user to access that data irrespective of the
access mechanism or network source. For more information, see [Data\
access control for Amazon OpenSearch Serverless](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-data-access.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::OpenSearchServerless::AccessPolicy",
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

Type: AWS::OpenSearchServerless::AccessPolicy
Properties:
  Description: String
  Name: String
  Policy: String
  Type: String

```

## Properties

`Description`

The description of the policy.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the policy.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z][a-z0-9-]{2,31}$`

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

The type of access policy. Currently the only option is `data`.

_Required_: Yes

_Type_: String

_Allowed values_: `data`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the name of the access policy. For more information
about using the `Ref` function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

## Examples

### Create an access policy that allows access to all collections and indexes

The following example specifies an OpenSearch Serverless access policy that
provides full access to the resources within `my-collection` to the user
`test-user`.

For a complete sample policy that creates network, encryption, and access
policies, as well as a matching collection, see [Using AWS CloudFormation to create Amazon OpenSearch Serverless\
collections](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-cfn.html) in the _Amazon OpenSearch Service Developer_
_Guide._

#### JSON

```json

{
   "Description":"OpenSearch Serverless access policy template",
   "Resources":{
      "TestAccessPolicy":{
         "Type":"AWS::OpenSearchServerless::AccessPolicy",
         "Properties":{
            "Name":"test-access-policy",
            "Type":"data",
            "Description":"Access policy for
                my-collection",
            "Policy":{
               "Fn::Sub":"[{\"Description\":\"Access for
                test-user\",\"Rules\":[{\"ResourceType\":\"index\",\"Resource\":[\"index/*/*\"],\"Permission\":[\"aoss:*\"]},
                {\"ResourceType\":\"collection\",\"Resource\":[\"collection/my-collection\"],\"Permission\":[\"aoss:*\"]}],
                \"Principal\":[\"arn:aws:iam::${AWS::AccountId}:user/test-user\"]}]"
            }
         }

```

#### YAML

```yaml

Description: 'OpenSearch Serverless access policy template'
  Resources:
  TestAccessPolicy:
    Type: 'AWS::OpenSearchServerless::AccessPolicy'
    Properties:
      Name: test-access-policy
      Type: data
      Description: Access policy for my-collection
      Policy:
        !Sub >-
         [{"Description":"Access for
         test-user","Rules":[{"ResourceType":"index","Resource":["index/*/*"],"Permission":["aoss:*"]},
         {"ResourceType":"collection","Resource":["collection/my-collection"],"Permission":["aoss:*"]}],
         "Principal":["arn:aws:iam::${AWS::AccountId}:user/test-user"]}]
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon OpenSearch Serverless

AWS::OpenSearchServerless::Collection
