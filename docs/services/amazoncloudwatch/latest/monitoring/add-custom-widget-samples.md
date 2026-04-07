# Sample custom widgets for a CloudWatch dashboard

AWS provides sample custom widgets in both JavaScript and
Python. You can create these sample widgets by using the link for each widget in
this list.

Alternatively,
you can create and customize a widget by using the CloudWatch console. The links in
this list open an AWS CloudFormation console and use
an CloudFormation quick-create link to create the custom widget.

You can also access the custom widget samples on [GitHub](https://github.com/aws-samples/cloudwatch-custom-widgets-samples).

Following this list, complete examples of the Echo widget are shown for
each language.

JavaScript

###### Sample custom widgets in JavaScript

- [Echo](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetEcho-js&template=customWidgets%2FcustomWidgetEcho-js.yaml) – A basic echoer that you can use
to test how HTML appears in a custom widget, without having
to write a new widget.

- [Hello world](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetHelloWorld-js&template=customWidgets%2FcustomWidgetHelloWorld-js.yaml) – A very basic starter
widget.

- [Custom widget debugger](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetDebugger-js&template=customWidgets%2FcustomWidgetDebugger-js.yaml) – A debugger widget that displays useful information about the Lambda runtime environment.

- [Query CloudWatch Logs Insights](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetLogsInsightsQuery-js&template=customWidgets%2FcustomWidgetLogsInsightsQuery-js.yaml) – Run and edit CloudWatch Logs Insights queries.

- [Run Amazon Athena queries](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetAthenaQuery-js&template=customWidgets%2FcustomWidgetAthenaQuery-js.yaml) – Run and edit Athena queries.

- [Call AWS API](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetAwsCall-js&template=customWidgets%2FcustomWidgetAwsCall-js.yaml) – Call any read-only AWS API and display the results in JSON format.

- [Fast CloudWatch bitmap graph](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetCloudWatchBitmapGraph-js&template=customWidgets%2FcustomWidgetCloudWatchBitmapGraph-js.yaml) – Render CloudWatch graphs using on the server side, for fast display.

- [Text widget from CloudWatch dashboard](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetIncludeTextWidget-js&template=customWidgets%2FcustomWidgetIncludeTextWidget-js.yaml) – Displays the first text widget from the specified CloudWatch dashboard.

- [CloudWatch metric data as a table](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetCloudWatchMetricDataTable-js&template=customWidgets%2FcustomWidgetCloudWatchMetricDataTable-js.yaml) – Displays raw CloudWatch metric data in a table.

- [Amazon EC2 table](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetEc2Table-js&template=customWidgets%2FcustomWidgetEc2Table-js.yaml) – Displays the top EC2 instances by CPU utilization. This widget also includes a Reboot button, which is disabled by default.

- [AWS CodeDeploy deployments](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetCodeDeploy-js&template=customWidgets%2FcustomWidgetCodeDeploy-js.yaml) – Displays CodeDeploy deployments.

- [AWS Cost Explorer report](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetCostExplorerReport-js&template=customWidgets%2FcustomWidgetCostExplorerReport-js.yaml) – Displays a report on the cost of each AWS service for a selected time range.

- [Display content of external URL](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetFetchURL-js&template=customWidgets%2FcustomWidgetFetchURL-js.yaml) – Displays the content of an externally accessible URL.

- [Display an Amazon S3 object](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetS3GetObject-js&template=customWidgets%2FcustomWidgetS3GetObject-js.yaml) – Displays an object in an Amazon S3 bucket in your account.

- [Simple SVG pie chart](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetSimplePie-js&template=customWidgets%2FcustomWidgetSimplePie-js.yaml) – Example of a graphical SVG-based widget.

Python

###### Sample custom widgets in Python

- [Echo](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetEcho-py&template=customWidgets%2FcustomWidgetEcho-py.yaml) – A basic echoer which you can use
to test how HTML appears in a custom widget, without having
to write a new widget.

- [Hello world](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetHelloWorld-py&template=customWidgets%2FcustomWidgetHelloWorld-py.yaml) – A very basic starter
widget.

- [Custom widget debugger](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetDebugger-py&template=customWidgets%2FcustomWidgetDebugger-py.yaml) – A debugger widget that displays useful information about the Lambda runtime environment.

- [Call AWS API](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetAwsCall-py&template=customWidgets%2FcustomWidgetAwsCall-py.yaml) – Call any read-only AWS API and display the results in JSON format.

- [Fast CloudWatch bitmap graph](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetCloudWatchBitmapGraph-py&template=customWidgets%2FcustomWidgetCloudWatchBitmapGraph-py.yaml) – Render CloudWatch graphs using on the server side, for fast display.

- [Send dashboard snapshot by email](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetEmailDashboardSnapshot-py&template=customWidgets%2FcustomWidgetEmailDashboardSnapshot-py.yaml) – Take a snapshot of the current dashboard and send it to email recipients.

- [Send dashboard snapshot to Amazon S3](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetSnapshotDashboardToS3-py&template=customWidgets%2FcustomWidgetSnapshotDashboardToS3-py.yaml) – Take a snapshot of the current dashboard and store it in Amazon S3.

- [Text widget from CloudWatch dashboard](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetIncludeTextWidget-py&template=customWidgets%2FcustomWidgetIncludeTextWidget-py.yaml) – Displays the first text widget from the specified CloudWatch dashboard.

- [Display content of external URL](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetFetchURL-py&template=customWidgets%2FcustomWidgetFetchURL-py.yaml) – Displays the content of an externally accessible URL.

- [RSS reader](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetRssReader-py&template=customWidgets%2FcustomWidgetRssReader-py.yaml) – Displays RSS feeds.

- [Display an Amazon S3 object](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetS3GetObject-py&template=customWidgets%2FcustomWidgetS3GetObject-py.yaml) – Displays an object in an Amazon S3 bucket in your account.

- [Simple SVG pie chart](https://console.aws.amazon.com/cloudwatch/cfn.js?action=create&param_DoCreateExampleDashboard=Yes&region=us-east-1&stackName=customWidgetSimplePie-py&template=customWidgets%2FcustomWidgetSimplePie-py.yaml) – Example of a graphical SVG-based widget.

**Echo widget in JavaScript**

The following is the Echo sample widget in JavaScript.

```javascript

const DOCS = `
## Echo
A basic echo script. Anything passed in the \`\`\`echo\`\`\` parameter is returned as the content of the custom widget.
### Widget parameters
Param | Description
---|---
**echo** | The content to echo back

### Example parameters
\`\`\` yaml
echo: <h1>Hello world</h1>
\`\`\`
`;

exports.handler = async (event) => {
    if (event.describe) {
        return DOCS;
    }

    let widgetContext = JSON.stringify(event.widgetContext, null, 4);
    widgetContext = widgetContext.replace(/</g, '&lt;');
    widgetContext = widgetContext.replace(/>/g, '&gt;');

    return `${event.echo || ''}<pre>${widgetContext}</pre>`;
};
```

**Echo widget in Python**

The following is the Echo sample widget in Python.

````python

import json

DOCS = """
## Echo
A basic echo script. Anything passed in the ```echo``` parameter is returned as the content of the custom widget.
### Widget parameters
Param | Description
---|---
**echo** | The content to echo back

### Example parameters
``` yaml
echo: <h1>Hello world</h1>
```"""

def lambda_handler(event, context):
    if 'describe' in event:
        return DOCS

    echo = event.get('echo', '')
    widgetContext = event.get('widgetContext')
    widgetContext = json.dumps(widgetContext, indent=4)
    widgetContext = widgetContext.replace('<', '&lt;')
    widgetContext = widgetContext.replace('>', '&gt;')

    return f'{echo}<pre>{widgetContext}</pre>'
````

**Echo widget in Java**

The following is the Echo sample widget in Java.

````java

package example;

import com.amazonaws.services.lambda.runtime.Context;
import com.amazonaws.services.lambda.runtime.RequestHandler;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;

public class Handler implements RequestHandler<Event, String>{

  static String DOCS = ""
    + "## Echo\n"
    + "A basic echo script. Anything passed in the ```echo``` parameter is returned as the content of the custom widget.\n"
    + "### Widget parameters\n"
    + "Param | Description\n"
    + "---|---\n"
    + "**echo** | The content to echo back\n\n"
    + "### Example parameters\n"
    + "```yaml\n"
    + "echo: <h1>Hello world</h1>\n"
    + "```\n";

  Gson gson = new GsonBuilder().setPrettyPrinting().create();

  @Override
  public String handleRequest(Event event, Context context) {

    if (event.describe) {
      return DOCS;
    }

    return (event.echo != null ? event.echo : "") + "<pre>" + gson.toJson(event.widgetContext) + "</pre>";
  }
}

class Event {

    public boolean describe;
    public String echo;
    public Object widgetContext;

    public Event() {}

    public Event(String echo, boolean describe, Object widgetContext) {
        this.describe = describe;
        this.echo = echo;
        this.widgetContext = widgetContext;
    }
}
````

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Creating a custom widget

Adding a text widget
