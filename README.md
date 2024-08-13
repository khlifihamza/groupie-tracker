# Groupie Tracker

Groupie Tracker is a web application that displays information about music bands and artists using data from a given API. The project is built using Go for the backend and HTML/CSS for the frontend.

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [API Structure](#api-structure)
- [Contributing](#contributing)

## Features

- Display information about bands and artists, including:
  - Name(s)
  - Image
  - Formation year
  - First album date
  - Members
- Show last and/or upcoming concert locations
- Display last and/or upcoming concert dates
- Implement data visualizations for easy information consumption
- Feature client-server interaction for dynamic data retrieval

## Technologies Used

- Backend: Go
- Frontend: HTML, CSS
- API Integration

## Getting Started

### Prerequisites

- Go (version 1.20 or later)
- Web browser

### Installation

1. Clone the repository:
   ```
   git clone https://learn.zone01oujda.ma/git/hkhlifi/groupie-tracker.git
   ```

2. Navigate to the project directory:
   ```
   cd groupie-tracker
   ```

3. Build and run the application:
   ```
   go run main.go
   ```

4. Open your web browser and visit `http://localhost:8080`.

## Usage

you can click on the card of any artist/band and you gonna get more information about it 

## API Structure

The project uses an API that consists of four main parts:

1. **Artists**: Contains information about bands and artists (names, image, formation year, first album date, members).
2. **Locations**: Provides last and/or upcoming concert locations.
3. **Dates**: Includes last and/or upcoming concert dates.
4. **Relation**: Links all other parts (artists, dates, and locations).


## Contributing

Contributions to the Groupie Tracker project are welcome. Please follow these steps to contribute:

1. Fork the repository
2. Create a new branch
3. Make your changes and commit them
4. Push to your fork and submit a pull request
