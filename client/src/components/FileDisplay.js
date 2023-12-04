import React, { useState, useEffect } from "react";
import { VStack, Text } from "@chakra-ui/react";
import axios from "axios";
import File from "./File";

function FileDisplay() {
  const [displayInfo, setDisplay] = useState([]);

  useEffect(() => {
    handleDisplay();
  }, []);

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
    <VStack spacing={4} align="baseline">
      <Text>Uploaded Files:</Text>
      {displayInfo?.map((fileInfo, index) => (
        <File key={index} info={fileInfo} />
      ))}
    </VStack>
  );
}

export default FileDisplay;
