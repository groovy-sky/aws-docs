This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::ApiDestination

Creates an API destination, which is an HTTP invocation endpoint configured as a target
for events.

When using ApiDesinations with OAuth authentication we recommend these best practices:

- Create a secret in Secrets Manager with your OAuth credentials.

- Reference that secret in your CloudFormation template for `AWS::Events::Connection` using
CloudFormation dynamic reference syntax. For more information, see [Secrets Manager secrets](../userguide/dynamic-references.md#dynamic-references-secretsmanager).

When the Connection resource is created the secret will be passed to EventBridge and stored in the customer account using “Service Linked Secrets,”
effectively creating two secrets. This will minimize the cost because the original secret is only accessed when a CloudFormation template is created or updated,
not every time an event is sent to the ApiDestination. The secret stored in the customer account by EventBridge is the one used for each event sent to the
ApiDestination and AWS is responsible for the fees.

###### Note

The secret stored in the customer account by EventBridge can’t be updated directly, only when a CloudFormation template is updated.

For examples of CloudFormation templates that use secrets, see [Examples](../userguide/aws-resource-events-connection.md#aws-resource-events-connection--examples).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Events::ApiDestination",
  "Properties" : {
      "ConnectionArn" : String,
      "Description" : String,
      "HttpMethod" : String,
      "InvocationEndpoint" : String,
      "InvocationRateLimitPerSecond" : Integer,
      "Name" : String
    }
}

```

### YAML

```yaml

Type: AWS::Events::ApiDestination
Properties:
  ConnectionArn: String
  Description: String
  HttpMethod: String
  InvocationEndpoint: String
  InvocationRateLimitPerSecond: Integer
  Name: String

```

## Properties

`ConnectionArn`

The ARN of the connection to use for the API destination. The destination endpoint must
support the authorization type specified for the connection.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws([a-z]|\-)*:events:([a-z]|\d|\-)*:([0-9]{12})?:connection/[\.\-_A-Za-z0-9]+/[\-A-Za-z0-9]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description for the API destination to create.

_Required_: No

_Type_: String

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HttpMethod`

The method to use for the request to the HTTP invocation endpoint.

_Required_: Yes

_Type_: String

_Allowed values_: `GET | HEAD | POST | OPTIONS | PUT | DELETE | PATCH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InvocationEndpoint`

The URL to the HTTP invocation endpoint for the API destination.

_Required_: Yes

_Type_: String

_Pattern_: `^((%[0-9A-Fa-f]{2}|[-()_.!~*';/?:@\x26=+$,A-Za-z0-9])+)([).!';/?:,])?$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InvocationRateLimitPerSecond`

The maximum number of requests per second to send to the HTTP invocation endpoint.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name for the API destination to create.

_Required_: No

_Type_: String

_Pattern_: `[\.\-_A-Za-z0-9]+`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the API destination that was created by the
request.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the API destination that was created by the request.

`ArnForPolicy`

Returns the Amazon Resource Name (ARN) of an API destination in resource format, so it can be used in the `Resource` element of IAM permission policy statements.
For more information, see [Resource types defined by Amazon EventBridge](../../../service-authorization/latest/reference/list-amazoneventbridge.md#amazoneventbridge-resources-for-iam-policies) in the _Service Authorization Reference_.

For example, the following resource defines an IAM policy that grants permission to update a specific API destination.

```

Resources:
  ExamplePolicy:
    Type: AWS::IAM::Policy
    Properties:
      PolicyName: ExamplePolicy
      PolicyDocument:
        Version: '2012-10-17'
        Statement:
          - Effect: Allow
            Action:
              - events:UpdateApiDestination
            Resource:
              - !GetAtt myApiDestination.ArnForPolicy
```

## Examples

- [Create an ApiDestination for PagerDuty](#aws-resource-events-apidestination--examples--Create_an_ApiDestination_for_PagerDuty)

- [Create an ApiDestination for Slack](#aws-resource-events-apidestination--examples--Create_an_ApiDestination_for_Slack)

### Create an ApiDestination for PagerDuty

The following example creates an ApiDestination connection to PagerDuty.

#### JSON

```json

"Parameters": {
   "PagerDutyAPIKeyParam": {
       "Type": "String",
       "Description": "API Key for the PagerDuty Environment",
       "NoEcho": true
},
"Resources": {
   "MyConnection": {
       "Type": "AWS::Events::Connection",
       "Properties": {
           "AuthorizationType": "API_KEY",
           "Description": "Connection to PagerDuty API",
           "AuthParameters": {
               "ApiKeyAuthParameters": {
                   "ApiKeyName": "PagerDuty Authorization",
                   "ApiKeyValue": {
                       "Ref": "PagerDutyAPIKeyParam"
                   }
               }
           }
       }
   },
   "MyApiDestination": {
       "Type": "AWS::Events::ApiDestination",
       "Properties": {
           "ConnectionArn": {
               "Fn::GetAtt": [
                   "MyConnection",
                   "Arn"
               ]
           },
           "Description": "API Destination to send events to PagerDuty",
           "HttpMethod": "POST",
           "InvocationEndpoint": "https://events.pagerduty.com/v2/enqueue"
      }
   },
}
```

#### YAML

```yaml

Parameters:
  PagerDutyAPIKeyParam:
    Type: String
    Description: API Key for the PagerDuty Environment
    NoEcho: True

Resources:
  MyConnection:
    Type: AWS::Events::Connection
    Properties:
      AuthorizationType: API_KEY
      Description: Connection to PagerDuty API
      AuthParameters:
        ApiKeyAuthParameters:
          ApiKeyName: PagerDuty Authorization
          ApiKeyValue: !Ref PagerDutyAPIKeyParam

  MyApiDestination:
    Type: AWS::Events::ApiDestination
    Properties:
      ConnectionArn: !GetAtt MyConnection.Arn
      Description: API Destination to send events to PagerDuty
      HttpMethod: POST
      InvocationEndpoint: https://events.pagerduty.com/v2/enqueue
```

### Create an ApiDestination for Slack

The following example creates an ApiDestination connection to Slack.

#### JSON

```json

"Parameters": {
   "pName": {
       "Type": "String"
    },
   "pEndpoint": {
       "Type": "String"
    },
   "pSecretId": {
       "Type": "String"
    }
 },
"Resources": {
   "SlackRole": {
       "Type": "AWS::IAM::Role",
       "Properties": {
           "AssumeRolePolicyDocument": {
               "Version": "2012-10-17",
               "Statement": [
                {
                   "Effect": "Allow",
                   "Principal": {
                       "Service": [
                           "events.amazonaws.com"
                        ]
                    },
                   "Action": [
                       "sts:AssumeRole"
                    ]
               }
               ]
           },
          "Path": "/service-role/",
          "Policies": [
           {
              "PolicyName": "eventbridge-api-destinations",
              "PolicyDocument": {
                  "Version": "2012-10-17",
                  "Statement": [
                   {
                      "Effect": "Allow",
                      "Action": [
                          "events:InvokeApiDestination"
                       ],
                      "Resource": {
                          "Fn::GetAtt": [
                              "SlackDestination",
                              "Arn"
                           ]
                      }
                  }
                  ]
              }
         }
        ]
      }
    },
   "SlackBus": {
       "Type": "AWS::Events::EventBus",
       "Properties": {
           "Name": {
               "Ref": "pName"
            }
        }
    },
   "SlackConnection": {
       "Type": "AWS::Events::Connection",
       "Properties": {
           "AuthorizationType": "API_KEY",
           "AuthParameters": {
               "ApiKeyAuthParameters": {
                   "ApiKeyName": "Authorization",
                   "ApiKeyValue": {
                       "Fn::Sub": "{{resolve:secretsmanager:arn:aws:secretsmanager:${AWS::Region}:${AWS::AccountId}:secret:${pSecretId}}}"
                    }
                }
            }
        }
    },
   "SlackDestination": {
       "Type": "AWS::Events::ApiDestination",
       "Properties": {
           "ConnectionArn": {
               "Fn::GetAtt": [
                   "SlackConnection",
                   "Arn"
                ]
            },
           "HttpMethod": "POST",
           "InvocationEndpoint": {
               "Ref": "pEndpoint"
            },
           "InvocationRateLimitPerSecond": 10
        }
    },
   "SlackRule": {
       "Type": "AWS::Events::Rule",
       "Properties": {
           "State": "ENABLED",
           "EventBusName": {
               "Ref": "SlackBus"
            },
           "EventPattern": {
               "detail": {
                   "channel": [
                       "aws-integration-testing"
                    ]
                }
           },
         "Targets": [
          {
              "Id": "slack-destination",
              "Arn": {
                  "Fn::GetAtt": [
                      "SlackDestination",
                      "Arn"
                   ]
              },
             "RoleArn": {
                 "Fn::GetAtt": [
                     "SlackRole",
                     "Arn"
                  ]
              },
            "HttpParameters": {
                "HeaderParameters": {
                    "Content-type": "application/json;charset=utf-8"
                 }
             },
            "InputTransformer": {
                "InputPathsMap": {
                    "channel": "$.detail.channel",
                    "text": "$.detail.text"
                 },
                "InputTemplate": "{\n  \"channel\": <channel>,\n  \"text\": <text>\n}\n"
            }
          }
          ]
        }
    }
 },
"Outputs": {
    "outRole": {
        "Value": {
            "Fn::GetAtt": [
                "SlackRole",
                "Arn"
             ]
         }
     },
    "outBus": {
        "Value": {
            "Fn::GetAtt": [
                "SlackBus",
                "Arn"
             ]
         }
     },
    "outConnection": {
        "Value": {
            "Fn::GetAtt": [
                "SlackConnection",
                "Arn"
             ]
         }
     },
    "outConnectionSecret": {
        "Value": {
            "Fn::GetAtt": [
                "SlackConnection",
                "SecretArn"
             ]
         }
     },
    "outDestination": {
        "Value": {
            "Fn::GetAtt": [
                "SlackDestination",
                "Arn"
             ]
         }
     },
    "outRule": {
        "Value": {
            "Fn::GetAtt": [
                "SlackRule",
                "Arn"
             ]
         }
     }
 }
}
```

#### YAML

```yaml

Parameters:
  pName:
    Type: String
  pEndpoint:
    Type: String
  pSecretId:
    Type: String

Resources:
  SlackRole:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: '2012-10-17'
        Statement:
          - Effect: Allow
            Principal:
              Service:
                - events.amazonaws.com
            Action:
              - sts:AssumeRole
      Path: '/service-role/'
      Policies:
        - PolicyName: eventbridge-api-destinations
          PolicyDocument:
            Version: '2012-10-17'
            Statement:
              - Effect: Allow
                Action:
                  - events:InvokeApiDestination
                Resource: !GetAtt SlackDestination.Arn
  SlackBus:
    Type: AWS::Events::EventBus
      Properties:
        Name: !Ref pName
  SlackConnection:
    Type: AWS::Events::Connection
    Properties:
      AuthorizationType: API_KEY
      AuthParameters:
        ApiKeyAuthParameters:
          ApiKeyName: Authorization
          ApiKeyValue: !Sub '{{resolve:secretsmanager:arn:aws:secretsmanager:${AWS::Region}:${AWS::AccountId}:secret:${pSecretId}}}'
  SlackDestination:
    Type: AWS::Events::ApiDestination
    Properties:
      ConnectionArn: !GetAtt SlackConnection.Arn
      HttpMethod: POST
      InvocationEndpoint: !Ref pEndpoint
      InvocationRateLimitPerSecond: 10
  SlackRule:
    Type: AWS::Events::Rule
    Properties:
      State: ENABLED
      EventBusName: !Ref SlackBus
      EventPattern:
        detail:
          channel: ["aws-integration-testing"]
      Targets:
        - Id: slack-destination
          Arn: !GetAtt SlackDestination.Arn
          RoleArn: !GetAtt SlackRole.Arn
          HttpParameters:
            HeaderParameters:
              Content-type: application/json;charset=utf-8
          InputTransformer:
            InputPathsMap:
              "channel": "$.detail.channel"
              "text": "$.detail.text"
          InputTemplate: |
            {
              "channel": <channel>,
              "text": <text>
            }
Outputs:
  outRole:
    Value: !GetAtt SlackRole.Arn
  outBus:
    Value: !GetAtt SlackBus.Arn
  outConnection:
    Value: !GetAtt SlackConnection.Arn
  outConnectionSecret:
    Value: !GetAtt SlackConnection.SecretArn
  outDestination:
    Value: !GetAtt SlackDestination.Arn
  outRule:
    Value: !GetAtt SlackRule.Arn
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon EventBridge

AWS::Events::Archive

All content copied from https://docs.aws.amazon.com/.
