# Treaty-DailyTask

<!--toc:start-->
- [Treaty-DailyTask](#treaty-dailytask)
  - [What to build](#what-to-build)
    - [Problem](#problem)
    - [Solution](#solution)
  - [Where to build](#where-to-build)
    - [Platforms](#platforms)
    - [API Service & Database](#api-service-database)
  - [System Architecture _HIGH LEVEL DESIGN_](#system-architecture-high-level-design)
  - [Planning & System Design](#planning-system-design)
<!--toc:end-->

I program in my spare time until i finish the project

## What to build

### Problem

When tracking money spending on certain categories user needs to input amount and description manually via notes and after collectively inserting all the amount and description at the end of the week user summarizes all the spending in a google sheet document wherein it computes the total spent plus the predictive saving for the current month.

### Solution

Build a mobile application that supports Android and iOS which collects all the inputs of the amount and description/category of the user and automatically summarize all the spending and savings prediction for the current month.

## Where to build

### Platforms

- **React Native (Android / IOS)**
- **Web** _future projects most likely for an Admin build_

### API Service & Database

- MongoDB Realm _Client side DB_
- PostgreSQL _Server side DB_
- Google Sheets API (_Needs more R&D_)

## System Architecture _HIGH LEVEL DESIGN_

- Can do CRUD operation on category **price**, **description**
- Separated by Category UID
- Debounce offline save _background thread_
- Sync Save operation _background service task_
- Sync Retrieve operation _can be triggered manually or via app launch_
- Local Push Notification on actions made **time based operation**

## Planning & System Design

1. Story planning
    1. Features
        - Create Task
        - Retrieve Task
        - Update Task
        - Delete Task
        - Complete Task
        - Debounce save operation _cache strategy_
        - Local Push Notification _time base reminder_
        - Sync Save _Service API_
        - Sync Retrieve _Service API_
2. Development Architecture (Code) Base design
    1. Redux Dependency
    2. Realm Dependency
    3. Local Notification Dependency
    4. Testing Library Dependency
3. Layout & Design

App Eraser (HIGH LEVEL DESIGN) <https://app.eraser.io/workspace/1kydQuEFbIxhYKqdvoqB>

Figma Design (LAYOUT)
<https://www.figma.com/design/OFVeJ0ssLEdo2cbSf9ZGk3/Treaty-Daily-Design>

1. CI/CD Deployment Status
