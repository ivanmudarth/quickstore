import React, { useState } from "react";
import { VStack, Spacer, Box, Center } from "@chakra-ui/react";
import UploadFile from "./UploadFile";
import FileDisplay from "./FileDisplay";
import UserInput from "./UserInput";
import TagSearch from "./TagSearch";

function Home() {
  // arbitrary value used as a dependency in useEffect in FileDisplay
  const [uploadComplete, setUploadComplete] = useState<{
    uploadComplete: number;
  }>({ uploadComplete: 0 });
  const [searchInput, setSearchInput] = useState<{ searchInput: string[] }>({
    searchInput: [],
  });

  function handleCallback(input: string[]) {
    setSearchInput({ searchInput: input });
  }

  return (
    <Center>
      <Box>
        <VStack spacing={4} align="baseline">
          <UserInput
            components={{
              uploadComponent: (
                <UploadFile
                  finishUpload={() =>
                    setUploadComplete({
                      uploadComplete: uploadComplete.uploadComplete + 1,
                    })
                  }
                />
              ),
              searchComponent: <TagSearch getSearchInput={handleCallback} />,
            }}
          />
          <Spacer />
          <FileDisplay
            uploadComplete={uploadComplete.uploadComplete}
            searchInput={searchInput.searchInput}
          />
        </VStack>
      </Box>
    </Center>
  );
}

export default Home;