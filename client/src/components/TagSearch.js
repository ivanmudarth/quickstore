import React, { useState } from "react";
import {
  Input,
  Button,
  VStack,
  FormLabel,
  FormControl,
} from "@chakra-ui/react";
import { processTagInput } from "../utils";

function TagSearch(props) {
  const [rawTagInput, setTagInput] = useState("");

  function handleSearch(event) {
    event.preventDefault();

    // clean input and send it to FileDisplay component
    const userTags = processTagInput(rawTagInput);
    props.getSearchInput(userTags);
  }

  function handleTagChange(event) {
    setTagInput(event.target.value);
  }

  return (
    <form onSubmit={handleSearch}>
      <VStack spacing={2} align="baseline">
        <FormControl isRequired>
          <FormLabel>Filter files by tag:</FormLabel>
          <Input
            htmlSize={39} // TODO: make same as upload input
            width="auto"
            placeholder="Enter a comma separated list"
            onChange={handleTagChange}
          ></Input>
        </FormControl>
        <Button type="submit" colorScheme="blue" variant="solid">
          Search
        </Button>
      </VStack>
    </form>
  );
}

export default TagSearch;
