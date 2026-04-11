# Watching a database log file

Watching a database log file is equivalent to tailing the file on a UNIX or Linux system. You can watch a log file by using the AWS Management Console. RDS
refreshes the tail of the log every 5 seconds.

###### To watch a database log file

1. Open the Amazon RDS console at
    [https://console.aws.amazon.com/rds/](https://console.aws.amazon.com/rds).

2. In the navigation pane, choose **Databases**.

3. Choose the name of the DB instance that has the log file that you want to view.

4. Choose the **Logs & events** tab.

![Choose the Logs & events tab](https://docs.aws.amazon.com/images/AmazonRDS/latest/AuroraUserGuide/images/Monitoring_logsEvents.png)

5. In the **Logs** section, choose a log file, and then choose
    **Watch**.

![Choose a log](https://docs.aws.amazon.com/images/AmazonRDS/latest/AuroraUserGuide/images/Monitoring_LogsEvents_watch.png)

RDS shows the tail of the log, as in the following MySQL example.

![Tail of a log file](https://docs.aws.amazon.com/images/AmazonRDS/latest/AuroraUserGuide/images/Monitoring_LogsEvents_watch_content.png)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Downloading a database log file

Publishing to CloudWatch Logs

All content copied from https://docs.aws.amazon.com/.
