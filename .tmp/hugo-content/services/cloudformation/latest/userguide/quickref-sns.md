# Amazon SNS template snippets

This example shows an Amazon SNS topic resource. It requires a valid email address.

## JSON

```json

"MySNSTopic" : {
    "Type" : "AWS::SNS::Topic",
    "Properties" : {
        "Subscription" : [ {
            "Endpoint" : "add valid email address",
            "Protocol" : "email"
        } ]
    }
}
```

## YAML

```yaml

MySNSTopic:
  Type: AWS::SNS::Topic
  Properties:
    Subscription:
    - Endpoint: "add valid email address"
      Protocol: email
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon S3

Amazon SQS

All content copied from https://docs.aws.amazon.com/.
