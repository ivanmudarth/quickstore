import "./App.css";
import React, { useState, useEffect } from "react";
import {
  ChakraProvider,
  Box,
  Center,
  Input,
  Button,
  VStack,
  FormLabel,
  FormControl,
  Text,
  Spacer,
} from "@chakra-ui/react";
import axios from "axios";

function App() {
  const [file, setFile] = useState();
  const [displayInfo, setDisplay] = useState([]);
  const [rawTagInput, setTagInput] = useState("");

  // TODO: input requirements
  // username must be unique (check)
  // filename, each tag must be less than / equal to 15 chars

  useEffect(() => {
    handleDisplay();
  }, []);

  function handleFileChange(event) {
    setFile(event.target.files[0]);
  }

  function handleTagChange(event) {
    setTagInput(event.target.value);
  }

  function processTagInput(rawInput) {
    return rawInput.split(/[ \t]*,[ \t]*/);
  }

  function handleUpload(event) {
    // TODO:
    // limit number of files to add
    // case where no file is selected
    // limit type of file to upload
    // upload progress bar / disable input/submit while uploading
    // display error message if error in upload
    // user generated tags
    // clear input values on submit
    // ensure user tag input is valid (no invalid chars)
    event.preventDefault();

    const formData = new FormData();
    const userTags = processTagInput(rawTagInput);

    formData.append("file", file);
    formData.append("fileName", file.name);
    userTags.forEach((tag) => {
      formData.append("tags[]", tag);
    });
    formData.append("tags", userTags);

    const url = "http://localhost:8080/upload";
    const config = { headers: { "content-type": "multipart/form-data" } };
    axios
      .post(url, formData, config)
      .then((response) => {
        console.log(response.data);
      })
      .catch((error) => {
        console.log(error);
      });
  }

  function handleDisplay() {
    const url = "http://localhost:8080/display";

    axios
      .get(url)
      .then((response) => {
        setDisplay(response.data);
      })
      .catch((error) => {
        console.log(error);
      });
  }

  return (
    <ChakraProvider>
      <Center>
        <Box>
          <VStack spacing={4} align="baseline">
            <form onSubmit={handleUpload}>
              <VStack spacing={2} align="baseline">
                <FormControl isRequired>
                  <FormLabel>Upload a file:</FormLabel>
                  <Input type="file" onChange={handleFileChange} />
                </FormControl>
                <FormControl>
                  <FormLabel>Add tags to your file:</FormLabel>
                  <Input
                    placeholder="Enter a comma separated list"
                    onChange={handleTagChange}
                  ></Input>
                </FormControl>
                <Button type="submit" colorScheme="blue" variant="solid">
                  Upload
                </Button>
              </VStack>
            </form>
            <Spacer />
            <VStack spacing={4} align="baseline">
              <Text>Uploaded Files:</Text>
              {displayInfo?.map((info) => (
                <Box key={info["Key"]}>
                  <img src={info["URL"]} width={200} />
                  <Text>{info["Name"]}</Text>
                  <Text>{info["Size"]} MB</Text>
                  <Text>User Tags: {info["UserTags"].join(", ")}</Text>
                </Box>
              ))}
            </VStack>
          </VStack>
        </Box>
      </Center>
    </ChakraProvider>
  );
}

export default App;
