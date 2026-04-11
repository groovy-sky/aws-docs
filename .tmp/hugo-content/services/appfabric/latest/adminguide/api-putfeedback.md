# PutFeedback

The AWS AppFabric for productivity feature is in preview and is subject to change.

Allows users to submit feedback for a given insight or action.

###### Topics

- [Request body](#API_PutFeedback_request)

- [Response elements](#API_PutFeedback_response)

## Request body

The request accepts the following data in JSON format.

ParameterDescription

**id**

The ID of the object for which feedback is being
submitted. This can be either the InsightId or the
ActionId.

**feedbackFor**

The insight type for which the feedback is being
submitted.

Possible values: `ACTIONABLE_INSIGHT |
                                            MEETING_INSIGHT | ACTION`

**feedbackRating**

Feedback Rating from `1` to `5`.
Higher rating the better.

## Response elements

If the action is successful, the service sends back an HTTP 201 response with
an empty HTTP body.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListMeetingInsights

Token

All content copied from https://docs.aws.amazon.com/.
