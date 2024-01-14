<p align="center">
<img width="781" alt="Screen Shot 2024-01-12 at 12 34 30 AM" src="https://github.com/ivanmudarth/quickstore/assets/33183884/17218b03-e84b-40d7-ae79-1e83e15929e5">
</p>

# Quickstore

Quickstore lets you upload files and website links then effortlessly find them later using auto-tagging functionality 

## Features

### Upload / Store

Upload files (image, PDF) or website URLs using the input form and add any tags to describe what you're uploading. These items are stored in AWS S3 with their metadata in MySQL. 

<img src="https://github.com/ivanmudarth/quickstore/assets/33183884/5828ccc2-82d2-4aad-a253-37c08a172643" width="800" />

### Tagging

In addition to user-inputted tags, the contents of these files are auto-tagged on upload. 

Images are tagged using computer vision technology from the Imagga API while PDFs and websites are scraped and tagged using the TextRank algorithm. 

User tags are displayed in grey while auto tags are in green:

<img width="800" alt="Screen Shot 2024-01-13 at 10 47 53 PM" src="https://github.com/ivanmudarth/quickstore/assets/33183884/bda8335a-2431-470d-95cb-cbe5fad9b137">


### Search

User and auto tags allow files and URLs to be searchable. A list of tags can be queried and items with matching tags will be returned. 

<img src="https://github.com/ivanmudarth/quickstore/assets/33183884/59adb367-3c0d-4641-9fdb-7cea3a7a0554" width="800" />


## Technologies

- React Typescript frontend
- Shadcn and Chakra UI for component styling
- Golang backend
- Imagga API for image tagging
- Opengraph.io API for website preview data
- AWS Localstack S3 for storing images and PDFs
- MySQL for storing metadata about files and websites

## Installation

### Prerequisites 

- [Localstack CLI](https://docs.localstack.cloud/getting-started/installation/#localstack-cli)
- MySQL Workbench 8.0+
- Golang version go1.21.4+
- Node v20.7.0+
- npm 10.1.0+
- [shadcn/ui CLI](https://ui.shadcn.com/docs/cli)

Complete the following steps

1. Install dependencies for the client:

```
cd client
npm install
```

2. Install dependencies for the server:
```
cd server
go get -d ./...
```

3. After installing MySQL Workbench, create a database with a name, user, and password that you set. Create the env file `server/.env` using the same structure as `server/.env.example`. Add your database name, user, and password here. 

4. Create a free account on [Imagga](https://imagga.com/) and [Opengraph.io](https://www.opengraph.io/). Add your API keys to the env file. 

## Run Locally

Complete the following steps, each from the separate terminal tab

1. Start up Localstack on Docker
```
localstack start
```

2. Start the Golang server
```
cd server
go run main.go
```

3. Run the client
```
cd client
npm start
```

