import React, { useState } from "react";
import { VStack, Spacer, Box, Center } from "@chakra-ui/react";
import UploadFile from "./UploadFile";
import FileDisplay from "./FileDisplay";
import UserInput from "./UserInput";
import TagSearch from "./TagSearch";

function Home() {
  // arbitrary value used as a dependency in useEffect in FileDisplay
  const [uploadComplete, setUploadComplete] = useState(0);
  const [searchInput, setSearchInput] = useState([]);

  function handleCallback(input) {
    setSearchInput(input);
  }

  return (
    <Center>
      <Box>
        <VStack spacing={4} align="baseline">
          <UserInput
            uploadComponent={
              <UploadFile finishUpload={() => setUploadComplete(1)} />
            }
            searchComponent={<TagSearch getSearchInput={handleCallback} />}
          />
          <Spacer />
          <FileDisplay
            uploadComplete={uploadComplete}
            searchInput={searchInput}
          />
        </VStack>
      </Box>
    </Center>
  );
}

export default Home;
