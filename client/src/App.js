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
} from "@chakra-ui/react";
import axios from "axios";

function App() {
  const [file, setFile] = useState();
  const [urls, setURLs] = useState([]);

  useEffect(() => {
    handleDisplay();
  }, []);

  function handleChange(event) {
    setFile(event.target.files[0]);
  }

  function handleUpload(event) {
    // TODO:
    // limit number of files to add
    // case where no file is selected
    // limit type of file to upload
    // upload progress bar / disable input/submit while uploading
    // display error message if error in upload
    event.preventDefault();
    const url = "http://localhost:8080/upload";
    const formData = new FormData();
    formData.append("file", file);
    formData.append("fileName", file.name);
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
        setURLs(response.data);
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
              <FormControl>
                <VStack spacing={2} align="baseline">
                  <FormLabel>Upload a File:</FormLabel>
                  <Input type="file" onChange={handleChange} />
                  <Button type="submit" colorScheme="blue" variant="solid">
                    Upload
                  </Button>
                </VStack>
              </FormControl>
            </form>
            <VStack spacing={4} align="baseline">
              <Text>Uploaded Files:</Text>
              {urls.map((url) => (
                <img key={url} src={url} width={200} />
              ))}
            </VStack>
          </VStack>
        </Box>
      </Center>
    </ChakraProvider>
  );
}

export default App;
