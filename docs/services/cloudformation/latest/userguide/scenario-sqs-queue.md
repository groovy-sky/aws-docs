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

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon SNS

Amazon Timestream
