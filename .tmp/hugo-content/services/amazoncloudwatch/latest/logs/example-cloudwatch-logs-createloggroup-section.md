# Use `CreateLogGroup` with an AWS SDK or CLI

The following code examples show how to use `CreateLogGroup`.

.NET

**SDK for .NET**

###### Note

There's more on GitHub. Find the complete example and learn how to set up and run in the
[AWS Code\
Examples Repository](https://github.com/awsdocs/aws-doc-sdk-examples/tree/main/dotnetv3/CloudWatchLogs).

```csharp

    using System;
    using System.Threading.Tasks;
    using Amazon.CloudWatchLogs;
    using Amazon.CloudWatchLogs.Model;

    /// <summary>
    /// Shows how to create an Amazon CloudWatch Logs log group.
    /// </summary>
    public class CreateLogGroup
    {
        public static async Task Main()
        {
            // This client object will be associated with the same AWS Region
            // as the default user on this system. If you need to use a
            // different AWS Region, pass it as a parameter to the client
            // constructor.
            var client = new AmazonCloudWatchLogsClient();

            string logGroupName = "cloudwatchlogs-example-loggroup";

            var request = new CreateLogGroupRequest
            {
                LogGroupName = logGroupName,
            };

            var response = await client.CreateLogGroupAsync(request);

            if (response.HttpStatusCode == System.Net.HttpStatusCode.OK)
            {
                Console.WriteLine($"Successfully create log group with ID: {logGroupName}.");
            }
            else
            {
                Console.WriteLine("Could not create log group.");
            }
        }
    }

```

- For API details, see
[CreateLogGroup](../../../../reference/goto/dotnetsdkv3/logs-2014-03-28/createloggroup.md)
in _AWS SDK for .NET API Reference_.

CLI

**AWS CLI**

The following command creates a log group named `my-logs`:

```nohighlight

aws logs create-log-group --log-group-name my-logs

```

- For API details, see
[CreateLogGroup](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/logs/create-log-group.html)
in _AWS CLI Command Reference_.

JavaScript

**SDK for JavaScript (v3)**

###### Note

There's more on GitHub. Find the complete example and learn how to set up and run in the
[AWS Code\
Examples Repository](https://github.com/awsdocs/aws-doc-sdk-examples/tree/main/javascriptv3/example_code/cloudwatch-logs).

```javascript

import { CreateLogGroupCommand } from "@aws-sdk/client-cloudwatch-logs";
import { client } from "../libs/client.js";

const run = async () => {
  const command = new CreateLogGroupCommand({
    // The name of the log group.
    logGroupName: process.env.CLOUDWATCH_LOGS_LOG_GROUP,
  });

  try {
    return await client.send(command);
  } catch (err) {
    console.error(err);
  }
};

export default run();

```

- For API details, see
[CreateLogGroup](../../../../reference/awsjavascriptsdk/v3/latest/client/cloudwatch-logs/command/createloggroupcommand.md)
in _AWS SDK for JavaScript API Reference_.

For a complete list of AWS SDK developer guides and code examples, see
[Using CloudWatch Logs with an AWS SDK](../../../../reference/amazoncloudwatch/latest/logs/sdk-general-information-section.md).
This topic also includes information about getting started and details about previous SDK versions.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateExportTask

CreateLogStream

All content copied from https://docs.aws.amazon.com/.
