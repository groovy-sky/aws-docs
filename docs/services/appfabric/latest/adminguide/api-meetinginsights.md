# MeetingInsights

The AWS AppFabric for productivity feature is in preview and is subject to change.

Contains a summary of the top 3 meetings along with meeting purpose, related
cross-app artifacts, and activities from tasks, emails, messages, and calendar
events.

ParameterDescription

**insightId**

The unique id for the generated insight.

**insightContent**

The description of the insight highlighting the details in a
string format. As in, why is this insight important.

**insightTitle**

The title of the generated insight.

**createdAt**

When the insight was generated.

**calendarEvent**

The important calendar event or meeting that the user should
focus on.

Calendar Event object:

- `startTime` — The start time of the
event.

- `endTime` — The end time of the
event.

- `eventUrl` — The URL for the
calendar event on the ISV app.

**resources**

The list containing the other resources related to the
generate the insight.

Resource object:

- `appName` — The app name to which
the resource belongs.

- `resourceTitle` — The resource
title.

- `resourceType` — The type of the
resource.

The possible values are: `EMAIL | EVENT | MESSAGE
                                                  | TASK`

- `resourceUrl` — The resource URL in
the app.

- `appIconUrl` — The image URL of the
app to which the resource belongs.

**nextToken**

The pagination token to fetch the next set of insights. It’s
an optional field which if returned null means there are no more
insights to load.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AppClientSummary

VerificationDetails

All content copied from https://docs.aws.amazon.com/.
