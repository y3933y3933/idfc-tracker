<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#getting-started">Getting Started</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

## About The Project

IDFC Tracker is a simple command-line tool to track progress towards a goal by adding points over time. The user can create a profile, set a goal, and incrementally add points. Every point added is logged with a reason, and users can view their history, set a new goal, or export their records.



### Built With

* [Go](https://golang.org/)
* [Cobra](https://github.com/spf13/cobra)
* [SQLite](https://www.sqlite.org/)
* [pterm](https://github.com/pterm/pterm)
* [sqlc](https://github.com/kyleconroy/sqlc)



## Getting Started

This guide will help you set up and run IDFC Tracker on your local machine.

### Prerequisites

Make sure you have [Go](https://golang.org/) installed on your system.


### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/y3933y3933/idfc-tracker.git
   ```
2. Navigate into the project directory:
   ```bash
   cd idfc-tracker
   ```
3. Install the dependencies:
   ```bash
   go install
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Usage

### Commands

#### `init` - Initialize User
Create a new user and set an initial goal.
```bash
$ idfc-tracker init
```
**Example**:
```bash
Enter your name: Joanne
Set your initial goal points: 100
```

#### `add` - Add Points
Add points to the current user's progress and record the reason.
```bash
$ idfc-tracker add
```
**Example**:
```bash
How many points would you like to add? 10
Reason for adding points: Completed a task
```

#### `clear` - Clear Points
Clear all points and history for the active user.
```bash
$ idfc-tracker clear
```

#### `export` - Export History
Export history records to JSON or CSV.
```bash
$ idfc-tracker export --json
$ idfc-tracker export --csv
```

#### `history` - View History
View the user's history records. Supports date filtering.
```bash
$ idfc-tracker history
$ idfc-tracker history --start "2025-01-01" --end "2025-01-31"
```

#### `set --goal` - Set Goal
Set or update the current goal points for the active user.
```bash
$ idfc-tracker set --goal 150
```

#### `status` - Show Current Progress
Show the current total points and goal progress as a progress bar.
```bash
$ idfc-tracker status
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Roadmap

- [ ] Feature 1: Add multi-user support
- [ ] Feature 2: Implement data visualization (charts, progress graphs)
- [ ] Feature 3: Integrate notifications for reminders

See the [open issues](https://github.com/your-username/idfc-tracker/issues) for a full list of proposed features and known issues.


## Acknowledgments

* Thanks to all open-source contributors for inspiration.
* Special thanks to [Cobra](https://github.com/spf13/cobra) for the CLI framework.

