This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::MultiRegionAccessPointPolicy

Applies an Amazon S3 access policy to an Amazon S3 Multi-Region Access Point.

It is not possible to delete an access policy for a Multi-Region Access Point from the
CloudFormation template. When you attempt to delete the policy, CloudFormation updates the
policy using `DeletionPolicy:Retain` and `UpdateReplacePolicy:Retain`.
CloudFormation updates the policy to only allow access to the account that created the
bucket.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3::MultiRegionAccessPointPolicy",
  "Properties" : {
      "MrapName" : String,
      "Policy" : Json
    }
}

```

### YAML

```yaml

Type: AWS::S3::MultiRegionAccessPointPolicy
Properties:
  MrapName: String
  Policy: Json

```

## Properties

`MrapName`

The name of the Multi-Region Access Point.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9][-a-z0-9]{1,48}[a-z0-9]$`

_Minimum_: `3`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Policy`

The access policy associated with the Multi-Region Access Point.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the Multi-Region Access Point.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Simple Multi-Region Access Point Policy

The following example grants access permissions to CloudWatch.

It is very important to note where you need to use the name versus the alias for the
Multi-Region Access Point. In the following example, the name is
`DOC-EXAMPLE-MULTI-REGION-ACCESS-POINT`, the alias of the Multi-Region Access
Point is `mfzwi23gnjvgw.mrap`, and the AWS account is
`123456789012`. For more information about how ARNs for Multi-Region Access
Points work, see [Making requests\
using a Multi-Region Access Point](../../../s3/latest/userguide/multiregionaccesspointrequests.md) in the in the _Amazon S3 User_
_Guide_.

#### JSON

```json

{
   "SampleMultiRegionAccessPointPolicy":{
      "Type":"AWS::S3::MultiRegionAccessPointPolicy",
      "DeletionPolicy":"Retain",
      "UpdateReplacePolicy":"Retain",
      "Properties":{
         "MrapName":{
            "Ref":"DOC-EXAMPLE-MULTI-REGION-ACCESS-POINT"
         },
         "Policy":{
            "Statement":[
               {
                  "Action":[
                     "s3:GetObject"
                  ],
                  "Effect":"Allow",
                  "Resource":{
                     "Fn::Sub":[
                        "arn:aws:s3::123456789012:accesspoint/mfzwi23gnjvgw.mrap/object/*",
                        {
                           "mrapalias":{
                              "Fn::GetAtt":[
                                 "mfzwi23gnjvgw.mrap",
                                 "Alias"
                              ]
                           }
                        }
                     ]
                  },
                  "Principal":{
                     "Service":"cloudwatch.amazonaws.com"
                  }
               }
            ]
         }
      }
   }
}
```

#### YAML

```yaml

SampleMultiRegionAccessPointPolicy:
  Type: 'AWS::S3::MultiRegionAccessPointPolicy'
  DeletionPolicy: Retain
  UpdateReplacePolicy: Retain
  Properties:
    MrapName:
      Ref: DOC-EXAMPLE-MULTI-REGION-ACCESS-POINT
    Policy:
      Statement:
        - Action:
            - 's3:GetObject'
          Effect: Allow
          Resource:
            'Fn::Sub':
              - 'arn:aws:s3::123456789012:accesspoint/mfzwi23gnjvgw.mrap/object/*'
              - mrapalias:
                  'Fn::GetAtt':
                    - mfzwi23gnjvgw.mrap
                    - Alias
          Principal:
            Service: cloudwatch.amazonaws.com
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Region

PolicyStatus
