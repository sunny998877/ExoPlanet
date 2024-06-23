# Exoplanet Service

This is a Golang microservice for managing exoplanets, with functionalities to add, list, retrieve, update, delete, and estimate fuel for trips to exoplanets.

## Features

- **Add an Exoplanet**: Add a new exoplanet with details like name, description, distance from Earth, radius, mass (for terrestrial planets), and type (GasGiant or Terrestrial).
- **List Exoplanets**: Retrieve a list of all available exoplanets.
- **Get Exoplanet by ID**: Retrieve information about a specific exoplanet by its unique ID.
- **Update Exoplanet**: Update the details of an existing exoplanet.
- **Delete Exoplanet**: Remove an exoplanet from the catalog.
- **Fuel Estimation**: Calculate the fuel cost estimation for a trip to a particular exoplanet for a given crew capacity.

## Setup

### Prerequisites

- Go 1.18 or higher
- Docker (optional, for containerization)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/exoplanet-service.git
   cd exoplanet-service
