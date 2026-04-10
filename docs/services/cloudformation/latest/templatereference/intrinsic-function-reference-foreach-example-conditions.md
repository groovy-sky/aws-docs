This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# Examples of `Fn::ForEach` in the `Conditions` section

These examples demonstrate using the `Fn::ForEach` intrinsic function in the
`Conditions` section. For more information about this section, see [Conditions](../userguide/conditions-section-structure.md)
in the _AWS CloudFormation User Guide_.

###### Important

`Conditions` must be the second property listed, or later. Stack creation
will fail if `Conditions` is the first property listed within the template
fragment parameter of `Fn::ForEach`.

```yaml

Resources:
  'Fn::ForEach::Topics':
    - LogicalId
    - !Ref TopicList
    - '${LogicalId}':
        Condition: !Sub 'TopicCondition${LogicalId}'
        Type: AWS::SNS::Topic
        Properties:
          TopicName: !Sub 'My${LogicalId}'
```

`Conditions` must be added as the second key, or later, for stack creation to
succeed:

```yaml

Resources:
  'Fn::ForEach::Topics':
    - LogicalId
    - !Ref TopicList
    - '${LogicalId}':
        Type: AWS::SNS::Topic
        Condition: !Sub 'TopicCondition${LogicalId}'
        Properties:
          TopicName: !Sub 'My${LogicalId}'
```

###### Topics

- [Replicate a single condition](#intrinsic-function-reference-foreach-example-replicated-single-condition)

## Replicate a single condition

This example uses the `Fn::ForEach` intrinsic function in the
`Conditions` section to replicate multiple similar conditions with different
properties.

### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Transform": "AWS::LanguageExtensions",
    "Parameters": {
        "ParamA": {
            "Type": "String",
            "AllowedValues": [
                "true",
                "false"
            ]
        },
        "ParamB": {
            "Type": "String",
            "AllowedValues": [
                "true",
                "false"
            ]
        },
        "ParamC": {
            "Type": "String",
            "AllowedValues": [
                "true",
                "false"
            ]
        },
        "ParamD": {
            "Type": "String",
            "AllowedValues": [
                "true",
                "false"
            ]
        }
    },
    "Conditions": {
        "Fn::ForEach::CheckTrue": [
            "Identifier",
            ["A", "B", "C", "D"],
            {
                "IsParam${Identifier}Enabled": {
                    "Fn::Equals": [
                        {"Ref": {"Fn::Sub": "Param${Identifier}"}},
                        "true"
                    ]
                }
            }
        ]
    },
    "Resources": {
        "WaitConditionHandle": {
            "Type": "AWS::CloudFormation::WaitConditionHandle"
        }
    }
}
```

### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Transform: AWS::LanguageExtensions
Parameters:
  ParamA:
    Type: String
    AllowedValues:
      - 'true'
      - 'false'
  ParamB:
    Type: String
    AllowedValues:
      - 'true'
      - 'false'
  ParamC:
    Type: String
    AllowedValues:
      - 'true'
      - 'false'
  ParamD:
    Type: String
    AllowedValues:
      - 'true'
      - 'false'
Conditions:
  'Fn::ForEach::CheckTrue':
    - Identifier
    - [A, B, C, D]
    - 'IsParam${Identifier}Enabled': !Equals
        - !Ref
          'Fn::Sub': 'Param${Identifier}'
        - 'true'
Resources:
  WaitConditionHandle:
    Type: AWS::CloudFormation::WaitConditionHandle
```

The transformed template will be equivalent to the following template:

```yaml

AWSTemplateFormatVersion: 2010-09-09
Parameters:
  ParamA:
    Type: String
    AllowedValues:
      - 'true'
      - 'false'
  ParamB:
    Type: String
    AllowedValues:
      - 'true'
      - 'false'
  ParamC:
    Type: String
    AllowedValues:
      - 'true'
      - 'false'
  ParamD:
    Type: String
    AllowedValues:
      - 'true'
      - 'false'
Conditions:
  IsParamAEnabled: !Equals
    - !Ref ParamA
    - 'true'
  IsParamBEnabled: !Equals
    - !Ref ParamB
    - 'true'
  IsParamCEnabled: !Equals
    - !Ref ParamC
    - 'true'
  IsParamDEnabled: !Equals
    - !Ref ParamD
    - 'true'
Resources:
  WaitConditionHandle:
    Type: AWS::CloudFormation::WaitConditionHandle
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Outputs section

Fn::GetAtt

All content copied from https://docs.aws.amazon.com/.
