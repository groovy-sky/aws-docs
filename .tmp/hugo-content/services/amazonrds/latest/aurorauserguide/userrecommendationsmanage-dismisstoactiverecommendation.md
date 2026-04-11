# Modifying dismissed Amazon Aurora recommendations to active recommendations

You can move one or more dismissed Amazon Aurora recommendations to active recommendations using the Amazon RDS console, AWS CLI, or Amazon RDS API.

###### To move one or more dismissed recommendations to active recommendations

1. Sign in to the AWS Management Console and open the Amazon RDS console at
    [https://console.aws.amazon.com/rds/](https://console.aws.amazon.com/rds).

2. In the navigation pane, perform any of the following:

- Choose **Recommendations**.

The **Recommendations** page displays a list of
recommendations sorted by the severity for all the resources in your
account.

- Choose **Databases** and then choose **Recommendations**
for a resource in the databases page.

The **Recommendations** tab displays the recommendations and its details for the selected resource.

3. Choose one or more dismissed recommendations from the list and then choose
    **Move to active**.

![A few dismissed recommendations selected and Move to active button highlighted in the console](https://docs.aws.amazon.com/images/AmazonRDS/latest/AuroraUserGuide/images/Recommendations_DismissToActive.png)

A banner displays a successful or failure message when the moving the
    selected recommendations from dismissed to active status.

The following example shows the banner with the successful message.

![a banner in the console showing the message with the number of resources moved successfully from dismissed to active recommendations](https://docs.aws.amazon.com/images/AmazonRDS/latest/AuroraUserGuide/images/Recommendation-DismissToActive-Banner.png)

The following example shows the banner with the failure message.

![a banner in the console showing the message with the resource that failed to move from dismissed to active recommendations](https://docs.aws.amazon.com/images/AmazonRDS/latest/AuroraUserGuide/images/Recommendation-DismissToActive-Banner-Failure.png)

###### To change a dismissed Aurora recommendation to active recommendation using the AWS CLI

1. Run the command `aws rds describe-db-recommendations --filters "Name=status,Values=dismissed"`.

The output provides a list of recommendations in `dismissed` status.

2. Find the `recommendationId` for the recommendation that you want to change the status from step 1.

3. Run the command `>aws rds modify-db-recommendation --status active --recommendationId <ID>` with the
    `recommendationId` from step 2 to change to active recommendation.

To change a dismissed Aurora recommendation to active recommendation using the Amazon RDS API, use the [ModifyDBRecommendation](../../../../reference/amazonrds/latest/apireference/api-modifydbrecommendation.md) operation.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Dismissing recommendations

Recommendations reference

All content copied from https://docs.aws.amazon.com/.
