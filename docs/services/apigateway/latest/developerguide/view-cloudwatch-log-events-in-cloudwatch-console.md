# View API Gateway log events in the CloudWatch console

The following section explains the necessary prerequisites and how to view API Gateway log events in the CloudWatch
console.

## Prerequisites

1. You must have an API created in API Gateway. Follow the instructions in [Develop REST APIs in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/rest-api-develop.html).

2. You must have the API deployed and invoked at least once. Follow the instructions in [Deploy REST APIs in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-deploy-api.html) and [Invoke REST APIs in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-call-api.html).

3. You must have CloudWatch Logs enabled for a stage. Follow the instructions in [Set up CloudWatch logging for REST APIs in API Gateway](set-up-logging.md).

## To view logged API requests and responses using the CloudWatch console

1. Open the CloudWatch console at
    [https://console.aws.amazon.com/cloudwatch/](https://console.aws.amazon.com/cloudwatch).

2. If necessary, change the AWS Region. From the navigation bar, select the
    Region where your AWS resources reside. For more information, see [Regions and\
    Endpoints](https://docs.aws.amazon.com/general/latest/gr/rande.html).

3. In the navigation pane, choose **Logs**, **Log groups**.

4. Under the **Log Groups** table, choose a log group of the
    **API-Gateway-Execution-Logs\_{rest-api-id}/{stage-name}**
    name.

5. Under the **Log Streams** table, choose a log stream.
    You can use the timestamp to help locate the log stream of your interest.

6. Choose **Text** to view raw text or choose
    **Row** to view the event row by row.

###### Important

CloudWatch lets you delete log groups or streams. Do not manually delete API Gateway API log
groups or streams; let API Gateway manage these resources. Manually deleting log
groups or streams may cause API requests and responses not to be logged. If that
happens, you can delete the entire log group for the API and redeploy the API.
This is because API Gateway creates log groups or log streams for an API stage
at the time when it is deployed.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

View metrics in the CloudWatch console

Monitoring tools in AWS for API Gateway
