import "./App.css";
import React, { useState } from "react";
import {
  ChakraProvider,
  Box,
  Center,
  Input,
  Button,
  VStack,
  FormLabel,
  FormControl,
} from "@chakra-ui/react";
import axios from "axios";

function App() {
  const [file, setFile] = useState();

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

  return (
    <ChakraProvider>
      <Center>
        <Box>
          <VStack spacing={4} align="baseline">
            <form onSubmit={handleUpload}>
              <FormControl>
                <FormLabel>Upload file</FormLabel>
                <Input type="file" onChange={handleChange} />
                <Button type="submit" colorScheme="blue" variant="solid">
                  Upload
                </Button>
              </FormControl>
            </form>
          </VStack>
        </Box>
      </Center>
    </ChakraProvider>
  );
}

export default App;
