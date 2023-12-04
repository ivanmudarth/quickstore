package database

import (
	"log"
	"os"
)

func createUserTable() {
	// One to many relationship with file
	_, err := DB.Exec(` 
		CREATE TABLE IF NOT EXISTS User (
			UserID INT NOT NULL AUTO_INCREMENT, 
			Username VARCHAR(15) NOT NULL UNIQUE,
			PRIMARY KEY (UserID)
		);`)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		log.Println("Database table User created or already exists")
	}
}

func createFileTable() {
	// TODO: make UserID required once users are supported
	// One to many relationship with tag
	_, err := DB.Exec(` 
		CREATE TABLE IF NOT EXISTS File (
			FileID INT NOT NULL AUTO_INCREMENT, 
			UserID INT,
			S3Key VARCHAR(36) NOT NULL,
			Name VARCHAR(30) NOT NULL,
			Size DECIMAL(10, 2) NOT NULL, -- NOTE: in MB
			Type VARCHAR(30),
			UploadTime DATETIME DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (FileID),
			FOREIGN KEY (UserID) REFERENCES User(UserID)
		);`)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		log.Println("Database table File created or already exists")
	}
}

func createTagTable() {
	_, err := DB.Exec(` 
		CREATE TABLE IF NOT EXISTS Tag (
			TagID INT NOT NULL AUTO_INCREMENT, 
			FileID INT NOT NULL,
			Name VARCHAR(40) NOT NULL,
			Type ENUM("User", "Auto") NOT NULL,
			PRIMARY KEY (TagID),
			FOREIGN KEY (FileID) REFERENCES File(FileID)
		);`)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		log.Println("Database table Tag created or already exists")
	}
}

func CreateAllTables() {
	createUserTable()
	createFileTable()
	createTagTable()
}
