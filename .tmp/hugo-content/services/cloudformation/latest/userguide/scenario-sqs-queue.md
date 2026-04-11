# Amazon SQS template snippets

This example shows an Amazon SQS queue.

## JSON

```json

"MyQueue" : {
    "Type" : "AWS::SQS::Queue",
    "Properties" : {
        "VisibilityTimeout" : "value"
    }
}
```

## YAML

```yaml

MyQueue:
  Type: AWS::SQS::Queue
  Properties:
    VisibilityTimeout: value
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon SNS

Amazon Timestream

All content copied from https://docs.aws.amazon.com/.
