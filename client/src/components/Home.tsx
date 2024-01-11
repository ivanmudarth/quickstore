import React, { useState } from "react";
import { VStack, Spacer, Box, Center } from "@chakra-ui/react";
import Upload from "./Upload";
import ItemDisplay from "./ItemDisplay";
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
    <div>
      <Center>
        <Box>
          <VStack spacing={4} align="baseline">
            <UserInput
              components={{
                uploadComponent: (
                  <Upload
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
          </VStack>
        </Box>
      </Center>

      <ItemDisplay
        uploadComplete={uploadComplete.uploadComplete}
        searchInput={searchInput.searchInput}
      />
    </div>
  );
}

export default Home;
