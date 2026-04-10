This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchServerless::SecurityPolicy

Creates an encryption or network policy to be used by one or more OpenSearch Serverless
collections.

Network policies specify access to a collection and its OpenSearch Dashboards endpoint
from public networks or specific VPC endpoints. For more information, see [Network access\
for Amazon OpenSearch Serverless](../../../opensearch-service/latest/developerguide/serverless-network.md).

Encryption policies specify a KMS encryption key to assign to particular collections.
For more information, see [Encryption at\
rest for Amazon OpenSearch Serverless](../../../opensearch-service/latest/developerguide/serverless-encryption.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::OpenSearchServerless::SecurityPolicy",
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

Type: AWS::OpenSearchServerless::SecurityPolicy
Properties:
  Description: String
  Name: String
  Policy: String
  Type: String

```

## Properties

`Description`

The description of the security policy.

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

The type of security policy. Can be either `encryption` or
`network`.

_Required_: Yes

_Type_: String

_Allowed values_: `encryption | network`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the name of the security policy. For more information
about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

## Examples

- [Create an encryption policy](#aws-resource-opensearchserverless-securitypolicy--examples--Create_an_encryption_policy)

- [Create a network policy](#aws-resource-opensearchserverless-securitypolicy--examples--Create_a_network_policy)

### Create an encryption policy

The following example specifies an OpenSearch Serverless encryption policy named
`logs-encryption-policy` with an AWS owned key. The
policy will apply to all future collections with names that begin with "logs".

For a complete sample policy that creates network, encryption, and access
policies, as well as a matching collection, see [Using AWS CloudFormation to create Amazon OpenSearch Serverless\
collections](../../../opensearch-service/latest/developerguide/serverless-cfn.md) in the _Amazon OpenSearch Service Developer_
_Guide._

#### JSON

```json

{
   "Description":"OpenSearch Serverless encryption policy template",
   "Resources":{
      "TestSecurityPolicy":{
         "Type":"AWS::OpenSearchServerless::SecurityPolicy",
         "Properties":{
            "Name":"logs-encryption-policy",
            "Type":"encryption",
            "Description":"Encryption policy for test collections",
            "Policy":"{\"Rules\":[{\"ResourceType\":\"collection\",\"Resource\":[\"collection/logs*\"]}],\"AWSOwnedKey\":true}"
         }
      }
   }
}
```

#### YAML

```yaml

Description: OpenSearch Serverless encryption policy template
Resources:
  TestSecurityPolicy:
    Type: 'AWS::OpenSearchServerless::SecurityPolicy'
    Properties:
      Name: logs-encryption-policy
      Type: encryption
      Description: Encryption policy for test collections
      Policy: >-
        {"Rules":[{"ResourceType":"collection","Resource":["collection/logs*"]}],"AWSOwnedKey":true}
```

### Create a network policy

The following example specifies an OpenSearch Serverless network policy named
`logs-network-policy`. It provides public access to OpenSearch
endpoints and OpenSearch Dashboards endpoints. The policy will apply to all
collections with names that begin with "logs".

#### JSON

```json

{
   "Description":"OpenSearch Serverless network policy template",
   "Resources":{
      "SecurityPolicy":{
         "Type":"AWS::OpenSearchServerless::SecurityPolicy",
         "Properties":{
            "Name":"logs-network-policy",
            "Type":"network",
            "Description":"Network policy for test collections",
            "Policy":"[{\"Rules\":[{\"ResourceType\":\"collection\",\"Resource\":[\"collection/logs*\"]},
                {\"ResourceType\":\"dashboard\",\"Resource\":[\"collection/logs*\"]}],\"AllowFromPublic\":true}]"
         }
      }
   }
}
```

#### YAML

```yaml

Description: OpenSearch Serverless network policy template
Resources:
  SecurityPolicy:
    Type: 'AWS::OpenSearchServerless::SecurityPolicy'
    Properties:
      Name: logs-network-policy
      Type: network
      Description: Network policy for test collections
      Policy >-
        [{"Rules":[{"ResourceType":"collection","Resource":["collection/logs*"]},
        {"ResourceType":"dashboard","Resource":["collection/logs*"]}],"AllowFromPublic":true}]
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SamlConfigOptions

AWS::OpenSearchServerless::VpcEndpoint

All content copied from https://docs.aws.amazon.com/.
