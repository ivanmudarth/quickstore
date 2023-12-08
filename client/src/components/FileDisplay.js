import React, { useState, useEffect } from "react";
import { VStack, Text } from "@chakra-ui/react";
import axios from "axios";
import File from "./File";

// TODO: bug - search a tag. upload a file. display won't update cause searchInput still has a value
// maybe have two different effects with different dependencies
function FileDisplay(props) {
  const [displayInfo, setDisplay] = useState([]);

  // is called every time uploadComplete prop is updated
  useEffect(() => {
    handleDisplay(props.searchInput);
  }, [props.uploadComplete, props.searchInput]);

  function handleDisplay(searchInput) {
    const params = { tags: searchInput };
    const url = "http://localhost:8080/display";
    axios
      .get(url, { params })
      .then((response) => {
        console.log(response.data);
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
