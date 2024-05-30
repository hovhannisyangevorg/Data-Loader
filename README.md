# Golang API Data Processing Client

This Golang application fetches data from the API Ninjas API, processes it, and saves the processed data to an AWS S3 bucket. The application is designed for efficiency and ease of configuration, ensuring seamless data handling and storage.

## Features

- Fetch data from the API Ninjas API endpoint
- Process the fetched data
- Save processed data to an AWS S3 bucket
- Easy configuration using environment variables

## Requirements

- Golang 1.16 or later
- AWS account with S3 access
- API Ninjas API key

## Installation

1. **Clone the repository:**

    ```sh
    git clone git@github.com:hovhannisyangevorg/Data-Loade.git
    cd Data-Procesor
    ```

2. **Install dependencies:**

    ```sh
    go mod tidy
    ```

3. **Set up environment variables:**

    Create a `.env` file in the root of the project with the following content:

    ```dotenv
    PATH="/Users/gevorghovhannisyan/Documents/myProject/resources/json"
    CORE_URL="https://api.api-ninjas.com/v1/animals?name=%s"
    API_KEY=""
    AWS_KEY=""
    AWS_SECRET=""
    AWS_REGION=""
    AWS_S3_BUCKET_NAME="your-aws-S3storage-name"
    ```

## Usage

1. **Run the application:**

    ```sh
    go run ./cmd/client/main.go
    ```

    The application will fetch animal data from the API Ninjas API, process it, and save the processed data to the configured S3 bucket.

## Configuration

- **PATH**: The path to the resources directory.
- **CORE_URL**: The API Ninjas endpoint to fetch data from, with a placeholder for the query parameter.
- **API_KEY**: Your API Ninjas API key.
- **AWS_KEY**: Your AWS access key ID.
- **AWS_SECRET**: Your AWS secret access key.
- **AWS_REGION**: The AWS region where your S3 bucket is located.
- **AWS_S3_BUCKET_NAME**: The name of your S3 bucket.

## Example

Below is an example of how to configure and run the application.

1. **Update the .env file:**

    ```dotenv
    PATH="/Users/gevorghovhannisyan/Documents/myProject/resources/json"
    CORE_URL="https://api.api-ninjas.com/v1/animals?name=%s"
    API_KEY="your-api-ninjas-api-key"
    AWS_KEY="your-aws-access-key"
    AWS_SECRET="your-aws-secret-key"
    AWS_REGION="your-aws-region"
    AWS_S3_BUCKET_NAME="your-aws-S3storage-name"
    ```

2. **Run the application:**

    ```sh
    go run ./cmd/client/main.go
    ```

     You should see logs indicating that the animal data has been successfully fetched from the API Ninjas API, processed, and uploaded to your S3 bucket.
